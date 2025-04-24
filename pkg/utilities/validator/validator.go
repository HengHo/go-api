package validator

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"time"
	"unicode"
	"unsafe"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

// โครงสร้างเพื่อเก็บ FieldError จาก go-playground validator
type fieldError struct {
	err validator.FieldError
}

// โครงสร้างสำหรับเก็บรายละเอียดของ validation error รายตัว
type ValidationError struct {
	Code    string `json:"code"`    // ชื่อ tag ที่ไม่ผ่าน เช่น REQUIRED, DATE
	Field   string `json:"field"`   // ชื่อฟิลด์ที่ไม่ผ่าน
	Message string `json:"message"` // ข้อความที่อธิบายข้อผิดพลาด
}

// โครงสร้างสำหรับตอบกลับเมื่อ validation ล้มเหลว
type ErrResponse struct {
	Code    int         `json:"code"`    // รหัส HTTP เช่น 406
	Message string      `json:"message"` // ข้อความรวม
	Errors  interface{} `json:"errors"`  // ลิสต์ของ ValidationError
}

/*
 * ฟังก์ชันหลักในการตรวจสอบข้อมูล
 * รับ input (struct) แล้วเรียก validator มาตรวจสอบ field ตาม tag ที่กำหนดไว้
 */
func Validate(input interface{}) error {
	validate = validator.New()

	// ลงทะเบียนการตรวจสอบแบบ custom เพิ่มเติม
	_ = validate.RegisterValidation("date", IsDate)
	_ = validate.RegisterValidation("is-boolean", isBoolean)
	_ = validate.RegisterValidation("isThaiOrEnglish", validateThaiOrEnglish)
	_ = validate.RegisterValidation("isASCII", validateASCII)
	_ = validate.RegisterValidation("isComplexPassword", validateComplexPassword)
	_ = validate.RegisterValidation("isSex", validateSex)

	return validate.Struct(input) // เริ่มตรวจสอบ
}

/*
 * ฟังก์ชันจัดรูปแบบ error ให้อยู่ในรูปแบบที่แสดงผลได้
 * แปลงจาก error เดิมให้กลายเป็น ErrResponse
 */
func FormatValidationErrors(err error) ErrResponse {
	errs := []ValidationError{}

	// วนลูปอ่าน error แต่ละรายการที่เกิดจาก validator
	for _, err := range err.(validator.ValidationErrors) {
		validationErr := ValidationError{
			Code:    strings.ToUpper(err.ActualTag()), // เช่น REQUIRED, DATE
			Field:   err.Field(),                      // ชื่อฟิลด์
			Message: fieldError{err}.String(),         // ข้อความอธิบาย
		}

		errs = append(errs, validationErr)
	}

	// สร้างโครงสร้าง ErrResponse สำหรับส่งกลับ
	serializedErr := ErrResponse{
		Code:    ErrorStatusCode[ERR_VALIDATION_FAILED],     // รหัส HTTP เช่น 406
		Message: ErrorMessage[ERR_VALIDATION_FAILED],        // ข้อความรวม
		Errors:  errs,                                       // ข้อผิดพลาดแต่ละรายการ
	}

	// ใช้ unsafe.Pointer เพื่อเร่งประสิทธิภาพ (หลีกเลี่ยง copy memory)
	return *(*ErrResponse)(unsafe.Pointer(&serializedErr))
}

/*
 * ตรวจสอบว่า string มีรูปแบบวันที่แบบ yyyy-mm-dd หรือไม่
 */
func IsDate(fl validator.FieldLevel) bool {
	regexString := `^\d{4}\-(0?[1-9]|1[012])\-(0?[1-9]|[12][0-9]|3[01])$`
	Regex := regexp.MustCompile(regexString)
	return Regex.MatchString(fl.Field().String())
}

/*
 * ตรวจสอบว่าค่าที่รับเข้ามาเป็นประเภท boolean
 */
