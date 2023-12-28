package profile

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/mrexmelle/connect-authx/internal/config"
)

type Controller struct {
	ConfigService  *config.Service
	ProfileService *Service
}

func NewController(cfg *config.Service, svc *Service) *Controller {
	return &Controller{
		ConfigService:  cfg,
		ProfileService: svc,
	}
}

// Get Profiles : HTTP endpoint to get a profile
// @Tags Profiles
// @Description Get a profile
// @Produce json
// @Param ehid path string true "Employee Hash ID"
// @Success 200 {object} ResponseDto "Success Response"
// @Failure 500 "InternalServerError"
// @Router /profiles/{ehid} [GET]
func (c *Controller) Get(w http.ResponseWriter, r *http.Request) {
	response := c.ProfileService.RetrieveByEhid(
		chi.URLParam(r, "ehid"),
	)

	if response.Status != "OK" {
		http.Error(w, "GET failure: "+response.Status, http.StatusInternalServerError)
		return
	}

	responseBody, _ := json.Marshal(
		&response,
	)
	w.Write([]byte(responseBody))
}

// Patch Profiles : HTTP endpoint to patch a profile
// @Tags Profiles
// @Description Patch a profile
// @Accept json
// @Produce json
// @Param ehid path string true "Employee Hash ID"
// @Param data body PatchRequestDto true "Profile Patch Request"
// @Success 200 {object} PatchResponseDto "Success Response"
// @Failure 500 "InternalServerError"
// @Router /profiles/{ehid} [PATCH]
func (c *Controller) Patch(w http.ResponseWriter, r *http.Request) {
	var requestBody PatchRequestDto
	json.NewDecoder(r.Body).Decode(&requestBody)

	if r.Body != nil {
		bodyBytes, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Printf("Body reading error: %v", err)
			return
		}
		defer r.Body.Close()
		fmt.Printf("Body: %s\n", bodyBytes)
	}

	for k, v := range requestBody.Fields {
		fmt.Printf("Key: %s, Value: %s\n", k, v)
	}

	response := c.ProfileService.UpdateByEhid(
		requestBody.Fields,
		chi.URLParam(r, "ehid"),
	)

	if response.Status != "OK" {
		http.Error(w, "PATCH failure: "+response.Status, http.StatusInternalServerError)
		return
	}

	responseBody, _ := json.Marshal(
		&response,
	)
	w.Write([]byte(responseBody))
}

// Delete Profiles : HTTP endpoint to delete a profile
// @Tags Profiles
// @Description Delete a profile
// @Produce json
// @Param ehid path string true "Employee Hash ID"
// @Success 200 {object} DeleteResponseDto "Success Response"
// @Failure 500 "InternalServerError"
// @Router /profiles/{ehid} [DELETE]
func (c *Controller) Delete(w http.ResponseWriter, r *http.Request) {
	response := c.ProfileService.DeleteByEhid(chi.URLParam(r, "ehid"))

	if response.Status != "OK" {
		http.Error(w, "DELETE failure: "+response.Status, http.StatusInternalServerError)
		return
	}

	responseBody, _ := json.Marshal(&response)
	w.Write([]byte(responseBody))
}
