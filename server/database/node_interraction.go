package database

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type ChildData struct {
	ID        uuid.UUID
	Name      string
	Type      string
	UpdatedAt time.Time
}

func InsertNode(name string, nodeType NodeType, parentId *uuid.UUID, ownerId uuid.UUID) error {
	var children []ChildData

	query := DB.Model(&Node{}).
		Where("name = ? AND owner_id = ?", name, ownerId)

	if parentId == nil {
		query = query.Where("parent_id IS NULL")
	} else {
		query = query.Where("parent_id = ?", *parentId)
	}

	if err := query.Find(&children).Error; err != nil {
		return err
	}

	if len(children) > 0 {
		return fmt.Errorf("%s %s already exists", nodeType, name)
	}

	node := Node{
		Name:     name,
		Type:     string(nodeType),
		ParentID: parentId,
		OwnerID:  ownerId,
	}

	if err := DB.Create(&node).Error; err != nil {
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
