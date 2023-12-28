package config

import (
	"os"

	"github.com/spf13/viper"
)

type Repository interface {
	GetProfile() string
	GetReadDsn() string
	GetWriteDsn() string
	GetJwtSecret() string
	GetJwtValidMinute() int
	GetPort() int
	GetDefaultUserPassword() string
}

type RepositoryImpl struct {
	Profile             string
	ReadDsn             string
	WriteDsn            string
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

	var readDsn = ""
	for key, value := range viper.GetStringMapString("app.datasource.read") {
		readDsn += string(key + "=" + value + " ")
	}

	var writeDsn = ""
	for key, value := range viper.GetStringMapString("app.datasource.write") {
		writeDsn += string(key + "=" + value + " ")
	}

	jwtSecret := viper.GetString("app.security.jwt.secret")
	jwtValidMinute := viper.GetInt("app.security.jwt.valid-minute")
	defaultUserPassword := viper.GetString("app.security.default-user-password")
	port := viper.GetInt("app.server.port")

	return &RepositoryImpl{
		Profile:             profile,
		ReadDsn:             readDsn,
		WriteDsn:            writeDsn,
		JwtSecret:           jwtSecret,
		JwtValidMinute:      jwtValidMinute,
		Port:                port,
		DefaultUserPassword: defaultUserPassword,
	}
}

func (r *RepositoryImpl) GetProfile() string {
	return r.Profile
}

func (r *RepositoryImpl) GetReadDsn() string {
	return r.ReadDsn
}

func (r *RepositoryImpl) GetWriteDsn() string {
	return r.WriteDsn
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
