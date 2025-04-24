package middlewares

import (
	"strings" // นำเข้าแพ็คเกจสำหรับจัดการกับสตริง
	"github.com/labstack/echo/v4" // นำเข้า Echo web framework
	"github.com/labstack/echo/v4/middleware" // นำเข้า middleware ของ Echo
)

// ฟังก์ชัน APIKeyMiddleware ที่รับ apiKey เป็นอาร์กิวเมนต์ และคืนค่า Middleware สำหรับการตรวจสอบ API Key
func APIKeyMiddleware(apiKey string) echo.MiddlewareFunc {
	// คืนค่า middleware ที่ทำงานก่อนและหลังการเรียก handler
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		// ฟังก์ชันที่ทำงานเมื่อมีการเรียก HTTP request เข้ามา
		return func(c echo.Context) error {
			// ข้ามการตรวจสอบ API key สำหรับเส้นทางที่เริ่มต้นด้วย /swagger/
			// ซึ่งมักจะใช้สำหรับการเข้าถึง Swagger UI ที่ไม่ต้องการตรวจสอบ API Key
			if strings.HasPrefix(c.Path(), "/swagger/") {
				return next(c) // หากเส้นทางตรงกับ Swagger ให้เรียก handler ถัดไปโดยไม่ต้องตรวจสอบ API key
			}

			// ใช้ middleware KeyAuth สำหรับเส้นทางอื่นๆ ที่ต้องการตรวจสอบ API key
			// ตรวจสอบจาก Header "X-API-Key" ว่าตรงกับ apiKey ที่ส่งเข้ามาหรือไม่
			return middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
				KeyLookup: "header:X-API-Key", // กำหนดว่า API Key จะถูกค้นหาจาก Header ชื่อ "X-API-Key"
				Validator: func(key string, c echo.Context) (bool, error) {
					// ตรวจสอบว่า API Key ที่ได้รับมาตรงกับ apiKey ที่กำหนดหรือไม่
					return key == apiKey, nil
				},
			})(next)(c) // เรียก middleware ที่สร้างขึ้นมาและส่งไปยัง handler ถัดไป
		}
	}
}
