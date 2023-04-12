// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// KcodeAdminAction is the golang structure for table kcode_admin_action.
type KcodeAdminAction struct {
	Id         uint   `json:"id"         description:""`
	Module     string `json:"module"     description:"所属模块名"`
	Name       string `json:"name"       description:"行为唯一标识"`
	Title      string `json:"title"      description:"行为标题"`
	Remark     string `json:"remark"     description:"行为描述"`
	Rule       string `json:"rule"       description:"行为规则"`
	Log        string `json:"log"        description:"日志规则"`
	Status     int    `json:"status"     description:"状态"`
	CreateTime uint   `json:"createTime" description:"创建时间"`
	UpdateTime uint   `json:"updateTime" description:"更新时间"`
}