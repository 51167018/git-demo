package service

import (
	"AutoSignIn/model"
	"AutoSignIn/utils"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
)

func Login(username string, password string, success chan bool) {
	//设置请求参数
	m := map[string]string{
		"uname":    username,
		"password": password,
	}
	postFormValue := utils.CreateUrlValue(m)
	//构建请求体
	request, err := utils.CreateRequest("POST",
		"http://passport2.chaoxing.com/fanyalogin",
		postFormValue,
	)
	if err != nil {
		fmt.Println(err.Error())
	}
	resp := utils.DoPost(request)
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
	}
	var loginModel = model.LoginModel{}
	fmt.Println(string(respBody))
	err = json.Unmarshal(respBody, &loginModel)
	if err != nil {
		fmt.Println(err.Error())
	}
	c := make(chan interface{})
	if loginModel.Status == true {
		go Sign(c)
		c <- resp.Cookies()
		exit := <-c
		if v, ok := exit.(bool); ok && v {
			success <- v
		}
	} else {
		success <- false
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println("文件关闭失败")
			return
		}
	}(resp.Body)
}
