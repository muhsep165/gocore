package database

import (
	"database/sql"
	"go-core/config"
	"time"
)

func GetMysqlConnection() *sql.DB {
	return dbMysql
}

func MySqlConnection() *sql.DB {
	cfg := config.Config()
	dbMysql, err := sql.Open("mysql", cfg.DB_USERNAME+":"+cfg.DB_PASSWORD+"@tcp("+cfg.DB_HOST+":"+cfg.DB_PORT+")/"+cfg.DB_NAME+"?parseTime=true&loc=Local")
	if err != nil {
		panic(err)
	}

	dbMysql.SetMaxIdleConns(10)
	dbMysql.SetMaxOpenConns(100)
	dbMysql.SetConnMaxIdleTime(5 * time.Minute)
	dbMysql.SetConnMaxLifetime(60 * time.Minute)
	return dbMysql
}
