package domain // กำหนด package ที่ชื่อว่า domain ซึ่งจะเก็บโครงสร้างข้อมูลหลักและ interface

// ExampleInterface เป็น interface ที่ประกอบด้วย method ที่ต้องการให้ repository รองรับ
// ในกรณีนี้ method `CreateStudent` ที่ใช้ในการบันทึกข้อมูลนักเรียน
type ExampleInterface interface {
	// CreateStudent รับข้อมูลของ student และส่งกลับ error หากการสร้างข้อมูลไม่สำเร็จ
	CreateStudent(student Student) error
}