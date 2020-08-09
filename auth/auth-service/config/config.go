package config


type Config struct {
	DBDriver   string
	DBHost     string
	DBPort     string
	DBName     string
	DBUser     string
	DBPassword string
	AutoMigrate bool
	UserSvcHost string
	ScopeSvcHost string
}