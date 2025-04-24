package interface_conv

import (
	"fmt"
	"reflect"
	"testing"
)

// ทดสอบฟังก์ชัน ToUint ซึ่งแปลงค่าต่าง ๆ (interface{}) ให้กลายเป็น uint64
func TestToUint(t *testing.T) {
	type args struct {
		iParem interface{} // อินพุตที่ต้องการทดสอบการแปลง
	}

	// ชุด test case ต่าง ๆ
	tests := []struct {
		name string // ชื่อเคสสำหรับดูตอนรัน
		args args   // อินพุต
		want uint64 // ค่าที่คาดหวังว่าจะได้
	}{
		// เคสแปลง string "1" เป็น uint64(1)
		{
			name: "string to int",
			args: args{iParem: "1"},
			want: uint64(1),
		},
		// แปลง float32 เป็น uint64 (ตัดเศษทิ้ง)
		{
			name: "float32 to int",
			args: args{iParem: float32(2.34)},
			want: uint64(2),
		},
		// แปลง float64 เป็น uint64 (ตัดเศษทิ้ง)
		{
			name: "float64 to int",
			args: args{iParem: float64(7.34)},
			want: uint64(7),
		},
		// แปลง int ปกติเป็น uint64
		{
			name: "int to int",
			args: args{iParem: int(3)},
			want: uint64(3),
		},
		// แปลง int32 เป็น uint64
		{
			name: "int32 to int",
			args: args{iParem: int32(4)},
			want: uint64(4),
		},
		// แปลง int64 เป็น uint64
		{
			name: "int64 to int",
			args: args{iParem: int64(5)},
			want: uint64(5),
		},
		// เคสที่แปลงไม่ได้ เช่น string "One"
		{
			name: "can't convert",
			args: args{iParem: "One"},
			want: uint64(0), // ควรคืนค่า default
		},
	}

	// วนทดสอบทุกเคส
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := ToUint(tt.args.iParem)
			xType := reflect.TypeOf(tt.args.iParem)
			// แสดงผล type และค่าที่ทดสอบในแต่ละรอบ
			fmt.Println("Type:", xType, " Parem:", tt.args.iParem, " Want:", tt.want, " Got:", got)
			// ตรวจสอบว่าผลลัพธ์ตรงกับที่ต้องการไหม
			if got != tt.want {
				t.Errorf("ToUint() = %v, want %v", got, tt.want)
			}
		})
	}
}

// ทดสอบฟังก์ชัน ToFloat ซึ่งแปลงค่าต่าง ๆ (interface{}) ให้กลายเป็น float64
func TestToFloat(t *testing.T) {
	type args struct {
		fParem interface{} // อินพุตที่ต้องการแปลง
	}

	// ชุด test case
	tests := []struct {
		name string
		args args
		want float64 // ค่าที่คาดหวัง
	}{
		// int → float64
		{
			name: "int to float64",
			args: args{fParem: 1},
			want: float64(1.00),
		},
		// int32 → float64
		{
			name: "int32 to float64",
			args: args{fParem: int32(1)},
			want: float64(1.00),
		},
		// int64 → float64
		{
			name: "int64 to float64",
			args: args{fParem: int64(1)},
			want: float64(1.00),
		},
		// string "2.34" → float64
		{
			name: "string to float64",
			args: args{fParem: "2.34"},
			want: float64(2.34),
		},
		// float32 → float64
		{
			name: "float32 to float64",
			args: args{fParem: float32(2.34)},
			want: float64(2.34),
		},
		// float64 → float64
		{
			name: "float64 to float64",
			args: args{fParem: float64(2.34)},
			want: float64(2.34),
		},
		// เคสที่แปลงไม่ได้ เช่น string "One"
		{
			name: "can't convert",
			args: args{fParem: "One"},
			want: float64(0.00),
		},
	}

	// วนทดสอบทุกเคส
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := ToFloat(tt.args.fParem)
			xType := reflect.TypeOf(tt.args.fParem)
			// แสดงผลการทดสอบ
			fmt.Println("Type:", xType, " Parem:", tt.args.fParem, " Want:", tt.want, " Got:", got)
			// เช็คว่าผลลัพธ์ตรงกับที่คาดหวังหรือไม่
			if got != tt.want {
				t.Errorf("ToFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}
