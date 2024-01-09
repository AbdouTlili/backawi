package config

import "os"

func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}

func LoadConfig() *Config {
	config := &Config{
		AWSRegion:          getEnv("AWS_REGION", "eu-west-3"),
		AWSAccessKeyID:     getEnv("AWS_ACCESS_KEY_ID", ""),
		AWSSecretAccessKey: getEnv("AWS_SECRET_ACCESS_KEY", ""),
		DynamoDBTableName:  getEnv("DYNAMODB_TABLE_NAME", "my-default-table"),
	}

	return config
}
