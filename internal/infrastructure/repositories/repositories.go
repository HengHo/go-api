package repositories

import (
	"gorm.io/gorm"
)

// Repository struct จะใช้เป็น container สำหรับ repository ต่าง ๆ ที่จะใช้ในการจัดการฐานข้อมูล
type Repository struct {
	ExampleRepo *Repo // ใช้สำหรับเก็บ Repo ที่เกี่ยวกับ "ExampleRepo"
	// สามารถเพิ่ม Repo อื่น ๆ ได้ที่นี่ เช่น UserRepo, ProductRepo เป็นต้น
	UserRepo *Repo
}

// Repo struct ใช้สำหรับเก็บการเชื่อมต่อกับฐานข้อมูลผ่าน GORM
type Repo struct {
	db *gorm.DB // ตัวแปรที่ใช้เชื่อมต่อกับฐานข้อมูล
}

// New function ใช้เพื่อสร้าง Repository ใหม่ โดยรับค่าเป็นการเชื่อมต่อฐานข้อมูล (db) และคืนค่ากลับเป็น Repository
func New(db *gorm.DB) *Repository {
	return &Repository{
		ExampleRepo: &Repo{ // สร้างตัว Repo ใหม่และเก็บไว้ใน ExampleRepo
			db: db, // ส่งค่าการเชื่อมต่อฐานข้อมูล (db) เข้าไป
		},
		UserRepo: &Repo{ // สร้างตัว Repo ใหม่และเก็บไว้ใน ExampleRepo
			db: db, // ส่งค่าการเชื่อมต่อฐานข้อมูล (db) เข้าไป
		},
	}
}
