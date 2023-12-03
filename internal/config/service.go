package config

import (
	"strings"

	"github.com/go-chi/jwtauth"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Service struct {
	ConfigRepository Repository
	Db               *gorm.DB
	TokenAuth        *jwtauth.JWTAuth
}

func NewService(
	cr Repository,
) *Service {
	db, err := gorm.Open(
		postgres.Open(strings.TrimSpace(cr.GetDsn())),
		&gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		},
	)
	if err != nil {
		panic(err)
	}

	jwta := jwtauth.New(
		"HS256",
		[]byte(cr.GetJwtSecret()),
		nil,
	)

	return &Service{
		ConfigRepository: cr,
		Db:               db,
		TokenAuth:        jwta,
	}
}

func (s *Service) GetProfile() string {
	return s.ConfigRepository.GetProfile()
}

func (s *Service) GetJwtValidMinute() int {
	return s.ConfigRepository.GetJwtValidMinute()
}

func (s *Service) GetPort() int {
	return s.ConfigRepository.GetPort()
}

func (s *Service) GetDefaultUserPassword() string {
	return s.ConfigRepository.GetDefaultUserPassword()
}
