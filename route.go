package main

import (
	"selfmade-webframework/framework/framework"
	"selfmade-webframework/framework/framework/middleware"
)

func registerRouter(core *framework.Core) {
	// HTTP方法+靜態路由匹配 area
	core.Get("/user/login", UserLoginController)
	
	// 批量通用前綴area
	subjectApi := core.Group("/subject") 
	{ 
		// 動態路由area 
		subjectApi.Delete("/:id", SubjectDelController) 
		subjectApi.Put("/:id", SubjectUpdateController) 
		// 在group中使用middleware.Test3() 為單個路由建立中間件   
		subjectApi.Get("/:id", middleware.Test3(), SubjectGetController)
		subjectApi.Get("/list/all", SubjectListController)
	}
	// subjectApi.Use(middleware.Test3())


}
