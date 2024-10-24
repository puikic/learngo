package main

import (
	"go.uber.org/zap"
	"net/http"
)

var logger *zap.Logger
var sugar *zap.SugaredLogger

func initLogger() {
	logger, _ = zap.NewDevelopment()
	//logger, _ = zap.NewProduction()
	sugar = logger.Sugar()
}

func simpleHttpGet(url string) {
	resp, err := http.Get(url)
	if err != nil {
		sugar.Error(
			"Error fetching url",
			zap.String("url", url),
			zap.Error(err),
		)
	} else {
		sugar.Info("Success...",
			zap.String("statusCode", resp.Status),
			zap.String("url", url),
		)
		resp.Body.Close()
	}
}

func main() {
	initLogger()
	defer logger.Sync()
	simpleHttpGet("https://www.baidu.com")
	simpleHttpGet("https://www.google.com")
	//simpleHttpGet("https://www.sogo.com")
}
