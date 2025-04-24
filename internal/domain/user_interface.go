package domain

type UserInterface interface {
	// CreateUser รับข้อมูลของ user และส่งกลับ error หากการสร้างข้อมูลไม่สำเร็จ
	CreateUser(user User) error
}