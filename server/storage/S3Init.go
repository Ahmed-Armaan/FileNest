package storage

import (
	"context"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var (
	s3Client   *s3.Client
	ctx        context.Context
	bucketName string
)

func S3Init() error {
	ctx = context.Background()

	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return err
	}

	client := s3.NewFromConfig(cfg)
	bucketName = os.Getenv("AWS_BUCKET_NAME")

	// testing if the client and variables can establish a connection
	if _, err := client.HeadBucket(ctx, &s3.HeadBucketInput{
		Bucket: &bucketName,
	}); err != nil {
		return err
	}

	s3Client = client
	return nil
}
