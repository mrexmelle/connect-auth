package session

import (
	"encoding/json"
	"net/http"

	"github.com/mrexmelle/connect-authx/internal/config"
)

type Controller struct {
	ConfigService  *config.Service
	SessionService *Service
}

func NewController(cfg *config.Service, svc *Service) *Controller {
	return &Controller{
		ConfigService:  cfg,
		SessionService: svc,
	}
}

// Post Sessions : HTTP endpoint to post new sessions
// @Tags Sessions
// @Description Post a new session
// @Accept json
// @Produce json
// @Param data body RequestDto true "Session Request"
// @Success 200 {object} ResponseDto "Success Response"
// @Failure 500 "InternalServerError"
// @Failure 401 "Unauthorized"
// @Router /sessions [POST]
func (c *Controller) Post(w http.ResponseWriter, r *http.Request) {
	var requestBody RequestDto
	json.NewDecoder(r.Body).Decode(&requestBody)

	queryResult, err := c.SessionService.Authenticate(requestBody)

	if err != nil {
		http.Error(w, "POST failure: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if !queryResult {
		http.Error(w, "POST failure: "+err.Error(), http.StatusUnauthorized)
		return
	}

	signingResult, exp, err := c.SessionService.GenerateJwt(requestBody.EmployeeId)
	if err != nil {
		http.Error(w, "POST failure: "+err.Error(), http.StatusInternalServerError)
		return
	}

	http.SetCookie(
		w, &http.Cookie{
			Name:     "jwt",
			Value:    signingResult,
			Path:     "/",
			Expires:  exp,
			HttpOnly: true,
			Secure:   true,
			SameSite: http.SameSiteNoneMode,
		},
	)
	responseBody, _ := json.Marshal(
		&ResponseDto{Token: signingResult},
	)
	w.Write([]byte(responseBody))
}
