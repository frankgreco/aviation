package main

import (
	"fmt"
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/Masterminds/squirrel"

	"github.com/frankgreco/aviation/internal/db"
	"github.com/frankgreco/aviation/types"
)

var (
	psq = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
)

type suggestionServer struct {
	config *suggestionServerConfig
}

type suggestionServerConfig struct {
	db *db.DB
}

func newSuggestionServer(config *suggestionServerConfig) types.SuggestionServiceServer {
	return &suggestionServer{
		config: config,
	}
}

func (s *suggestionServer) ListSuggestions(ctx context.Context, req *types.ListSuggestionsRequest) (*types.ListSuggestionsReply, error) {
	toReturn := &types.ListSuggestionsReply{
		Type: req.Requested.Type,
		Suggestions: []string{},
	}

	columnToReturn, err := typeToColumn(req.Requested.Type)
	if err != nil {
		return nil, err
	}

	existingFilters := map[string]interface{}{}
	{
		for _, filter := range req.Existing {
			column, err := typeToColumn(filter.Type)
			if err != nil {
				// TODO: log
				continue
			}
			existingFilters[column] = filter.Value
		}
	}

	column, err := typeToColumn(req.Requested.Type)
	if err != nil {
		// TODO: log
		return nil, err
	}
	if _, err := s.config.db.QueryRowsTx(ctx, nil, db.QueryScan{
		Name: "get suggestions",
		Query: psq.Select(columnToReturn).
			From("aviation.registration r").
			Join("aviation.aircraft a ON r.aircraft_id = a.id").
			Where(db.When(squirrel.Like{column: fmt.Sprintf("%s%%", req.Requested.Value)}, req.Requested.Value != "")).
			Where(db.When(squirrel.Eq(existingFilters), req.Existing != nil && len(req.Existing) > 0)),
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
	}); err != nil {
		if err == db.ErrNotFound {
			return toReturn, nil
		}
		return nil, err
	}

	return toReturn, nil
}

func typeToColumn(t types.FilterType) (string, error) {
	switch t {
	case types.FilterType_N_NUMBER:
		return "r.tail_number", nil
	case types.FilterType_MAKE:
		return "a.make", nil
	case types.FilterType_MODEL:
		return "a.model", nil
	case types.FilterType_AIRLINE:
		return "a.model", nil // todo
	default:
		return "", fmt.Errorf("%s", "todo")
	}
}
