package models

import "time"

type Image struct {
	Name      string
	Data      []byte
	CreatedAt *time.Time
	UpdatedAt *time.Time
}
