package middlewares

import (
	"backend-service/pkg/utilities/responses" // นำเข้าแพ็คเกจที่เกี่ยวข้องกับการจัดการ response ของแอปพลิเคชัน
	"errors" // นำเข้าแพ็คเกจที่ใช้สำหรับการจัดการกับ error
	"net/http" // นำเข้าแพ็คเกจที่เกี่ยวข้องกับ HTTP status codes
	"github.com/labstack/echo/v4" // นำเข้า Echo web framework
	"github.com/sirupsen/logrus" // นำเข้า Logrus สำหรับการบันทึก log
)

// ฟังก์ชัน CustomHTTPErrorHandler สร้าง error handler สำหรับจัดการกับข้อผิดพลาดของ HTTP
// โดยใช้ Logrus ในการบันทึก log ของข้อผิดพลาดที่เกิดขึ้น
func CustomHTTPErrorHandler(logger *logrus.Logger) echo.HTTPErrorHandler {
	// คืนค่าฟังก์ชันที่เป็น handler สำหรับจัดการ error ของ HTTP
	return func(err error, c echo.Context) {
		var appErr *responses.ApplicationError

		// หาก error ที่ได้รับเป็น ApplicationError จาก responses.ApplicationError
		if errors.As(err, &appErr) {
			// บันทึก log ของ error โดยระบุ code และข้อความของ error
			logger.WithError(err).WithField("code", appErr.Code).Error(err.Error())

			// ส่ง response JSON โดยใช้ข้อมูลจาก error ที่เกิดขึ้น
			if err := c.JSON(responses.GetHttpStatusForCode(appErr.Code), responses.ErrorResponse{
				Code: appErr.Code, // ใช้ code ของ error
				Error: responses.ErrorDetail{
					Message: appErr.Message, // ข้อความของ error
					Stack:   appErr.Error(), // stack trace ของ error
				},
			}); err != nil {
				// หากเกิดข้อผิดพลาดในการส่ง response ให้บันทึก log
				logger.Error(err)
			}
		}

		// บันทึก log ของ error ที่เกิดขึ้น
		logger.WithError(err).Error(err.Error())

		// หากไม่ใช่ ApplicationError ให้ส่ง response เป็น InternalServerError (HTTP 500)
		if err := c.JSON(http.StatusInternalServerError, responses.ErrorResponse{
			Code: responses.UNEXPECTED_EXCEPTION, // ใช้รหัสข้อผิดพลาดทั่วไป
			Error: responses.ErrorDetail{
				Message: responses.DEFAULT_ERROR_MESSAGE, // ข้อความผิดพลาดมาตรฐาน
				Stack:   err.Error(), // stack trace ของ error
			},
		}); err != nil {
			// หากเกิดข้อผิดพลาดในการส่ง response ให้บันทึก log
			logger.Error(err)
		}
	}
}
