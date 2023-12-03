package session

type RequestDto struct {
	EmployeeId string `json:"employee_id"`
	Password   string `json:"password"`
}

type ResponseDto struct {
	Token string `json:"token"`
}
