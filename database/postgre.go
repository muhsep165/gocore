package database

import (
	"database/sql"
	"fmt"
	"go-core/config"
	"go-core/helper"
	"time"
)

func GetPostgreSqlConnection() *sql.DB {
	return dbPostgreSql
}

func PostgreSqlConnection() *sql.DB {
	cfg := config.Config()
	postgresqlDbInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		cfg.DB_HOST, helper.ToInt(cfg.DB_PORT), cfg.DB_USERNAME, cfg.DB_PASSWORD, cfg.DB_NAME)

	dbPostgreSql, err := sql.Open("postgres", postgresqlDbInfo)
	if err != nil {
		panic(err)
	}

	dbPostgreSql.SetMaxIdleConns(10)
	dbPostgreSql.SetMaxOpenConns(100)
	dbPostgreSql.SetConnMaxIdleTime(5 * time.Minute)
	dbPostgreSql.SetConnMaxLifetime(60 * time.Minute)
	return dbPostgreSql
}
