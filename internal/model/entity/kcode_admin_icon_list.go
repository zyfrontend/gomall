// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// KcodeAdminIconList is the golang structure for table kcode_admin_icon_list.
type KcodeAdminIconList struct {
	Id     uint64 `json:"id"     description:""`
	IconId uint   `json:"iconId" description:"所属图标id"`
	Title  string `json:"title"  description:"图标标题"`
	Class  string `json:"class"  description:"图标类名"`
	Code   string `json:"code"   description:"图标关键词"`
}
