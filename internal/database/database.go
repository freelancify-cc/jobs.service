package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/freelancify/jobs/config"
	"github.com/freelancify/jobs/internal/models"
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
		&models.BaseModel{},
		&models.JobStatusModel{},
		&models.PaymentPlanModel{},
		&models.JobModel{},
	); err != nil {
		log.Fatal(err.Error())
	}
}

func SelectAllJobs() ([]models.JobModel, error) {
	jobs := []models.JobModel{}
	results := Db.Find(&jobs)
	return jobs, results.Error
}

func InsertJob(job *models.JobModel) (int, error) {
	result := Db.Create(job)
	return int(result.RowsAffected), result.Error
}

func SelectJobById(id int) (models.JobModel, error) {
	job := models.JobModel{}
	result := Db.First(&job, id)
	return job, result.Error
}

func DeleteRequestById(id int) error {
	var job models.JobModel
	job.Id = id
	result := Db.Delete(&job)
	return result.Error
}
