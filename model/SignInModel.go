package model

type SignInModel struct {
	Success bool   `json:"success"`
	Msg     string `json:"msg"`
	Info    struct {
		ID     int `json:"id"`
		Status int `json:"status"`
	} `json:"info"`
}
