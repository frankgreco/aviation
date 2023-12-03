module github.com/frankgreco/aviation/download

go 1.21

toolchain go1.21.0

require (
	github.com/aws/aws-lambda-go v1.17.0
	github.com/aws/aws-sdk-go v1.32.7
	github.com/frankgreco/aviation v0.0.0
	github.com/sirupsen/logrus v1.6.0
)

require (
	github.com/ghodss/yaml v1.0.0 // indirect
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b // indirect
	github.com/golang/protobuf v1.4.3 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.1.0 // indirect
	github.com/jmespath/go-jmespath v0.3.0 // indirect
	github.com/konsorten/go-windows-terminal-sequences v1.0.3 // indirect
	golang.org/x/sys v0.0.0-20200803210538-64077c9b5642 // indirect
	google.golang.org/genproto v0.0.0-20210106152847-07624b53cd92 // indirect
	google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.0.1 // indirect
	google.golang.org/protobuf v1.25.0 // indirect
	gopkg.in/yaml.v2 v2.2.3 // indirect
)

replace github.com/frankgreco/aviation => ../
