package session

import (
	"github.com/mrexmelle/connect-authx/internal/config"
	"github.com/mrexmelle/connect-authx/internal/credential"
	"github.com/mrexmelle/connect-authx/internal/errmap"
	"github.com/mrexmelle/connect-authx/internal/security"
)

type Service struct {
	ConfigService        *config.Service
	SecurityService      *security.Service
	CredentialRepository credential.Repository
	ErrorMapper          errmap.Class
}

func NewService(cfg *config.Service, ss *security.Service, cr credential.Repository) *Service {
	return &Service{
		ConfigService:        cfg,
		SecurityService:      ss,
		CredentialRepository: cr,
		ErrorMapper:          *errmap.New(&ErrorMap),
	}
}

func (s *Service) Authenticate(req PostRequestDto) (*SigningResult, error) {
	exists, err := s.CredentialRepository.ExistsByEmployeeIdAndPassword(
		req.EmployeeId,
		req.Password,
	)

	if err != nil {
		return nil, err
	} else if !exists {
		return nil, ErrAuthentication
	}

	jwt, exp, err := s.SecurityService.GenerateJwt(req.EmployeeId)
	if err != nil {
		return nil, err
	}

	return &SigningResult{
			Token:   jwt,
			Expires: exp,
		},
		nil
}
