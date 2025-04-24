package validator // ชื่อแพ็คเกจ ใช้จัดการการตรวจสอบความถูกต้องของข้อมูลและจัดรูปแบบการตอบกลับ

import (
	"net/http"
	"reflect"

	"github.com/labstack/echo" // ใช้ framework Echo สำหรับสร้าง API
)

// ResponseMessage เป็นโครงสร้างข้อมูลสำหรับจัดรูปแบบการตอบกลับ (response) ที่มีทั้ง code, message และข้อมูลที่เกี่ยวข้อง
type ResponseMessage struct {
	Code        int         `json:"code,omitempty"`        // รหัส HTTP หรือรหัสภายในของระบบ
	MessageCode string      `json:"messageCode,omitempty"` // รหัสข้อความที่ใช้ภายใน
	Message     string      `json:"message,omitempty"`     // ข้อความที่อธิบาย error หรือผลลัพธ์
	Data        interface{} `json:"data,omitempty"`        // ข้อมูลที่ส่งกลับ (เช่น ข้อมูลผลลัพธ์หรือรายละเอียด error)
}

// Response เป็น struct ที่ฝัง echo.Context ไว้ภายใน และเพิ่ม method สำหรับจัดการผลลัพธ์
type Response struct {
	echo.Context
	Message func(string) ResponseMessage // ฟังก์ชันสำหรับดึงข้อความ error จาก messageCode
}

// HandleError ใช้จัดการ error ทั่วไป (ไม่สามารถระบุประเภทได้แน่ชัด)
func (e Response) HandleError(err error) error {
	errMessage := e.Message(err.Error()) // สร้าง ResponseMessage จาก error message

	// เช็คว่ามี message ที่ส่งกลับมาหรือไม่ (ถ้าไม่มีให้ส่งข้อความ error ปกติกลับไป)
	if reflect.ValueOf(errMessage).IsZero() {
		errMessage := struct {
			Code    int    `json:"code,omitempty"`
			Message string `json:"message,omitempty"`
		}{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		return e.JSON(http.StatusInternalServerError, errMessage)
	}

	// ถ้ามี error message ที่จัดรูปแบบแล้ว ให้ส่งกลับเป็น JSON
	return e.JSON(http.StatusInternalServerError, errMessage)
}

// HandleSuccess ใช้สำหรับส่ง response กลับเมื่อสำเร็จ (HTTP 200)
func (e Response) HandleSuccess(response interface{}) error {
	return e.JSON(http.StatusOK, response)
}

// HandleCreated ใช้เมื่อสร้าง resource ใหม่สำเร็จ (HTTP 201)
func (e Response) HandleCreated(response interface{}) error {
	return e.JSON(http.StatusCreated, response)
}

// HandleBadRequest ใช้ส่ง error เมื่อผู้ใช้ส่งคำขอที่ไม่ถูกต้อง (HTTP 400)
func (e Response) HandleBadRequest(err error) error {
	return e.JSON(http.StatusBadRequest, FormatValidationErrors(err)) // แปลง error เป็นรูปแบบที่อ่านง่ายก่อนส่งกลับ
}

// NotFound ใช้จัดการกรณีไม่พบข้อมูล (HTTP 404)
func (e Response) NotFound(err error) error {
	message := e.Message(err.Error()) // พยายามแปลง error เป็นข้อความที่มี format

	// ถ้าไม่มี message ที่กำหนดไว้ ให้ส่ง error ตรง ๆ กลับ
	if reflect.ValueOf(message).IsZero() {
		return e.JSON(http.StatusNotFound, err)
	}
	// ถ้ามี message ที่จัดรูปแบบแล้ว ให้ส่งกลับ
	return e.JSON(http.StatusNotFound, message)
}
