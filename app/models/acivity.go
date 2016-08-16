package models

import "time"

type Activity struct  {
	ID                 int                    `gorm:"AUTO_INCREMENT, primary_key"`
	Discription        string                 `json:"discription"`
	CreatedAt          time.Time              `json:"createdAt"`
	UpdatedAt          time.Time

}
