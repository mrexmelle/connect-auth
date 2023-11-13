package session

type RequestDto struct {
	EmployeeId string `json:"employeeId"`
	Password   string `json:"password"`
}

type ResponseDto struct {
	Token string `json:"token"`
}
