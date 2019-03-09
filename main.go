package main

import (
	"blog/controllers"
	"blog/controllers/admin"
	_ "blog/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func main() {
	orm.Debug = true
	addRouters()
	//beego.AddFuncMap("array",array)
	beego.Run()
}

func addRouters() {
	//log4go.LoadConfiguration("log4go.xml")
	beego.ErrorController(&controllers.Error{})
	beego.SetStaticPath("/uploads", "uploads")
	//beego.Router("/admin/login", &admin.LoginController{})
	//beego.Router("/admin/index", &admin.IndexController{})
	ns := beego.NewNamespace("/admin",
		beego.NSRouter("/login", &admin.Login{}),
		beego.NSRouter("/logout", &admin.Logout{}),
		beego.NSRouter("/my_password", &admin.MyPassword{}),
		//blog
		//使用注解路由
		//beego.NSRouter("/blog", &admin.BlogController{}, "get:List"),
		//beego.NSRouter("/blog", &admin.BlogController{}, ),
		//beego.NSRouter("/blog/detail/:id", &admin.BlogController{}, "get:Detail"),
		//beego.NSRouter("/blog/:id", &admin.BlogController{}, "put:Put"), //, "put:Put"
		//beego.NSRouter("/blog/:id", &admin.BlogController{}, "get:Get"),
		//beego.NSRouter("/blog/add", &admin.BlogController{}, "get:Add"),
		//type
		//使用注解路由
		//beego.NSRouter("/types", &admin.TypeController{},"get:List"),
		//beego.NSRouter("/types/add", &admin.TypeController{},"get:Add"),
		//beego.NSRouter("/types/add/:id", &admin.TypeController{},"get:Add"),
		//beego.NSRouter("/types/:id", &admin.TypeController{},"get:ListChild"),
		//beego.NSRouter("/type/:id", &admin.TypeController{},"get:Get"),// "get:Get"
		//beego.NSRouter("/type", &admin.TypeController{}),
		beego.NSInclude(
			&admin.Index{},
			&admin.Area{},
			&admin.Attachment{},
			&admin.Type{},
			&admin.BlogTag{},
			&admin.BlogCat{},
			&admin.Blog{},
			&admin.BlogSync{},
			&admin.Select{},
			&admin.Upload{},
			&admin.AdminUser{},
			&admin.AdminRole{},
			&admin.Member{},
			&admin.MemberGroup{},
			&admin.Site{},
			&admin.Oauth{},
		),
	)
	beego.AddNamespace(ns)
	//首页
	beego.Router("/", &controllers.Blog{}, "get:GetAll")
	beego.Router("/search/", &controllers.Blog{}, "get:GetAll")
	beego.Router("/page/:page", &controllers.Blog{}, "get:GetAll")
	beego.Router("/article/:id", &controllers.Blog{}, "get:Get")
	beego.Router("/tag/:tag", &controllers.Tag{}, "get:GetAll")
	//API
	beego.Router("/api/blog/create", &controllers.Blog{}, "post:Create")
	beego.Router("/api/blog/cat", &controllers.BlogCat{}, "get:GetAll")
	beego.Router("/api/blog/tag", &controllers.BlogTag{}, "get:GetAll")
}

//func array(m map[string]interface{}, s int) []string {
//	if m != nil {
//		a := m["tag_"+strconv.Itoa(s)]
//		if a != nil {
//			b := a.([]string)
//			if len(b) > 0 {
//				return b
//			}
//		}
//	}
//
//	return []string{}
//}
