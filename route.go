package main

import (
	"selfmade-webframework/framework/framework/gin"
	"selfmade-webframework/framework/framework/middleware"
)

func registerRouter(core *gin.Engine) {
	// HTTP方法+靜態路由匹配 area
	core.GET("/user/login", UserLoginController)
	
	// 批量通用前綴area
	subjectApi := core.Group("/subject") 
	{ 
		// 動態路由area 
		subjectApi.DELETE("/:id", SubjectDelController) 
		subjectApi.PUT("/:id", SubjectUpdateController) 
		// 在group中使用middleware.Test3() 為單個路由建立中間件   
		subjectApi.GET("/:id", middleware.Test3(), SubjectGetController)
		subjectApi.GET("/list/all", SubjectListController)
	}
	// subjectApi.Use(middleware.Test3())


}
