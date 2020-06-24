package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"strings"

	"github.com/Masterminds/squirrel"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/jmoiron/sqlx"
	"github.com/kyokomi/emoji"
	log "github.com/sirupsen/logrus"
	"golang.org/x/text/language"
	"golang.org/x/text/message"

	"github.com/frankgreco/aviation"
	"github.com/frankgreco/aviation/api"
	"github.com/frankgreco/aviation/utils/db"
)

var (
	psq   = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
	dbURL = fmt.Sprintf("postgres://%s:%s@%s:5432/aviation", os.Getenv("RDS_USERNAME"), os.Getenv("RDS_PASSWORD"), os.Getenv("RDS_ENDPOINT"))
)

func main() {
	lambda.Start(do)
}

func do(ctx context.Context) error {
	// Twitter client
	client := twitter.NewClient(oauth1.NewConfig(os.Getenv("TWITTER_CONSUMER_KEY"), os.Getenv("TWITTER_CONSUMER_SECRET")).Client(
		oauth1.NoContext,
		oauth1.NewToken(os.Getenv("TWITTER_ACCESS_TOKEN"), os.Getenv("TWITTER_ACCESS_SECRET")),
	))

	dbase, err := db.New(dbURL, log.New().WithFields(log.Fields{
		"app": "database",
	}))
	if err != nil {
		return err
	}
	defer dbase.Close()

	newStuff := []struct {
		aircraft     api.Aircraft
		registration api.Registration
	}{}

	if _, err := dbase.QueryRowsTx(context.Background(), nil, db.QueryScan{
		Name: "retrieve all new registrations",
		Query: psq.Select("a.manufactuer_name, a.model_name, a.series, r.serial_number, r.year_manufactured, r.registrant_name, a.num_engines, a.num_seats, a.weight, a.cruising_speed").
			From("aviation.registration r").
			JoinClause("NATURAL JOIN aviation.aircraft a").
			Where(squirrel.Eq{
				"created::date": aviation.Date,
			}),
		Callback: db.ScanFunc(func(rows *sqlx.Rows) (arr []interface{}, err error) {
			for {
				var stuff struct {
					aircraft     api.Aircraft
					registration api.Registration
				}
				if err = rows.Scan(&stuff.aircraft.ManufacturerName, &stuff.aircraft.ModelName, &stuff.aircraft.Series, &stuff.registration.SerialNumber, &stuff.registration.YearManufactured, &stuff.registration.RegistrantName, &stuff.aircraft.NumEngines, &stuff.aircraft.NumSeats, &stuff.aircraft.Weight, &stuff.aircraft.CruisingSpeed); err == sql.ErrNoRows {
					return nil, nil
				}
				if err != nil {
					return
				}
				newStuff = append(newStuff, stuff)
				if !rows.Next() {
					break
				}
			}
			return
		}),
	}); err != nil && err != db.ErrNotFound {
		dbase.Close()
		return err
	}

	p := message.NewPrinter(language.English)

	tweet, _, err := client.Statuses.Update(
		p.Sprintf(
			"There were %d new registrations with the #FAA today! Details %s", len(newStuff), emoji.Sprintf(":backhand_index_pointing_down:"),
		), nil)
	if err != nil {
		log.Error("could not send tweet")
		return err
	}

	for _, stuff := range newStuff {
		_, _, err = client.Statuses.Update(
			p.Sprintf(
				"Make: %s\nModel: %s\nSeries: %s\nSerial Number: %s\nOwner: %s\nEngines: %s\nSeats: %s",
				stuff.aircraft.ManufacturerName,
				stuff.aircraft.ModelName,
				stuff.aircraft.Series,
				stuff.registration.SerialNumber,
				func() string {
					owner := strings.Title(strings.ToLower(stuff.registration.RegistrantName))
					if owner == "" {
						return "unknown"
					}
					return owner
				}(),
				strings.TrimLeft(stuff.aircraft.NumEngines, "0"),
				strings.TrimLeft(stuff.aircraft.NumSeats, "0"),
			), &twitter.StatusUpdateParams{
				InReplyToStatusID: tweet.ID,
			})
		if err != nil {
			log.Error("could not send tweet")
			return err
		}
	}
	return nil
}
