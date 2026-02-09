package database

import (
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"time"
)

type DatabaseStore interface {
	CreateNode(name string, nodeType NodeType, parentId *uuid.UUID, ownerId uuid.UUID, size *int64, objectKey ...string) error
	GetRootNode(googleId string) (*Node, error)
	ListChildren(parentId *uuid.UUID, googleId string) ([]ChildData, error)
	ListChildrenForDeletion(parentId *uuid.UUID) ([]ChildData, error)
	GetNodeObjectInfo(Id uuid.UUID) (*Node, error)
	MarkNodeDeleted(googleId string, nodeId uuid.UUID) error
	ListDeletedNodes(count int) ([]DeletedNodeData, error)
	DeleteNodePermanently(id uuid.UUID) error
	InsertUser(userName string, googleID string, email string, profileImage string) (*User, error)
	GetUserByGoogleID(googleID string) (*User, error)
	UserIDByGoogleIDQuery(googleId string, columns ...string) *gorm.DB
	GetUserDataByGoogleId(googleId string, columns ...string) (*User, error)
	ShareNode(nodeId uuid.UUID, password string, googleId string) (string, error)
	GetSharedPasswordStatus(code string) (bool, error)
	GetSharedNode(code string, password ...string) ([]ChildData, error)
	GetAllSharedNodes(googleId string) ([]SharedNode, error)
}

type DatabaseHolder struct {
	DB *gorm.DB
}

func DbInit() (DatabaseStore, error) {
	dsn := os.Getenv("DB_URL")
	databaseStore := &DatabaseHolder{}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return databaseStore, err
	}

	if err := db.AutoMigrate(&User{}, &Node{}, &Share{}); err != nil {
		return databaseStore, err
	}

	if err := setConstraints(db); err != nil {
		return databaseStore, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return databaseStore, err
	}

	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(time.Hour)

	databaseStore.DB = db
	return databaseStore, nil
}
