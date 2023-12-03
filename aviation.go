package aviation

import (
	"fmt"
	"time"
)

var (
	FaaDatabaseURL  = "https://registry.faa.gov/database/ReleasableAircraft.zip"
	AwsRegion       = "us-west-2"
	AwsS3BucketName = "aircraft-registry"
	Now             = time.Now()
	Date            = fmt.Sprintf("%d-%d-%d", Now.Month(), Now.Day(), Now.Year())
)
