package service

import (
	"AutoSignIn/model"
	"AutoSignIn/utils"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

func YunDongLogin() *http.Response {
	m := map[string]string{
		"email":  "email",
		"passwd": "passwd",
	}
	value := utils.CreateUrlValue(m)
	//构建请求体
	request, err := utils.CreateRequest("POST",
		"https://ccave.org/auth/login",
		value)
	if err != nil {
		utils.Logger.Error(err)
	}
	resp := utils.DoPost(request)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			utils.Logger.Error(err)
		}
	}(resp.Body)
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		utils.Logger.Error(err)
	}
	var yunDongModel = model.YunDongModel{}
	err = json.Unmarshal(respBody, &yunDongModel)
	if err != nil {
		utils.Logger.Error(err)
	}
	if yunDongModel.Ret == 1 {
		return resp
	} else {
		return nil
	}
}

func YunDongSigIn(cookies string) bool {
	//构建请求体
	request, err := utils.CreateRequest("POST",
		"https://ccave.org/user/checkin",
		"")
	if err != nil {
		utils.Logger.Error(err)
	}
	request.Header.Set("cookie", cookies)
	resp := utils.DoPost(request)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			utils.Logger.Error(err)
		}
	}(resp.Body)
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		utils.Logger.Error(err)
	}
	var yunDongModel = model.YunDongModel{}
	err = json.Unmarshal(respBody, &yunDongModel)
	if err != nil {
		utils.Logger.Error(err)
	}
	if yunDongModel.Ret == 1 {
		return true
	} else {
		return false
	}
}
