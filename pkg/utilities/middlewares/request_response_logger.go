package middlewares

import (
	"bytes"               // ใช้สำหรับการจัดการกับ byte buffer
	"io"                  // ใช้สำหรับการอ่านข้อมูลจาก input stream
	"net/http"            // ใช้สำหรับการทำงานกับ HTTP
	"github.com/labstack/echo/v4" // ใช้สำหรับ Echo web framework
	"github.com/sirupsen/logrus" // ใช้สำหรับบันทึก log
)

// CustomResponseWriter wraps the original http.ResponseWriter and captures the response body
// ประกาศ CustomResponseWriter ที่ทำหน้าที่ห่อหุ้ม http.ResponseWriter และจับข้อมูล body ของ response
type CustomResponseWriter struct {
	http.ResponseWriter // ฝัง ResponseWriter เดิมไว้ใน CustomResponseWriter
	body *bytes.Buffer   // ใช้ Buffer เพื่อจับข้อมูลที่ส่งกลับใน response
}

// Write method ของ CustomResponseWriter ทำหน้าที่เขียนข้อมูลลงใน buffer และส่งข้อมูลไปยัง ResponseWriter เดิม
func (w *CustomResponseWriter) Write(b []byte) (int, error) {
	w.body.Write(b) // เขียนข้อมูลลงใน buffer
	return w.ResponseWriter.Write(b) // เขียนข้อมูลลง ResponseWriter เดิม
}

// RequestResponseLogger logs the request and response bodies
// ฟังก์ชันนี้เป็น middleware สำหรับบันทึก log ของ request และ response bodies
func RequestResponseLogger(logger *logrus.Logger) echo.MiddlewareFunc {
	// คืนค่าฟังก์ชันที่ทำหน้าที่เป็น middleware
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		// คืนค่าฟังก์ชันที่ทำหน้าที่เป็น handler ที่รับ request และ response
		return func(c echo.Context) error {
			// อ่าน request body
			req := c.Request()
			bodyBytes, err := io.ReadAll(req.Body) // อ่านเนื้อหาของ request body
			if err != nil {
				return err // หากเกิดข้อผิดพลาดในการอ่าน body ให้คืน error
			}
			req.Body = io.NopCloser(bytes.NewBuffer(bodyBytes)) // รีเซ็ต body กลับเป็นข้อมูลเดิม เพื่อให้สามารถใช้งานต่อได้
			bodyRequest := string(bodyBytes) // แปลงข้อมูล request body เป็น string

			// เตรียม buffer สำหรับจับข้อมูล response body
			res := c.Response()
			resBody := new(bytes.Buffer)
			customResWriter := &CustomResponseWriter{ResponseWriter: res.Writer, body: resBody} // สร้าง CustomResponseWriter ใหม่
			res.Writer = customResWriter // ใช้ CustomResponseWriter เป็น writer ของ response

			// เรียก next handler เพื่อให้ process request ไปต่อ
			err = next(c)
			if err != nil {
				c.Error(err) // หากเกิดข้อผิดพลาดในการทำงานของ handler ให้ส่ง error กลับไป
			}

			// บันทึก log ของ request body และ response body
			logger.Infof("Request body: %s, Response body: %s", bodyRequest, resBody.String())
			return nil // คืนค่า nil เพื่อให้การทำงานดำเนินต่อไป
		}
	}
}
