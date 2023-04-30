package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DataBase struct {
		Dialect  string // Database type
		Host     string // Name of the host
		Port     int    // Port to connect
		Username string // Username value
		Password string // Password value
		Name     string // Database name
		Charset  string // Coding system
	}

	Server struct {
		Host string // Server address
		Port int    // Port address
	}

	configName string
}

func NewConfig(cfgName string) *Config {
	cfg := &Config{configName: cfgName}
	cfg.SetDefaultConfig()
	if cfg.IsOpened() {
		cfg.ExtractConfigData()
	}
	return cfg
}

func (cfg *Config) ExtractConfigData() {
	// Database info
	cfg.DataBase.Dialect = viper.GetString("database.dialect")
	cfg.DataBase.Host = viper.GetString("database.host")
	cfg.DataBase.Port = viper.GetInt("database.port")
	cfg.DataBase.Username = viper.GetString("database.username")
	cfg.DataBase.Password = viper.GetString("database.password")
	cfg.DataBase.Name = viper.GetString("database.name")
	cfg.DataBase.Charset = viper.GetString("database.charset")

	// Server info
	cfg.Server.Host = viper.GetString("server.host")
	cfg.Server.Port = viper.GetInt("server.port")

}

func (cfg *Config) SetDefaultConfig() {
	viper.SetConfigName(cfg.configName)
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./data/config/")
}

func (cfg *Config) IsOpened() bool {
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Error: Can't open config file:", err)
	}
	return true
}
