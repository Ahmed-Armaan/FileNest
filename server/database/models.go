package database

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type NodeType string

const (
	NodeTypeFile      NodeType = "file"
	NodeTypeDirectory NodeType = "directory"
)

type Node struct {
	ID           uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name         string    `gorm:"not null"`
	Type         NodeType  `gorm:"type node_type;not null"`
	ParentId     *uuid.UUID
	createdAt    time.Time
	lastModified time.Time
}

type Users struct {
	Id           uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	googleId     string    `gorm:"not null"`
	UserName     string
	Email        string
	profileImage string
}

type FileChunks struct {
	ID     uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	FileId uuid.UUID `gorm:"foreignKey;NodeRefer"`
	Link   string
}
