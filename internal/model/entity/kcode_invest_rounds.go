// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// KcodeInvestRounds is the golang structure for table kcode_invest_rounds.
type KcodeInvestRounds struct {
	Id        int     `json:"id"        description:""`
	StartTime int     `json:"startTime" description:"开始时间"`
	EndTime   int     `json:"endTime"   description:"结束时间"`
	IsLock    int     `json:"isLock"    description:"1默认锁住 排队处理完再恢复0"`
	AddTime   int     `json:"addTime"   description:"添加时间"`
	MaxUsdt   float64 `json:"maxUsdt"   description:"每轮剂量"`
}