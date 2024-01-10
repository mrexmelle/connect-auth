package profile

import (
	"github.com/mrexmelle/connect-authx/internal/config"
)

type Service struct {
	ConfigService     *config.Service
	ProfileRepository Repository
}

func NewService(cfg *config.Service, r Repository) *Service {
	return &Service{
		ConfigService:     cfg,
		ProfileRepository: r,
	}
}

func (s *Service) RetrieveByEhid(ehid string) (*Entity, error) {
	return s.ProfileRepository.FindByEhid(ehid)
}

func (s *Service) UpdateByEhid(fields map[string]string, ehid string) error {
	return s.ProfileRepository.UpdateByEhid(fields, ehid)
}

func (s *Service) DeleteByEhid(ehid string) error {
	return s.ProfileRepository.DeleteByEhid(ehid)
}

func (s *Service) RetrieveEmployeeIdByEhid(ehid string) (string, error) {
	return s.ProfileRepository.FindEmployeeIdByEhid(ehid)
}
