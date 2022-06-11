package service

import (
	"AutoSignIn/model"
	"AutoSignIn/utils"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func Sign(c chan interface{}) {
	//构建请求参数
	m := map[string]string{
		"address":     "定位失败",
		"geolocation": "null",
		"type":        "0",
		"recruitId":   "461414",
		"pcid":        "4142",
		"pcmajorid":   "2459402",
		"allowOffset": "2000",
		"offset":      "NaN",
		"offduty":     "0",
	}
	signFormValue := utils.CreateUrlValue(m)
	//创建cookie
	var cookieString string
	cookies := <-c
	if v, ok := cookies.([]*http.Cookie); ok {
		for _, cookie := range v {
			cookieString += fmt.Sprintf("%v=%v;", cookie.Name, cookie.Value)
		}
	}
	//构建请求体
	request, err := utils.CreateRequest("POST",
		"http://fzjyxy.dgsx.chaoxing.com/mobile/clockin/addclockin2",
		signFormValue)
	if err != nil {
		fmt.Println(err.Error())
	}
	request.Header.Set("cookie", cookieString)
	resp := utils.DoPost(request)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err.Error())
		}
	}(resp.Body)
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
	}
	var signInModel = model.SignInModel{}
	err = json.Unmarshal(respBody, &signInModel)
	if err != nil {
		fmt.Println(err.Error())
	}
	if signInModel.Success {
		c <- true
	}
}
