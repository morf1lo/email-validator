package emailvalidator

import (
	"strings"
	"errors"
	"regexp"
)

var (
	err 	= errors.New("Invalid email")
	regex	= regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
)

func Valid(email string) error {
	if !regex.MatchString(strings.ToLower(email)) {
		return err
	}
	
	return nil
}
