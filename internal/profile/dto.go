package profile

import (
	"github.com/mrexmelle/connect-authx/internal/dto/dtorespwithdata"
	"github.com/mrexmelle/connect-authx/internal/dto/dtorespwithoutdata"
)

type PatchRequestDto struct {
	Fields map[string]string `json:"fields"`
}

type GetResponseDto = dtorespwithdata.Class[Entity]
type GetEhidResponseDto = dtorespwithdata.Class[string]
type PatchResponseDto = dtorespwithoutdata.Class
type DeleteResponseDto = dtorespwithoutdata.Class
