package database

import (
	"fmt"
	"log"
	"time"

	"backend-service/internal/infrastructure/database/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Config struct เก็บข้อมูลการตั้งค่าการเชื่อมต่อกับฐานข้อมูล PostgreSQL
type Config struct {
	Host     string `env:"THIRDPARTY_POS_HOST"`     // ที่อยู่ของเซิร์ฟเวอร์ฐานข้อมูล
	Username string `env:"THIRDPARTY_POS_USERNAME"` // ชื่อผู้ใช้งานสำหรับเชื่อมต่อฐานข้อมูล
	Password string `env:"THIRDPARTY_POS_PASSWORD"` // รหัสผ่านของผู้ใช้งาน
	Database string `env:"THIRDPARTY_POS_DATABASE"` // ชื่อฐานข้อมูลที่ต้องการเชื่อมต่อ
	Port     int    `env:"THIRDPARTY_POS_PORT"`     // พอร์ตที่ใช้ในการเชื่อมต่อ (โดยปกติ PostgreSQL ใช้พอร์ต 5432)
}

// buildDSN สร้าง Data Source Name (DSN) สำหรับเชื่อมต่อกับฐานข้อมูล PostgreSQL
func buildDSN(config *Config) string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Host,          // ชื่อหรือ IP ของเซิร์ฟเวอร์ฐานข้อมูล
		config.Port,          // พอร์ตการเชื่อมต่อ
		config.Username,      // ชื่อผู้ใช้งาน
		config.Password,      // รหัสผ่านของผู้ใช้งาน
		config.Database,      // ชื่อฐานข้อมูล
	)
}

// ConnectDB เชื่อมต่อกับฐานข้อมูล PostgreSQL โดยใช้ GORM
func ConnectDB(dbConf *Config) *gorm.DB {
	// สร้าง DSN จากการตั้งค่าที่ได้รับ
	dsn := buildDSN(dbConf) 
	log.Printf("dsn: %s\n", dsn) // แสดง DSN ใน log (ไม่ควรแสดงรหัสผ่านใน production)
	
	// เชื่อมต่อฐานข้อมูล PostgreSQL ผ่าน GORM
	dbConn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		// ใช้ logger ของ GORM สำหรับบันทึกข้อมูลการ query ระดับ Info
		Logger: logger.Default.LogMode(logger.Info),
		// กำหนดฟังก์ชันที่ใช้ในการดึงเวลา (UTC) สำหรับการบันทึกข้อมูล
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	})

	// หากไม่สามารถเชื่อมต่อฐานข้อมูลได้ ให้หยุดการทำงานและแสดงข้อความผิดพลาด
	if err != nil {
		panic(`fatal error: cannot connect to database`)
	}

	// ทำการ auto-migrate ตาราง StudentModel (สร้างหรืออัปเดตตารางตามที่กำหนดใน model)
	err = dbConn.AutoMigrate(&models.StudentModel{},&models.UserModel{},  )
	if err != nil {
		// หาก auto-migration ล้มเหลว ให้แสดงข้อความผิดพลาด
		panic(fmt.Sprintf("fatal error: auto-migration failed: %v", err))
	}

	// คืนค่าการเชื่อมต่อฐานข้อมูล
	return dbConn
}
