package profile

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/mrexmelle/connect-auth/internal/config"
)

type Controller struct {
	Config         *config.Config
	ProfileService *Service
}

func NewController(cfg *config.Config, svc *Service) *Controller {
	return &Controller{
		Config:         cfg,
		ProfileService: svc,
	}
}

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

func (c *Controller) Delete(w http.ResponseWriter, r *http.Request) {
	response := c.ProfileService.DeleteByEhid(chi.URLParam(r, "ehid"))

	if response.Status != "OK" {
		http.Error(w, "DELETE failure: "+response.Status, http.StatusInternalServerError)
		return
	}

	responseBody, _ := json.Marshal(&response)
	w.Write([]byte(responseBody))
}
