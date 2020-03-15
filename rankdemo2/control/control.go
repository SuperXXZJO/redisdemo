package control

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"rankdemo2/model"
	"strconv"
	"time"
)

//login 登录
func Login (c echo.Context)error{
	phone,_:= strconv.Atoi(c.FormValue("phone"))
	password :=c.FormValue("password")
	mod,err :=model.Login(phone)
	if err != nil{
		return c.JSON(300,"请输入正确的手机号！")
	}
	if mod.Password != password{
		return c.JSON(300,"密码错误")
	}
	//生成token
	claims := model.UserToken{
		Userid: mod.Userid,
		Username: mod.Username,
		StandardClaims: jwt.StandardClaims{ExpiresAt:time.Now().Add(2 * time.Hour).Unix()},
	}
	token :=jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	ss,err := token.SignedString([]byte("123"))
	return c.JSON(200,ss)
}
//signuo 注册
func Signup (c echo.Context)error{
	inf := model.User{}
	err := c.Bind(&inf)
	if err != nil{
		return c.JSON(300,"输入数据有误")
	}
	if len(strconv.Itoa(inf.Phone)) !=11 {
		return c.JSON(300,"请输入正确的手机号码！")
	}
	if  inf.Password == ""{
		return c.JSON(300,"密码不能为空！")
	}

	err = model.Signup(&inf)
	if err != nil {
		return c.JSON(301,err.Error())
	}
	return c.JSON(200,"注册成功！")
}

//参赛
func Join (c echo.Context) error {
	inf := c.Param("userid")
	res := model.Join(inf)
	return c.JSON(200,res)
}


//退赛
func Cancel (c echo.Context)error{
	inf :=c.Param("userid")
	_ = model.Cancel(inf)
	return c.JSON(200,"退赛成功")
}

//用户投票
func Vote (c echo.Context)error{
	inf2:=c.Param("id")//你的id
	inf,_ := strconv.Atoi(c.Param("userid"))//你要投的人的id
	res := model.Vote(inf,inf2)
	if res==1 {
		return c.JSON(200,"投票成功")
	}
	if res == 3{
		return c.JSON(300,"用户未参赛")
	}
	return c.JSON(200,"没票了")
}

//查看排行榜
func Rank (c echo.Context)error{
	res:=model.GetRank()
	return c.JSON(200,res)
}

