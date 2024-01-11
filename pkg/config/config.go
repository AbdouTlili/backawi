package config

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Warnf("Failed to Lookup the Environment variable %v, loading default value", key)
		return defaultValue
	}
	return value
}

func LoadConfig() *Config {
	config := &Config{
		AWSRegion:          getEnv("AWS_REGION", "eu-west-3"),
		AWSAccessKeyID:     getEnv("AWS_ACCESS_KEY_ID", ""),
		AWSSecretAccessKey: getEnv("AWS_SECRET_ACCESS_KEY", ""),
		DynamoDBTableName:  getEnv("DYNAMODB_TABLE_NAME", ""),
	}

	return config
}
