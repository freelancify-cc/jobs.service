package models

import (
    "time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseModel struct {
	Id        uuid.UUID      `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}

type JobModel struct {
    BaseModel 
    JobName string `json:"job_name"`
    JobDescription string `json:"job_description"`
    PostedEmployer uuid.UUID `json:"posted_employer" gorm:"type:uuid"`
    JobStatus int `json:"job_status"`
    PaymentPlan int `jsob:"payment_plan"`
}
