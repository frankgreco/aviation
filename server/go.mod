module github.com/frankgreco/aviation/server

go 1.15

require (
	github.com/Masterminds/squirrel v1.5.0
	github.com/frankgreco/aviation/internal v0.0.0
	github.com/frankgreco/aviation/types v0.0.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.1.0
	github.com/jackc/pgx/v4 v4.10.1
	github.com/jmoiron/sqlx v1.2.0
	github.com/oklog/run v1.1.0 // indirect
	go.uber.org/zap v1.10.0
	google.golang.org/grpc v1.34.0
)

replace (
	github.com/frankgreco/aviation/internal => ../internal
	github.com/frankgreco/aviation/types => ../types
)
