// presigner actions as recommended by the aws sdk.
// each function creates a http request which can be executed using
// HTTP client to get the corresponding presigner URL.

package helper

//import (
//	"context"
//	"time"
//
//	"github.com/aws/aws-sdk-go-v2/aws"
//	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
//	"github.com/aws/aws-sdk-go-v2/service/s3"
//)
//
//type Presigner struct {
//	PresignClient *s3.PresignClient
//}
//
//func NewPresignerClient(client *s3.Client) *Presigner {
//	presignerClient := s3.NewPresignClient(client)
//	return &Presigner{
//		PresignClient: presignerClient,
//	}
//}
//
//func (presigner Presigner) GetObject(
//	ctx context.Context, bucketName string, objectKey string, lifetimeSecs time.Duration) (*v4.PresignedHTTPRequest, error) {
//	request, err := presigner.PresignClient.PresignGetObject(ctx, &s3.GetObjectInput{
//		Bucket: aws.String(bucketName),
//		Key:    aws.String(objectKey),
//	}, func(opts *s3.PresignOptions) {
//		opts.Expires = lifetimeSecs
//	})
//	if err != nil {
//		return nil, err
//	}
//	return request, err
//}
//
//func (presigner Presigner) PutObject(
//	ctx context.Context, bucketName string, objectKey string, lifetimeSecs time.Duration) (*v4.PresignedHTTPRequest, error) {
//	request, err := presigner.PresignClient.PresignPutObject(ctx, &s3.PutObjectInput{
//		Bucket: aws.String(bucketName),
//		Key:    aws.String(objectKey),
//	}, func(opts *s3.PresignOptions) {
//		opts.Expires = lifetimeSecs
//	})
//	if err != nil {
//		return nil, err
//	}
//	return request, err
//}
//
//func (presigner Presigner) DeleteObject(ctx context.Context, bucketName string, objectKey string) (*v4.PresignedHTTPRequest, error) {
//	request, err := presigner.PresignClient.PresignDeleteObject(ctx, &s3.DeleteObjectInput{
//		Bucket: aws.String(bucketName),
//		Key:    aws.String(objectKey),
//	})
//	if err != nil {
//		return nil, err
//	}
//	return request, err
//}
//
//func (presigner Presigner) PresignPostObject(ctx context.Context, bucketName string, objectKey string, lifetimeSecs time.Duration) (*s3.PresignedPostRequest, error) {
//	request, err := presigner.PresignClient.PresignPostObject(ctx, &s3.PutObjectInput{
//		Bucket: aws.String(bucketName),
//		Key:    aws.String(objectKey),
//	}, func(options *s3.PresignPostOptions) {
//		options.Expires = lifetimeSecs
//	})
//	if err != nil {
//		return nil, err
//	}
//	return request, nil
//}
