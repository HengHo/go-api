package usecase

import (
	"fmt"
	"log"

	"backend-service/internal/domain"
	"backend-service/pkg/utilities/mailer"

)

func (uc *Usecase) CreateUser(user domain.User) error {
	// เรียกเมธอด CreateUser จาก repository เพื่อบันทึกข้อมูลนักเรียนลงในฐานข้อมูล
	err := uc.userRepo.CreateUser(user)
	// ถ้าเกิด error จากการบันทึก ให้ return error กลับไป
	if err != nil {
		return err
	}
	// ส่งอีเมลต้อนรับผู้ใช้ใหม่
	subject := "Welcome to Our Platform"
	body := fmt.Sprintf("Hi %s,\n\nThank you for signing up! ", user.FirstName)

	err = mailer.SendEmail(user.Email, subject, body)
	if err != nil {
		log.Printf("Failed to send email: %v", err)
		// ตอบกลับโดยไม่หยุดการทำงานหลัก
		return err
	}

	log.Println("User added and email sent successfully!")
	// ถ้าไม่มี error ใด ๆ ให้ return nil แสดงว่าสำเร็จ
	return nil
}
