package credential

import (
	"github.com/mrexmelle/connect-authx/internal/dto/dtorespwithoutdata"
)

type PatchPasswordRequestDto struct {
	CurrentPassword string `json:"current_password"`
	NewPassword     string `json:"new_password"`
}

type PostRequestDto struct {
	EmployeeId string `json:"employee_id"`
	Password   string `json:"password"`
}

type PostResponseDto = dtorespwithoutdata.Class
type DeleteResponseDto = dtorespwithoutdata.Class
type PatchPasswordResponseDto = dtorespwithoutdata.Class
type DeletePasswordResponseDto = dtorespwithoutdata.Class
