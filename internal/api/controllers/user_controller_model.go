package controllers

import "time"

type UserDTO struct {
	FirstName string     `json:"firstName" example:"Jimmy" validate:"required,max=50,min=2"`
	LastName   string     `json:"lastName" example:"Karuture" validate:"required,max=50,min=2"`
	Age        int        `json:"age" example:"20" validate:"required,min=1,max=120"`
	Sex        string     `json:"sex" example:"mele" validate:"isSex"`
	Email      string     `json:"email" example:"Jimmy@gmail.com" validate:"required,email"`
	Phone      string     `json:"phone" example:"0800000000" validate:"required,isThaiPhone"`
	Address    string     `json:"address" example:"Bangkok" validate:"required,max=100"`
	Role	   string     `json:"role" example:"user" validate:"required"`
	Password   string     `json:"password" example:"123456@test" validate:"required,isComplexPassword"`
	CreatedAt  *time.Time `json:"createdAt" example:"2023-05-17 23:50:50" `
	UpdatedAt  *time.Time `json:"updatedAt" example:"2023-05-17 23:50:50" `

}