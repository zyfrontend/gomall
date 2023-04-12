// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// KcodeUserRecharge is the golang structure for table kcode_user_recharge.
type KcodeUserRecharge struct {
	Id             int     `json:"id"             description:""`
	UserId         int     `json:"userId"         description:""`
	PayAddress     string  `json:"payAddress"     description:"客户钱包地址"`
	WalletAddress  string  `json:"walletAddress"  description:"商家收款钱包地址"`
	Hash           string  `json:"hash"           description:"hash"`
	CreateTime     int     `json:"createTime"     description:"创建时间"`
	Status         int     `json:"status"         description:"是否已经处理 0 未处理 1 已经处理 2处理失败"`
	Money          float64 `json:"money"          description:"充值金额"`
	SetStatus      int     `json:"setStatus"      description:"脚本执行装0未执行， 1已执行"`
	UpdateTime     int     `json:"updateTime"     description:"审核时间"`
	Content        string  `json:"content"        description:"备注"`
	PayImg         string  `json:"payImg"         description:"打款凭证"`
	RechargeStatus int     `json:"rechargeStatus" description:"充值状态（0=待确认,1'已到账',2'到账失败'）"`
	MoneyType      int     `json:"moneyType"      description:"币种类型"`
}
