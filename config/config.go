package config

import (
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
)

type PublicConfig struct {
	ENVIRONMENT string
	LOG_FILE    string
	DB_USERNAME string
	DB_PASSWORD string
	DB_HOST     string
	DB_PORT     string
	DB_NAME     string
}

const ENV_DEVELOPMENT = "DEV"
const ENV_PRODUCTION = "PROD"

func Config() PublicConfig {
	var conf PublicConfig
	var env map[string]string
	_, b, _, _ := runtime.Caller(0)
	ProjectRootPath := filepath.Join(filepath.Dir(b), "../")
	env, err := godotenv.Read(ProjectRootPath + "/.env")

	if err != nil {
		panic("Unreadable config")
	} else {
		conf.ENVIRONMENT = env["ENVIRONMENT"]
		conf.LOG_FILE = env["LOG_FILE"]
		conf.DB_USERNAME = env["DB_USERNAME"]
		conf.DB_PASSWORD = env["DB_PASSWORD"]
		conf.DB_HOST = env["DB_HOST"]
		conf.DB_PORT = env["DB_PORT"]
		conf.DB_NAME = env["DB_NAME"]
	}
	return conf
}
