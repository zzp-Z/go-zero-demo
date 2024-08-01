package model

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Name   string     `gorm:"unique;size:20;not null"`         // 角色名，唯一，长度限制为20
	Status Status     `gorm:"type:enum('0', '1');default:'0'"` // 状态
	Users  []UserRole // 用户角色关系
}

// TableName 返回 User 结构体对应的表名。
func (Role) TableName() string {
	return "role" // 设置表名为 "user"
}
