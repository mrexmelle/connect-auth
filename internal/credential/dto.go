package credential

type PatchPasswordRequestDto struct {
	CurrentPassword string `json:"currentPassword"`
	NewPassword     string `json:"newPassword"`
}

type PostRequestDto struct {
	EmployeeId string `json:"employeeId"`
	Password   string `json:"password"`
}

type ResponseDto struct {
	Status string `json:"status"`
}
