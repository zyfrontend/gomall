// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"learn_goframe01/internal/dao/internal"
)

// internalKcodeAdminAttachmentDao is internal type for wrapping internal DAO implements.
type internalKcodeAdminAttachmentDao = *internal.KcodeAdminAttachmentDao

// kcodeAdminAttachmentDao is the data access object for table kcode_admin_attachment.
// You can define custom methods on it to extend its functionality as you wish.
type kcodeAdminAttachmentDao struct {
	internalKcodeAdminAttachmentDao
}

var (
	// KcodeAdminAttachment is globally public accessible object for table kcode_admin_attachment operations.
	KcodeAdminAttachment = kcodeAdminAttachmentDao{
		internal.NewKcodeAdminAttachmentDao(),
	}
)

// Fill with you ideas below.