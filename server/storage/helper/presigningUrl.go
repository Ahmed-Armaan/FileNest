package helper

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func PresignPart(ctx context.Context, s3Client *s3.Client, bucketName string, objectKey string, uploadId string, partNumber int32) (string, error) {
	presigner := s3.NewPresignClient(s3Client)

	req, err := presigner.PresignUploadPart(ctx, &s3.UploadPartInput{
		Bucket:     aws.String(bucketName),
		Key:        aws.String(objectKey),
		UploadId:   aws.String(uploadId),
		PartNumber: &partNumber,
	}, s3.WithPresignExpires(5*time.Minute))

	if err != nil {
		return "", err
	}

	return req.URL, nil
}

func GetPresignDownloadUrl(ctx context.Context, s3Client *s3.Client, bucketName string, objectKey string, size int64) (string, error) {
	presigner := s3.NewPresignClient(s3Client)

	req, err := presigner.PresignGetObject(ctx, &s3.GetObjectInput{
		Bucket: &bucketName,
		Key:    &objectKey,
	}, s3.WithPresignExpires(presignExpiryForSize(size)))

	if err != nil {
		return "", err
	}

	return req.URL, nil
}

func presignExpiryForSize(size int64) time.Duration {
	const MB = 1024 * 1024
	const GB = 1024 * MB

	switch {
	case size <= 50*MB:
		return 10 * time.Minute
	case size <= 200*MB:
		return 30 * time.Minute
	case size <= 1*GB:
		return 2 * time.Hour
	case size <= 5*GB:
		return 6 * time.Hour
	default:
		return 12 * time.Hour
	}
}
