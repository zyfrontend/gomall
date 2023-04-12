// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"learn_goframe01/internal/dao/internal"
)

// internalKcodeAdminIconDao is internal type for wrapping internal DAO implements.
type internalKcodeAdminIconDao = *internal.KcodeAdminIconDao

// kcodeAdminIconDao is the data access object for table kcode_admin_icon.
// You can define custom methods on it to extend its functionality as you wish.
type kcodeAdminIconDao struct {
	internalKcodeAdminIconDao
}

var (
	// KcodeAdminIcon is globally public accessible object for table kcode_admin_icon operations.
	KcodeAdminIcon = kcodeAdminIconDao{
		internal.NewKcodeAdminIconDao(),
	}
)

// Fill with you ideas below.
