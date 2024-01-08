package session

import (
	"time"
)

type SigningResult struct {
	Token   string    `json:"token"`
	Expires time.Time `json:"expires"`
}
