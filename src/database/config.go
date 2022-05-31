package database

import (
	"os"
)

type PsqlDbConnection struct {
	Host     string
	Port     string
	Database string
	Username string
	Password string
	SslMode  string
}

func DbConfig() PsqlDbConnection {
	dbConfig := PsqlDbConnection{}
	dbConfig.Host = os.Getenv("DB_HOST")
	dbConfig.Port = os.Getenv("DB_PORT")
	dbConfig.Database = os.Getenv("DB_NAME")
	dbConfig.Username = os.Getenv("DB_USER")
	dbConfig.Password = os.Getenv("DB_PASSWORD")
	dbConfig.SslMode = os.Getenv("DB_SSL_MODE")

	return dbConfig
}