package repositories

import (
	"log"		

	
	"backend-service/internal/domain"
	"backend-service/internal/infrastructure/database/models"
)

func (r *Repo) CreateUser(user domain.User) error {
	// บันทึกข้อมูลนักเรียนที่รับเข้ามาใน log
	log.Printf("CreateUser: %+v", user)

	// สร้าง UserModel จากข้อมูลที่รับมาใน domain.User
	dbRepo := models.UserModel{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Age: 	user.Age,
		Sex: 	models.SexType(user.Sex),
		Email:     user.Email,
		Phone:     user.Phone,
		Address:   user.Address,
		Role:      user.Role,
		Password: user.Password,
		CreatedAt: user.CreatedAt,

	}

	// สร้างแถวใหม่ในฐานข้อมูล โดยใช้ข้อมูลจาก dbRepo
	if err := r.db.Create(&dbRepo).Error; err != nil {
		// ถ้ามีข้อผิดพลาดในการสร้างข้อมูล จะส่งข้อผิดพลาดนั้นกลับไป
		return err
	}

	// ถ้าการสร้างข้อมูลสำเร็จ จะคืนค่า nil เพื่อบอกว่าไม่มีข้อผิดพลาด
	return nil
}