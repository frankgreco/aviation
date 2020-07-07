module github.com/frankgreco/aviation/search

require (
	github.com/Masterminds/squirrel v1.4.0
	github.com/coreos/go-systemd v0.0.0-20190719114852-fd7a80b32e1f // indirect
	github.com/frankgreco/aviation/api v0.0.0
	github.com/frankgreco/aviation/utils v0.0.0
	github.com/jackc/pgx/v4 v4.7.0 // indirect
	github.com/jmoiron/sqlx v1.2.0
	github.com/kr/pty v1.1.8 // indirect
	github.com/lib/pq v1.7.0 // indirect
	github.com/sirupsen/logrus v1.6.0
	github.com/spf13/cobra v1.0.0
	github.com/stretchr/objx v0.2.0 // indirect
	go.uber.org/multierr v1.5.0 // indirect
	golang.org/x/net v0.0.0-20190813141303-74dc4d7220e7 // indirect
)

replace (
	github.com/frankgreco/aviation/api => ../api
	github.com/frankgreco/aviation/utils => ../utils
)
