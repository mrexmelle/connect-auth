package mapper

import (
	"crypto/sha256"
	"fmt"
)

func ToEhid(employeeId string) string {
	hasher := sha256.New()
	hasher.Write([]byte(employeeId))

	return fmt.Sprintf("u%x", hasher.Sum(nil))
}

func ToStatus(err error) string {
	if err != nil {
		return err.Error()
	}
	return "OK"
}
