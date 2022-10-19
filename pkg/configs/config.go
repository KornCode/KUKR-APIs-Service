package config

import (
	"os"

	"github.com/KornCode/KUKR-APIs-Service/pkg/logs"
	"github.com/joho/godotenv"
)

type Configs struct {
	Server  Server
	RedisDB RedisDB
	MySQLDB MySQLDB
}

type Server struct {
	Host        string
	Port        string
	ReadTimeout string
}

type RedisDB struct {
	Host     string
	Port     string
	Password string
}

type MySQLDB struct {
	Host     string
	Port     string
	DBName   string
	Username string
	Password string
}

func NewConfigs() Configs {
	if err := godotenv.Load(".env"); err != nil {
		logs.Error(err)
	}

	return Configs{
		Server: Server{
			Host:        os.Getenv("SERVER_HOST"),
			Port:        os.Getenv("SERVER_PORT"),
			ReadTimeout: os.Getenv("SERVER_READ_TIMEOUT"),
		},
		RedisDB: RedisDB{
			Host:     os.Getenv("REDIS_HOST"),
			Port:     os.Getenv("REDIS_PORT"),
			Password: os.Getenv("REDIS_PASSWORD"),
		},
		MySQLDB: MySQLDB{
			Host:     os.Getenv("MYSQL_HOST"),
			Port:     os.Getenv("MYSQL_PORT"),
			DBName:   os.Getenv("MYSQL_DBNAME"),
			Username: os.Getenv("MYSQL_USERNAME"),
			Password: os.Getenv("MYSQL_PASSWORD"),
		},
	}
}
