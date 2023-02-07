package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

// Database settings struct
type Database struct {
	Envion     string
	DbName     string
	DbPort     int
	DbUser     string
	DbPassword string
	DbHost     string
}

type Settings struct {
	Debug   bool
	Environ int
	APIPort uint
	DB      Database
}

var environment int

var (
	// No settings error when .env file is not found
	NoSettingsError = errors.New("No settings found")

	// When the .env file is not found
	NoDotEnvError = errors.New("The .env file in local is required")
)

const (
	Debug = iota
	Prod
	Test
)

// define config type
func ConfigType() string {
	if environment == Debug {
		return "env"
	}

	return "config"
}

func (env *Settings) GetCurrentEnvironment() (string, int, *error) {
	var environ string
	var err *error
	if env.Environ == Debug {
		environ = "DEBUG"
	} else if env.Environ == Prod {
		environ = "PRODUCTION"
	} else if env.Environ == Test {
		environ = "Testing"
	} else {
		err = &NoDotEnvError
	}

	return environ, env.Environ, err
}

func (env *Settings) GetDatabaseConfig() string {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", env.DB.DbHost, env.DB.DbUser, env.DB.DbPassword, env.DB.DbName, env.DB.DbPort)

	return dsn
}

func setViperEnvironment() {
	workDir, _ := os.Getwd()

	viper.SetConfigType(ConfigType())
	viper.AddConfigPath(filepath.Dir(workDir))
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

}

func NewConfig(environmentType int) (*Settings, error) {
	var settings Settings
	environment = environmentType
	setViperEnvironment()

	settings.Debug = environmentType == Debug

	settings.DB.DbName = viper.GetString("DB_NAME")
	settings.DB.DbPassword = viper.GetString("DB_PASSWORD")
	settings.DB.DbPort = viper.GetInt("DB_PORT")
	settings.DB.DbUser = viper.GetString("DB_USER")
	settings.DB.DbHost = viper.GetString("DB_HOST")
	settings.APIPort = viper.GetUint("APP_PORT")

	return &settings, nil
}
