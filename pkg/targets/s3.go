package targets

import (
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/pibblokto/backlokto/pkg/types"
)

func createFilename(oldName string) string {
	filenameSlice := strings.Split(path.Base(oldName), ".")
	return filenameSlice[0] + "_" + strconv.FormatInt(time.Now().Unix(), 10) + ".sql"
}

func S3Target(target *types.Target, articats *types.Artifacts) {
	// Extract values from target and artifacts struct
	var access_key string
	var secret_key string
	var aws_region string
	var bucket_name string = target.S3BucketName
	var filepath string = articats.Filepath

	if target.AccessKey == "" {
		access_key = os.Getenv("AWS_ACCESS_KEY_ID")
	} else {
		access_key = target.AccessKey
	}

	if target.SecretKey == "" {
		secret_key = os.Getenv("AWS_SECRET_ACCESS_KEY")
	} else {
		secret_key = target.SecretKey
	}

	if target.Region == "" {
		aws_region = os.Getenv("AWS_DEFAULT_REGION")
	} else {
		aws_region = target.Region
	}
	// Create an S3 session
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(aws_region),
		Credentials: credentials.NewStaticCredentials(access_key, secret_key, ""),
	})
	if err != nil {
		fmt.Println(fmt.Errorf("failed to create AWS session: %v", err))
		return
	}

	// Open the file to be uploaded
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println(fmt.Errorf("failed to open file: %v", err))
		return
	}
	defer file.Close()

	// Create an S3 uploader
	uploader := s3.New(sess)
	var trailing_slash string = ""
	if target.S3BucketKey != "" {
		if target.S3BucketKey[len(target.S3BucketKey)-1] != '/' {
			trailing_slash = "/"
		}
	}

	// Upload the file to S3
	newFilename := createFilename(file.Name())
	_, err = uploader.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucket_name),
		Key:    aws.String(target.S3BucketKey + trailing_slash + newFilename),
		Body:   file,
	})
	if err != nil {
		fmt.Println(fmt.Errorf("failed to upload file to S3: %v", err))
		return
	}

	fmt.Printf("File %s uploaded to S3 bucket %s\n", newFilename, bucket_name)
}
