package model

import "gorm.io/gorm"

type UserRole struct {
	gorm.Model
	UserID uint `gorm:"index;uniqueIndex:idx_user_role"` // 外键，用户ID
	RoleID uint `gorm:"index;uniqueIndex:idx_user_role"` // 外键，角色ID
	User   User `gorm:"foreignKey:UserID"`               // 关联的用户
	Role   Role `gorm:"foreignKey:RoleID"`               // 关联的角色
}

// TableName 返回 User 结构体对应的表名。
func (UserRole) TableName() string {
	return "user_role" // 设置表名为 "user"
}
