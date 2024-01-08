package session

import (
	"encoding/json"
	"net/http"

	"github.com/mrexmelle/connect-authx/internal/config"
	"github.com/mrexmelle/connect-authx/internal/dtoresponse"
	"github.com/mrexmelle/connect-authx/internal/security"
)

type Controller struct {
	ConfigService   *config.Service
	SecurityService *security.Service
	SessionService  *Service
}

func NewController(cfg *config.Service, ss *security.Service, svc *Service) *Controller {
	return &Controller{
		ConfigService:   cfg,
		SecurityService: ss,
		SessionService:  svc,
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
	json.NewDecoder(r.Body).Decode(&requestBody)
	data, err := c.SessionService.Authenticate(requestBody)
	dtoresponse.NewWithData[SigningResult](data, err).
		WithErrorMap(&ErrorMap).
		WithPrewriteHook(func(s *SigningResult) {
			cookie := c.SecurityService.GenerateJwtCookie(s.Token, s.Expires)
			http.SetCookie(w, &cookie)
		}).
		RenderTo(w)
}
