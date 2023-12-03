package credential

type PatchPasswordRequestDto struct {
	CurrentPassword string `json:"current_password"`
	NewPassword     string `json:"new_password"`
}

type PostRequestDto struct {
	EmployeeId string `json:"employee_id"`
	Password   string `json:"password"`
}

type ResponseDto struct {
	Status string `json:"status"`
}
