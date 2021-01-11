package main

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"

	"github.com/frankgreco/aviation/internal/db"
	"github.com/frankgreco/aviation/internal/log"
	"github.com/frankgreco/aviation/types"
)

var (
	psq = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
)

type suggestionServer struct {
	*suggestionServerConfig
	types.UnimplementedSuggestionServiceServer
}

type suggestionServerConfig struct {
	db     *db.DB
	logger log.Logger
}

func newSuggestionServer(config *suggestionServerConfig) types.SuggestionServiceServer {
	return &suggestionServer{
		suggestionServerConfig: config,
	}
}

func (s *suggestionServer) ListSuggestions(ctx context.Context, req *types.ListSuggestionsRequest) (*types.ListSuggestionsReply, error) {
	toReturn := &types.ListSuggestionsReply{
		Suggestions: []string{},
		Size:        req.Size,
	}

	requestedType, columnToReturn, requestedValue, err := filter(req.Requested)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	toReturn.Type = requestedType

	existingFilters := map[string]interface{}{}
	{
		for _, f := range req.Existing {
			_, column, value, err := filter(f)
			if err != nil {
				s.logger.Warn(err.Error())
				continue
			}
			existingFilters[column] = value
		}
	}

	query := psq.Select(columnToReturn).
		From("aviation.registration r").
		Join("aviation.aircraft a ON r.aircraft_id = a.id").
		Where(db.When(squirrel.Like{columnToReturn: fmt.Sprintf("%s%%", requestedValue)}, requestedValue != "")).
		Where(db.When(squirrel.Eq(existingFilters), req.Existing != nil && len(req.Existing) > 0))

	if req.Size > 0 {
		query = query.Limit(uint64(req.Size))
	}

	if _, err := s.db.QueryRowsTx(ctx, nil, db.QueryScan{
		Name:  "get suggestions",
		Query: query,
		Callback: func(out *[]string) db.ScanFunc {
			return db.ScanFunc(func(rows *sqlx.Rows) (arr []interface{}, err error) {
				if out == nil {
					out = &[]string{}
				}
				for {
					var suggestion string
					if err = rows.Scan(&suggestion); err == sql.ErrNoRows {
						err = db.ErrNotFound
					}
					if err != nil {
						return
					}
					arr = append(arr, suggestion)
					*out = append(*out, suggestion)
					if !rows.Next() {
						break
					}
				}
				return
			})
		}(&toReturn.Suggestions),
	}); err != nil && err != db.ErrNotFound {
		s.logger.Error(err.Error())
		return nil, err
	}

	if req.Size <= 0 {
		toReturn.Size = int32(len(toReturn.Suggestions))
	}

	return toReturn, nil
}

func filter(f string) (t types.FilterType, column, value string, err error) {
	arr := strings.Split(f, "=")
	if arr == nil || len(arr) != 2 {
		err = fmt.Errorf("filter [%s] is malformed", f)
		return
	}
	switch arr[0] {
	case types.FilterType_N_NUMBER.String():
		column = "r.tail_number"
		t = types.FilterType_N_NUMBER
	case types.FilterType_MAKE.String():
		column = "a.make"
		t = types.FilterType_MAKE
	case types.FilterType_MODEL.String():
		column = "a.model"
		t = types.FilterType_MODEL
	case types.FilterType_AIRLINE.String():
		column = "a.model" // todo
		t = types.FilterType_AIRLINE
	default:
		err = fmt.Errorf("unknown filter [%s]", arr[0])
		return
	}
	value = arr[1]
	return
}
