package session

import "github.com/mrexmelle/connect-authx/internal/dto"

type PostRequestDto struct {
	EmployeeId string `json:"employee_id"`
	Password   string `json:"password"`
}

type PostResponseDto = dto.HttpResponseWithData[SigningResult]
