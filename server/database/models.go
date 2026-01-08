package database

import (
	"time"

	"github.com/google/uuid"
)

type NodeType string

const (
	NodeTypeFile      NodeType = "file"
	NodeTypeDirectory NodeType = "directory"
)

type Node struct {
	ID        uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name      string     `gorm:"not null"`
	Type      string     `gorm:"not null"`
	ParentID  *uuid.UUID `gorm:"type:uuid;index"`
	OwnerID   uuid.UUID  `gorm:"type:uuid;not null;index"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type User struct {
	ID           uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	GoogleID     string    `gorm:"unique;not null"`
	UserName     string
	Email        string
	ProfileImage string
}

type FileChunk struct {
	ID         uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	FileID     uuid.UUID `gorm:"type:uuid;not null;index"`
	ChunkIndex int       `gorm:"not null"`
	Link       string    `gorm:"not null"`
}
