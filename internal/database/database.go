package database 

import (
    "log"

    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "gorm.io/gorm/logger"

    "github.com/freelancing/jobs/config"
    "github.com/freelancing/jobs/internal/models"
)

var Db *gorm.DB

func Initialize(config *config.Config) {
	var err error
    println(config.DbString)
	Db, err = gorm.Open(postgres.Open(config.DbString), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}
	if config.Environment == "prod" {
		Db.Logger.LogMode(logger.Error)
	}

	if err = Db.AutoMigrate(
		// migrating models
		&models.Job{},
	); err != nil {
		log.Fatal(err.Error())
	}
}


