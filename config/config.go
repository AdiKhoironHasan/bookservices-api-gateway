package config

import (
	"os"
	"strconv"
	"strings"
)

type DBConfig struct {
	DBDriver   string
	DBHost     string
	DBPort     string
	DBUser     string
	DBName     string
	DBPassword string
	DBTimeZone string
	DBLog      bool
}

type Config struct {
	AppName     string
	AppPort     int
	AppEnv      string
	AppLang     string
	AppTimeZone string
	GRPCPort    int
	AppKey
	DBConfig
	Dependency
}

type AppKey struct {
	AppSecretKey  string
	AppServiceKey string
}

type Dependency struct {
	BookServices DependBookServices
	UserServices DependUserServices
}

type DependBookServices struct {
	AppPortGRPC int
	AppHostGRPC string
}

type DependUserServices struct {
	AppPortGRPC int
	AppHostGRPC string
}

func New() *Config {
	return &Config{
		AppName:     getEnv("APP_NAME", ""),
		AppPort:     getEnvAsInt("APP_PORT", 8080),
		AppEnv:      getEnv("APP_ENV", "development"),
		AppLang:     getEnv("APP_LANG", "en"),
		AppTimeZone: getEnv("APP_TIMEZONE", ""),
		GRPCPort:    getEnvAsInt("GRPC_PORT", 9000),
		DBConfig: DBConfig{
			DBDriver:   getEnv("DB_DRIVER", "postgres"),
			DBHost:     getEnv("DB_HOSt", "localhost"),
			DBPort:     getEnv("DB_PORT", "5432"),
			DBName:     getEnv("DB_NAME", "db"),
			DBUser:     getEnv("DB_USER", "postgres"),
			DBPassword: getEnv("DB_PASS", ""),
			DBTimeZone: getEnv("APP_TIMEZONE", "Asia/Jakarta"),
			DBLog:      getEnvAsBool("ENABLE_LOGGER", true),
		},
		AppKey: AppKey{
			AppSecretKey:  getEnv("APP_SECRET_KEY", ""),
			AppServiceKey: getEnv("APP_SERVICE_KEY", ""),
		},
		Dependency: Dependency{
			BookServices: DependBookServices{
				AppPortGRPC: getEnvAsInt("BOOK_GRPC_PORT", 9000),
				AppHostGRPC: getEnv("BOOK_GRPC_HOST", "localhost"),
			},
			UserServices: DependUserServices{
				AppPortGRPC: getEnvAsInt("USER_GRPC_PORT", 9000),
				AppHostGRPC: getEnv("USER_GRPC_HOST", "localhost"),
			},
		},
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exist := os.LookupEnv(key); exist {
		return value
	}

	if nextValue := os.Getenv(key); nextValue != "" {
		return nextValue
	}

	return defaultVal
}

func getEnvAsInt(name string, defaultVal int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return defaultVal
}

func getEnvAsBool(name string, defaultVal bool) bool {
	valueStr := getEnv(name, "")
	if value, err := strconv.ParseBool(valueStr); err == nil {
		return value
	}

	return defaultVal
}

func getEnvAsSlice(name string, defaultVal []string, sep string) []string {
	valueStr := getEnv(name, "")
	if valueStr == "" {
		return defaultVal
	}

	value := strings.Split(valueStr, sep)

	return value
}
