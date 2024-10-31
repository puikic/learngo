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
	// 方式1：直接指定配置文件路径
	viper.SetConfigFile("./src/viper/config.yaml") // 相对路径：相对可执行文件的相对路径，本项目即web_app
	//viper.SetConfigType("/home/cpq/go/src/learngo/src/viper/config.yaml") // 绝对路径
	// 方式2：指定配置文件名和配置文件的位置，viper自行查找可用的配置文件
	//viper.SetConfigName("config") // 配置文件名（不需要带后缀）
	//viper.AddConfigPath("./src/viper")

	//viper.SetConfigType("yaml")   // 配合远程配置中心使用的
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
