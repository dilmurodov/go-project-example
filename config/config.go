package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	PostgresHost     string
	PostgresPort     int
	PostgresUser     string
	PostgresPassword string
	PostgresDatabase string

	PostgresMaxConnections int
	ServerPort             string
	ServerHost             string
}

func Load() Config {

	if err := godotenv.Load(); err != nil {
		fmt.Println(ErrEnvNotFound)
	}

	config := Config{}

	config.PostgresDatabase = cast.ToString(getOrReturnDefaultValue("POSTGRES_DB", "pb_project"))
	config.PostgresPassword = cast.ToString(getOrReturnDefaultValue("POSTGRES_PASSWORD", "admin"))
	config.PostgresPort = cast.ToInt(getOrReturnDefaultValue("POSTGRES_PORT", 5432))
	config.PostgresHost = cast.ToString(getOrReturnDefaultValue("POSTGRES_HOST", "0.0.0.0"))
	config.PostgresUser = cast.ToString(getOrReturnDefaultValue("POSTGRES_USER", "postgres"))
	config.PostgresMaxConnections = cast.ToInt(getOrReturnDefaultValue("POSTGRES_MAX_CONNECTIONS", 10))
	config.ServerPort = cast.ToString(getOrReturnDefaultValue("SERVER_PORT", ":80"))
	config.ServerHost = cast.ToString(getOrReturnDefaultValue("SERVER_HOST", "goservice"))

	return config
}

func getOrReturnDefaultValue(key string, defaultValue interface{}) interface{} {
	val, exists := os.LookupEnv(key)
	if exists {
		return val
	}

	return defaultValue
}
