package controllers

import (
	"net/http" // สำหรับใช้งานสถานะ HTTP เช่น 200, 400, 500

	"backend-service/internal/domain"         // โดเมนหลักของระบบ ใช้สำหรับจัดการข้อมูลของ student
	"backend-service/pkg/utilities/responses" // แพ็กเกจที่จัดการรูปแบบการตอบกลับ (response)

	"github.com/labstack/echo/v4" // Echo framework สำหรับสร้าง REST API
)

// @Summary      Create Student                             // สรุปความสามารถของ endpoint นี้
// @Description  Create a new student record                // คำอธิบายเพิ่มเติมของ API
// @Tags         Students                                   // แสดงกลุ่มของ API นี้ใน Swagger
// @Accept       json                                       // ข้อมูลที่รับเข้ามาจะอยู่ในรูปแบบ JSON
// @Produce      json                                       // ข้อมูลที่ตอบกลับจะเป็น JSON
// @Param        student  body      StudentDTO  true  "Student details" // ตัวแปร student อยู่ใน request body และต้องมีค่า
// @Success      201      {object}  responses.Response       // ถ้าสำเร็จจะได้ response ที่เป็น object และสถานะ 201
// @Failure      400      {object}  responses.ErrorResponse  // ถ้า client ส่งข้อมูลผิด จะได้ 400 และ error object
// @Failure      500      {object}  responses.ErrorResponse  // ถ้า server มีปัญหา จะได้ 500 และ error object
// @Router       /v1/students [post]                        // เส้นทางของ endpoint นี้คือ POST /v1/students
// @Security     X-API-Key                                 // ใช้ API Key เพื่อยืนยันความปลอดภัยของ API

// CreateStudent เป็น method ของ Controller ที่ทำหน้าที่รับ request สำหรับสร้างข้อมูล student ใหม่
func (h *Controller) CreateStudent(c echo.Context) error {
	var studentDTO StudentDTO // ประกาศตัวแปรเพื่อรับข้อมูลที่เข้ามาในรูปแบบ JSON แล้ว map เป็น struct

	// Bind JSON request body ไปยัง studentDTO
	// ถ้าเกิด error ในการแปลงข้อมูลจะตอบกลับด้วย HTTP 400 (Bad Request)
	if err := c.Bind(&studentDTO); err != nil {
		return c.JSON(http.StatusBadRequest, responses.Error(http.StatusBadRequest, err.Error()))
	}

	// สร้าง object student จากข้อมูลที่ได้จาก client
	student := domain.Student{
		FirstName: studentDTO.FirstName, // กำหนดชื่อจริง
		LastName:  studentDTO.LastName,  // กำหนดนามสกุล
		Email:     studentDTO.Email,     // กำหนดอีเมล
	}

	// เรียกใช้ usecase layer เพื่อบันทึกข้อมูล student ลงฐานข้อมูล
	err := h.uc.CreateStudent(student)

	// ถ้ามี error ระหว่างบันทึกข้อมูล (เช่น error จากฐานข้อมูล)
	if err != nil {
		// ตอบกลับด้วย HTTP 500 (Internal Server Error)
		return c.JSON(http.StatusInternalServerError, responses.Error(http.StatusInternalServerError, err.Error()))
	} else {
		// ถ้าบันทึกสำเร็จ ตอบกลับด้วย HTTP 201 (Created)
		// ข้อความตอบกลับอาจปรับให้เหมาะสมเช่น "Successfully created student"
		return c.JSON(http.StatusCreated, responses.Ok(http.StatusCreated, "Successfully save selling items", nil))
	}
}
