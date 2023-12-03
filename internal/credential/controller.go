package credential

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/mrexmelle/connect-auth/internal/config"
)

type Controller struct {
	ConfigService     *config.Service
	CredentialService *Service
}

func NewController(cfg *config.Service, svc *Service) *Controller {
	return &Controller{
		ConfigService:     cfg,
		CredentialService: svc,
	}
}

// Post Credentials : HTTP endpoint to post new credentials
// @Tags Credentials
// @Description Post a new credential
// @Accept json
// @Produce json
// @Param data body PostRequestDto true "Credential Request"
// @Success 200 {object} ResponseDto "Success Response"
// @Failure 500 "InternalServerError"
// @Router /credentials [POST]
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

// Delete Credentials : HTTP endpoint to delete credentials
// @Tags Credentials
// @Description Delete a credential
// @Produce json
// @Param employee_id path string true "Employee ID"
// @Success 200 {object} ResponseDto "Success Response"
// @Failure 500 "InternalServerError"
// @Router /credentials/{employee_id} [DELETE]
func (c *Controller) Delete(w http.ResponseWriter, r *http.Request) {

	err := c.CredentialService.DeleteByEmployeeId(chi.URLParam(r, "employee_id"))

	if err != nil {
		http.Error(w, "DELETE failure: "+err.Error(), http.StatusInternalServerError)
		return
	}

	responseBody, _ := json.Marshal(
		&ResponseDto{Status: "OK"},
	)
	w.Write([]byte(responseBody))
}

// Patch Password : HTTP endpoint to patch password
// @Tags Credentials
// @Description Patch a credential's password
// @Accept json
// @Produce json
// @Param employee_id path string true "Employee ID"
// @Param data body PatchPasswordRequestDto true "Password Patch Request"
// @Success 200 {object} ResponseDto "Success Response"
// @Failure 500 "InternalServerError"
// @Router /credentials/{employee_id}/password [PATCH]
func (c *Controller) PatchPassword(w http.ResponseWriter, r *http.Request) {
	var requestBody PatchPasswordRequestDto
	json.NewDecoder(r.Body).Decode(&requestBody)

	response, err := c.CredentialService.UpdatePasswordByEmployeeId(
		chi.URLParam(r, "employee_id"),
		requestBody)

	if err != nil {
		http.Error(w, "PATCH failure: "+err.Error(), http.StatusInternalServerError)
		return
	}

	responseBody, _ := json.Marshal(
		&response,
	)
	w.Write([]byte(responseBody))
}

// Reset Password : HTTP endpoint to reset password
// @Tags Credentials
// @Description Reset a credential's password
// @Produce json
// @Param employee_id path string true "Employee ID"
// @Success 200 {object} ResponseDto "Success Response"
// @Failure 500 "InternalServerError"
// @Router /credentials/{employee_id}/password [DELETE]
func (c *Controller) DeletePassword(w http.ResponseWriter, r *http.Request) {
	err := c.CredentialService.ResetPasswordByEmployeeId(chi.URLParam(r, "employee_id"))

	if err != nil {
		http.Error(w, "DELETE failure: "+err.Error(), http.StatusInternalServerError)
		return
	}

	responseBody, _ := json.Marshal(
		&ResponseDto{Status: "OK"},
	)
	w.Write([]byte(responseBody))
}
