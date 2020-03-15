package router

import (
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	"log"
	"rankdemo2/control"
)


//连接数据库
var Db *sqlx.DB

func init(){
	db,err := sqlx.Open("mysql","root:root@tcp(127.0.0.1:3306)/zhihu?charset=utf8")
	if err != nil{
		log.Fatal(err.Error())
	}
	if err = db.Ping() ;err != nil{
		log.Fatal(err.Error())
	}
	Db=db

}

func Run (){
	rank:= echo.New()
	rank.POST("/login",control.Login)//登录
	rank.POST("/signup",control.Signup)//注册
	rank.GET("/rank",control.Rank)//查看排行榜
	api := rank.Group("/api",control.ServerHeader)
	api.POST("/join/:userid",control.Join)//参赛
	api.GET("/cancel/:userid",control.Cancel)//退赛
	api.GET("/:id/vote/:userid",control.Vote)//投票
	rank.Start(":8080")

}
