package models

type Users struct {
	UUID     string `json:"uuid" gorm:"column:user_id;primaryKey;default:gen_random_uuid()"`
	Username string `json:"username" binding:"required" gorm:"column:name"`
	Email    string `json:"email" binding:"required,email" gorm:"column:email"`
	Age      int    `json:"age" binding:"required,gt=18" gorm:"column:age"`
	Password string `json:"password" binding:"required" gorm:"column:password"`
	Status   int    `json:"status" binding:"required,oneof=1 2" gorm:"column:status"`
	Level    int    `json:"level" binding:"required,oneof=1 2" gorm:"column:level"`
}
