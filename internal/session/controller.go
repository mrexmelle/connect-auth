package session

import (
	"encoding/json"
	"net/http"

	"github.com/mrexmelle/connect-authx/internal/config"
	"github.com/mrexmelle/connect-authx/internal/dto/dtorespwithdata"
	"github.com/mrexmelle/connect-authx/internal/localerror"
	"github.com/mrexmelle/connect-authx/internal/security"
)

type Controller struct {
	ConfigService     *config.Service
	LocalErrorService *localerror.Service
	SecurityService   *security.Service
	SessionService    *Service
}

func NewController(
	cfg *config.Service,
	les *localerror.Service,
	ss *security.Service,
	svc *Service,
) *Controller {
	return &Controller{
		ConfigService:     cfg,
		LocalErrorService: les,
		SecurityService:   ss,
		SessionService:    svc,
	}
}

// Post Sessions : HTTP endpoint to post new sessions
// @Tags Sessions
// @Description Post a new session
// @Accept json
// @Produce json
// @Param data body PostRequestDto true "Session Request"
// @Success 200 {object} PostResponseDto "Success Response"
// @Failure 400 "BadRequest"
// @Failure 401 "Unauthorized"
// @Failure 500 "InternalServerError"
// @Router /sessions [POST]
func (c *Controller) Post(w http.ResponseWriter, r *http.Request) {
	var requestBody PostRequestDto
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		dtorespwithdata.NewError(
			localerror.ErrBadJson.Error(),
			err.Error(),
		).RenderTo(w, http.StatusBadRequest)
		return
	}
	data, err := c.SessionService.Authenticate(requestBody)
	info := c.LocalErrorService.Map(err)
	dtorespwithdata.New(
		data,
		info.ServiceErrorCode,
		info.ServiceErrorMessage,
	).WithPrewriteHook(func(s *SigningResult) {
		cookie := c.SecurityService.GenerateJwtCookie(s.Token, s.Expires)
		http.SetCookie(w, &cookie)
	}).RenderTo(w, info.HttpStatusCode)
}
