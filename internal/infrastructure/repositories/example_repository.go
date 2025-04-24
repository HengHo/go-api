package repositories

import (
	"log"

	"backend-service/internal/domain"
	"backend-service/internal/infrastructure/database/models"
)

// CreateStudent ใช้เพื่อสร้างข้อมูลนักเรียนใหม่ในฐานข้อมูล
func (r *Repo) CreateStudent(student domain.Student) error {
	// บันทึกข้อมูลนักเรียนที่รับเข้ามาใน log
	log.Printf("CreateStudent: %+v", student)

	// สร้าง StudentModel จากข้อมูลที่รับมาใน domain.Student
	dbRepo := models.StudentModel{
		FirstName: student.FirstName,
		LastName:  student.LastName,
		Email:     student.Email,
	}

	// สร้างแถวใหม่ในฐานข้อมูล โดยใช้ข้อมูลจาก dbRepo
	if err := r.db.Create(&dbRepo).Error; err != nil {
		// ถ้ามีข้อผิดพลาดในการสร้างข้อมูล จะส่งข้อผิดพลาดนั้นกลับไป
		return err
	}

	// ถ้าการสร้างข้อมูลสำเร็จ จะคืนค่า nil เพื่อบอกว่าไม่มีข้อผิดพลาด
	return nil
}
