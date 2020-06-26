module github.com/frankgreco/aviation/load

require (
	github.com/Masterminds/squirrel v1.4.0
	github.com/aws/aws-lambda-go v1.17.0
	github.com/aws/aws-sdk-go v1.32.7
	github.com/frankgreco/aviation v0.0.0
	github.com/frankgreco/aviation/api v0.0.0
	github.com/frankgreco/aviation/utils v0.0.0-20200624030258-6c014fac6407
	github.com/jmoiron/sqlx v1.2.0
	github.com/sirupsen/logrus v1.6.0
)

replace (
	github.com/frankgreco/aviation => ../
	github.com/frankgreco/aviation/api => ../api
	github.com/frankgreco/aviation/utils => ../utils
)
