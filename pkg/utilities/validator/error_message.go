package validator // ชื่อ package สำหรับกลุ่มโค้ดที่ใช้จัดการเรื่องการ validate ข้อมูล

import (
	"net/http" // ใช้สำหรับระบุ HTTP Status Code
)

// ประกาศ constant สำหรับ error code ที่ใช้เมื่อ validation ล้มเหลว
const (
	ERR_VALIDATION_FAILED = `VALIDATION_FAILED`
)

// ErrorMessage เป็น map ที่ใช้จับคู่ error code กับข้อความ error ที่เข้าใจง่าย
var ErrorMessage = map[string]string{
	`VALIDATION_FAILED`: `Failed for validate the input`, // เมื่อ validate input ไม่ผ่าน จะแสดงข้อความนี้
}

// ErrorStatusCode เป็น map ที่ใช้จับคู่ error code กับ HTTP status code ที่จะส่งกลับไปยัง client
var ErrorStatusCode = map[string]int{
	`VALIDATION_FAILED`: http.StatusNotAcceptable, // HTTP 406 – ไม่สามารถยอมรับค่าที่ส่งมาได้
}
