package helper

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
)

func CreateNewUpload(ctx context.Context, s3Client *s3.Client, bucketName string) (uploadId string, objectKey string, err error) {
	objectKey = "data/" + uuid.NewString()

	out, err := s3Client.CreateMultipartUpload(ctx, &s3.CreateMultipartUploadInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
	})
	if err != nil {
		return "", "", err
	}

	return *out.UploadId, objectKey, nil
}
