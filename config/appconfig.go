package config

import (
	"context" // ใช้สำหรับการจัดการ context ใน Go
	"backend-service/internal/infrastructure/database" // นำเข้า package สำหรับการจัดการฐานข้อมูล
	"github.com/sethvargo/go-envconfig" // นำเข้า package สำหรับการโหลดค่าตัวแปรจาก environment variables
)

var conf *AppConfig // ประกาศตัวแปร conf เป็น pointer ของ AppConfig ใช้เก็บค่า configuration

// AppConfig struct เก็บค่าคอนฟิกที่ใช้ในแอปพลิเคชัน
type AppConfig struct {
	Client   *database.Config // ค่าคอนฟิกสำหรับการเชื่อมต่อฐานข้อมูล
	ApiKey   string           `env:"API_KEY"` // ค่าคอนฟิกสำหรับ API Key ที่เก็บใน environment variable
	Port     string           `env:"API_PORT"` // ค่าคอนฟิกสำหรับ Port ที่เก็บใน environment variable
	Env      string           `env:"APP_ENV"` // ค่าคอนฟิกสำหรับ environment เช่น development, production
	Basepath string           `env:"BASE_PATH"` // ค่าคอนฟิกสำหรับ base path ของ API
}

// GetAppconfig function ใช้โหลดค่าคอนฟิกจาก environment variable
func GetAppconfig() *AppConfig {
	var config AppConfig // สร้างตัวแปร config เป็นตัวเก็บค่าคอนฟิก

	// โหลดค่าจาก environment variable ลงใน config โดยใช้ envconfig
	if err := envconfig.Process(context.Background(), &config); err != nil {
		panic(err) // หากเกิดข้อผิดพลาดในการโหลดค่าคอนฟิก ให้หยุดโปรแกรม
	}

	conf = &config // ตั้งค่า conf ให้ชี้ไปที่ config ที่โหลดมา

	return conf // คืนค่า conf ซึ่งเป็น pointer ของ AppConfig
}
