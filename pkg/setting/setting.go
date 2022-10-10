package setting

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

// Server :
type Server struct {
	RunMode      string        `mapstructure:"run_mode"`
	HTTPPort     int           `mapstructure:"http_port"`
	ReadTimeout  time.Duration `mapstructure:"read_timeout"`
	WriteTimeout time.Duration `mapstructure:"write_timeout"`
}

// Database:
type Database struct {
	Type        string `mapstructure:"type"`
	Host        string `mapstructure:"host"`
	Port        string `mapstructure:"port"`
	User        string `mapstructure:"user"`
	Password    string `mapstructure:"password"`
	Name        string `mapstructure:"name"`
	TablePrefix string `mapstructure:"table_prefix"`
}

type FileConfig struct {
	Server   *Server   `mapstructure:"server"`
	Database *Database `mapstructure:"database"`
}

var FileConfigSetting = &FileConfig{}

func Setup() {
	now := time.Now()
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'config.json': %v", err)
	}
	err = viper.Unmarshal(FileConfigSetting)
	if err != nil {
		log.Fatalf("setting.Setup, fail to Unmarshal 'config.json': %v", err)
	}
	timeSpent := time.Since(now)
	log.Printf("Config setting is ready in %v", timeSpent)

}
