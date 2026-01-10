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

func InsertNode(name string, nodeType NodeType, parentId *uuid.UUID, ownerId uuid.UUID) error {
	//	var children []ChildData
	//	query := DB.Model(&Node{}).
	//		Where("name = ? AND owner_id = ?", name, ownerId)
	//	if parentId == nil {
	//		query = query.Where("parent_id IS NULL")
	//	} else {
	//		query = query.Where("parent_id = ?", *parentId)
	//	}
	//	if err := query.Find(&children).Error; err != nil {
	//		return err
	//	}
	//	if len(children) > 0 {
	//		return fmt.Errorf("%s %s already exists", nodeType, name)
	//	}

	node := Node{
		Name:     name,
		Type:     string(nodeType),
		ParentID: parentId,
		OwnerID:  ownerId,
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
		Name:     "/",
		Type:     string(NodeTypeDirectory),
		ParentID: nil,
		OwnerID:  ownerId,
	}).Error; err != nil {
		return err
	}

	return nil
}

func GetAllChild(parentId *uuid.UUID, ownerId uuid.UUID) ([]ChildData, error) {
	var children []ChildData

	query := DB.Model(&Node{}).
		Select("id, name, type, updated_at").
		Where("owner_id = ?", ownerId)

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
