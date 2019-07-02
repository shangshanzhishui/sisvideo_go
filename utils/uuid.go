package utils

import "github.com/satori/go.uuid"

func CreateUUID() string{
	u := uuid.Must(uuid.NewV4()).String()
	return u
}
