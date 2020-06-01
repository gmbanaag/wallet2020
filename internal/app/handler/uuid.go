package handler

import (
	uuid "github.com/satori/go.uuid"
)

func generateUUID() string {
	id := uuid.NewV4()

	return id.String()
}
