package model

import (
	"github.com/go-redis/redis"
	"github.com/jmoiron/sqlx"
	"log"
)

//连接数据库
var Db *sqlx.DB

func init(){
	db,err := sqlx.Open("mysql","root:root@tcp(127.0.0.1:3306)/rankdemo?charset=utf8")
	if err != nil{
		log.Fatal(err.Error())
	}
	if err = db.Ping() ;err != nil{
		log.Fatal(err.Error())
	}
	Db=db

}

// 定义redis链接池
var client *redis.Client
//连接redis
func RedisClient() {
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := client.Ping().Result()
	if err !=nil {
		panic(err)
	}

}