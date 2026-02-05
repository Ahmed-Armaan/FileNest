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

// An index on ownerId, parentId, deletedAt: For active children fetch
// A unique index on name, parentId, ownerId, deletedAt: To ensure unique siblings (excluding soft-deleted nodes)
// A unique index on owner_id when parent_id is NULL to ensure only one root node per user
// A partial index exist on id when type is directory and tagged as deleted for better performance during hard deletes
type Node struct {
	ID        uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name      string     `gorm:"not null;uniqueIndex:unique_siblings"`
	Type      string     `gorm:"not null"`
	ParentID  *uuid.UUID `gorm:"type:uuid;index:parent_owner_active;uniqueIndex:unique_siblings"`
	OwnerID   uuid.UUID  `gorm:"type:uuid;not null;index:parent_owner_active;uniqueIndex:unique_siblings"`
	ObjectKey *string
	SizeBytes *int64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index:parent_owner_active;uniqueIndex:unique_siblings"`
}

type User struct {
	ID           uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	GoogleID     string    `gorm:"unique;not null;index"`
	UserName     string    `gorm:"not null"`
	Email        string
	ProfileImage string
}
