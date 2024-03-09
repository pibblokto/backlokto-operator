package providers

import (
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rds"
	"github.com/pibblokto/backlokto/pkg/types"
)

func RdsSnapshot(job *types.BackupJob) {
	var access_key string
	var secret_key string
	var aws_region string
	var dbInstanceIdentifier string = job.Spec.RdsIdentifier
	var dbSnapshotIdentifier string = fmt.Sprintf(`%v-%v`, job.Spec.RdsSnapshotPrefix, time.Now().Unix())

	if job.Spec.AccessKey == "" {
		access_key = os.Getenv("AWS_ACCESS_KEY_ID")
	} else {
		access_key = job.Spec.AccessKey
	}

	if job.Spec.SecretKey == "" {
		secret_key = os.Getenv("AWS_SECRET_ACCESS_KEY")
	} else {
		secret_key = job.Spec.SecretKey
	}

	if job.Spec.RdsRegion == "" {
		aws_region = os.Getenv("AWS_DEFAULT_REGION")
	} else {
		aws_region = job.Spec.RdsRegion
	}
	// Create a new sessions

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(aws_region),
		Credentials: credentials.NewStaticCredentials(access_key, secret_key, ""),
	})
	if err != nil {
		fmt.Println(fmt.Errorf("failed to create AWS session: %v", err))
		return
	}

	svc := rds.New(sess)
	input := &rds.CreateDBSnapshotInput{
		DBInstanceIdentifier: aws.String(dbInstanceIdentifier),
		DBSnapshotIdentifier: aws.String(dbSnapshotIdentifier),
	}

	result, err := svc.CreateDBSnapshot(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case rds.ErrCodeDBSnapshotAlreadyExistsFault:
				fmt.Println(rds.ErrCodeDBSnapshotAlreadyExistsFault, aerr.Error())
			case rds.ErrCodeInvalidDBInstanceStateFault:
				fmt.Println(rds.ErrCodeInvalidDBInstanceStateFault, aerr.Error())
			case rds.ErrCodeDBInstanceNotFoundFault:
				fmt.Println(rds.ErrCodeDBInstanceNotFoundFault, aerr.Error())
			case rds.ErrCodeSnapshotQuotaExceededFault:
				fmt.Println(rds.ErrCodeSnapshotQuotaExceededFault, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}
