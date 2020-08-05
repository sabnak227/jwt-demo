package config

import "os"

type Configurations struct {
	Port     string
	Database DatabaseConfigurations
}

type DatabaseConfigurations struct {
	DBDriver   string
	DBHost     string
	DBPort     string
	DBName     string
	DBUser     string
	DBPassword string
}

func NewDatabaseConfigurations(dbDriver string, dbHost string, dbPort string, dbName string, dbUser string, dbPassword string) *DatabaseConfigurations {
	return &DatabaseConfigurations{
		DBDriver:   dbDriver,
		DBHost:     dbHost,
		DBPort:     dbPort,
		DBName:     dbName,
		DBUser:     dbUser,
		DBPassword: dbPassword,
	}
}

func NewConfigurations(port string, db *DatabaseConfigurations) *Configurations {
	return &Configurations{
		Port:     port,
		Database: *db,
	}
}

func GetConfigFromEnv(key string, defaultVal string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return defaultVal
	} else {
		return val
	}
}
