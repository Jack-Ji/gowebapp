package utils

import (
	"errors"
	"github.com/satori/go.uuid"
	"strings"
)

func GenerateUUID() string {
	u := uuid.NewV4()
	//fmt.Printf("UUIDv4: %s\n", u)
	return u.String()
}

func GetServiceUUID() (string, error) {
	s := GenerateUUID()
	if s == "" {
		return "", errors.New("uuid generator failed")
	}
	return strings.Replace(s, "-", "", -1), nil
}
