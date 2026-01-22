package helper

import (
	"context"
	"sort"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

type CompetedPartsData struct {
	Etag       string `json:"etag"`
	PartNumber int32  `json:"partNumber"`
}

func CompleteMultipartUpload(ctx context.Context, s3Client *s3.Client, bucketName string, objectKey string, uploadId string, completedPartsData []CompetedPartsData) error {
	completedParts := make([]types.CompletedPart, 0, len(completedPartsData))

	for _, part := range completedPartsData {
		completedParts = append(completedParts, types.CompletedPart{
			ETag:       &part.Etag,
			PartNumber: &part.PartNumber,
		})
	}

	sort.Slice(completedParts, func(i, j int) bool {
		return *completedParts[i].PartNumber < *completedParts[j].PartNumber
	})

	_, err := s3Client.CompleteMultipartUpload(ctx, &s3.CompleteMultipartUploadInput{
		Bucket:   &bucketName,
		Key:      &objectKey,
		UploadId: &uploadId,
		MultipartUpload: &types.CompletedMultipartUpload{
			Parts: completedParts,
		},
	})
	if err != nil {
		s3Client.AbortMultipartUpload(ctx, &s3.AbortMultipartUploadInput{
			Bucket:   &bucketName,
			Key:      &objectKey,
			UploadId: &uploadId,
		})
		return err
	}

	return nil
}
