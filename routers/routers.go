package routers

import (
	"bubble/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	//加载模板文件路径
	r.LoadHTMLGlob("templates/*")
	//加载静态模板文件路径
	r.Static("/static", "static")

	//路由
	r.GET("/index", controller.IndexHandle)

	//v1 todo
	v1Group := r.Group("v1")
	{
		//待办事项
		//添加
		v1Group.POST("/todo", controller.CreateTodo)
		//查看素有待办事项
		v1Group.GET("/todo", controller.TodoListAll)
		/*//查看某一个待办事项
		v1Group.GET("/todo/:id", func(ctx *gin.Context) {

		})*/
		//修改
		v1Group.PUT("/todo/:id", controller.ModifyTodo)
		//删除
		v1Group.DELETE("/todo/:id", controller.DeleteATodoByID)

	}
	return r
}
