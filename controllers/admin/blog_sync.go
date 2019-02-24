package admin

import (
	"blog/app/csdn"
	"blog/fox/array"
	"blog/model"
	"blog/service/blog"
	"fmt"
)

//博客同步
type BlogSync struct {
	Base
}

func (c *BlogSync) URLMapping() {
	c.Mapping("List", c.List)
	c.Mapping("Go", c.Go)
}

//列表
// @router /auth_csdn [get]
func (c *BlogSync) List() {

	//web:=csdn.NewAuthorizeWeb()
	//ok,err:=web.SetConfig()
	//if err !=nil{
	//	fmt.Println(err)
	//}
	//fmt.Println("status:",ok);
	//web.SetRedirectUri("http://www.foxwho.com:8080/admin/auth_token")
	//
	//c.Data["url"] = web.GetAuthorizeUrl()
	c.TplName = "admin/auth/list.html"
}

// @router /blog_sync/go [get]
func (c *BlogSync) Go() {
	blogId, _ := c.GetInt("blog_id")
	typeId, err := c.GetInt("type_id")
	if err == nil {
		id, err := c.GetInt("id")
		if err == nil {
			//调用CSDN 获取文章接口
			b, err := csdn.Get(typeId, id, blogId)
			if err != nil {
				c.Error(err.Error())
			} else {
				m, _ := array.ObjToMap(b)
				c.Success("ok", m)
			}
		} else {
			c.Error("id 不能为空")
		}
	} else {
		c.Error("type_id 不能为空")
	}

}

// @router /blog_sync/go [post]
func (c *BlogSync) Post() {
	blogId, err := c.GetInt("blog_id")
	if err != nil {
		c.Error("blog_id 不能为空")
	} else {
		typeId, err := c.GetInt("type_id")
		if err != nil {
			c.Error("type_id 不能为空")
		} else {
			mod := model.NewBlog()
			data, err := mod.GetById(blogId)
			if err == nil {
				//调用CSDN 文章创建接口
				sync := csdn.NewCsdnBlogApp()
				m := blog.NewBlogService()
				m.Blog = data
				id, err := sync.Create(m, typeId)
				if err == nil {
					maps := make(map[string]interface{})
					maps["id"] = id
					c.Success("成功!", maps)
				} else {
					c.Error(err.Error())
				}
			} else {
				c.Error(err.Error())
			}
		}
	}

}

// @router /blog_sync/go [put]
func (c *BlogSync) Put() {
	blogId, err := c.GetInt("blog_id")
	if err != nil {
		c.Error("blog_id 不能为空")
	} else {
		typeId, err := c.GetInt("type_id")
		if err != nil {
			c.Error("type_id 不能为空")
		} else {
			id, err := c.GetInt("id")
			if err == nil {
				mod := model.NewBlog()
				data, err := mod.GetById(blogId)
				if err == nil {
					//调用CSDN 文章更新接口
					sync := csdn.NewCsdnBlogApp()
					m := blog.NewBlogService()
					m.Blog = data
					err := sync.Update(m, typeId, id)
					if err == nil {
						c.Success("成功!")
					} else {
						c.Error(err.Error())
					}
				} else {
					fmt.Println(err)
					c.Error(err.Error())
				}
			} else {
				c.Error("id 不能为空")
			}
		}
	}

}
