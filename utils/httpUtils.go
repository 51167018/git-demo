package utils

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func DoPost(request *http.Request) *http.Response {
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
	}
	return resp
}

func CreateRequest(method string, url string, body string) (*http.Request, error) {
	request, err := http.NewRequest(
		method,
		url,
		strings.NewReader(body),
	)

	//设置请求头默认参数
	request.Header.Add("accept-language", "zh-CN,zh;q=0.8")
	request.Header.Add("user-agent", "Mozilla/5.0 (Linux; U; Android 4.4.2; en-us; Android SDK built for x86 Build/KK) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30")
	request.Header.Add("cache-control", "no-cache")
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	return request, err
}

func CreateUrlValue(parameterMap map[string]string) string {
	result := url.Values{}
	for k, v := range parameterMap {
		result.Set(k, v)
	}
	return result.Encode()
}
