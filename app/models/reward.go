package models

import "time"

type Reward struct  {
	ID                 int                   `gorm:"AUTO_INCREMENT, primary_key"`
	Name               string                `json:"name"`
	Bones              int                    `json:"bones"`
	Sallery            int                     `json:"sallery"`
	CreatedAt          time.Time              `json:"createdAt"`
	UpdatedAt          time.Time


}
