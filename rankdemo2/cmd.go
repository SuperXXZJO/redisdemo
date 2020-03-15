package main

import (
	_ "github.com/go-sql-driver/mysql"
	"rankdemo2/control"
	"rankdemo2/model"
	"rankdemo2/router"
)



func main (){
	model.RedisClient()
	control.Update()
	router.Run()
}

