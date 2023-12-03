module github.com/frankgreco/aviation/load

go 1.21

toolchain go1.21.0

require (
	github.com/Masterminds/squirrel v1.5.0
	github.com/aws/aws-lambda-go v1.41.0
	github.com/aws/aws-sdk-go v1.48.11
	github.com/frankgreco/aviation v0.0.0
	github.com/frankgreco/aviation/api v0.0.0
	github.com/frankgreco/aviation/internal v0.0.0-20210108003342-f47da5804365
	github.com/jmoiron/sqlx v1.2.0
	github.com/sirupsen/logrus v1.6.0
)

require (
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/konsorten/go-windows-terminal-sequences v1.0.3 // indirect
	github.com/lann/builder v0.0.0-20180802200727-47ae307949d0 // indirect
	github.com/lann/ps v0.0.0-20150810152359-62de8c46ede0 // indirect
	github.com/lib/pq v1.7.0 // indirect
	github.com/oklog/run v1.1.0 // indirect
	golang.org/x/sys v0.5.0 // indirect
)

replace (
	github.com/frankgreco/aviation => ../
	github.com/frankgreco/aviation/api => ../api
	github.com/frankgreco/aviation/internal => ../internal
)
