package profile

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
	"github.com/mrexmelle/connect-authx/internal/config"
	"github.com/mrexmelle/connect-authx/internal/dto/dtorespwithdata"
	"github.com/mrexmelle/connect-authx/internal/dto/dtorespwithoutdata"
	"github.com/mrexmelle/connect-authx/internal/localerror"
)

type Controller struct {
	ConfigService     *config.Service
	LocalErrorService *localerror.Service
	ProfileService    *Service
}

func NewController(cfg *config.Service, les *localerror.Service, svc *Service) *Controller {
	return &Controller{
		ConfigService:     cfg,
		LocalErrorService: les,
		ProfileService:    svc,
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
		dtorespwithdata.NewError(
			localerror.ErrAuthentication.Error(),
			err.Error(),
		).RenderTo(w, http.StatusUnauthorized)
		return
	}

	data := claims["sub"].(string)
	if data == "" {
		dtorespwithdata.NewError(
			localerror.ErrAuthentication.Error(),
			"sub claim empty or not found",
		).RenderTo(w, http.StatusBadRequest)
		return
	}

	dtorespwithdata.New[string](
		&data,
		localerror.ErrSvcCodeNone,
		"OK",
	).RenderTo(w, http.StatusOK)
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
	info := c.LocalErrorService.Map(err)
	dtorespwithdata.New[Entity](
		data,
		info.ServiceErrorCode,
		info.ServiceErrorMessage,
	).RenderTo(w, info.HttpStatusCode)
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
		dtorespwithoutdata.New(
			localerror.ErrBadJson.Error(),
			err.Error(),
		).RenderTo(w, http.StatusBadRequest)
		return
	}
	err = c.ProfileService.UpdateByEhid(requestBody.Fields, chi.URLParam(r, "ehid"))
	info := c.LocalErrorService.Map(err)
	dtorespwithoutdata.New(
		info.ServiceErrorCode,
		info.ServiceErrorMessage,
	).RenderTo(w, info.HttpStatusCode)
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
	info := c.LocalErrorService.Map(err)
	dtorespwithoutdata.New(
		info.ServiceErrorCode,
		info.ServiceErrorMessage,
	).RenderTo(w, info.HttpStatusCode)
}
