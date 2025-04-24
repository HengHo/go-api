package http_request

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

// Result struct ใช้สำหรับเก็บผลลัพธ์จากการทำ HTTP request
type Result struct {
	Error error         // เก็บ error ที่เกิดขึ้นระหว่าง request
	Data  interface{}   // เก็บข้อมูลที่ response กลับมา ถ้าไม่ error
}

// ErrorResponse struct ใช้สำหรับเก็บ error ที่ได้จาก API ในรูปแบบ JSON
type ErrorResponse struct {
	ErrorCode   interface{} `json:"code"`        // อาจเป็น string หรือ number แล้วแต่ API
	MessageCode string      `json:"messageCode"` // รหัสข้อความ error
	Message     string      `json:"message"`     // ข้อความ error ปกติ
}

// other_http_header เป็น alias ของ map[string]string ใช้สำหรับส่ง header เพิ่มเติม
type other_http_header map[string]string

// SuccessFn คือฟังก์ชันที่รับ body (byte array) แล้วแปลงข้อมูลเมื่อ response สำเร็จ
type SuccessFn func(bodyByte []byte) interface{}

// FailureFn คือฟังก์ชันที่รับ response และ body แล้ว return เป็น error ถ้า response ไม่สำเร็จ
type FailureFn func(response *http.Response, bodyByte []byte) error

// HandleResponse ส่ง request และจัดการผลลัพธ์ โดยใช้ successFn หรือ failureFn ตามสถานะของ response
func HandleResponse(request *http.Request, successFn SuccessFn, failureFn FailureFn) *Result {
	result := &Result{}

	// ส่ง HTTP request
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		result.Error = err
		return result
	}
	defer response.Body.Close() // ปิด body หลังใช้งาน

	// อ่าน body ทั้งหมดจาก response
	bodyByte, err := io.ReadAll(response.Body)
	if err != nil {
		result.Error = err
		return result
	}

	// ถ้าสถานะ response อยู่ในช่วง 200-299 ถือว่าสำเร็จ
	if response.StatusCode >= http.StatusOK && response.StatusCode < http.StatusMultipleChoices {
		// เรียก successFn เพื่อแปลงข้อมูล และเก็บไว้ใน result.Data
		result.Data = successFn(bodyByte)
	} else {
		// ถ้าไม่สำเร็จ เรียก failureFn และเก็บ error ที่ได้
		result.Error = failureFn(response, bodyByte)
	}

	return result
}

// BuildRequest สร้าง HTTP request พร้อมแนบ API key, body และ header เพิ่มเติมถ้ามี
func BuildRequest(method, url string, ApiKey string, rawBody interface{}, header ...other_http_header) (*http.Request, error) {
	var request *http.Request
	var err error

	// ถ้ามี body ให้ marshal เป็น JSON แล้วแนบใน request
	if rawBody != nil {
		jsonStr, err := json.Marshal(rawBody)
		if err != nil {
			return nil, err
		}

		body := bytes.NewBuffer(jsonStr)
		request, err = http.NewRequest(method, url, body)
		if err != nil {
			return nil, err
		}

		// ตั้ง content-type เป็น application/json
		request.Header.Set("Content-Type", "application/json")
	} else {
		// ถ้าไม่มี body สร้าง request แบบไม่มี payload
		request, err = http.NewRequest(method, url, nil)
		if err != nil {
			return nil, err
		}
	}

	// เพิ่ม header x-api-key
	request.Header.Add("x-api-key", ApiKey)

	// ถ้ามี header เสริม ให้เพิ่มเข้าไปใน request
	if len(header) > 0 {
		for k, v := range header[0] {
			request.Header.Set(k, v)
		}
	}

	return request, nil
}

// GetErrorResponse ใช้เพื่อดึงข้อความ error ที่เหมาะสมจาก ErrorResponse
func GetErrorResponse(errorResponse ErrorResponse) string {
	// ถ้ามี messageCode ให้ใช้ messageCode
	errorMessage := errorResponse.MessageCode
	if errorMessage == "" {
		// ถ้าไม่มี messageCode ใช้ message ปกติแทน
		errorMessage = errorResponse.Message
	}
	return errorMessage
}
