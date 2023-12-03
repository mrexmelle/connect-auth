package config

import (
	"os"

	"github.com/spf13/viper"
)

type Repository interface {
	GetProfile() string
	GetDsn() string
	GetJwtSecret() string
	GetJwtValidMinute() int
	GetPort() int
	GetDefaultUserPassword() string
}

type RepositoryImpl struct {
	Profile             string
	Dsn                 string
	JwtSecret           string
	JwtValidMinute      int
	Port                int
	DefaultUserPassword string
}

func NewRepository() Repository {
	profile := os.Getenv("APP_PROFILE")
	if profile == "" {
		profile = "local"
	}
	viper.SetConfigName("application-" + profile)
	viper.SetConfigType("yml")
	for _, cp := range []string{
		"/etc/conf",
		"./config",
	} {
		viper.AddConfigPath(cp)
	}
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	var dsn = ""
	for key, value := range viper.GetStringMapString("app.datasource") {
		dsn += string(key + "=" + value + " ")
	}

	jwtSecret := viper.GetString("app.security.jwt.secret")
	jwtValidMinute := viper.GetInt("app.security.jwt.valid-minute")
	defaultUserPassword := viper.GetString("app.security.default-user-password")
	port := viper.GetInt("app.server.port")

	return &RepositoryImpl{
		Profile:             profile,
		Dsn:                 dsn,
		JwtSecret:           jwtSecret,
		JwtValidMinute:      jwtValidMinute,
		Port:                port,
		DefaultUserPassword: defaultUserPassword,
	}
}

func (r *RepositoryImpl) GetProfile() string {
	return r.Profile
}

func (r *RepositoryImpl) GetDsn() string {
	return r.Dsn
}

func (r *RepositoryImpl) GetJwtSecret() string {
	return r.JwtSecret
}

func (r *RepositoryImpl) GetJwtValidMinute() int {
	return r.JwtValidMinute
}

func (r *RepositoryImpl) GetPort() int {
	return r.Port
}

func (r *RepositoryImpl) GetDefaultUserPassword() string {
	return r.DefaultUserPassword
}
