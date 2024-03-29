package initializers

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"

	"github.com/lib/pq"
	"github.com/spf13/viper"
)

type Environment string

const (
	DevelopmentEnv Environment = "development"
	ProductionEnv  Environment = "production"
)

type Config struct {
	PORT                      string         `mapstructure:"PORT"`
	ENV                       Environment    `mapstructure:"ENV"`
	DB_HOST                   string         `mapstructure:"DB_HOST"`
	DB_PORT                   string         `mapstructure:"DB_PORT"`
	DB_NAME                   string         `mapstructure:"DB_NAME"`
	DB_USER                   string         `mapstructure:"DB_USER"`
	DB_PASSWORD               string         `mapstructure:"DB_PASSWORD"`
	REDIS_HOST                string         `mapstructure:"REDIS_HOST"`
	REDIS_PORT                string         `mapstructure:"REDIS_PORT"`
	REDIS_PASSWORD            string         `mapstructure:"REDIS_PASSWORD"`
	BACKEND_URL               string         `mapstructure:"BACKEND_URL"`
	FRONTEND_URL              string         `mapstructure:"FRONTEND_URL"`
	YOUTUBE_API_KEYS_LOCATION string         `mapstructure:"YOUTUBE_API_KEYS_LOCATION"`
	YOUTUBE_API_KEYS          pq.StringArray `mapstructure:"YOUTUBE_API_KEYS"`
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

	CONFIG.YOUTUBE_API_KEYS, err = loadAPIKeysFromFile(CONFIG.YOUTUBE_API_KEYS_LOCATION)

	if err != nil {
		err := fmt.Errorf("environment variable %s not found: %s", "YOUTUBE_API_KEYS", err.Error())
		log.Fatal(err)
	} else if len(CONFIG.YOUTUBE_API_KEYS) == 0 {
		err := fmt.Errorf("environment variable %s not found: file %s is empty", "YOUTUBE_API_KEYS", CONFIG.YOUTUBE_API_KEYS_LOCATION)
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

func loadAPIKeysFromFile(filename string) (pq.StringArray, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("file %s not found", filename)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var apiKeys pq.StringArray
	for scanner.Scan() {
		apiKeys = append(apiKeys, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return apiKeys, nil
}
