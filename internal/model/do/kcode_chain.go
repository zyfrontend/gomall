// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// KcodeChain is the golang structure of table kcode_chain for DAO operations like Where/Data.
type KcodeChain struct {
	g.Meta `orm:"table:kcode_chain, do:true"`
	Id     interface{} //
	Name   interface{} // 链名称
	TypeId interface{} // 链类型（1=erc20，2=trc20）
}
