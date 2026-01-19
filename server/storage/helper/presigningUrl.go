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
