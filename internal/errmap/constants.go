package errmap

import (
	"database/sql"
	"net/http"

	"gorm.io/gorm"
)

const (
	ErrSvcCodeDuplicatedKey      = "duplicated_key"
	ErrSvcCodeForeignKeyViolated = "foreign_key_violation"
	ErrSvcCodeRecordNotFound     = "record_not_found"

	ErrSvcCodeUnregistered = "unregistered"
	ErrSvcCodeNone         = "success"
)

var DefaultErrorMap = map[error]CodePair{
	gorm.ErrDuplicatedKey:      NewCodePair(http.StatusBadRequest, ErrSvcCodeDuplicatedKey),
	gorm.ErrForeignKeyViolated: NewCodePair(http.StatusBadRequest, ErrSvcCodeForeignKeyViolated),
	gorm.ErrRecordNotFound:     NewCodePair(http.StatusNotFound, ErrSvcCodeRecordNotFound),
	sql.ErrNoRows:              NewCodePair(http.StatusNotFound, ErrSvcCodeRecordNotFound),
}
