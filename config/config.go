package config

import (
	"github.com/jinzhu/configor"
	"log"
)

type Log struct {
	Driver    string `env:"driver"`
	InfoPath  string `env:"info_path"`
	ErrPath   string `env:"err_path"`
	HttpDebug bool   `env:"http_debug"`
}

type BaiduTranslateConfig struct {
	Env       string
	AppID     string `required:"true" env:"app_id"`
	AppSecret string `required:"true" env:"app_secret"`
	BaseUri   string `env:"base_uri"`
	Log       `env:"log"`
}

// LoadConfig 从配置文件中加载应用配置
func LoadConfig() (*BaiduTranslateConfig, error) {
	// 从配置文件中加载配置，你可以使用 Viper 等配置库
	// 返回加载后的配置结构体

	return nil, nil
}

func configFiles() []string {
	return []string{"config.yml"}
}

func Get() *BaiduTranslateConfig {
	conf := new(BaiduTranslateConfig)
	err := configor.New(&configor.Config{}).Load(conf, configFiles()...)
	if err != nil {
		log.Printf("%#v", conf)
		panic(err)
	}
	return conf
}
