package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

type Config struct {
	DBDriver  string
	DBDSN     string
	JWTSecret string
}

func LoadConfig() Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("未找到 .evn 文件")
	}
	return Config{
		DBDriver:  getEnv("DB_DRIVER", "mysql"),
		DBDSN:     getEnv("DB_DSN", ""),
		JWTSecret: getEnv("JWT_SECRET", ""),
	}
}

func getEnv(key string, defaultValue string) string {
	v := os.Getenv(key)
	if v == "" {
		return defaultValue
	}
	return v
}

func getEnvInt(key string, defaultValue int) int {
	v := os.Getenv(key)
	if v == "" {
		return defaultValue
	}
	i, err := strconv.Atoi(v)
	if err != nil {
		log.Printf("环境变量%s的值配置错误，不是有效的整数", key)
		return defaultValue
	}
	return i
}
