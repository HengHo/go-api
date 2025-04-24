package usecase

import (
	"backend-service/internal/domain" // นำเข้า struct Student จาก package domain
)

// CreateStudent เป็นเมธอดของ Usecase สำหรับสร้างข้อมูลนักเรียนใหม่
func (uc *Usecase) CreateStudent(student domain.Student) error {
	// เรียกเมธอด CreateStudent จาก repository เพื่อบันทึกข้อมูลนักเรียนลงในฐานข้อมูล
	err := uc.exampleRepo.CreateStudent(student)

	// ถ้าเกิด error จากการบันทึก ให้ return error กลับไป
	if err != nil {
		return err
	}

	// ถ้าไม่มี error ใด ๆ ให้ return nil แสดงว่าสำเร็จ
	return nil
}
