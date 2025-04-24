package logger

import (
	"github.com/sirupsen/logrus"         // ใช้ logrus เป็น logging library หลัก
	"go.elastic.co/ecslogrus"            // ใช้สำหรับแปลง logrus formatter ให้อยู่ในรูปแบบ ECS (Elastic Common Schema)
)

// GetLogger สร้าง logger ใหม่ และตั้งค่าให้ใช้ ECS formatter
func GetLogger() *logrus.Logger {
	logger := logrus.New()                   // สร้าง logger ตัวใหม่จาก logrus
	logger.SetFormatter(&ecslogrus.Formatter{}) // ตั้งค่าให้ logger ใช้ ECS formatter แทน formatter ปกติ
	return logger                            // ส่ง logger กลับไป
}
