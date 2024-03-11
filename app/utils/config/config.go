package config

import (
	"github.com/spf13/viper"
	"go-backend/app/global/my_errors"
	"go-backend/app/global/variable"
	"log"
)

// CreateYamlFactory 创建一个yaml配置文件工厂
// 参数设置为可变参数的文件名，这样参数就可以不需要传递，如果传递了多个，我们只取第一个参数作为配置文件名
func CreateYamlFactory(fileName ...string) map[string]any {
	yamlConfig := viper.New()
	// 配置文件所在目录
	yamlConfig.AddConfigPath(variable.BasePath + "/config")
	// 需要读取的文件名,默认为：config
	if len(fileName) == 0 {
		yamlConfig.SetConfigName("config")
	} else {
		yamlConfig.SetConfigName(fileName[0])
	}
	//设置配置文件类型(后缀)为 yml
	yamlConfig.SetConfigType("yml")

	if err := yamlConfig.ReadInConfig(); err != nil {
		log.Fatal(my_errors.ErrorsConfigInitFail + err.Error())
	}

	return yamlConfig.AllSettings()
}
