// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// KcodeSmsEmailLogDao is the data access object for table kcode_sms_email_log.
type KcodeSmsEmailLogDao struct {
	table   string                  // table is the underlying table name of the DAO.
	group   string                  // group is the database configuration group name of current DAO.
	columns KcodeSmsEmailLogColumns // columns contains all the column names of Table for convenient usage.
}

// KcodeSmsEmailLogColumns defines and stores column names for table kcode_sms_email_log.
type KcodeSmsEmailLogColumns struct {
	Id           string // 记录ID
	Email        string // 用户邮箱
	Scene        string // 发送场景 1:用户注册,2:找回登录密码,3:找回资金密码,3:客户下单,4:客户支付,5:商家发货,6:实名验证' 7:退款成功
	Status       string // 发送状态
	Code         string // 验证码
	Msg          string // 发送短信内容
	AddTime      string // 添加时间
	ValidityTime string // 有效期
	Ip           string //
	ErrorMsg     string // 发送短信异常内容
}

// kcodeSmsEmailLogColumns holds the columns for table kcode_sms_email_log.
var kcodeSmsEmailLogColumns = KcodeSmsEmailLogColumns{
	Id:           "id",
	Email:        "email",
	Scene:        "scene",
	Status:       "status",
	Code:         "code",
	Msg:          "msg",
	AddTime:      "add_time",
	ValidityTime: "validity_time",
	Ip:           "ip",
	ErrorMsg:     "error_msg",
}

// NewKcodeSmsEmailLogDao creates and returns a new DAO object for table data access.
func NewKcodeSmsEmailLogDao() *KcodeSmsEmailLogDao {
	return &KcodeSmsEmailLogDao{
		group:   "default",
		table:   "kcode_sms_email_log",
		columns: kcodeSmsEmailLogColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *KcodeSmsEmailLogDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *KcodeSmsEmailLogDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *KcodeSmsEmailLogDao) Columns() KcodeSmsEmailLogColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *KcodeSmsEmailLogDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *KcodeSmsEmailLogDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *KcodeSmsEmailLogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}