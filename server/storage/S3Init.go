package storage

import (
	"context"
	"os"

	"github.com/Ahmed-Armaan/FileNest/database"
	"github.com/Ahmed-Armaan/FileNest/storage/helper"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type StorageStore interface {
	DeleteFileById(id uuid.UUID, db database.DatabaseStore) error
	DeleteFileByObjectKey(objectKey *string) error
	GetNewUploadUrl(c *gin.Context)
	GetUploadUrl(c *gin.Context)
	CompleteUploadS3(objectKey string, uploadId string, completdPartsData []helper.CompetedPartsData) error
	DownloadInit(db database.DatabaseStore) gin.HandlerFunc
}

type StorageHolder struct {
	s3Client   *s3.Client
	ctx        context.Context
	bucketName string
}

func S3Init() (StorageStore, error) {
	ctx := context.Background()

	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, err
	}

	client := s3.NewFromConfig(cfg)
	bucket := os.Getenv("AWS_BUCKET_NAME")

	if _, err := client.HeadBucket(ctx, &s3.HeadBucketInput{
		Bucket: &bucket,
	}); err != nil {
		return nil, err
	}

	return &StorageHolder{
		s3Client:   client,
		ctx:        ctx,
		bucketName: bucket,
	}, nil
}
