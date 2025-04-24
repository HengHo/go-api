package usecase

import (
	// นำเข้าโครงสร้างข้อมูลของ domain และ interface
	"backend-service/internal/domain"

	// นำเข้า repository ที่ใช้ติดต่อกับฐานข้อมูล
	"backend-service/internal/infrastructure/repositories"
)

// Usecase struct เป็นเลเยอร์ที่รวม business logic หรือกฎการทำงานของแอปพลิเคชัน
type Usecase struct {
	// exampleRepo คือ interface สำหรับเรียกใช้งาน repository
	exampleRepo domain.ExampleInterface
	userRepo	domain.UserInterface
	
}

// New คือฟังก์ชัน constructor สำหรับสร้าง Usecase ใหม่
// รับ repository หลักที่รวม repo ย่อย ๆ ไว้ภายใน
func New(repo *repositories.Repository) *Usecase {
	return &Usecase{
		// ดึง ExampleRepo จาก repository แล้วเก็บไว้ใน Usecase
		exampleRepo: repo.ExampleRepo,
		userRepo: repo.UserRepo,
	}
}
