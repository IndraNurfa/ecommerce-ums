package models

import "time"

type Health struct {
	Message string    `json:"message"`
	Time    time.Time `json:"time"`
}
