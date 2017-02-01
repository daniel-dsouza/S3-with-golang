package common

import (
	"log"

	"io"

	"bytes"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

var awsSession *session.Session

func getSession() *session.Session {
	if awsSession != nil {
		return awsSession
	}

	creds := credentials.NewEnvCredentials()
	_, err := creds.Get()
	if err != nil {
		log.Panicln(err)
	}

	awsSession = session.New(&aws.Config{
		Credentials: creds,
		Region:      aws.String("us-west-2"),
	})

	return awsSession
}

//UploadToS3 a file to S3
func UploadToS3(file io.Reader, key string) {
	sess := getSession()

	uploader := s3manager.NewUploader(sess)
	_, err1 := uploader.Upload(&s3manager.UploadInput{
		ACL:         aws.String("public-read"),
		Bucket:      aws.String("elasticbeanstalk-us-west-2-422535611875"),
		Key:         aws.String(key),
		ContentType: aws.String("image/jpeg"),
		Body:        file,
	})

	if err1 != nil {
		log.Println("err")

	}
}

// DownloadFromS3 does shit
func DownloadFromS3(key string) *bytes.Buffer {
	sess := getSession()
	downloader := s3manager.NewDownloader(sess)

	// var buffer bytes.Buffer

	var buffer2 aws.WriteAtBuffer
	_, err1 := downloader.Download(&buffer2, &s3.GetObjectInput{
		Bucket: aws.String("elasticbeanstalk-us-west-2-422535611875"),
		Key:    aws.String(key),
	})

	if err1 != nil {
		log.Println("err")
	}

	return bytes.NewBuffer(buffer2.Bytes())
}
