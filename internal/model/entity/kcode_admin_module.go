// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// KcodeAdminModule is the golang structure for table kcode_admin_module.
type KcodeAdminModule struct {
	Id           uint   `json:"id"           description:""`
	Name         string `json:"name"         description:"模块名称（标识）"`
	Title        string `json:"title"        description:"模块标题"`
	Icon         string `json:"icon"         description:"图标"`
	Description  string `json:"description"  description:"描述"`
	Author       string `json:"author"       description:"作者"`
	AuthorUrl    string `json:"authorUrl"    description:"作者主页"`
	Config       string `json:"config"       description:"配置信息"`
	Access       string `json:"access"       description:"授权配置"`
	Version      string `json:"version"      description:"版本号"`
	Identifier   string `json:"identifier"   description:"模块唯一标识符"`
	SystemModule uint   `json:"systemModule" description:"是否为系统模块"`
	CreateTime   uint   `json:"createTime"   description:"创建时间"`
	UpdateTime   uint   `json:"updateTime"   description:"更新时间"`
	Sort         int    `json:"sort"         description:"排序"`
	Status       int    `json:"status"       description:"状态"`
}
