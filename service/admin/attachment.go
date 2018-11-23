package admin

import (
	"blog/fox/db"
	"blog/model"
	"fmt"
)

//附件
type Attachment struct {
	*model.Attachment
}

//快速初始化
func NewAttachmentService() *Attachment {
	return new(Attachment)
}

//列表
func (c *Attachment) GetAll(q map[string]interface{}, fields []string, orderBy string, page int, limit int) (*db.Paginator, error) {
	mode := model.NewAttachment()
	data, err := mode.GetAll(q, fields, orderBy, page, 20)
	if err != nil {
		return nil, err
	}
	return data, nil
}

//更新添加的数据
//@type_id 模块
//@id      id
//@aid     管理员ID
func (c *Attachment) UpdateByTypeIdId(typeId, aid, id int) (bool, error) {
	maps := make(map[string]interface{})
	maps["type_id"] = typeId
	maps["id"] = 0
	maps["aid"] = aid
	mod := model.NewAttachment()
	mod.Id = id
	//fmt.Println("save",mod)
	num, err := db.Filter(maps).Update(mod)
	if err != nil {
		return false, err
	}
	fmt.Println("更新条数", num)
	return true, nil
}
