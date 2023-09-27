package infrastructure

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/digicon-trap1-2023/backend/infrastructure/migration"
	"github.com/digicon-trap1-2023/backend/util"
	sqldriver "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewGormDB() (*gorm.DB, error) {
	engine, err := gorm.Open(
		mysql.New(mysql.Config{DSNConfig: newDsnConfig()}),
		newGormConfig(),
	)
	if err != nil {
		return nil, err
	}

	db, err := engine.DB()
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(2)
	db.SetMaxOpenConns(16)
	if err := initDB(engine); err != nil {
		return nil, err
	}

	return engine, nil
}

func initDB(db *gorm.DB) error {
	_, err := migration.Migrate(db, migration.AllTables())
	if err != nil {
		return err
	}
	return nil
}

func newGormConfig() *gorm.Config {
	return &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold:             200 * time.Millisecond,
				LogLevel:                  logger.Warn,
				IgnoreRecordNotFoundError: true,
				Colorful:                  true,
			},
		),
		NowFunc: func() time.Time {
			return time.Now().Truncate(time.Microsecond)
		},
	}
}

func newDsnConfig() *sqldriver.Config {
	if os.Getenv("NEOSHOWCASE") == "true" {
		return &sqldriver.Config{
			User:   util.ReadEnvs("NS_MARIADB_USER"),
			Passwd: util.ReadEnvs("NS_MARIADB_PASSWORD"),
			Net:    "tcp",
			Addr: fmt.Sprintf(
				"%s:%s",
				util.ReadEnvs("NS_MARIADB_HOSTNAME"),
				util.ReadEnvs("NS_MARIADB_PORT"),
			),
			DBName:    util.ReadEnvs("NS_MARIADB_DATABASE"),
			Collation: "utf8mb4_general_ci",
			ParseTime:            true,
			AllowNativePasswords: true,
			Params: map[string]string{
				"charset": "utf8mb4",
			},
		}
	}
	return &sqldriver.Config{
		User:                 util.ReadEnvs("DB_USERNAME"),
		Passwd:               util.ReadEnvs("DB_PASSWORD"),
		Net:                  "tcp",
		Addr:                 fmt.Sprintf("%s:%s", util.ReadEnvs("DB_HOSTNAME"), util.ReadEnvs("DB_PORT")),
		DBName:               util.ReadEnvs("DB_DATABASE"),
		Collation:            "utf8mb4_general_ci",
		ParseTime:            true,
		AllowNativePasswords: true,
		Params: map[string]string{
			"charset": "utf8mb4",
		},
	}
}
