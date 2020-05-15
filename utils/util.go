package utils

import (
	"log"

	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

func GetEnvVar(key string) string {
	viper.SetConfigFile(".env")
	viper.AddConfigPath("../")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Failed to read in config file. Error: %v", err)
	}

	value, ok := viper.Get(key).(string)
	if !ok {
		log.Fatal("Invalid type assertion")
	}
	return value
}

func HashPassword(string rawPassword) (string, error) {
	// Hash raw password
	hashedPwdBytes, err := bcrypt.GenerateFromPassword([]byte(rawPassword), bcrypt.DefaultCost)
	return string(hashedPwdBytes), err
}
