package configs

import (
	"log"
	"time"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	// DB
	DBDriver 	string `mapstructure:"DB_DRIVER"`
	DBHost		string `mapstructure:"DB_HOST"`
	DBPort		string `mapstructure:"DB_PORT"`
	DBName		string `mapstructure:"DB_NAME"`
	DBPsw		string `mapstructure:"DB_PSW"`
	DBSslmode	string `mapstructure:"DB_SSLMODE"`
	DBEncoding	string `mapstructure:"ENCODING"`

	// Server
	ServerAddr	string `mapstructure:"SERVER_ADDR"`
	ReadTimeout	time.Duration `mapstructure:"READ_TIMEOUT"`
	WriteTimeout time.Duration `mapstructure:"WRITE_TIMEOUT"`
	MaxHeaderBytes	int `mapstructure:"MAX_HEADER_BYTES"`
}

const (
	configDir = "configs"
	configName = "config"
	configType = "env"
)

func NewConfig() (*Config, error) {
	configEnvPath := configDir+"/"+configName+"."+configType
	err := godotenv.Load(configEnvPath)
	
	if  err != nil {
		log.Println("No .env file found, using system ENV")
	}

	viper.AutomaticEnv()

	viper.AddConfigPath(configDir)
	viper.SetConfigName(configName)
	viper.SetConfigType(configType)

	err = viper.ReadInConfig()
	if err != nil {
		log.Println("No .env file found, using system ENV")
	}

	var cfg Config
	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}