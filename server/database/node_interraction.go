package database

import (
	"errors"
	"time"

	"github.com/Ahmed-Armaan/FileNest/database/helper"
	"github.com/Ahmed-Armaan/FileNest/utils"
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

var (
	ErrNoPasswordProvided = errors.New("No password provided")
	ErrWrongPassword      = errors.New("Wrong password provided")
)

func (db *DatabaseHolder) CreateNode(name string, nodeType NodeType, parentId *uuid.UUID, ownerId uuid.UUID, size *int64, objectKey ...string) error {
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

	if err := db.DB.Create(&node).Error; err != nil {
		if helper.ResolvePostgresError(err) == helper.ErrUniqueViolation {
			return errors.New("duplicate node")
		}
		return err
	}

	return nil
}

// private helper — receiver optional, but this is cleaner
func (db *DatabaseHolder) insertRootNode(tx *gorm.DB, ownerId uuid.UUID) error {
	return tx.Create(&Node{
		Name:      "/",
		Type:      string(NodeTypeDirectory),
		ParentID:  nil,
		OwnerID:   ownerId,
		ObjectKey: nil,
		SizeBytes: nil,
	}).Error
}

func (db *DatabaseHolder) GetRootNode(googleId string) (*Node, error) {
	var nodeData Node
	subQuery := db.UserIDByGoogleIDQuery(googleId)

	if err := db.DB.Model(&Node{}).
		Select("id, updated_at").
		Where("owner_id = (?) AND parent_id IS NULL", subQuery).
		Take(&nodeData).Error; err != nil {
		return nil, err
	}

	return &nodeData, nil
}

func (db *DatabaseHolder) ListChildren(parentId *uuid.UUID, googleId string) ([]ChildData, error) {
	var children []ChildData

	query := db.DB.Model(&Node{}).
		Select("id, name, type, updated_at").
		Where(
			"owner_id = (?) AND deleted_at IS NULL",
			db.UserIDByGoogleIDQuery(googleId, UserDbColums.ID),
		)

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

func (db *DatabaseHolder) ListChildrenForDeletion(parentId *uuid.UUID) ([]ChildData, error) {
	var children []ChildData

	if err := db.DB.Model(&Node{}).
		Select("id", "type").
		Where("parent_id = ?", parentId).
		Find(&children).Error; err != nil {
		return children, err
	}

	return children, nil
}

func (db *DatabaseHolder) GetNodeObjectInfo(id uuid.UUID) (*Node, error) {
	node := Node{}

	if err := db.DB.Model(&Node{}).
		Select("object_key, size_bytes, name").
		Where("id = ?", id).
		Take(&node).Error; err != nil {
		return &node, err
	}

	return &node, nil
}

func (db *DatabaseHolder) MarkNodeDeleted(googleId string, nodeId uuid.UUID) error {
	subQuery := db.UserIDByGoogleIDQuery(googleId, UserDbColums.ID)

	return db.DB.Model(&Node{}).
		Where("id = ? AND owner_id = (?)", nodeId, subQuery).
		Update("deleted_at", time.Now()).
		Error
}

func (db *DatabaseHolder) ListDeletedNodes(count int) ([]DeletedNodeData, error) {
	deletedNodes := []DeletedNodeData{}

	if err := db.DB.Model(&Node{}).
		Select("id, type").
		Where("deleted_at IS NOT NULL").
		Limit(count).
		Find(&deletedNodes).
		Error; err != nil {
		return deletedNodes, err
	}

	return deletedNodes, nil
}

func (db *DatabaseHolder) DeleteNodePermanently(id uuid.UUID) error {
	return db.DB.Delete(&Node{}, id).Error
}

func (db *DatabaseHolder) ShareNode(nodeId uuid.UUID, password string, googleId string) (string, error) {
	code, err := utils.GenerateCode(8)
	if err != nil {
		return "", errors.New("Failed to genberate code")
	}

	hashedPassWord, err := utils.HashAndSalt(password)
	if err != nil {
		return "", err
	}

	user, err := db.GetUserDataByGoogleId(googleId, UserDbColums.ID)
	if err != nil {
		return "", err
	}

	share := Share{
		Code:      code,
		NodeId:    nodeId,
		OwnerId:   user.ID,
		Password:  &hashedPassWord,
		RevokedAt: nil,
	}

	err = db.DB.Create(&share).Error
	if err != nil {
		return "", err
	}
	return code, nil
}

func (db *DatabaseHolder) GetSharedPasswordStatus(code string) (bool, error) {
	shared_node := &Share{}
	if err := db.DB.Model(&Share{}).
		Select("password").
		Where("code = ? AND revoked_at IS NULL", code).
		Take(shared_node).Error; err != nil {
		return true, err
	}
	return (shared_node.Password != nil), nil
}

func (db *DatabaseHolder) GetSharedNode(code string, password ...string) ([]ChildData, error) {
	var childData []ChildData
	sharedNode := &Share{}

	if err := db.DB.Model(&Share{}).
		Select("node_id, password").
		Where("code = ? AND revoked_at IS NULL", code).
		Take(sharedNode).Error; err != nil {
		return childData, err
	}

	if sharedNode.Password != nil {
		if len(password) == 0 {
			return childData, ErrNoPasswordProvided
		}
		if !utils.ComparePassword(*sharedNode.Password, password[0]) {
			return childData, ErrWrongPassword
		}
	}

	type nodeMeta struct {
		Type string
	}

	var meta nodeMeta
	if err := db.DB.Model(&Node{}).
		Select("type").
		Where("id = ? AND deleted_at IS NULL", sharedNode.NodeId).
		Take(&meta).Error; err != nil {
		return childData, err
	}

	if meta.Type == string(NodeTypeFile) {
		var node ChildData

		if err := db.DB.Model(&Node{}).
			Select("id, name, type, updated_at").
			Where("id = ? AND deleted_at IS NULL", sharedNode.NodeId).
			Take(&node).Error; err != nil {
			return childData, err
		}

		childData = []ChildData{node}
		return childData, nil
	}

	if err := db.DB.Model(&Node{}).
		Select("id, name, type, updated_at").
		Where("parent_id = ? AND deleted_at IS NULL", sharedNode.NodeId).
		Find(&childData).Error; err != nil {
		return childData, err
	}

	return childData, nil
}
