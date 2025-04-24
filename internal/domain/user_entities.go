package domain

import "time"

type User struct {
	FirstName string     `json:"firstName"`
	LastName   string     `json:"lastName"`
	Age        int        `json:"age"`
	Sex        string     `json:"sex"`
	Email      string     `json:"email"`
	Phone      string     `json:"phone"`
	Address    string     `json:"address"`
	Role	   string     `json:"role"`
	Password   string     `json:"password"`
	CreatedAt  *time.Time `json:"createdAt"`
	UpdatedAt  *time.Time `json:"updatedAt"`
}
