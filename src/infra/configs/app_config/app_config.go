package app_config

import (
	"os"
	"strconv"

	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	AppConfig  AppConfig
	HttpConfig HttpConfig
	DBConfigs  DBConfig
}

type AppConfig struct {
	AppName string
	RunMode string
}

type HttpConfig struct {
	HttpPort    string
	HttpTimeout int
}

type DBConfig struct {
	Host                   string
	Username               string
	Password               string
	Name                   string
	Port                   string
	SSLMode                string
	Schema                 string
	MaxOpenConnection      int
	MaxIddleConnection     int
	MaxIddleTimeConnection int64
	MaxLifeTimeConnection  int64
}

var osGetenv = os.Getenv

func InitAppConfig() Config {

	appConfig := AppConfig{
		AppName: osGetenv("APP_NAME"),
		RunMode: osGetenv("RUN_MODE"),
	}

	dbConfig := DBConfig{
		Host:     osGetenv("DB_HOST"),
		Username: osGetenv("DB_USERNAME"),
		Password: osGetenv("DB_PASSWORD"),
		Name:     osGetenv("DB_NAME"),
		Port:     osGetenv("DB_PORT"),
		SSLMode:  osGetenv("DB_SSL_MODE"),
		Schema:   osGetenv("DB_SCHEMA"),
	}
	dBMaxOpenConn, err := strconv.Atoi(osGetenv("MAX_OPEN_CONNECTION"))
	if err == nil {
		dbConfig.MaxOpenConnection = dBMaxOpenConn
	}

	dBMaxIdleConn, err := strconv.Atoi(osGetenv("MAX_IDDLE_CONNECTION"))
	if err == nil {
		dbConfig.MaxIddleConnection = dBMaxIdleConn
	}

	dBMaxIdleTimeConn, err := strconv.Atoi(osGetenv("MAX_IDLE_TIME_CONN_SECONDS"))
	if err == nil {
		dbConfig.MaxIddleTimeConnection = int64(dBMaxIdleTimeConn)
	}
	dBMaxLifeTimeConn, err := strconv.Atoi(osGetenv("MAX_LIFE_TIME_CONN_SECONDS"))
	if err == nil {
		dbConfig.MaxLifeTimeConnection = int64(dBMaxLifeTimeConn)
	}

	portDefault := "8080"
	getPort := osGetenv("HTTP_PORT")
	if getPort != "" {
		portDefault = getPort
	}

	httConfig := HttpConfig{
		HttpPort:    portDefault,
		HttpTimeout: 120, // default in second
	}
	httpTimeout, err := strconv.Atoi(osGetenv("HTTP_TIMEOUT"))
	if err == nil {
		httConfig.HttpTimeout = httpTimeout
	}

	return Config{
		AppConfig:  appConfig,
		HttpConfig: httConfig,
		DBConfigs:  dbConfig,
	}
}
