package credential

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/mrexmelle/connect-auth/internal/config"
)

type Controller struct {
	Config            *config.Config
	CredentialService *Service
}

func NewController(cfg *config.Config, svc *Service) *Controller {
	return &Controller{
		Config:            cfg,
		CredentialService: svc,
	}
}

func (c *Controller) Post(w http.ResponseWriter, r *http.Request) {
	var requestBody PostRequestDto
	json.NewDecoder(r.Body).Decode(&requestBody)

	response, err := c.CredentialService.Create(requestBody)

	if err != nil {
		http.Error(w, "POST failure: "+err.Error(), http.StatusInternalServerError)
		return
	}

	responseBody, _ := json.Marshal(
		&response,
	)
	w.Write([]byte(responseBody))
}

func (c *Controller) Delete(w http.ResponseWriter, r *http.Request) {

	err := c.CredentialService.DeleteByEmployeeId(chi.URLParam(r, "ehid"))

	if err != nil {
		http.Error(w, "DELETE failure: "+err.Error(), http.StatusInternalServerError)
		return
	}

	responseBody, _ := json.Marshal(
		&ResponseDto{Status: "OK"},
	)
	w.Write([]byte(responseBody))
}

func (c *Controller) PatchPassword(w http.ResponseWriter, r *http.Request) {
	var requestBody PatchPasswordRequestDto
	json.NewDecoder(r.Body).Decode(&requestBody)

	response, err := c.CredentialService.UpdatePassword(requestBody)

	if err != nil {
		http.Error(w, "PATCH failure: "+err.Error(), http.StatusInternalServerError)
		return
	}

	responseBody, _ := json.Marshal(
		&response,
	)
	w.Write([]byte(responseBody))
}

func (c *Controller) DeletePassword(w http.ResponseWriter, r *http.Request) {
	err := c.CredentialService.ResetPassword(chi.URLParam(r, "ehid"))

	if err != nil {
		http.Error(w, "DELETE failure: "+err.Error(), http.StatusInternalServerError)
		return
	}

	responseBody, _ := json.Marshal(
		&ResponseDto{Status: "OK"},
	)
	w.Write([]byte(responseBody))
}
