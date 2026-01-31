package storage

import (
	"github.com/Ahmed-Armaan/FileNest/database"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
)

func DeleteFileFromS3(id uuid.UUID) error {
	fileData, err := database.GetObjectKey_Size_Name(id)
	if err != nil {
		return err
	}

	if _, err := s3Client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: &bucketName,
		Key:    fileData.ObjectKey,
	}); err != nil {
		return err
	}
	return nil
}
