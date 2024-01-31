package entity

type User struct {
	StudentID	string valid:"required~กรุณากรอก StudentID ให้ถูกต้อง, matches(^[BMD]\\d{7}$)~ขึ้นต้นด้วย BMD เท่านั้น"
	Password	string valid:"required~กรุณากรอกรหัสผ่าน, stringlength(8|30)~ความยาวรหัสผ่านไม่เกิน 8-30 ตัวอักษร"
	Age			int valid:"range(15|50)~กรุณากรอกอายุระหว่าง 15-50"
	FirstName	string valid:"required~กรุณากรอกชื่อจริง, stringlength(8|30)~ความยาวรหัสผ่านไม่เกิน 8-30 ตัวอักษร"
	Email		string valid:"required~กรุณากรอกอีเมล, stringlength(8|30)~ความยาวรหัสผ่านไม่เกิน 8-30 ตัวอักษร"
	Profile		string	gorm:"type:longtext"
	Phone		string	valid:"required~กรุณากรอกเบอร์โทร, stringlength(10|10)~เบอร์โทรต้อง 10 character"
	URL			string valid:"required~กรุณาเชื่อมURL, minstringlength(2)~ใส่URL"
}