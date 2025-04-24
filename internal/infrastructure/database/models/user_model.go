package models

import "time"

type SexType string

const (
    Male   SexType = "Male"
    Female SexType = "Female"
    Other  SexType = "Other"
)

type UserModel struct {
	Id        int       `gorm:"column:id;primary_key;type:serial"`   // Primary key with serial type
	FirstName string    `gorm:"column:first_name;type:varchar(255)"` // First name with varchar(255)
	LastName  string    `gorm:"column:last_name;type:varchar(255)"`  // Last name with varchar(255)
	Age       int       `gorm:"column:age;type:int"`                // Age with int type
    Sex       SexType   `gorm:"column:sex;type:varchar(10)"`   
	Email     string    `gorm:"column:email;type:varchar(255)"`      // Email with varchar(255)
	Phone     string    `gorm:"column:phone;type:varchar(255)"`      // Phone with varchar(255)
	Address   string    `gorm:"column:address;type:varchar(255)"`    // Address with varchar(255)
	Role      string    `gorm:"column:role;type:varchar(255)"`
	Password	string 	 `gorm:"column:role;type:varchar(255)"`
	CreatedAt *time.Time `gorm:"column:created_at;type:timestamp"`    // CreatedAt with timestamp type
	UpdatedAt *time.Time `gorm:"column:updated_at;type:timestamp"`    // UpdatedAt with timestamp type
}

func (e *UserModel) TableName() string {
	return "user"
}