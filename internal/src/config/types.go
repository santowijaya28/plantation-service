package config

type Config struct {
	DatabaseConfig DatabaseType
	Token          TokenType
}

type DatabaseType struct {
	Host       string
	Port       string
	DbName     string
	DbUserName string
	DbPwd      string
}

type TokenType struct {
	SecretKey string
}
