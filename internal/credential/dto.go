package credential

import "github.com/mrexmelle/connect-authx/internal/dtoresponse"

type PatchPasswordRequestDto struct {
	CurrentPassword string `json:"current_password"`
	NewPassword     string `json:"new_password"`
}

type PostRequestDto struct {
	EmployeeId string `json:"employee_id"`
	Password   string `json:"password"`
}

type PostResponseDto = dtoresponse.HttpResponseWithoutData
type DeleteResponseDto = dtoresponse.HttpResponseWithoutData
type PatchPasswordResponseDto = dtoresponse.HttpResponseWithoutData
type DeletePasswordResponseDto = dtoresponse.HttpResponseWithoutData
