package credential

import (
	"github.com/mrexmelle/connect-authx/internal/config"
	"github.com/mrexmelle/connect-authx/internal/profile"
	"github.com/mrexmelle/connect-authx/internal/security"
)

type Service struct {
	ConfigService        *config.Service
	SecurityService      *security.Service
	CredentialRepository Repository
	ProfileRepository    profile.Repository
}

func NewService(
	cfg *config.Service,
	ss *security.Service,
	cr Repository,
	pr profile.Repository,
) *Service {
	return &Service{
		ConfigService:        cfg,
		SecurityService:      ss,
		CredentialRepository: cr,
		ProfileRepository:    pr,
	}
}

func (s *Service) Create(req PostRequestDto) error {
	trx := s.ConfigService.WriteDb.Begin()
	defer func() {
		if r := recover(); r != nil {
			trx.Rollback()
		}
	}()
	if trx.Error != nil {
		return trx.Error
	}

	err := s.CredentialRepository.CreateWithDb(
		trx,
		req.EmployeeId,
		req.Password,
	)
	if err != nil {
		trx.Rollback()
		return err
	}

	err = s.ProfileRepository.CreateWithDb(
		trx,
		req.EmployeeId,
		s.SecurityService.GenerateHash(req.EmployeeId),
	)
	if err != nil {
		trx.Rollback()
		return err
	}

	err = trx.Commit().Error
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) DeleteByEmployeeId(employeeId string) error {
	return s.CredentialRepository.DeleteByEmployeeId(employeeId)
}

func (s *Service) UpdatePasswordByEmployeeId(employeeId string, req PatchPasswordRequestDto) error {
	return s.CredentialRepository.UpdatePasswordByEmployeeIdAndPassword(
		req.NewPassword,
		employeeId,
		req.CurrentPassword,
	)
}

func (s *Service) ResetPasswordByEmployeeId(employeeId string) error {
	return s.CredentialRepository.ResetPasswordByEmployeeId(employeeId)
}
