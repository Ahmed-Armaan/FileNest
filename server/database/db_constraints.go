package database

import (
	"gorm.io/gorm"
)

func setConstraints(db *gorm.DB) error {
	// only one node with no parent can exist
	if err := db.Exec(`CREATE UNIQUE INDEX IF NOT EXISTS uniq_root_per_user
		ON nodes (owner_id)
		WHERE parent_id IS NULL;`).Error; err != nil {
		return err
	}

	// index all soft deleted nodes, for better performance while hard deleting
	if err := db.Exec(`CREATE INDEX IF NOT EXISTS deleted_nodes
	ON nodes (id)
	WHERE deleted_at IS NOT NULL`).Error; err != nil {
		return err
	}

	// objectKey for directory are always null
	// However for files, it can never be null
	if err := db.Exec(`DO $$
	BEGIN
	ALTER TABLE nodes
	ADD CONSTRAINT objectkey_null_for_dir CHECK (
		(type = 'file' AND object_key IS NOT NULL) OR
		(type = 'directory' AND object_key IS NULL)
	);
	EXCEPTION
	WHEN duplicate_object THEN
	NULL;
	END $$;
	`).Error; err != nil {
		return err
	}

	// indedx to fetch all shared nodes information by the owner
	if err := db.Exec(`
	CREATE INDEX IF NOT EXISTS shares_owner_node_active
	ON shares (owner_id, node_id)
	WHERE revoked_at IS NULL`).Error; err != nil {
		return err
	}

	// index all revoked shared nodes, for better performance while hard deleting
	if err := db.Exec(`CREATE INDEX IF NOT EXISTS deleted_sharing_nodes
	ON shares (id)
	WHERE revoked_at IS NOT NULL`).Error; err != nil {
		return err
	}

	return nil
}
