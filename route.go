package main

import (
	"selfmade-webframework/framework"
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
		subjectApi.Get("/:id", SubjectGetController) 
		subjectApi.Get("/list/all", SubjectListController)
	}
	// subjectApi.Use(middleware.Test3())


}
