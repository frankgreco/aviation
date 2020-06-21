module github.com/frankgreco/aviation/load

require (
	github.com/Masterminds/squirrel v1.4.0
	github.com/coreos/go-systemd v0.0.0-20190719114852-fd7a80b32e1f // indirect
	github.com/frankgreco/aviation/api v0.0.0
	github.com/frankgreco/aviation/utils v0.0.0
	github.com/go-kit/kit v0.10.0 // indirect
	github.com/jackc/pgx/v4 v4.6.0 // indirect
	github.com/jmoiron/sqlx v1.2.0 // indirect
	github.com/konsorten/go-windows-terminal-sequences v1.0.2 // indirect
	github.com/kr/pty v1.1.8 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/stretchr/objx v0.2.0 // indirect
)

replace (
	github.com/frankgreco/aviation/api => ../api
	github.com/frankgreco/aviation/utils => ../utils
)
