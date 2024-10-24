package main

import (
	"go.uber.org/zap"
	"net/http"
)

var logger *zap.Logger

func initLogger() {
	logger, _ = zap.NewProduction()
}
func simpleHttpGet(url string) {
	resp, err := http.Get(url)
	if err != nil {
		logger.Error(
			"Error fetching url",
			zap.String("statusCode", resp.Status),
			zap.String("url", url),
		)
	} else {
		logger.Info("Success...",
			zap.String("statusCode", resp.Status),
			zap.String("url", url),
		)
		resp.Body.Close()
	}
}
func main() {
	initLogger()
	defer logger.Sync()
	//simpleHttpGet("https://www.baidu.com")
	simpleHttpGet("https://www.google.com")
}
