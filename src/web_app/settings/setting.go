package settings

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Init() (err error) {
	viper.SetConfigName("config") // 配置文件名（无拓展名）
	viper.SetConfigType("yaml")   //若配置文件的名称中没有拓展名，则需要配置此项
	//viper.SetConfigFile("./src/viper/config.yaml")
	viper.AddConfigPath("./src/web_app/")
	err = viper.ReadInConfig()
	if err != nil {
		fmt.Printf("viper.ReadInConfig() failed with %v\n", err)
		return
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
	return
}
