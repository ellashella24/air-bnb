package configs

import (
	"os"
	"sync"

	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
)

type AppConfig struct {
	Port     int
	Driver   string
	Name     string
	Address  string
	DB_Port  int
	Username string
	Password string
}

// var SecretKey = "secret123"
var lock = &sync.Mutex{}
var appConfig *AppConfig

func GetConfig() *AppConfig {
	lock.Lock()
	defer lock.Unlock()

	if appConfig == nil {
		appConfig = initConfig()
	}

	return appConfig
}

func initConfig() *AppConfig {

	err := godotenv.Load(".env")
	if err != nil {
		log.Infof("can't read file env: %s", err)
	}

	var defaultConfig AppConfig
	defaultConfig.Port = 8000
	defaultConfig.Driver = getEnv("DRIVER", "mysql")
	defaultConfig.Name = getEnv("DB_NAME", "airbnb")
	defaultConfig.Address = getEnv("ADDRESS", "localhost")
	defaultConfig.DB_Port = 3306
	defaultConfig.Username = getEnv("DB_USERNAME", "root")
	defaultConfig.Password = getEnv("DB_PASSWORD", "")

	return &defaultConfig
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// Config yaml
// type AppConfig struct {
// 	Port     int `yaml:"port"`
// 	Database struct {
// 		Driver   string `yaml:"driver"`
// 		Name     string `yaml:"name"`
// 		Address  string `yaml:"address"`
// 		Port     int    `yaml:"port"`
// 		Username string `yaml:"username"`
// 		Password string `yaml:"password"`
// 	}
// }

// var SecretKey = "secret123"
// var lock = &sync.Mutex{}
// var appConfig *AppConfig

// func GetConfig() *AppConfig {
// 	lock.Lock()
// 	defer lock.Unlock()

// 	if appConfig == nil {
// 		appConfig = initConfig()
// 	}

// 	return appConfig
// }

// func initConfig() *AppConfig {
// 	var defaultConfig AppConfig
// 	defaultConfig.Port = 8000
// 	defaultConfig.Database.Driver = "mysql"
// 	defaultConfig.Database.Name = "airbnb_test"
// 	defaultConfig.Database.Address = "localhost"
// 	defaultConfig.Database.Port = 3306
// 	defaultConfig.Database.Username = "root"
// 	defaultConfig.Database.Password = ""

// 	viper.SetConfigType("yaml")
// 	viper.SetConfigName("config")
// 	viper.AddConfigPath("./configs/")
// 	if err := viper.ReadInConfig(); err != nil {
// 		log.Info("failed to open file")
// 		return &defaultConfig
// 	}

// 	var finalConfig AppConfig
// 	err := viper.Unmarshal(&finalConfig)
// 	if err != nil {
// 		log.Info("failed to extract external config, use default value")
// 		return &defaultConfig
// 	}
// 	return &finalConfig
// }
