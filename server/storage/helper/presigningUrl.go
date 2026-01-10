package helper

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func GetPresignedUrl(method string, objectKey string, s3Client *s3.Client, ctx context.Context, bucketName string) (string, error) {
	if s3Client == nil {
		return "", errors.New("s3 not initialized")
	}

	presigner := NewPresignerClient(s3Client)
	expiry := 5 * time.Minute

	switch method {

	case http.MethodGet:
		req, err := presigner.GetObject(ctx, bucketName, objectKey, expiry)
		if err != nil {
			return "", err
		}
		return req.URL, nil

	case http.MethodPut:
		req, err := presigner.PutObject(ctx, bucketName, objectKey, expiry)
		if err != nil {
			return "", err
		}
		return req.URL, nil

	case http.MethodDelete:
		req, err := presigner.DeleteObject(ctx, bucketName, objectKey)
		if err != nil {
			return "", err
		}
		return req.URL, nil

	default:
		return "", errors.New("invalid HTTP method for presigned URL")
	}
}
