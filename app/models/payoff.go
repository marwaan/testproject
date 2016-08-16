package models

type Payoff struct  {
	ID                  int                   `gorm:"AUTO_INCREMENT, primary_key"json:"id"`
	Emp_firstName       string                 `json:"emp_firstName"`
	Monthly_salary      int                   `json:"monthly_salary"`
	Employee            Employee                  `json:"employee"`

}