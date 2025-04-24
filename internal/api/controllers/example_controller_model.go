package controllers

import "time" // ใช้สำหรับจัดการข้อมูลวันและเวลา (time.Time)

// StudentDTO คือ Data Transfer Object สำหรับรับ-ส่งข้อมูลของ student จาก/ไป client
type StudentDTO struct {
	// CreatedAt คือวันเวลาที่สร้างข้อมูล (อาจใช้แสดงผลหรือจัดเก็บ)
	// เป็น pointer เพื่อให้สามารถตรวจสอบว่า client ส่งค่ามาหรือไม่
	// json:"createdAt" ระบุชื่อ key ที่จะใช้ใน JSON
	// example ใช้สำหรับแสดงตัวอย่างใน Swagger
	CreatedAt *time.Time `json:"createdAt" example:"2023-05-17 23:50:50"`

	// UpdatedAt คือวันเวลาที่มีการอัปเดตข้อมูลล่าสุด
	// ใช้ pointer เช่นกัน เหมือน CreatedAt
	UpdatedAt *time.Time `json:"updatedAt" example:"2023-05-17 23:50:50"`

	// FirstName คือชื่อจริงของ student
	// เป็น string ธรรมดา (ไม่ใช่ pointer เพราะมักจะต้องมีค่า)
	FirstName string `json:"firstName" example:"Jimmy"`

	// LastName คือนามสกุลของ student
	LastName string `json:"lastName" example:"Karuture"`

	// Email คืออีเมลของ student
	Email string `json:"email" example:"jimmy@hiso.com"`
}
