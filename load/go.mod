module github.com/frankgreco/aviation/load

require (
	github.com/Masterminds/squirrel v1.4.0
	github.com/aws/aws-sdk-go v1.32.7
	github.com/coreos/go-systemd v0.0.0-20190719114852-fd7a80b32e1f // indirect
	github.com/frankgreco/aviation v0.0.0
	github.com/frankgreco/aviation/api v0.0.0
	github.com/frankgreco/aviation/utils v0.0.0
	github.com/jackc/pgx/v4 v4.6.0 // indirect
	github.com/jmoiron/sqlx v1.2.0
	github.com/kr/pty v1.1.8 // indirect
	github.com/sirupsen/logrus v1.6.0
	github.com/stretchr/objx v0.2.0 // indirect
	golang.org/x/tools v0.0.0-20190823170909-c4a336ef6a2f // indirect
)

replace (
	github.com/frankgreco/aviation => ../
	github.com/frankgreco/aviation/api => ../api
	github.com/frankgreco/aviation/utils => ../utils
)
