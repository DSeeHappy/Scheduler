package models

import "time"

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
	CustomerID   string `json:"customer_id" form:"customer_id"`
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

type AddressDTO struct {
	Address1 string `json:"address1" form:"address1"`
	Address2 string `json:"address2" form:"address2"`
	City     string `json:"city" form:"city"`
	State    string `json:"state" form:"state"`
	ZipCode  string `json:"zip_code" form:"zip_code"`
}

type AppointmentDTO struct {
	ID          string    `json:"id" form:"id"`
	ShopID      string    `json:"shop_id" form:"shop_id"`
	CustomerID  string    `json:"customer_id" form:"customer_id"`
	VehicleID   string    `json:"vehicle_id" form:"vehicle_id"`
	StartTime   time.Time `json:"start_time" form:"start_time"`
	EndTime     time.Time `json:"end_time" form:"end_time"`
	Title       string    `json:"title" form:"title"`
	Description string    `json:"description" form:"description"`
	Color       string    `json:"color" form:"color"`
	PickupTime  time.Time `json:"pickup_time" form:"pickup_time"`
	RideOption  string    `json:"ride_option" form:"ride_option"`
}

type VehicleDTO struct {
	ID           string `json:"id" form:"id"`
	CustomerID   string `json:"customer_id" form:"customer_id"`
	Make         string `json:"make" form:"make"`
	Model        string `json:"model" form:"model"`
	Year         string `json:"year" form:"year"`
	LicensePlate string `json:"license_plate" form:"license_plate"`
	VIN          string `json:"vin" form:"vin"`
	Engine       string `json:"engine" form:"engine"`
	Color        string `json:"color" form:"color"`
	State        string `json:"state" form:"state"`
	Notes        string `json:"notes" form:"notes"`
	UnitNumber   string `json:"unit_number" form:"unit_number"`
	SubModel     string `json:"sub_model" form:"sub_model"`
}

type CustomerDTO struct {
	ID             string     `json:"id" form:"id"`
	ShopID         string     `json:"shop_id" form:"shop_id"`
	CustomerTypeId string     `json:"customer_type_id" form:"customer_type_id"`
	FirstName      string     `json:"first_name" form:"first_name"`
	LastName       string     `json:"last_name" form:"last_name"`
	Email          string     `json:"email" form:"email"`
	Phone          string     `json:"phone" form:"phone"`
	Address        AddressDTO `json:"address" form:"address"`
	OkForMarketing bool       `json:"ok_for_marketing" form:"ok_for_marketing"`
	Notes          string     `json:"notes" form:"notes"`
}
