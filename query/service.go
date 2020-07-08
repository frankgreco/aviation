package main

import (
	"context"
	"database/sql"

	"github.com/frankgreco/aviation/api"
	"github.com/frankgreco/aviation/search"
)

type Service interface {
	Search(context.Context, string, int) ([]api.SearchResult, error)
}

type service struct {
	db *sql.DB
}

func NewService(db *sql.DB) Service {
	return &service{db}
}

func (s *service) Search(ctx context.Context, query string, limit int) ([]api.SearchResult, error) {
	return search.New(append(search.Parse(query), search.Database(s.db), search.Limit(limit))...)
}
