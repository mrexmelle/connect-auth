package config

import (
	"os"
	"strings"

	"github.com/go-chi/jwtauth"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Config struct {
	Profile             string
	Db                  *gorm.DB
	TokenAuth           *jwtauth.JWTAuth
	JwtValidMinute      int
	Port                int
	DefaultUserPassword string
}

func New(
	plainConfigName string,
	configType string,
	configPaths []string,
) (Config, error) {
	profile := DetermineProfile()
	viper.SetConfigName(plainConfigName + "-" + profile)
	viper.SetConfigType(configType)
	for _, cp := range configPaths {
		viper.AddConfigPath(cp)
	}
	err := viper.ReadInConfig()
	if err != nil {
		return Config{}, err
	}

	db, err := CreateDb("app.datasource")
	if err != nil {
		return Config{}, err
	}

	jwta := CreateTokenAuth("app.security.jwt.secret")
	jwtvm := viper.GetInt("app.security.jwt.valid-minute")
	port := viper.GetInt("app.server.port")
	defaultUserPassword := viper.GetString("app.security.default-user-password")

	return Config{
		Profile:             profile,
		Db:                  db,
		TokenAuth:           jwta,
		JwtValidMinute:      jwtvm,
		Port:                port,
		DefaultUserPassword: defaultUserPassword,
	}, nil
}

func DetermineProfile() string {
	profile := os.Getenv("APP_PROFILE")
	if profile == "" {
		profile = "local"
	}
	return profile
}

func CreateDb(configKey string) (*gorm.DB, error) {
	var dsn = ""
	for key, value := range viper.GetStringMapString(configKey) {
		dsn += string(key + "=" + value + " ")
	}
	return gorm.Open(
		postgres.Open(strings.TrimSpace(dsn)),
		&gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		},
	)
}

func CreateTokenAuth(configKey string) *jwtauth.JWTAuth {
	return jwtauth.New(
		"HS256",
		[]byte(viper.GetString(configKey)),
		nil,
	)
}
