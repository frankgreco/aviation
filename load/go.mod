module github.com/frankgreco/aviation/load

go 1.21

toolchain go1.21.0

require (
	github.com/Masterminds/squirrel v1.5.0
	github.com/aws/aws-lambda-go v1.17.0
	github.com/aws/aws-sdk-go v1.32.7
	github.com/frankgreco/aviation v0.0.0
	github.com/frankgreco/aviation/api v0.0.0
	github.com/frankgreco/aviation/internal v0.0.0-20210108003342-f47da5804365
	github.com/jmoiron/sqlx v1.2.0
	github.com/sirupsen/logrus v1.6.0
)

require (
	github.com/frankgreco/aviation/utils v0.0.0-20200624030258-6c014fac6407 // indirect
	github.com/jackc/chunkreader/v2 v2.0.1 // indirect
	github.com/jackc/pgconn v1.8.0 // indirect
	github.com/jackc/pgio v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgproto3/v2 v2.0.6 // indirect
	github.com/jackc/pgservicefile v0.0.0-20200714003250-2b9c44734f2b // indirect
	github.com/jackc/pgtype v1.6.2 // indirect
	github.com/jackc/pgx/v4 v4.10.1 // indirect
	github.com/jmespath/go-jmespath v0.3.0 // indirect
	github.com/konsorten/go-windows-terminal-sequences v1.0.3 // indirect
	github.com/lann/builder v0.0.0-20180802200727-47ae307949d0 // indirect
	github.com/lann/ps v0.0.0-20150810152359-62de8c46ede0 // indirect
	github.com/lib/pq v1.7.0 // indirect
	github.com/oklog/run v1.1.0 // indirect
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9 // indirect
	golang.org/x/sys v0.0.0-20200803210538-64077c9b5642 // indirect
	golang.org/x/text v0.3.3 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
)

replace (
	github.com/frankgreco/aviation => ../
	github.com/frankgreco/aviation/api => ../api
	github.com/frankgreco/aviation/internal => ../internal
)
