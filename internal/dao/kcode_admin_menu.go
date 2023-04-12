// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"learn_goframe01/internal/dao/internal"
)

// internalKcodeAdminMenuDao is internal type for wrapping internal DAO implements.
type internalKcodeAdminMenuDao = *internal.KcodeAdminMenuDao

// kcodeAdminMenuDao is the data access object for table kcode_admin_menu.
// You can define custom methods on it to extend its functionality as you wish.
type kcodeAdminMenuDao struct {
	internalKcodeAdminMenuDao
}

var (
	// KcodeAdminMenu is globally public accessible object for table kcode_admin_menu operations.
	KcodeAdminMenu = kcodeAdminMenuDao{
		internal.NewKcodeAdminMenuDao(),
	}
)

// Fill with you ideas below.
