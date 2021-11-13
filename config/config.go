package config

type Config struct {
	Database struct {
		Username string `env:"MYSQL_USERNAME", required`
		Password string `env:"MYSQL_ROOT_PASSWORD"", required`
		Host     string `env:"MYSQL_HOST", default=localhost`
		Port     string `env:"MYSQL_PORT, required"`
		Name     string `env:"MYSQL_DATABASE, required"`
	}
}