package main

import (
	"context"
	"fmt"
	"os"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	_ "github.com/jackc/pgx/v4/stdlib"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	"github.com/frankgreco/aviation/internal/db"
	grpcserver "github.com/frankgreco/aviation/internal/grpc"
	httpserver "github.com/frankgreco/aviation/internal/http"
	"github.com/frankgreco/aviation/internal/run"
	"github.com/frankgreco/aviation/types"
)

func main() {
	ctx := context.Background()
	logger := zap.NewExample().Sugar()

	var dbRunnable run.Runnable
	{
		dbRunnable = db.Prepare(&db.Options{
			ConnectionString: os.Getenv("CONNECTION_STRING"),
			Logger:           logger,
		})
	}

	var grpcRunnable run.Runnable
	{
		s := grpc.NewServer()
		types.RegisterSuggestionServiceServer(s, newSuggestionServer(&suggestionServerConfig{
			db: &db.DB{
				DB:     dbRunnable.(*db.DB).DB,
				Logger: logger,
			},
			logger: logger,
		}))
		grpcRunnable = grpcserver.Prepare(&grpcserver.Options{
			Server:     s,
			Network:    "tcp",
			Address:    "0.0.0.0:8082",
			Reflection: true,
			Logger:     logger,
		})
	}

	var httpRunnable run.Runnable
	{
		mux := runtime.NewServeMux()
		opts := []grpc.DialOption{grpc.WithInsecure()}
		if err := types.RegisterSuggestionServiceHandlerFromEndpoint(ctx, mux, ":8082", opts); err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
		httpRunnable = httpserver.Prepare(&httpserver.Options{
			Name:         "http server",
			Handler:      mux,
			InsecurePort: 9092,
			Logger:       logger,
		})
	}

	var g run.Group

	g.Add(run.Always, grpcRunnable)
	g.Add(run.Always, httpRunnable)
	g.Add(run.Always, dbRunnable)
	g.Add(run.Always, run.NewMonitor(ctx, logger))

	if err := g.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	os.Exit(0)
}
