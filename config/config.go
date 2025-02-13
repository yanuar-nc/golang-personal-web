package config

type Config struct {
	MySQL MySQLConfig
}

type MySQLConfig struct {
	Host     string
	DBName   string
	Port     int
	Username string
	Password string
}
