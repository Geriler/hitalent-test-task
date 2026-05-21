package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewPostgresDB(dsn string, silent bool) (*gorm.DB, error) {
	cfg := &gorm.Config{}
	if silent {
		cfg.Logger = logger.Default.LogMode(logger.Silent)
	}
	return gorm.Open(postgres.Open(dsn), cfg)
}
