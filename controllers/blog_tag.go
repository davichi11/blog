package controllers

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

//
type BlogTag struct {
	BaseNoLogin
}

func (c *BlogTag) GetAll() {
	var result map[string]interface{}
	s := "{\"code\":1,\"info\":\"ok\",\"data\":[{\"tag_id\":1,\"name\":\"PHP\"},{\"tag_id\":2,\"name\":\"JAVA\"}]}"
	err := json.Unmarshal([]byte(s), &result)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(result)
	c.Data["json"] = result
	c.ServeJSON()
}
