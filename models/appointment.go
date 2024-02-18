package models

type Appointment struct {
	ID            string  `json:"id" form:"id"`
	JobID         string  `json:"job_id" form:"job_id"`
	EmployeeID    string  `json:"employee_id" form:"employee_id"`
	ScheduledDate string  `json:"scheduled_date" form:"scheduled_date"`
	IsCompleted   bool    `json:"is_completed" form:"is_completed"`
	Vehicle       Vehicle `json:"vehicle" form:"vehicle"`
	User          User    `json:"user" form:"user"`
}

type Vehicle struct {
	ID           string `json:"id" form:"id"`
	Make         string `json:"make" form:"make"`
	Model        string `json:"model" form:"model"`
	Year         string `json:"year" form:"year"`
	LicensePlate string `json:"license_plate" form:"license_plate"`
	VIN          string `json:"vin" form:"vin"`
}

type User struct {
	ID        string `json:"id" form:"id"`
	Username  string `json:"username" form:"username"`
	Password  string `json:"password" form:"password"`
	Role      string `json:"role" form:"role"`
	FirstName string `json:"first_name" form:"first_name"`
	LastName  string `json:"last_name" form:"last_name"`
	Email     string `json:"email" form:"email"`
	Phone     string `json:"phone" form:"phone"`
	Address   string `json:"address" form:"address"`
	City      string `json:"city" form:"city"`
	State     string `json:"state" form:"state"`
	ZipCode   string `json:"zip_code" form:"zip_code"`
	Country   string `json:"country" form:"country"`
}
