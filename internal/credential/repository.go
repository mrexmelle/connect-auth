package credential

import (
	"database/sql"
	"errors"
	"time"

	"github.com/mrexmelle/connect-authx/internal/config"
	"gorm.io/gorm"
)

type Repository interface {
	CreateWithDb(db *gorm.DB, employeeId string, password string) error
	ExistsByEmployeeIdAndPassword(employeeId string, password string) (bool, error)
	DeleteByEmployeeId(employeeId string) error
	UpdatePasswordByEmployeeIdAndPassword(newPassword string, employeeId string, currentPassword string) error
	ResetPasswordByEmployeeId(employeeId string) error
}

type RepositoryImpl struct {
	ConfigService *config.Service
	TableName     string
}

func NewRepository(cfg *config.Service) Repository {
	return &RepositoryImpl{
		ConfigService: cfg,
		TableName:     "credentials",
	}
}

func (r *RepositoryImpl) CreateWithDb(
	db *gorm.DB,
	employeeId string,
	password string,
) error {
	result := db.Exec(
		"INSERT INTO "+r.TableName+"(employee_id, password_hash, "+
			"created_at, updated_at) "+
			"VALUES(?, CRYPT(?, GEN_SALT('bf', 8)), NOW(), NOW())",
		employeeId,
		password,
	)
	return result.Error
}

func (r *RepositoryImpl) ExistsByEmployeeIdAndPassword(
	employeeId string,
	password string,
) (bool, error) {
	var idResult string
	err := r.ConfigService.ReadDb.
		Select("employee_id").
		Table(r.TableName).
		Where(
			"employee_id = ? AND password_hash = CRYPT(?, password_hash) "+
				"AND deleted_at IS NULL",
			employeeId,
			password,
		).
		Row().
		Scan(&idResult)
	if errors.Is(err, sql.ErrNoRows) {
		return false, nil
	}

	return (idResult == employeeId), err
}

func (r *RepositoryImpl) DeleteByEmployeeId(employeeId string) error {
	now := time.Now()
	result := r.ConfigService.WriteDb.
		Table(r.TableName).
		Where("employee_id = ? AND deleted_at IS NULL", employeeId).
		Updates(
			map[string]interface{}{
				"deleted_at": now,
				"updated_at": now,
			},
		)
	return result.Error
}

func (r *RepositoryImpl) UpdatePasswordByEmployeeIdAndPassword(
	newPassword string,
	employeeId string,
	currentPassword string,
) error {
	result := r.ConfigService.WriteDb.Exec(
		"UPDATE "+r.TableName+" SET "+
			"password_hash = CRYPT(?, GEN_SALT('bf', 8)), "+
			"updated_at = NOW() "+
			"WHERE employee_id = ? AND password_hash = CRYPT(?, password_hash)",
		newPassword,
		employeeId,
		currentPassword,
	)
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return result.Error
}

func (r *RepositoryImpl) ResetPasswordByEmployeeId(
	employeeId string,
) error {
	result := r.ConfigService.WriteDb.Exec(
		"UPDATE "+r.TableName+" SET "+
			"password_hash = CRYPT(?, GEN_SALT('bf', 8)), "+
			"updated_at = NOW() "+
			"WHERE employee_id = ?",
		r.ConfigService.GetDefaultUserPassword(),
		employeeId,
	)
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return result.Error
}
