package blog

import (
	"blog/fox"
	"blog/fox/db"
	"blog/model"
)

//同步第三方博客
type BlogSyncMapping struct {
}

//快速初始化
func NewBlogSyncMappingService() *BlogSyncMapping {
	return new(BlogSyncMapping)
}

//根据博客ID和类别ID 获取 第三方ID数据
func (t *BlogSyncMapping) GetAppId(typeId, blogId int) (*model.BlogSyncMapping, error) {
	maps := make(map[string]interface{})
	maps["blog_id"] = blogId
	maps["type_id"] = typeId
	m := model.NewBlogSyncMapping()
	_, err := db.Filter(maps).Get(m)
	if err != nil {
		return nil, fox.NewError("不存在")
	}
	if m.MapId < 1 {
		return nil, fox.NewError("不存在")
	}
	return m, nil
}
