package security

import (
	"crypto/sha256"
	"fmt"
	"net/http"
	"time"

	"github.com/mrexmelle/connect-authx/internal/config"
)

type Service struct {
	ConfigService *config.Service
}

func NewService(cfg *config.Service) *Service {
	return &Service{
		ConfigService: cfg,
	}
}

func (s *Service) GenerateJwt(employeeId string) (string, time.Time, error) {
	now := time.Now()
	exp := now.Add(
		time.Minute * time.Duration(s.ConfigService.GetJwtValidMinute()),
	)
	_, token, err := s.ConfigService.TokenAuth.Encode(
		map[string]interface{}{
			"aud": "connect-web",
			"exp": exp.Unix(),
			"iat": now.Unix(),
			"iss": "connect-authx",
			"sub": s.GenerateHash(employeeId),
		},
	)

	if err != nil {
		return "", time.Time{}, err
	}

	return token, exp, nil
}

func (s *Service) GenerateHash(employeeId string) string {
	hasher := sha256.New()
	hasher.Write([]byte(employeeId))

	return fmt.Sprintf("u%x", hasher.Sum(nil))
}

func (s *Service) GenerateJwtCookie(token string, expires time.Time) (cookie http.Cookie) {
	return http.Cookie{
		Name:     "jwt",
		Value:    token,
		Path:     "/",
		Expires:  expires,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteNoneMode,
	}
}
