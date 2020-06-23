module github.com/frankgreco/aviation/download

require (
	github.com/aws/aws-lambda-go v1.17.0
	github.com/aws/aws-sdk-go v1.32.7
	github.com/frankgreco/aviation v0.0.0
	github.com/sirupsen/logrus v1.6.0
)

replace (
	github.com/frankgreco/aviation => ../
	github.com/frankgreco/aviation/utils => ../utils
)
