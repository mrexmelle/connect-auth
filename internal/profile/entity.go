package profile

type Entity struct {
	Ehid         string `json:"ehid"`
	EmployeeId   string `json:"employee_id"`
	Name         string `json:"name"`
	EmailAddress string `json:"email_address"`
	Dob          string `json:"dob"`
}
