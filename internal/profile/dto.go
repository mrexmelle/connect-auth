package profile

import "github.com/mrexmelle/connect-authx/internal/dtoresponse"

type PatchRequestDto struct {
	Fields map[string]string `json:"fields"`
}

type GetResponseDto = dtoresponse.HttpResponseWithData[Entity]
type GetEmployeeIdResponseDto = dtoresponse.HttpResponseWithData[string]
type PatchResponseDto = dtoresponse.HttpResponseWithoutData
type DeleteResponseDto = dtoresponse.HttpResponseWithoutData
