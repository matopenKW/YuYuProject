package dto

import (
	"time"
)

type Product struct {
	ProductName string     `json:",omitempty"`
	ProductNo   string     `json:",omitempty"`
	Timestamp   *time.Time `json:",omitempty"`
}
