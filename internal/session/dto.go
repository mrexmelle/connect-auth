package session

import "github.com/mrexmelle/connect-authx/internal/dtoresponse"

type PostRequestDto struct {
	EmployeeId string `json:"employee_id"`
	Password   string `json:"password"`
}

type PostResponseDto = dtoresponse.HttpResponseWithData[SigningResult]
