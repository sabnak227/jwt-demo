package config


type Config struct {
	DBDriver   string
	DBHost     string
	DBPort     string
	DBName     string
	DBUser     string
	DBPassword string
	AutoMigrate bool
	RedisHost   string
	RedisPassword   string
	RedisDB   int
	UserSvcHost string
	ScopeSvcHost string
	AmqpDsn     string
}