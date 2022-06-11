package main

import (
	"AutoSignIn/router"
	"AutoSignIn/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
)

var (
	AppVersion = "BETA"
)

func init() {
	//初始化日志框架
	cfg := zap.Config{
		Level:            zap.NewAtomicLevel(),
		Encoding:         "json",
		OutputPaths:      []string{"stdout", "./logs/log.log"},
		ErrorOutputPaths: []string{"stderr", "./logs/error.log"},
		Development:      true,
		InitialFields:    map[string]interface{}{"version": AppVersion},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:     "msg",
			LevelKey:       "level",
			TimeKey:        "time",
			NameKey:        "name",
			CallerKey:      "caller",
			EncodeName:     zapcore.FullNameEncoder,
			EncodeLevel:    zapcore.CapitalLevelEncoder,
			EncodeTime:     zapcore.RFC3339TimeEncoder,
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
	}

	l, err := cfg.Build()
	if err != nil {
		log.Fatal(err)
	}

	utils.Logger = l.Sugar()
}
func main() {
	engine := gin.Default()
	engine.GET("/ChaoXingAutoSignIn", router.SignIn)
	engine.GET("/CheckAutoSignIn", router.YunDong)
	err := engine.Run(":8099")
	if err != nil {
		utils.Logger.Error(err)
		return
	}
}
