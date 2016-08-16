package models

type Token struct  {
	ID          int              `gorm:"AUTO_INCREMENT, primary_key"`
	Name        string           `json:"name"`
	Email       string           `json:"email"`
	Token        string           `json:"token"`

}