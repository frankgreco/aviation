package main

import (
	"archive/zip"
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	log "github.com/sirupsen/logrus"
)

const (
	faaDatabaseURL  = "http://registry.faa.gov/database/ReleasableAircraft.zip"
	awsRegion       = "us-west-2"
	awsS3BucketName = "aircraft-registry"
)

func main() {
	lambda.Start(do)
}

func do(ctx context.Context) error {

	log.WithFields(log.Fields{
		"url": faaDatabaseURL,
	}).Info("retrieving archive file from url")

	// download zip file
	beginDownload := time.Now()
	resp, err := http.Get(faaDatabaseURL)
	if err != nil {
		log.WithFields(log.Fields{
			"url": faaDatabaseURL,
			"err": err.Error(),
		}).Error("could not retrieve file from url")
		return err
	}
	defer resp.Body.Close()

	log.WithFields(log.Fields{
		"time": time.Since(beginDownload),
		"url":  faaDatabaseURL,
	}).Info("successfully retrieved archive file from url")

	// read response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("could not read response body")
		return err
	}

	log.Info("unzipping archive")

	// unzip response body
	reader, err := zip.NewReader(bytes.NewReader(body), int64(len(body)))
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("could not unzip response body")
		return err
	}

	// this will hold all of the files in the archive
	files := map[string][]byte{}

	// iterate through files in archive
	for _, file := range reader.File {
		log.WithFields(log.Fields{
			"file": file.Name,
		}).Info("found file in archive")

		data, err := readFileFromArchive(file)
		if err != nil {
			log.WithFields(log.Fields{
				"file": file.Name,
			}).Error("could not process file in archive")
			return err
		}
		// add file data to data structure
		files[file.Name] = data
	}

	log.Info("uploading files to aws s3")

	sesh, err := session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Region: aws.String(awsRegion),
		},
	})
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("could not create aws session")
		return err
	}

	now := time.Now()
	date := fmt.Sprintf("%d-%d-%d", now.Month(), now.Day(), now.Year())

	beginUpload := time.Now()
	for name, data := range files {
		key := fmt.Sprintf("%s/%s", date, name)

		if _, err := s3.New(sesh).PutObject(&s3.PutObjectInput{
			Bucket:        aws.String(awsS3BucketName),
			Key:           aws.String(key),
			Body:          bytes.NewReader(data),
			ContentLength: aws.Int64(int64(len(data))),
			ContentType:   aws.String(http.DetectContentType(data)),
		}); err != nil {
			log.WithFields(log.Fields{
				"err": err.Error(),
			}).Error("could not write files to aws s3")
			return err
		}
		log.WithFields(log.Fields{
			"file": name,
			"key":  key,
		}).Info("successfully wrote file to aws s3")
	}

	log.WithFields(log.Fields{
		"time": time.Since(beginUpload),
	}).Info("successfully uploaded all files to aws s3")

	return nil
}

func readFileFromArchive(file *zip.File) ([]byte, error) {
	// open file in archive
	f, err := file.Open()
	if err != nil {
		log.WithFields(log.Fields{
			"file": file.Name,
		}).Error("could not open file in archive")
		return nil, err
	}
	defer f.Close()

	// read file in archive
	data, err := ioutil.ReadAll(f)
	if err != nil {
		log.WithFields(log.Fields{
			"file": file.Name,
		}).Error("could not read file in archive")
		return nil, err
	}
	return data, err
}
