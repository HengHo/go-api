package controllers

import (
	"backend-service/internal/application/usecase" // ดึงเข้ามาใช้สำหรับ business logic layer
	"github.com/labstack/echo/v4"                 // Echo framework สำหรับจัดการ HTTP Routing
)

// Controller คือโครงสร้างหลักที่เก็บ usecase ไว้สำหรับเรียกใช้งานภายใน controller
type Controller struct {
	uc *usecase.Usecase // uc คือ instance ของ usecase (business logic) ที่ controller จะเรียกใช้
}

// InitController ทำหน้าที่เชื่อมโยง route กับ handler function ของ controller
// พารามิเตอร์:
// - e: instance ของ Echo สำหรับกำหนด route
// - usecase: business logic ที่ส่งเข้ามาให้ controller ใช้งาน
func InitController(e *echo.Echo, usecase *usecase.Usecase) {
	// สร้าง instance ของ Controller และเก็บ usecase ไว้ในฟิลด์ uc
	controller := &Controller{uc: usecase}

	// สร้าง route group ด้วย prefix /v1 เช่น /v1/students
	group := e.Group("/v1")

	// กำหนดให้เมื่อ client ส่ง POST มาที่ /v1/students
	// จะเรียก method controller.CreateStudent เป็น handler
	group.POST("/students", controller.CreateStudent)
	// กำหนดให้เมื่อ client ส่ง POST มาที่ /v1/users
	// จะเรียก method controller.CreateUser เป็น handler
	group.POST("/users", controller.CreateUser)
}
