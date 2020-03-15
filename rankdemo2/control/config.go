package control

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/robfig/cron"
	"rankdemo2/model"
)

//中间件
func ServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "Echo/3.0")
		tokenstring := c.FormValue("token")
		claims := model.UserToken{}
		token,err := jwt.ParseWithClaims(tokenstring,&claims, func(token *jwt.Token) (i interface{}, err error) {
			return []byte("123"),nil
		})
		if err == nil && token.Valid{
			return next(c)
		}else {
			return c.JSON(300,"验证失败！请先登录！")
		}

	}
}

//定时任务 更新票数
func Update (){
	c := cron.New()
	c.AddFunc("0 0 0 * * *",model.UpdateVotes)
	c.Start()
}