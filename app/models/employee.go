package models

import "time"

type Employee struct  {
	ID              int                    `gorm:"AUTO_INCREMENT, primary_key"json:"id"`
	First_name      string                 `json:"first_name"`
	Last_name       string                 `json:"last_name"`
	Hire_date       time.Time              `json:"hire_date"`
	Martialstatus   string                  `json:"martialstatus"`
	User             User                     `json:"user"`
	UserID           int64
	CreatedAt       time.Time              `json:"createdAt"`
	UpdatedAt       time.Time
	DeletedAt       *time.Time              `sql:"index"`



}