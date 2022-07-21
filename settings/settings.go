package settings

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Init() (err error) {
	viper.SetConfigFile("./conf/config.yaml") // 指定配置文件路径
	err = viper.ReadInConfig()                // 读取配置信息
	if err != nil {                           // 读取配置信息失败
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	// 监控配置文件变化(热更新)
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件修改了")
	})
	return
}
