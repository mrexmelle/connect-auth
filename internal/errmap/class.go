package errmap

import "net/http"

type Class struct {
	ErrorMap *map[error]CodePair
}

func New(errorMap *map[error]CodePair) *Class {
	return &Class{
		ErrorMap: errorMap,
	}
}

func (s *Class) Map(err error) StatusInfo {
	if err == nil {
		return NewStatusInfo(http.StatusOK, ErrSvcCodeNone, "OK")
	}

	if s.ErrorMap != nil {
		codePair, exists := (*s.ErrorMap)[err]
		if exists {
			return NewStatusInfo(codePair.HttpStatusCode, codePair.ServiceErrorCode, err.Error())
		}
	}

	codePair, exists := DefaultErrorMap[err]
	if exists {
		return NewStatusInfo(codePair.HttpStatusCode, codePair.ServiceErrorCode, err.Error())
	}

	return NewStatusInfo(http.StatusBadRequest, ErrSvcCodeUnregistered, err.Error())
}
