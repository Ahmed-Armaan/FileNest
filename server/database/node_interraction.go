package database

import (
	"errors"
	"time"

	"github.com/Ahmed-Armaan/FileNest/database/helper"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ChildData struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Type      string    `json:"type"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type DeletedNodeData struct {
	ID   uuid.UUID `json:"id"`
	Type string    `json:"type"`
}

func InsertNode(name string, nodeType NodeType, parentId *uuid.UUID, ownerId uuid.UUID, size *int64, objectKey ...string) error {
	node := Node{
		Name:      name,
		Type:      string(nodeType),
		ParentID:  parentId,
		OwnerID:   ownerId,
		DeletedAt: nil,
	}

	if nodeType == NodeTypeDirectory {
		node.ObjectKey = nil
		node.SizeBytes = nil
	} else {
		node.ObjectKey = &objectKey[0]
		node.SizeBytes = size
	}

	if err := DB.Create(&node).Error; err != nil {
		if helper.ResolvePostgresError(err) == helper.ErrUniqueViolation {
			return errors.New("duplicate node")
		}
		return err
	}

	return nil
}

// insert root node for new users, part of a transaction
func insertRootNode(tx *gorm.DB, ownerId uuid.UUID) error {
	if err := tx.Create(&Node{
		Name:      "/",
		Type:      string(NodeTypeDirectory),
		ParentID:  nil,
		OwnerID:   ownerId,
		ObjectKey: nil,
		SizeBytes: nil,
	}).Error; err != nil {
		return err
	}

	return nil
}

func GetRootNodeId(googleId string) (*Node, error) {
	var nodeData Node

	subQuery := GetUserIdByGoogleIdSubQuery(googleId)

	err := DB.Model(&Node{}).
		Select("id, updated_at").
		Where(
			"owner_id = (?)",
			gorm.Expr("(?)", subQuery),
		).
		Where("parent_id IS NULL").
		Take(&nodeData).
		Error

	if err != nil {
		return nil, err
	}

	return &nodeData, nil
}

func GetAllChild(parentId *uuid.UUID, googleId string) ([]ChildData, error) {
	var children []ChildData

	query := DB.Model(&Node{}).
		Select("id, name, type, updated_at").
		Where(
			"owner_id = (?) AND deleted_at IS NULL",
			GetUserIdByGoogleIdSubQuery(googleId, UserDbColums.ID))

	if parentId == nil {
		query = query.Where("parent_id IS NULL")
	} else {
		query = query.Where("parent_id = ?", *parentId)
	}

	if err := query.Find(&children).Error; err != nil {
		return nil, err
	}

	return children, nil
}

func GetAllChildToDelete(parentId *uuid.UUID) ([]ChildData, error) {
	var children []ChildData

	if err := DB.Model(&Node{}).
		Select("id", "type").
		Where("parent_id = ?", parentId).
		Find(&children).Error; err != nil {
		return children, err
	}
	return children, nil
}

func GetObjectKey_Size_Name(Id uuid.UUID) (*Node, error) {
	node := Node{}

	if err := DB.Model(&Node{}).
		Select("object_key, size_bytes, name").
		Where("id = ?", Id).
		Take(&node).Error; err != nil {
		return &node, err
	}

	return &node, nil
}

func DeleteNode(googleId string, nodeId uuid.UUID) error {
	subquery := GetUserIdByGoogleIdSubQuery(googleId, UserDbColums.ID)
	if err := DB.Model(&Node{}).Where("id = ? AND owner_id = (?)",
		nodeId,
		subquery).
		Update("deleted_at", time.Now()).Error; err != nil {
		return err
	}

	return nil
}

func GetDeletedNodes(count int) ([]DeletedNodeData, error) {
	deletedNodes := []DeletedNodeData{}

	if err := DB.Model(&Node{}).
		Select("id, type").
		Where("deleted_at IS NOT NULL").
		Limit(count).Find(&deletedNodes).
		Error; err != nil {
		return deletedNodes, err
	}
	return deletedNodes, nil
}

func HardDeletion(id uuid.UUID) error {
	if err := DB.Delete(&Node{}, id).Error; err != nil {
		return err
	}
	return nil
}
