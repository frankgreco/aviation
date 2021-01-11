package aviation

import (
	"fmt"
	"time"

	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2"
	_ "google.golang.org/grpc/cmd/protoc-gen-go-grpc"
	_ "google.golang.org/protobuf/cmd/protoc-gen-go"
)

var (
	FaaDatabaseURL  = "http://registry.faa.gov/database/ReleasableAircraft.zip"
	AwsRegion       = "us-west-2"
	AwsS3BucketName = "aircraft-registry"
	Now             = time.Now()
	Date            = fmt.Sprintf("%d-%d-%d", Now.Month(), Now.Day(), Now.Year())
)
