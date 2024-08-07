package model

import "gorm.io/gorm"

type Permissions struct {
	gorm.Model
	Name        string            `gorm:"unique;size:20;not null"` // 角色名，唯一，长度限制为20
	Description string            `gorm:"size:100"`
	Status      Status            `gorm:"type:enum('0', '1');default:'0'"` // 状态
	Roles       []RolePermissions // 角色
}

// TableName 返回 Permissions 结构体对应的表名。
func (Permissions) TableName() string {
	return "permissions" // 设置表名为 "Permissions"
}
