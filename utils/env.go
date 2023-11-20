package utils

import (
	"fmt"
	"github.com/joho/godotenv"
)

func LoadEnv() error {
	envFilePath := "C:\\Users\\ghazal\\GolandProjects\\startProject\\.env"
	err := godotenv.Load(envFilePath)
	if err != nil {
		fmt.Println("some error happend")
		return err
	}
	return nil
}
