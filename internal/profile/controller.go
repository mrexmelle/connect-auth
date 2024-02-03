package profile

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
	"github.com/mrexmelle/connect-authx/internal/config"
	"github.com/mrexmelle/connect-authx/internal/dto/dtobuilderwithdata"
	"github.com/mrexmelle/connect-authx/internal/dto/dtobuilderwithoutdata"
	"github.com/mrexmelle/connect-authx/internal/localerror"
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

// Get My EHID : HTTP endpoint to get current user's EHID
// @Tags Profiles
// @Description Get current user's EHID
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Success 200 {object} GetEhidResponseDto "Success Response"
// @Failure 400 "BadRequest"
// @Failure 500 "InternalServerError"
// @Router /profiles/me/ehid [GET]
func (c *Controller) GetMyEhid(w http.ResponseWriter, r *http.Request) {
	_, claims, err := jwtauth.FromContext(r.Context())
	if err != nil {
		dtobuilderwithdata.New[string](new(string), localerror.ErrAuthentication).RenderTo(w)
		return
	}
	data := claims["sub"].(string)
	dtobuilderwithdata.New[string](&data, err).RenderTo(w)
}

// Get Profiles : HTTP endpoint to get a profile
// @Tags Profiles
// @Description Get a profile
// @Produce json
// @Param ehid path string true "Employee Hash ID"
// @Success 200 {object} GetResponseDto "Success Response"
// @Failure 400 "BadRequest"
// @Failure 500 "InternalServerError"
// @Router /profiles/{ehid} [GET]
func (c *Controller) Get(w http.ResponseWriter, r *http.Request) {
	data, err := c.ProfileService.RetrieveByEhid(
		chi.URLParam(r, "ehid"),
	)
	dtobuilderwithdata.New[Entity](data, err).RenderTo(w)
}

// Patch Profiles : HTTP endpoint to patch a profile
// @Tags Profiles
// @Description Patch a profile
// @Accept json
// @Produce json
// @Param ehid path string true "Employee Hash ID"
// @Param data body PatchRequestDto true "Profile Patch Request"
// @Success 200 {object} PatchResponseDto "Success Response"
// @Failure 400 "BadRequest"
// @Failure 500 "InternalServerError"
// @Router /profiles/{ehid} [PATCH]
func (c *Controller) Patch(w http.ResponseWriter, r *http.Request) {
	var requestBody PatchRequestDto
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		dtobuilderwithoutdata.New(localerror.ErrBadJson).RenderTo(w)
		return
	}
	err = c.ProfileService.UpdateByEhid(requestBody.Fields, chi.URLParam(r, "ehid"))
	dtobuilderwithoutdata.New(err).RenderTo(w)
}

// Delete Profiles : HTTP endpoint to delete a profile
// @Tags Profiles
// @Description Delete a profile
// @Produce json
// @Param ehid path string true "Employee Hash ID"
// @Success 200 {object} DeleteResponseDto "Success Response"
// @Failure 400 "BadRequest"
// @Failure 500 "InternalServerError"
// @Router /profiles/{ehid} [DELETE]
func (c *Controller) Delete(w http.ResponseWriter, r *http.Request) {
	err := c.ProfileService.DeleteByEhid(chi.URLParam(r, "ehid"))
	dtobuilderwithoutdata.New(err).RenderTo(w)
}
