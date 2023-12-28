package profile

import (
	"github.com/mrexmelle/connect-authx/internal/config"
	"github.com/mrexmelle/connect-authx/internal/mapper"
)

type Service struct {
	ConfigService     *config.Service
	ProfileRepository Repository
}

func NewService(
	cfg *config.Service,
	r Repository) *Service {
	return &Service{
		ConfigService:     cfg,
		ProfileRepository: r,
	}
}

func (s *Service) RetrieveByEhid(ehid string) ResponseDto {
	result, err := s.ProfileRepository.FindByEhid(ehid)
	return ResponseDto{
		Profile: result,
		Status:  mapper.ToStatus(err),
	}
}

func (s *Service) UpdateByEhid(
	fields map[string]string,
	ehid string,
) PatchResponseDto {
	err := s.ProfileRepository.UpdateByEhid(
		fields,
		ehid,
	)
	return PatchResponseDto{
		Status: mapper.ToStatus(err),
	}
}

func (s *Service) DeleteByEhid(
	ehid string,
) DeleteResponseDto {
	err := s.ProfileRepository.DeleteByEhid(ehid)
	return DeleteResponseDto{
		Status: mapper.ToStatus(err),
	}
}
