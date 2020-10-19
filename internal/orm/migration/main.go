package migration

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
	log "yogan.dev/go-gql-server/internal/logger"
	"yogan.dev/go-gql-server/internal/orm/migration/jobs"
	"yogan.dev/go-gql-server/internal/orm/models"
)

func updateMigration(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.User{},
	).Error
}

// ServiceAutoMigration migrates tables and structure
func ServiceAutoMigration(db *gorm.DB) error {
	m := gormigrate.New(db, gormigrate.DefaultOptions, nil)
	m.InitSchema(func(db *gorm.DB) error {
		log.Info("Migrations.InitSchema | Initializing Database Schema")
		switch db.Dialect().GetName() {
		case "postgres":
			db.Exec("create extension \"uuid-ossp\";")
		}
		if err := updateMigration(db); err != nil {
			return fmt.Errorf("Migration.InitSchema | %v", err)
		}

		return nil
	})
	m.Migrate()

	if err := updateMigration(db); err != nil {
		return err
	}

	m = gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		jobs.SeedUsers,
	})

	return m.Migrate()
}
