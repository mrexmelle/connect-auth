package credential

import (
	"time"

	"github.com/mrexmelle/connect-auth/internal/config"
	"gorm.io/gorm"
)

type RepositoryImpl struct {
	Config    *config.Config
	TableName string
}

type Repository interface {
	CreateWithDb(db *gorm.DB, employeeId string, password string) error
	ExistsByEmployeeIdAndPassword(employeeId string, password string) (bool, error)
	DeleteByEmployeeId(employeeId string) error
	UpdatePasswordByEmployeeIdAndPassword(newPassword string, employeeId string, currentPassword string) error
	ResetPasswordByEmployeeId(employeeId string) error
}

func NewRepository(cfg *config.Config) Repository {
	return &RepositoryImpl{
		Config:    cfg,
		TableName: "credentials",
	}
}

func (r *RepositoryImpl) CreateWithDb(
	db *gorm.DB,
	employeeId string,
	password string,
) error {
	res := db.Exec(
		"INSERT INTO "+r.TableName+"(employee_id, password_hash, "+
			"created_at, updated_at) "+
			"VALUES(?, CRYPT(?, GEN_SALT('bf', 8)), NOW(), NOW())",
		employeeId,
		password,
	)
	return res.Error
}

func (r *RepositoryImpl) ExistsByEmployeeIdAndPassword(
	employeeId string,
	password string,
) (bool, error) {
	var idResult string
	err := r.Config.Db.
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
	return (idResult == employeeId), err
}

func (r *RepositoryImpl) DeleteByEmployeeId(employeeId string) error {
	now := time.Now()
	result := r.Config.Db.
		Table(r.TableName).
		Where("employee_id = ? AND deleted_at IS NULL", employeeId).
		Updates(
			map[string]interface{}{
				"deleted_at": now,
				"updated_at": now,
			},
		)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *RepositoryImpl) UpdatePasswordByEmployeeIdAndPassword(
	newPassword string,
	employeeId string,
	currentPassword string,
) error {
	result := r.Config.Db.Exec(
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
	result := r.Config.Db.Exec(
		"UPDATE "+r.TableName+" SET "+
			"password_hash = CRYPT(?, GEN_SALT('bf', 8)), "+
			"updated_at = NOW() "+
			"WHERE employee_id = ?",
		r.Config.DefaultUserPassword,
		employeeId,
	)
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return result.Error
}
