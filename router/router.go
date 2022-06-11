package router

import (
	"AutoSignIn/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SignIn(c *gin.Context) {
	var isSuccessSignIn = make(chan bool)
	//获取参数
	username := c.Query("username")
	password := c.Query("password")
	go service.Login(username, password, isSuccessSignIn)
	ok := <-isSuccessSignIn
	if ok {
		c.JSON(http.StatusOK, gin.H{
			"code": "200",
			"msg":  "打卡成功",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": "200",
			"msg":  "打卡失败,请检查用户名或密码",
		})
	}
	return
}

func YunDong(c *gin.Context) {
	resp := service.YunDongLogin()
	//创建cookie
	var cookieString string
	for _, cookie := range resp.Cookies() {
		cookieString += fmt.Sprintf("%v=%v;", cookie.Name, cookie.Value)
	}
	ok := service.YunDongSigIn(cookieString)
	if ok {
		c.JSON(http.StatusOK, gin.H{
			"code": "200",
			"msg":  "签到成功",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": "500",
			"msg":  "签到失败",
		})
	}
	return
}
