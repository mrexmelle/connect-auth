package session

import (
	"errors"
	"net/http"

	"github.com/mrexmelle/connect-authx/internal/errmap"
)

var (
	ErrAuthentication = errors.New("authentication_error")
)

var ErrorMap = map[error]errmap.CodePair{
	ErrAuthentication: errmap.NewCodePair(http.StatusUnauthorized, ErrAuthentication.Error()),
}
