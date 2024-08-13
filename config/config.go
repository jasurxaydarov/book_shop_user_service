package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

const (
	// DebugMode indicates service mode is debug.
	DebugMode string = "debug"
	// TestMode indicates service mode is test.
	TestMode string = "test"
	// ReleaseMode indicates service mode is release.
	ReleaseMode string = "release"
)

type AppConfig struct {
	ProjectName string
	Version     string
	ServiceName string
	Host        string
	RpcPort     string
	Environment string
}

type PgSQLConfig struct {
	Username string
	Password string
	Host     string
	Port     int
	Database string
}

type Services struct {
}

type Config struct {
	AppConfig   AppConfig
	PgSQLConfig PgSQLConfig
}

func Load() Config {

	err := godotenv.Load("./.env")
	if err == nil {
		log.Println("configs loaded from .env")
	}

	var (
		appConfig   AppConfig
		pgSQLConfig PgSQLConfig
	)

	//app
	appConfig.ProjectName = cast.ToString(getDefaultValue("PROJECT_NAME", "book_shop"))
	appConfig.Version = cast.ToString(getDefaultValue("VERSION", "1.0.0"))

	// db config: "postgres://postgres:%21%28%28%281001%23%21gO@54.93.101.96:5432/product_service"
	pgSQLConfig.Username = cast.ToString(getDefaultValue("DB_USER", "jasur"))
	pgSQLConfig.Password = cast.ToString(getDefaultValue("DB_USER_PASSWORD", "1001"))
	pgSQLConfig.Host = cast.ToString(getDefaultValue("DB_HOST", "localhost"))
	pgSQLConfig.Port = cast.ToInt(getDefaultValue("DB_PORT", 5432))
	pgSQLConfig.Database = cast.ToString(getDefaultValue("DB_NAME", "book_shop"))

	//service
	appConfig.ServiceName = cast.ToString(getDefaultValue("SERVICE_NAME", "grpc_todo"))
	appConfig.RpcPort = cast.ToString(getDefaultValue("SERVICE_HOST", "54.93.101.96"))
	appConfig.RpcPort = cast.ToString(getDefaultValue("RPC_PORT", ":8081"))
	appConfig.Environment = "debug"

	return Config{
		AppConfig:   appConfig,
		PgSQLConfig: pgSQLConfig,
	}

}

func GetPgURL() string {
	username := cast.ToString(getDefaultValue("DB_USER", "jasur"))
	password := cast.ToString(getDefaultValue("DB_USER_PASSWORD", "1001"))
	host := cast.ToString(getDefaultValue("DB_HOST", "0.0.0.0"))
	port := cast.ToInt(getDefaultValue("DB_PORT", 5432))
	database := cast.ToString(getDefaultValue("DB_NAME", "book_shop"))

	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s", username, password, host, port, database)
}

func getDefaultValue(key string, defaultValue interface{}) interface{} {

	val, exists := os.LookupEnv(key)
	if exists {
		return val
	}

	return defaultValue
}
