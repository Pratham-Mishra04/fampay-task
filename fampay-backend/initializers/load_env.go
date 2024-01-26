package initializers

import (
	"fmt"
	"log"
	"reflect"

	"github.com/spf13/viper"
)

type Environment string

const (
	DevelopmentEnv Environment = "development"
	ProductionEnv  Environment = "production"
)

type Config struct {
	PORT            string      `mapstructure:"PORT"`
	ENV             Environment `mapstructure:"ENV"`
	DB_HOST         string      `mapstructure:"DB_HOST"`
	DB_PORT         string      `mapstructure:"DB_PORT"`
	DB_NAME         string      `mapstructure:"DB_NAME"`
	DB_USER         string      `mapstructure:"DB_USER"`
	DB_PASSWORD     string      `mapstructure:"DB_PASSWORD"`
	REDIS_HOST      string      `mapstructure:"REDIS_HOST"`
	REDIS_PORT      string      `mapstructure:"REDIS_PORT"`
	REDIS_PASSWORD  string      `mapstructure:"REDIS_PASSWORD"`
	BACKEND_URL     string      `mapstructure:"BACKEND_URL"`
	FRONTEND_URL    string      `mapstructure:"FRONTEND_URL"`
	YOUTUBE_API_KEY string      `mapstructure:"YOUTUBE_API_KEY"`
}

var CONFIG Config

func LoadEnv() {
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	err = viper.Unmarshal(&CONFIG)
	if err != nil {
		log.Fatal(err)
	}

	requiredKeys := getRequiredKeys(CONFIG)
	missingKeys := checkMissingKeys(requiredKeys, CONFIG)

	if len(missingKeys) > 0 {
		err := fmt.Errorf("following environment variables not found: %v", missingKeys)
		log.Fatal(err)
	}

	if CONFIG.ENV != DevelopmentEnv && CONFIG.ENV != ProductionEnv {
		err := fmt.Errorf("invalid ENV value: %s", CONFIG.ENV)
		log.Fatal(err)
	}
}

func getRequiredKeys(config Config) []string {
	requiredKeys := []string{}
	configType := reflect.TypeOf(config)

	for i := 0; i < configType.NumField(); i++ {
		field := configType.Field(i)
		tag := field.Tag.Get("mapstructure")
		if tag != "" {
			requiredKeys = append(requiredKeys, tag)
		}
	}

	return requiredKeys
}

func checkMissingKeys(requiredKeys []string, config Config) []string {
	missingKeys := []string{}

	configValue := reflect.ValueOf(config)
	for _, key := range requiredKeys {
		value := configValue.FieldByName(key).String()
		if value == "" {
			missingKeys = append(missingKeys, key)
		}
	}

	return missingKeys
}
