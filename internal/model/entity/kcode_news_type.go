// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// KcodeNewsType is the golang structure for table kcode_news_type.
type KcodeNewsType struct {
	Id          uint   `json:"id"          description:""`
	Name        string `json:"name"        description:"分类名称"`
	Pid         uint   `json:"pid"         description:"父id"`
	Status      uint   `json:"status"      description:"状态 0 不显示  1显示"`
	Sort        int    `json:"sort"        description:"排序序号"`
	SeoTitle    string `json:"seoTitle"    description:"seo"`
	Keywords    string `json:"keywords"    description:"keywords"`
	Description string `json:"description" description:"description"`
	Path        string `json:"path"        description:"节点"`
	Pic         string `json:"pic"         description:"图片"`
}
