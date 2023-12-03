package profile

import (
	"errors"
	"time"

	"github.com/mrexmelle/connect-auth/internal/config"
	"github.com/mrexmelle/connect-auth/internal/mapper"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Repository interface {
	PrepareWithDb(db *gorm.DB, employeeId string) error
	UpdateByEhid(fields map[string]string, ehid string) error
	FindByEhid(ehid string) (Entity, error)
	DeleteByEhid(ehid string) error
}

type RepositoryImpl struct {
	ConfigService *config.Service
	TableName     string
}

func NewRepository(cfg *config.Service) Repository {
	return &RepositoryImpl{
		ConfigService: cfg,
		TableName:     "profiles",
	}
}

func (r *RepositoryImpl) PrepareWithDb(
	db *gorm.DB,
	employeeId string,
) error {
	res := db.Exec(
		"INSERT INTO "+r.TableName+"(ehid, employee_id, name, email_address, dob, "+
			"created_at, updated_at) "+
			"VALUES(?, ?, ?, ?, ?, NOW(), NOW())",
		mapper.ToEhid(employeeId),
		employeeId,
		"",
		"",
		nil,
	)
	return res.Error
}

func (r *RepositoryImpl) UpdateByEhid(
	fields map[string]string,
	ehid string,
) error {
	dbFields := map[string]interface{}{}
	name, ok := fields["name"]
	if ok {
		dbFields["name"] = name
	}

	emailAddress, ok := fields["email_address"]
	if ok {
		dbFields["email_address"] = emailAddress
	}

	dob, ok := fields["dob"]
	if ok {
		ts, err := time.Parse("2006-01-02", dob)
		if err == nil {
			dbFields["dob"] = datatypes.Date(ts)
		} else {
			return err
		}
	}

	if len(dbFields) > 0 {
		dbFields["updated_at"] = time.Now()
		result := r.ConfigService.Db.
			Table(r.TableName).
			Where("ehid = ?", ehid).
			Updates(dbFields)

		if result.Error != nil {
			return result.Error
		}

		if result.RowsAffected == 0 {
			return errors.New("request invalid")
		}
	}

	return nil
}

func (r *RepositoryImpl) FindByEhid(ehid string) (Entity, error) {
	response := Entity{
		Ehid: ehid,
	}
	var dob time.Time
	err := r.ConfigService.Db.
		Select("employee_id, name, email_address, dob").
		Table(r.TableName).
		Where("ehid = ?", ehid).
		Row().
		Scan(&response.EmployeeId, &response.Name, &response.EmailAddress, &dob)
	if err != nil {
		return Entity{}, err
	}

	response.Dob = dob.Format("2006-01-02")
	return response, nil
}

func (r *RepositoryImpl) DeleteByEhid(ehid string) error {
	now := time.Now()
	result := r.ConfigService.Db.
		Table(r.TableName).
		Where("ehid = ?", ehid).
		Updates(
			map[string]interface{}{
				"email_address": "",
				"updated_at":    now,
			},
		)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
