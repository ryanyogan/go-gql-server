package orm

import (
	log "yogan.dev/go-gql-server/internal/logger"
	"yogan.dev/go-gql-server/internal/orm/migration"

	"github.com/jinzhu/gorm"
	// Imports the database dialect for PG
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"yogan.dev/go-gql-server/pkg/utils"
)

var autoMigrate, logMode, seedDB bool
var dsn, dialect string

// ORM struct to hold the gorm pointer to DB
type ORM struct {
	DB *gorm.DB
}

func init() {
	dialect = utils.MustGet("GORM_DIALECT")
	dsn = utils.MustGet("GORM_CONNECTION_DSN")
	seedDB = utils.MustGetBool("GORM_SEED_DB")
	logMode = utils.MustGetBool("GORM_LOGMODE")
	autoMigrate = utils.MustGetBool("GORM_AUTOMIGRATE")
}

// Factory creates a db connection with the selected dialect
func Factory() (*ORM, error) {
	db, err := gorm.Open(dialect, dsn)
	if err != nil {
		log.Panic("[ORM] err: ", err)
	}

	orm := &ORM{
		DB: db,
	}

	db.LogMode(logMode)

	if autoMigrate {
		err = migration.ServiceAutoMigration(orm.DB)
	}
	log.Info("[ORM} Database connection initialized.")
	return orm, err
}
