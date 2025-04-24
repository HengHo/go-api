package models

// StudentModel คือโครงสร้างข้อมูลที่ใช้ใน GORM เพื่อแมปกับตาราง "students" ในฐานข้อมูล
type StudentModel struct {
	// Id คือ primary key ที่ใช้ในตาราง students โดยใช้ประเภท serial ซึ่งจะเป็นการเพิ่มค่าอัตโนมัติ
	Id        int    `gorm:"column:id;primary_key;type:serial"`   // Primary key with serial type
	// FirstName คือชื่อจริงของนักเรียน โดยใช้ประเภท varchar(255)
	FirstName string `gorm:"column:first_name;type:varchar(255)"` // First name with varchar(255)
	// LastName คือนามสกุลของนักเรียน โดยใช้ประเภท varchar(255)
	LastName  string `gorm:"column:last_name;type:varchar(255)"`  // Last name with varchar(255)
	// Email คืออีเมลของนักเรียน โดยใช้ประเภท varchar(255)
	Email     string `gorm:"column:email;type:varchar(255)"`      // Email with varchar(255)
}

// TableName เป็นเมธอดที่บอก GORM ว่าตารางในฐานข้อมูลที่จะใช้สำหรับการแมปกับ StudentModel คือ "students"
func (e *StudentModel) TableName() string {
	return "students"
}
