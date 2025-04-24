package interface_conv

import (
	"errors"
	"strconv"
	"strings"
)

// ToUint แปลงค่าแบบ interface{} ให้กลายเป็น uint64
func ToUint(iParem interface{}) (uint64, error) {
	switch v := iParem.(type) {
	case int:
		// ถ้าเป็น int → แปลงเป็น uint64 ได้โดยตรง
		return uint64(v), nil
	case int32:
		// ถ้าเป็น int32 → แปลงเป็น uint64 ได้โดยตรง
		return uint64(v), nil
	case int64:
		// ถ้าเป็น int64 → แปลงเป็น uint64 ได้โดยตรง
		return uint64(v), nil
	case float32:
		// ถ้าเป็น float32 → แปลงเป็น uint64 (ตัดเศษทศนิยมทิ้ง)
		return uint64(v), nil
	case float64:
		// ถ้าเป็น float64 → แปลงเป็น uint64 (ตัดเศษทศนิยมทิ้ง)
		return uint64(v), nil
	case string:
		// ถ้าเป็น string → ลอง parse เป็น uint64 โดยใช้ base 10 และ bit size 32
		// เช่น "123" หรือ "  456  "
		intVal, err := strconv.ParseUint(strings.TrimSpace(v), 10, 32)
		if err != nil {
			// ถ้าแปลงไม่สำเร็จ (เช่น "abc") → คืน error กลับไป
			return 0, err
		}
		return intVal, nil
	default:
		// ถ้า type ไม่รู้จัก → ส่ง error กลับ
		return 0, errors.New("can't convert")
	}
}

// ToFloat แปลงค่าแบบ interface{} ให้กลายเป็น float64
func ToFloat(fParem interface{}) (float64, error) {
	var floatVal float64
	var err error
	switch v := fParem.(type) {
	case int:
		// int → float64 ได้โดยตรง
		return float64(v), nil
	case int32:
		// int32 → float64 ได้โดยตรง
		return float64(v), nil
	case int64:
		// int64 → float64 ได้โดยตรง
		return float64(v), nil
	case float32:
		// float32 → float64 ได้โดยตรง
		return float64(v), nil
	case float64:
		// float64 → float64 (ไม่ต้องแปลง)
		return float64(v), nil
	case string:
		// string → พยายามแปลงเป็น float64
		floatVal, err = strconv.ParseFloat(strings.TrimSpace(fParem.(string)), 64)
		if err != nil {
			// ถ้าแปลงไม่สำเร็จ (เช่น "One") → คืน error
			return 0, err
		}
		return floatVal, nil
	default:
		// type อื่นที่ไม่รู้จัก → คืน error
		return 0, errors.New("can't convert")
	}
}
