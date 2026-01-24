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
	Name      string     `gorm:"not null;uniqueIndex:unique_siblings"`
	Type      string     `gorm:"not null"`
	ParentID  *uuid.UUID `gorm:"type:uuid;index;uniqueIndex:unique_siblings"`          // recursive in nature, build virtual file system
	OwnerID   uuid.UUID  `gorm:"type:uuid;not null;index;uniqueIndex:unique_siblings"` // foreign key for User
	ObjectKey *string
	SizeBytes *int64
	CreatedAt time.Time
	UpdatedAt time.Time
}

type User struct {
	ID           uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	GoogleID     string    `gorm:"unique;not null"`
	UserName     string    `gorm:"not null"`
	Email        string
	ProfileImage string
}
