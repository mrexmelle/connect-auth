package profile

import "github.com/mrexmelle/connect-authx/internal/dto"

type PatchRequestDto struct {
	Fields map[string]string `json:"fields"`
}

type GetResponseDto = dto.HttpResponseWithData[Entity]
type GetEhidResponseDto = dto.HttpResponseWithData[string]
type PatchResponseDto = dto.HttpResponseWithoutData
type DeleteResponseDto = dto.HttpResponseWithoutData
