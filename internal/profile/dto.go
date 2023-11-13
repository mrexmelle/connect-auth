package profile

type ResponseDto struct {
	Profile Entity `json:"profile"`
	Status  string `json:"status"`
}

type PatchRequestDto struct {
	Fields map[string]string `json:"fields"`
}

type PatchResponseDto struct {
	Status string `json:"status"`
}

type DeleteResponseDto struct {
	Status string `json:"status"`
}
