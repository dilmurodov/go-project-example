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
}

func Load() Config {

	if err := godotenv.Load(); err != nil {
		fmt.Println(ErrEnvNotFound)
	}

	config := Config{}

	config.PostgresDatabase = cast.ToString(getOrReturnDefaultValue("POSTGRESS_DATABASE", "project_db"))
	config.PostgresPassword = cast.ToString(getOrReturnDefaultValue("POSTGRESS_PASSWORD", "admin"))
	config.PostgresPort = cast.ToInt(getOrReturnDefaultValue("POSTGRESS_PORT", 5432))
	config.PostgresHost = cast.ToString(getOrReturnDefaultValue("POSTGRESS_HOST", "0.0.0.0"))
	config.PostgresUser = cast.ToString(getOrReturnDefaultValue("POSTGRESS_USER", "postgres"))
	config.PostgresMaxConnections = cast.ToInt(getOrReturnDefaultValue("POSTGRESS_MAX_CONNECTIONS", 10))
	config.ServerPort = cast.ToString(getOrReturnDefaultValue("SERVER_PORT", ":5000"))

	return config
}

func getOrReturnDefaultValue(key string, defaultValue interface{}) interface{} {
	val, exists := os.LookupEnv(key)
	if exists {
		return val
	}

	return defaultValue
}
