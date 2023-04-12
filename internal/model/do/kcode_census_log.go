// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// KcodeCensusLog is the golang structure of table kcode_census_log for DAO operations like Where/Data.
type KcodeCensusLog struct {
	g.Meta        `orm:"table:kcode_census_log, do:true"`
	LogId         interface{} //
	UserId        interface{} //
	MoneyType     interface{} // 币种
	AddTime       interface{} // 添加时间
	RechargeMoney interface{} // 充值金额
	AtmMoney      interface{} // 提现金额
	ZhituiBonus   interface{} // 直推奖励
	TeamBonus     interface{} // 团队奖励
	WholeBonus    interface{} // 全网加权分红
	EquityBonus   interface{} // 权益分红
	ReleaseBonus  interface{} // 释放收益
	PjBonus       interface{} // 平级奖
}