package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"net/http"
	"os"
)

var logger *zap.Logger
var sugar *zap.SugaredLogger

//	func InitLogger() {
//		logger, _ = zap.NewDevelopment()
//		//logger, _ = zap.NewProduction()
//		sugar = logger.Sugar()
//	}
func initLogger() {
	writeSyncer := getLogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)
	logger = zap.New(core)
	sugar = logger.Sugar()
}
func getEncoder() zapcore.Encoder {
	return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
}

func getLogWriter() zapcore.WriteSyncer {
	file, _ := os.Create("./test.log")
	return zapcore.AddSync(file)
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
	//simpleHttpGet("https://www.google.com")
	//simpleHttpGet("https://www.sogo.com")
}
