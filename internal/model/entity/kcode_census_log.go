// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// KcodeCensusLog is the golang structure for table kcode_census_log.
type KcodeCensusLog struct {
	LogId         uint    `json:"logId"         description:""`
	UserId        uint    `json:"userId"        description:""`
	MoneyType     int     `json:"moneyType"     description:"币种"`
	AddTime       uint    `json:"addTime"       description:"添加时间"`
	RechargeMoney float64 `json:"rechargeMoney" description:"充值金额"`
	AtmMoney      float64 `json:"atmMoney"      description:"提现金额"`
	ZhituiBonus   float64 `json:"zhituiBonus"   description:"直推奖励"`
	TeamBonus     float64 `json:"teamBonus"     description:"团队奖励"`
	WholeBonus    float64 `json:"wholeBonus"    description:"全网加权分红"`
	EquityBonus   float64 `json:"equityBonus"   description:"权益分红"`
	ReleaseBonus  float64 `json:"releaseBonus"  description:"释放收益"`
	PjBonus       float64 `json:"pjBonus"       description:"平级奖"`
}
