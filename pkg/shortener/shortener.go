package shortener

import (
	uuid "github.com/satori/go.uuid"
)

// Generate - generates string
func Generate() string {
	id := uuid.NewV4()
	return id.String()
}
