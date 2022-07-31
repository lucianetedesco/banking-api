package settings

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"time"
)

type environment struct {
	Database struct {
		DBHost      string
		DBPort      int
		DBUser      string
		DBPPassword string
		DBName      string
	}
	Auth struct {
		SecretKey string
		ExpiresAt time.Duration
	}
}

var Environment environment

func setDefaults() {
	viper.SetDefault("Database.DBHost", "db")
	viper.SetDefault("Database.DBPort", 5432)
	viper.SetDefault("Database.DBUser", "postgres")
	viper.SetDefault("Database.DBPPassword", "postgres")
	viper.SetDefault("Database.DBName", "banking-api")

	viper.SetDefault("Auth.SecretKey", "511829610b5f42f0ca27a9e22f00abb2")
	viper.SetDefault("Auth.ExpiresAt", 1)
}

func InitConfig() {
	setDefaults()

	viper.BindEnv("Database.DBHost", "DATABASE_HOST")
	viper.BindEnv("Database.DBPort", "DATABASE_PORT")
	viper.BindEnv("Database.DBUser", "DATABASE_USER")
	viper.BindEnv("Database.DBPPassword", "DATABASE_PASSWORD")
	viper.BindEnv("Database.DBName", "DATABASE_NAME")

	viper.BindEnv("Auth.SecretKey", "AUTH_SECRET_KEY")
	viper.BindEnv("Auth.ExpiresAt", "AUTH_TOKEN_EXPIRES_AT")

	viper.AutomaticEnv()

	if err := viper.Unmarshal(&Environment); err != nil {
		log.Error("Error unmarshal: ", err)
	}
}
