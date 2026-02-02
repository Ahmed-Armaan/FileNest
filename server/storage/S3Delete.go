package storage

import (
	"github.com/Ahmed-Armaan/FileNest/database"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
)

func deleteFile(objectKey *string, s *StorageHolder) error {
	if _, err := s.s3Client.DeleteObject(s.ctx, &s3.DeleteObjectInput{
		Bucket: &s.bucketName,
		Key:    objectKey,
	}); err != nil {
		return err
	}
	return nil
}

func (s *StorageHolder) DeleteFileById(id uuid.UUID, db database.DatabaseStore) error {
	fileData, err := db.GetNodeObjectInfo(id)
	if err != nil {
		return err
	}

	return deleteFile(fileData.ObjectKey, s)
}

func (s *StorageHolder) DeleteFileByObjectKey(objectKey *string) error {
	return deleteFile(objectKey, s)
}
