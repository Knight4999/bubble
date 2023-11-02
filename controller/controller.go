package controller

import (
	"bubble/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// IndexHandle 加载index模板
func IndexHandle(cx *gin.Context) {
	cx.HTML(http.StatusOK, "index.html", nil)
}

// CreateTodo 添加一个待办事项
func CreateTodo(ctx *gin.Context) {
	//前端页面，填写一个待办事项，点击提交 会到这里
	//1.把请求中的数据拿出来
	var todo models.Todo
	ctx.BindJSON(&todo)
	//2.将数据添加到数据库
	err := models.CreateATodo(&todo)
	//3.返回响应
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return //如果添加失败，则发挥错误信息
	} else {
		ctx.JSON(http.StatusOK, todo) //添加成功则返回todo
		/*ctx.JSON(http.StatusOK, gin.H{
			"code": 20001,
			"msg":  "success",
			"data": todo,
		})*/
	}
}

// TodoListAll 查看所有的待办事项
func TodoListAll(ctx *gin.Context) {
	todos, err := models.GetTodoList()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, todos)
	}
}

// ModifyTodo 修改待办事项状态
func ModifyTodo(ctx *gin.Context) {
	id, ok := ctx.Params.Get("id")
	if !ok {
		ctx.JSON(http.StatusOK, gin.H{"err": "无效的ID"})
		return
	}
	todo, err := models.GetATodoByID(id)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"err": err.Error()})
		return
	}
	ctx.BindJSON(&todo)

	if err = models.UpdateTodo(todo); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"err": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, todo)
	}
}

// DeleteATodoByID 删除一个待办事项
func DeleteATodoByID(ctx *gin.Context) {
	id, ok := ctx.Params.Get("id")
	if !ok {
		ctx.JSON(http.StatusOK, gin.H{"err": "无效的ID"})
		return
	}

	if err := models.DeleteTodo(id); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"err": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{id: "Deleted"})
	}
}
