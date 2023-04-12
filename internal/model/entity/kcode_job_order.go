// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// KcodeJobOrder is the golang structure for table kcode_job_order.
type KcodeJobOrder struct {
	OrderId    int    `json:"orderId"    description:""`
	OrderSn    string `json:"orderSn"    description:""`
	UserId     int    `json:"userId"     description:""`
	JobId      int    `json:"jobId"      description:"节点ID"`
	AddTime    int    `json:"addTime"    description:"添加时间"`
	Status     int    `json:"status"     description:"0排队中1待审核2已通过"`
	UpdateTime int    `json:"updateTime" description:"修改时间"`
}