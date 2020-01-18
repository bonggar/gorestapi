package model

//User : user struct
type User struct {
	Model
	//required fields below
	Name   string `gorm:"column:name;size:100;not null" form:"name" json:"name" binding:"required,max=100"`
	Email  string `gorm:"column:email;type:varchar(100);unique;not null" form:"email" json:"email" binding:"required,email,max=100"`
	Phone  string `gorm:"column:phone;type:varchar(15);unique;not null" form:"phone" json:"phone" binding:"required,numeric,min=10,max=15"`
	Dob    string `gorm:"column:dob;type:date;not null" form:"dob" json:"dob" binding:"required" time_format:"2006-01-02"` //dob field have to validated outside gorm
	Gender string `gorm:"column:gender;type:char(1);not null" form:"gender" json:"gender" binding:"required,oneof=m f"`
	//non required fields below (allow null value)
	Address *string `gorm:"column:address;type:varchar(255);default:null" form:"address" json:"address" binding:"omitempty,max=255"`
}
