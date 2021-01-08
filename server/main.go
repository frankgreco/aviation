package main

import (
	"os"
	"context"
	"fmt"
	"os"

	"google.golang.org/grpc"
	"github.com/jmoiron/sqlx"
	_ "github.com/jackc/pgx/v4/stdlib"
	"go.uber.org/zap"

	grpcserver "github.com/frankgreco/aviation/internal/grpc"
	"github.com/frankgreco/aviation/internal/run"
	"github.com/frankgreco/aviation/internal/db"
	"github.com/frankgreco/aviation/types"
)

func main() {
	var g run.Group

	conn, err := sqlx.Open("pgx", os.Getenv("CONNECTION_STRING"))
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	var grpcRunnable run.Runnable
	{
		s := grpc.NewServer()
		types.RegisterSuggestionServiceServer(s, newSuggestionServer(&suggestionServerConfig{
			db: &db.DB{
				DB: conn,
				Logger: zap.NewExample().Sugar(),
			},
		}))
		grpcRunnable = grpcserver.Prepare(&grpcserver.Options{
			Server:     s,
			Network:    "tcp",
			Address:    "0.0.0.0:8082",
			Reflection: true,
		})
	}

	g.Add(run.Always, grpcRunnable)
	g.Add(run.Always, run.NewMonitor(context.Background()))

	if err := g.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	os.Exit(0)
}
