// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// KcodeAdPosition is the golang structure for table kcode_ad_position.
type KcodeAdPosition struct {
	Id            uint   `json:"id"            description:"表id"`
	PositionName  string `json:"positionName"  description:"广告位置名称"`
	AdWidth       uint   `json:"adWidth"       description:"广告位宽度"`
	AdHeight      uint   `json:"adHeight"      description:"广告位高度"`
	PositionDesc  string `json:"positionDesc"  description:"广告描述"`
	PositionStyle string `json:"positionStyle" description:"模板"`
	IsOpen        int    `json:"isOpen"        description:"0关闭1开启"`
}
