package config

type Config struct {
	Port     int `env:"PORT,default=8080"`
	Database struct {
		Username string `env:"DATABASE_USERNAME,required"`
		Password string `env:"DATABASE_PASSWORD,required"`
		Host     string `env:"DATABASE_HOST,default=localhost"`
		Port     string `env:"DATABASE_PORT,required"`
		Name     string `env:"DATABASE_NAME,required"`
	}
}
