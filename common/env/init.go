package env

import (
	"log"

	"github.com/spf13/viper"
)

type Env string

// kek buat type string literal di typescript (buat 2 variabel dengan namo EnvProd dan EnvDev)
const (
	EnvProd Env = "PRODUCTION"
	EnvDev  Env = "DEVELOPMENT"
)

type Config struct {
	Env         Env    `mapstructure:"ENV"`
	AppHost     string `mapstructure:"APP_HOST"`
	AppPort     string `mapstructure:"APP_PORT"`
	AppStaticPath   string `mapstructure:"APP_STATIC_PATH"`
	IsStaticCloud bool `mapstructure:"IS_STATIC_CLOUD"`
	IsMigrateInit bool `mapstructure:"IS_MIGRATE_INIT"`
	PostgresUrl string `mapstructure:"POSTGRES_CONNECTION_URL"`
	PostgresMigratePath string `mapstructure:"POSTGRES_MIGRATE_PATH"`
	AdminPhone string `mapstructure:"ADMIN_PHONE"`
	AdminEmail string `mapstructure:"ADMIN_EMAIL"`
	AdminPassword string `mapstructure:"ADMIN_PASSWORD"`

	AllowedOrigins []string `mapstructure:"ALLOWED_ORIGINS"`
}

func LoadConfig(envPath string) Config {
	var config Config

	viper.SetConfigFile(envPath)

	// Timpa nilai dari file kalau ada Env Var di OS
	viper.AutomaticEnv()

	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Error occured while .env reading: %v", err)
	}

	viper.SetDefault("APP_HOST", "0.0.0.0")
	viper.SetDefault("APP_PORT", "8000")
	viper.SetDefault("APP_STATIC_PATH", "./static")
	viper.SetDefault("IS_STATIC_CLOUD", false)
	viper.SetDefault("IS_MIGRATE_INIT", true)
	viper.SetDefault("ENV", EnvDev)
	viper.SetDefault("POSTGRES_MIGRATE_PATH", "./config/postgres/migration")
	viper.SetDefault("ADMIN_PHONE", "08121234567890")
	viper.SetDefault("ADMIN_EMAIL", "fa@fa.fa")
	viper.SetDefault("ADMIN_PASSWORD", "123456")
	viper.SetDefault("ALLOWED_ORIGINS", []string{"http://localhost:4321", "http://127.0.0.1:4321"})

	viper.Unmarshal(&config)

	return config
}
