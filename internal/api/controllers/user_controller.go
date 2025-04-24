package controllers

import (
	"backend-service/internal/domain"
	
	"backend-service/pkg/utilities/responses"
	"backend-service/pkg/utilities/validator"
	
	"net/http"


	"github.com/labstack/echo/v4"
)

// @Summary      Create User
// @Description  Create a new user record
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        user  body      UserDTO  true  "User details" // ตัวแปร user อยู่ใน request body และต้องมีค่า
// @Success      201      {object}  response.Response       // ถ้าสำเร็จจะได้ response ที่เป็น object และสถานะ 201
// @Failure      400      {object}  response.ErrorResponse  // ถ้า client ส่งข้อมูลผิด จะได้ 400 และ error object
// @Failure      500      {object}  response.ErrorResponse  // ถ้า server มีปัญหา จะได้ 500 และ error object
// @Router       /v1/users [post]
// @Security     X-API-Key

func (h *Controller) CreateUser(c echo.Context) error {
	var userDTO UserDTO // ประกาศตัวแปรเพื่อรับข้อมูลที่เข้ามาในรูปแบบ JSON แล้ว map เป็น struct

	// Bind JSON request body ไปยัง userDTO
	// ถ้าเกิด error ในการแปลงข้อมูลจะตอบกลับด้วย HTTP 400 (Bad Request)
	if err := c.Bind(&userDTO); err != nil {
		return c.JSON(http.StatusBadRequest, responses.Error(http.StatusBadRequest, err.Error()))
	}

	if err := c.Validate(userDTO); err != nil {
		validationErrors := validator.FormatValidationErrors(err)
		return c.JSON(http.StatusBadRequest, validationErrors)
	}
	// ตรวจสอบว่า userDTO มีข้อมูลครบถ้วนตามที่ต้องการหรือไม่
	// สร้าง object user จากข้อมูลที่ได้จาก client
	user := domain.User{
		FirstName: userDTO.FirstName, // กำหนดชื่อจริง
		LastName:  userDTO.LastName,  // กำหนดนามสกุล
		Age: 	 userDTO.Age,       // กำหนดอายุ
		Sex: 	 userDTO.Sex,        // กำหนดเพศ
		Email:     userDTO.Email,     // กำหนดอีเมล
		Phone:     userDTO.Phone,     // กำหนดเบอร์โทรศัพท์
		Address:   userDTO.Address,   // กำหนดที่อยู่
		Role:      userDTO.Role,      // กำหนดบทบาท
		Password:  userDTO.Password,  // กำหนดรหัสผ่าน
		CreatedAt: userDTO.CreatedAt, // กำหนดวันที่สร้าง
		UpdatedAt: userDTO.UpdatedAt, // กำหนดวันที่อัปเดต
	}

	// เรียกใช้ usecase layer เพื่อบันทึกข้อมูล user ลงฐานข้อมูล
	err := h.uc.CreateUser(user)

	// ถ้ามี error ระหว่างบันทึกข้อมูล (เช่น error จากฐานข้อมูล)
	if err != nil {
		// ตอบกลับด้วย HTTP 500 (Internal Server Error)
		return c.JSON(http.StatusInternalServerError, responses.Error(http.StatusInternalServerError, err.Error()))
	} else {
		
		return c.JSON(http.StatusCreated, responses.Ok(http.StatusCreated, "User created successfully", nil))

	}
}