package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name      string     `gorm:"unique;size:20;not null"`         // 用户名，唯一，长度限制为20
	Password  string     `gorm:"not null;size:200"`               // 密码 不能是空的
	Email     string     `gorm:"unique;not null"`                 // 邮箱，唯一 不能是空的
	Status    Status     `gorm:"type:enum('0', '1');default:'0'"` // 状态
	AvatarUrl string     // 头像
	Roles     []UserRole // 角色
}

// TableName 返回 User 结构体对应的表名。
func (User) TableName() string {
	return "user" // 设置表名为 "user"
}
