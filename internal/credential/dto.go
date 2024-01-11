package credential

import "github.com/mrexmelle/connect-authx/internal/dto"

type PatchPasswordRequestDto struct {
	CurrentPassword string `json:"current_password"`
	NewPassword     string `json:"new_password"`
}

type PostRequestDto struct {
	EmployeeId string `json:"employee_id"`
	Password   string `json:"password"`
}

type PostResponseDto = dto.HttpResponseWithoutData
type DeleteResponseDto = dto.HttpResponseWithoutData
type PatchPasswordResponseDto = dto.HttpResponseWithoutData
type DeletePasswordResponseDto = dto.HttpResponseWithoutData
