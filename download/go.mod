module github.com/frankgreco/aviation/download

go 1.21

toolchain go1.21.0

require (
	github.com/aws/aws-lambda-go v1.41.0
	github.com/aws/aws-sdk-go v1.48.11
	github.com/frankgreco/aviation v0.0.0
	github.com/sirupsen/logrus v1.6.0
)

require (
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/konsorten/go-windows-terminal-sequences v1.0.3 // indirect
	golang.org/x/sys v0.0.0-20200803210538-64077c9b5642 // indirect
)

replace github.com/frankgreco/aviation => ../
