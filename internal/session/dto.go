package session

import "github.com/mrexmelle/connect-authx/internal/dto/dtorespwithdata"

type PostRequestDto struct {
	EmployeeId string `json:"employee_id"`
	Password   string `json:"password"`
}

type PostResponseDto = dtorespwithdata.Class[SigningResult]
