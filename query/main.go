package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/akrylysov/algnhsa"
)

func main() {
	var s Service
	{
		db, err := sql.Open("pgx", fmt.Sprintf("postgres://%s:%s@%s:5432/aviation", os.Getenv("RDS_USERNAME"), os.Getenv("RDS_PASSWORD"), os.Getenv("RDS_ENDPOINT")))
		if err != nil {
			panic(err) // TODO: fix me!
		}
		defer db.Close()
		db.SetMaxOpenConns(1)

		s = NewService(db)
	}

	var h http.Handler
	{
		h = MakeHTTPHandler(s)
	}

	algnhsa.ListenAndServe(h, nil)
}
