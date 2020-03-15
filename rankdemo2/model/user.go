package model

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"strconv"
)

type LOGIN struct {
	Phone int
	Password string
}

type User struct {
	Userid int        `json:"Userid"` //id
	Username string    `json:"Username"` //昵称
	Phone int             `json:"Phone"` //手机号
	Password string        `json:"Password"` //密码
	Votes   int         `json:"Votes"` //票数


}

//token
type UserToken struct {
	Userid int
	Username string
	jwt.StandardClaims
}


//登录
func Login(phone int) (User,error){
	mod := User{}
	err := Db.Get(&mod,"select * from user where phone = ?",phone)
	return mod,err
}

//注册 signup
func Signup(mod *User)error{
	//开启事务
	sp,err :=Db.Begin()
	if err != nil{
		return err
	}
	result,err := sp.Exec("insert into user(phone,password)values (?,?)",mod.Phone,mod.Password)
	if err != nil{
		//回滚
		sp.Rollback()
		return err
	}
	rows,_ := result.RowsAffected()
	if rows <1{
		//回滚
		sp.Rollback()
		return errors.New("rows affected < 1")
	}
	//提交
	sp.Commit()
	return nil
}

//参赛
func Join (userid string) string{
	err := client.Set(userid, 0,0 ).Err()
	if err != nil {
		panic(err)
	}
	res,_ := client.Get(userid).Result()
	return res
}

//退赛
func Cancel (userid string)error{
	err := client.Del(userid).Err()
	if err != nil {
		panic(err)
	}
	return nil
}





//用户投票(每人每天只有三票)
func Vote (userid int,id string) int64{
	newuserid :=strconv.Itoa(userid)
	//查询用户是否参赛
	res,_ :=client.Get(newuserid).Result()
	if res == ""{
		return 3
	}

	//查票数
	mod := User{}
	err := Db.Get(&mod,"select  * from user where userid = ? limit 1",id)
	if err != nil {
		panic(err)
	}

	if mod.Votes > 0 {
		//投票
		_,err :=client.Incr(newuserid).Result()//增加票数
		client.ZIncrBy("rank",1,newuserid)//增加排行榜票数
		if err != nil {
			panic(err)
		}
		//减票数
		_,err1 := Db.Exec("update user set votes = votes-1 where userid=?",id)
		if err1 != nil {
			panic(err1)
		}
		return 1
	}
	return 0

}

//查看排行榜
func GetRank ()[]string{
	res,err:=client.ZRevRange("rank",0,-1).Result()
	if err != nil {
		panic(err)
	}
	return res
}

//更新票数
func UpdateVotes (){
	_,err := Db.Exec("update user set votes = 3 ")
	if err != nil{
		panic(err)
	}
	return
}
