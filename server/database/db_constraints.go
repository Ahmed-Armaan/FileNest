package database

import (
	//"github.com/Ahmed-Armaan/FileNest/database/helper"
	"gorm.io/gorm"
)

func setConstraints(db *gorm.DB) error {
	// only one node with no parent can exist
	if err := db.Exec(`CREATE UNIQUE INDEX IF NOT EXISTS uniq_root_per_user
		ON nodes (owner_id)
		WHERE parent_id IS NULL;`).Error; err != nil {
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
	//if err := db.Exec(`ALTER TABLE nodes
	//ADD CONSTRAINT objectkey_null_for_dir CHECK(
	//	(type = 'file' AND object_key IS NOT NULL) OR
	//	(type = 'directory' AND object_key IS NULL)
	//);`).Error; err != nil {
	//	if helper.ResolvePostgresError(err) != helper.ErrDuplicateObject {
	//		return err
	//	}
	//}

	return nil
}
