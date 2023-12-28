package session

import (
	"time"

	"github.com/mrexmelle/connect-authx/internal/config"
	"github.com/mrexmelle/connect-authx/internal/credential"
	"github.com/mrexmelle/connect-authx/internal/mapper"
)

type Service struct {
	ConfigService        *config.Service
	CredentialRepository credential.Repository
}

func NewService(cfg *config.Service, cr credential.Repository) *Service {
	return &Service{
		ConfigService:        cfg,
		CredentialRepository: cr,
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
		time.Minute * time.Duration(s.ConfigService.GetJwtValidMinute()),
	)
	_, token, err := s.ConfigService.TokenAuth.Encode(
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
