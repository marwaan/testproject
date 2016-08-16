package models

//import "github.com/jinzhu/gorm"
type User struct  {
	ID          int              `gorm:"AUTO_INCREMENT, primary_key"json:"id"`
	Name        string           `json:"name"`
	Email       string           `json:"email"`
	Password    string
	Employees   []Employee


}
