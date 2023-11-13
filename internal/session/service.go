package session

import (
	"time"

	"github.com/mrexmelle/connect-auth/internal/config"
	"github.com/mrexmelle/connect-auth/internal/credential"
	"github.com/mrexmelle/connect-auth/internal/mapper"
)

type Service struct {
	Config               *config.Config
	CredentialRepository *credential.Repository
}

func NewService(cfg *config.Config, repo *credential.Repository) *Service {
	return &Service{
		Config:               cfg,
		CredentialRepository: repo,
	}
}

func (s *Service) Authenticate(req RequestDto) (bool, error) {
	return s.CredentialRepository.ExistsByEmployeeIdAndPassword(
		req.EmployeeId,
		req.Password,
	)
}

func (s *Service) GenerateJwt(employeeId string) (string, time.Time, error) {
	now := time.Now()
	exp := now.Add(
		time.Minute * time.Duration(s.Config.JwtValidMinute),
	)
	_, token, err := s.Config.TokenAuth.Encode(
		map[string]interface{}{
			"aud": "connect-web",
			"exp": exp.Unix(),
			"iat": now.Unix(),
			"iss": "connect-auth",
			"sub": mapper.ToEhid(employeeId),
		},
	)

	if err != nil {
		return "", time.Time{}, err
	}

	return token, exp, nil
}
