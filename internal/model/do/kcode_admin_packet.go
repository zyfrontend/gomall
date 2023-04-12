// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// KcodeAdminPacket is the golang structure of table kcode_admin_packet for DAO operations like Where/Data.
type KcodeAdminPacket struct {
	g.Meta     `orm:"table:kcode_admin_packet, do:true"`
	Id         interface{} //
	Name       interface{} // 数据包名
	Title      interface{} // 数据包标题
	Author     interface{} // 作者
	AuthorUrl  interface{} // 作者url
	Version    interface{} //
	Tables     interface{} // 数据表名
	CreateTime interface{} // 创建时间
	UpdateTime interface{} // 更新时间
	Status     interface{} // 状态
}
