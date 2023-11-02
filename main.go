package main

//小清单项目
import (
	"bubble/dao"
	"bubble/models"
	"bubble/routers"
	"fmt"
)

func main() {
	//创建数据库
	//SQL：Create Database bubble
	//连接数据库
	err := dao.InitMysql()
	if err != nil {
		fmt.Printf("connect database faild,err: %v\n", err)
		panic(err)
	}
	//建立连接关系
	dao.DB.AutoMigrate(&models.Todo{})

	//获取路由
	r := routers.SetupRouter()
	r.Run(":23450")
}
