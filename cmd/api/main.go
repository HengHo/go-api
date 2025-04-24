package main

// นำเข้า packages ที่ใช้
import (
	"fmt"
	"log"

	// นำเข้า config ต่าง ๆ ของแอป เช่น .env และค่าคอนฟิก
	"backend-service/config"
	"backend-service/docs"
	"backend-service/internal/api/controllers"
	"backend-service/internal/application/usecase"
	"backend-service/internal/infrastructure/database"
	"backend-service/internal/infrastructure/repositories"
	"backend-service/pkg/utilities/logger"
	"backend-service/pkg/utilities/middlewares"

	// โหลด environment variables จากไฟล์ .env
	"github.com/joho/godotenv"
	// Echo คือ web framework ที่ใช้สร้าง REST API
	"github.com/labstack/echo/v4"
	// ใช้สร้าง Swagger UI สำหรับดูเอกสาร API
	echoSwagger "github.com/swaggo/echo-swagger"
)

// ส่วน Metadata สำหรับ Swagger API Documentation
// (ใช้สำหรับ auto generate เอกสาร)
 
// @contact.name        บริษัท มัลติ อินโนเวชั่น เอนยิเนียริ่ง จำกัด
// @contact.url         https://multiinno.com/
// @contact.email       akekapon.s@multiinno.com
//
// @title               Go
// @version             v1
// @description         Go API document
//
// @securityDefinitions.apikey X-API-Key
// @in                  header
// @name                X-API-Key

func main() {
	// โหลดค่าจากไฟล์ .env (เช่น PORT, DB_HOST, API_KEY เป็นต้น)
	if err := godotenv.Load(); err != nil {
		log.Fatal(err) // ถ้าโหลดไม่สำเร็จให้หยุดทำงานและแสดง error
	}

	// ดึงค่าคอนฟิกทั้งหมดจากฟังก์ชัน config.GetAppconfig()
	conf := config.GetAppconfig()

	// สร้าง logger สำหรับใช้พิมพ์ log ในระบบ
	logger := logger.GetLogger()

	// ตั้งค่า base path ของ Swagger (เช่น /api หรือ /v1) ให้ตรงกับ config
	docs.SwaggerInfo.BasePath = "/" + conf.Basepath

	// สร้าง Echo instance สำหรับสร้าง HTTP server
	e := echo.New()

	// ตั้งค่าฟังก์ชันสำหรับจัดการ error ทั่วทั้งระบบ (custom error handler)
	e.HTTPErrorHandler = middlewares.CustomHTTPErrorHandler(logger)

	// Middleware สำหรับ log ทุก request และ response
	e.Use(middlewares.RequestResponseLogger(logger))

	// เชื่อมต่อฐานข้อมูลตาม config ที่โหลดมา
	dbClient := database.ConnectDB(conf.Client)

	// สร้าง repository ที่ทำงานกับฐานข้อมูล
	repo := repositories.New(dbClient)

	// สร้าง usecase ที่รวม business logic โดยใช้ repo ที่สร้าง
	usecase := usecase.New(repo)

	// ถ้าอยู่ใน environment SIT หรือ UAT ให้แสดง Swagger UI
	if conf.Env == "sit" || conf.Env == "uat" {
		e.GET("/swagger/*", echoSwagger.WrapHandler) // เส้นทางนี้จะแสดงหน้า Swagger
	}

	// Middleware สำหรับ log path ของทุก request (ไว้ debug route ที่ถูกเรียก)
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			logger.Info("Request path: ", req.URL.Path)
			return next(c)
		}
	})

	// ใช้ middleware ตรวจสอบ API Key ที่แนบมาใน header
	apiKey := conf.ApiKey
	e.Use(middlewares.APIKeyMiddleware(apiKey))

	// ลงทะเบียน route ทั้งหมดจาก controllers (เช่น /users, /products, etc.)
	controllers.InitController(e, usecase)

	// วนลูปแสดง route ที่ถูกลงทะเบียนไว้ เพื่อใช้ debug
	for _, route := range e.Routes() {
		logger.Info(fmt.Sprintf("Method: %s, Path: %s", route.Method, route.Path))
	}

	// เริ่มทำงาน HTTP server ที่ port ตาม config เช่น ":8080"
	e.Logger.Fatal(e.Start(":" + conf.Port))
}
