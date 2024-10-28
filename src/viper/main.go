package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	Port        int    `mapstructure:"port"`
	Version     string `mapstructure:"version"`
	MySQLConfig `mapstructure:"mysql"`
}
type MySQLConfig struct {
	Host   string `mapstructure:"host"`
	Port   int    `mapstructure:"port"`
	Dbname string `mapstructure:"dbname"`
}

func main() {
	// 设置默认值
	//viper.SetDefault("fileDir", "./src/viper")
	// 读取配置文件
	viper.SetConfigName("config") // 配置文件名（无拓展名）
	viper.SetConfigType("yaml")   //若配置文件的名称中没有拓展名，则需要配置此项
	//viper.SetConfigFile("./src/viper/config.yaml")
	viper.AddConfigPath("./src/viper")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	viper.WatchConfig() // 运行时实时监控配置文件的变化
	viper.OnConfigChange(func(e fsnotify.Event) {
		// 配置文件发生变更之后会调用的回调函数
		fmt.Println("Config file changed:", e.Name)
	})
	var c Config
	if err := viper.Unmarshal(&c); err != nil {
		panic(err)
	}
	fmt.Println(c)
	fmt.Printf("%v\n", c)
	//r := gin.Default()
	//r.GET("/version", func(c *gin.Context) {
	//	c.String(http.StatusOK, viper.GetString("version"))
	//})
	//r.Run()
}
