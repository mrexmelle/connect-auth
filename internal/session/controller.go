package session

import (
	"encoding/json"
	"net/http"

	"github.com/mrexmelle/connect-auth/internal/config"
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
