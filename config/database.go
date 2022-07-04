package config

import "os"

type Database interface{}

type PsqlDbConnection struct {
	DbHost     string
	DbPort     string
	DbDatabase string
	DbUsername string
	DbPassword string
}

type MysqlDbConnection struct {
	DbHost     string
	DbPort     string
	DbDatabase string
	DbUsername string
	DbPassword string
}

type DatabaseConfig struct {
	Psql  PsqlDbConnection
	Mysql MysqlDbConnection
}

func DatabaseNew() Database {
	return &DatabaseConfig{
		Psql: PsqlDbConnection{
			DbHost:     os.Getenv("PSQL_DB_HOST"),
			DbPort:     os.Getenv("PSQL_DB_PORT"),
			DbDatabase: os.Getenv("PSQL_DB_DATABASE"),
			DbUsername: os.Getenv("PSQL_DB_USERNAME"),
			DbPassword: os.Getenv("PSQL_DB_PASSWORD"),
		},
		Mysql: MysqlDbConnection{
			DbHost:     os.Getenv("MYSQL_DB_HOST"),
			DbPort:     os.Getenv("MYSQL_DB_PORT"),
			DbDatabase: os.Getenv("MYSQL_DB_DATABASE"),
			DbUsername: os.Getenv("MYSQL_DB_USERNAME"),
			DbPassword: os.Getenv("MYSQL_DB_PASSWORD"),
		},
	}
}
