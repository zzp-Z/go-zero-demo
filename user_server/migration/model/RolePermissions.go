package model

import "gorm.io/gorm"

type RolePermissions struct {
	gorm.Model
	PermissionsID uint        `gorm:"index;uniqueIndex:idx_role_permissions"` // 外键，权限ID
	RoleID        uint        `gorm:"index;uniqueIndex:idx_role_permissions"` // 外键，角色ID
	Permissions   Permissions `gorm:"foreignKey:PermissionsID"`               // 关联的权限
	Role          Role        `gorm:"foreignKey:RoleID"`                      // 关联的角色
}

// TableName 返回 RolePermissions 结构体对应的表名。
func (RolePermissions) TableName() string {
	return "role_permissions" // 设置表名为 "role_permissions"
}