func isBoolean(fl validator.FieldLevel) bool {
	return reflect.TypeOf(fl.Field().Interface()).Kind() == reflect.Bool
}

/*
 * ฟังก์ชันแปลงข้อผิดพลาดจาก validator ให้เป็นข้อความอธิบาย
 */
func (q fieldError) String() string {
	var sb strings.Builder

	sb.WriteString("Validation failed on field '" + q.err.Field() + "'")
	sb.WriteString(", condition: " + q.err.ActualTag())

	// ถ้ามีพารามิเตอร์เพิ่มเติม เช่น max=10
	if q.err.Param() != "" {
		sb.WriteString(" { " + q.err.Param() + " }")
	}

	// แสดงค่าจริงที่ไม่ผ่าน
	if q.err.Value() != nil && q.err.Value() != "" {
		sb.WriteString(fmt.Sprintf(", actual: %v", q.err.Value()))
	}

	return sb.String()
}

/*
 * แปลง string วันที่ให้อยู่ในรูปแบบ time.Time
 * ใช้รูปแบบ RFC3339 เช่น 2023-01-01T00:00:00Z
 */
func ConvertStringToTime(date string) (time.Time, error) {
	var response error
	result, err := time.Parse(time.RFC3339, date)
	if err != nil {
		response = errors.New("Date is wrong format")
		return result, response
	}

	return result, response
}

/*
 * ตรวจสอบว่า endDate ไม่อยู่ก่อน startDate
 */
func ValidateDateRange(startDate time.Time, endDate time.Time) error {
	if endDate.Before(startDate) {
		return errors.New("End date must be after start date")
	}
	return nil
}

/*
 * ตรวจสอบว่า string มีเฉพาะภาษาไทย หรืออังกฤษ และช่องว่างเท่านั้น
 * ไม่อนุญาตตัวเลข
 */
func validateThaiOrEnglish(fl validator.FieldLevel) bool {
	for _, char := range fl.Field().String() {
		if unicode.IsDigit(char) {
			return false
		}
		if !(unicode.Is(unicode.Thai, char) || unicode.IsLetter(char) || unicode.IsSpace(char)) {
			return false
		}
	}
	return true
}

/*
 * ตรวจสอบว่า string เป็น ASCII ล้วน ไม่มีตัวอักษรภาษาพิเศษ
 */
func validateASCII(fl validator.FieldLevel) bool {
	for _, char := range fl.Field().String() {
		if char > unicode.MaxASCII {
			return false
		}
	}
	return true
}

/*
 * ตรวจสอบว่า password มีความซับซ้อน
 * เงื่อนไข:
 * - ความยาว 8 ถึง 32 ตัวอักษร
 * - มีตัวเลข
 * - มีอักษรพิมพ์ใหญ่
 * - มีอักษรพิมพ์เล็ก
 * - มีอักขระพิเศษ เช่น !@#$%^&*
 */
func validateComplexPassword(fl validator.FieldLevel) bool {
	var (
		hasMinLen  = false
		hasMaxLen  = true
		hasNumber  = false
		hasUpper   = false
		hasLower   = false
		hasSpecial = false
	)
	pass := fl.Field().String()

	if len(pass) >= 8 {
		hasMinLen = true
	}

	if len(pass) > 32 {
		hasMaxLen = false
	}

	// ตรวจสอบเงื่อนไขต่างๆ ทีละตัวอักษร
	for _, char := range pass {
		switch {
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case strings.ContainsRune("!@#$%^&*()-_+=", char):
			hasSpecial = true
		}
	}

	// ต้องผ่านทุกเงื่อนไข
	return hasMinLen && hasMaxLen && hasNumber && hasUpper && hasLower && hasSpecial
}
func validateSex(fl validator.FieldLevel) bool {
	sex := fl.Field().String()
	defaultSex := []string{"male", "female", "other"}
	for _, s := range defaultSex {
		if sex == s {
			return true		
			
		}
	}

	return false
}