package domain // กำหนด package ที่ชื่อว่า domain ซึ่งใช้เก็บโครงสร้างข้อมูลหลัก

import (
	// นำเข้าชุดคำสั่งที่ใช้จัดการกับเวลา
	"time"
)

// Student คือ struct ที่ใช้เก็บข้อมูลของนักเรียน
// จะใช้ struct นี้ในฐานะตัวแทนข้อมูลนักเรียนที่ต้องการบันทึก
type Student struct {
	// CreatedAt คือเวลาที่สร้างข้อมูลนักเรียน โดยเป็น pointer ที่ชี้ไปที่ time.Time (เพื่อให้สามารถเป็น null ได้)
	// นำไปใช้เพื่อการบันทึกว่าเมื่อไรข้อมูลนักเรียนนี้ถูกสร้าง
	CreatedAt *time.Time `json:"createdAt"`

	// UpdatedAt คือเวลาที่ข้อมูลนักเรียนถูกอัปเดตครั้งล่าสุด
	// ใช้เป็น pointer เช่นเดียวกับ CreatedAt เพื่อให้สามารถเป็น null ได้
	UpdatedAt *time.Time `json:"updatedAt"`

	// FirstName คือชื่อของนักเรียน
	FirstName string `json:"firstName"`

	// LastName คือนามสกุลของนักเรียน
	LastName string `json:"lastName"`

	// Email คืออีเมลของนักเรียน
	Email string `json:"email"`
}
