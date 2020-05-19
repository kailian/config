package config

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	vconfig *viper.Viper
}

func New(opt *Options) *Config {
	if opt == nil {
		opt = NewOptions()
	}
	NewConfig := viper.New()
	NewConfig.SetConfigName(opt.FileName)
	NewConfig.AddConfigPath(opt.FilePath)
	p := new(Config)
	p.vconfig = NewConfig
	return p
}

func (c *Config) ReadInConfig() (*Config, error) {
	if err := c.vconfig.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Errorf("Config Not Found: %s \n", err)
		}
		return nil, err
	}
	return c, nil
}

func (c *Config) GetConfig() *viper.Viper {
	return c.vconfig
}

// 监控配置变化
func (c *Config) WatchConfig() {
	c.vconfig.WatchConfig()
	c.vconfig.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
}
