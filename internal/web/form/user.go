package form

import (
	"regexp"
	"strings"
	"unicode/utf8"
)

type SignupUser struct {
	Username string
	Email    string
	Password string
	Errors   map[string]string
}

type SigninUser struct {
	Email    string
	Password string
	Errors   map[string]string
}

var rxEmail = regexp.MustCompile(`^[a-zA-Z\\_\\-\\.]+@[a-zA-Z\\_\\-\\.]+$`)

func (f *SignupUser) Valid() bool {

	f.Errors = make(map[string]string)

	if strings.TrimSpace(f.Username) == "" {
		f.Errors["username"] = "username is required."
	}

	if strings.TrimSpace(f.Email) == "" {
		f.Errors["email"] = "email is required."
	} else if len(f.Email) > 254 || !rxEmail.MatchString(f.Email) {
		f.Errors["email"] = "email is not valid."
	}

	if utf8.RuneCountInString(f.Password) < 8 {
		f.Errors["password"] = "password is too short."
	}

	return len(f.Errors) == 0
}

func (f *SigninUser) Valid() bool {

	f.Errors = make(map[string]string)

	if strings.TrimSpace(f.Email) == "" {
		f.Errors["email"] = "email is required"
	}
	if strings.TrimSpace(f.Password) == "" {
		f.Errors["password"] = "password is required"
	}
	println("what", len(strings.TrimSpace(f.Password)))
	if strings.TrimSpace(f.Password) != "" && len(strings.TrimSpace(f.Password)) < 8 {

		f.Errors["password"] = "password length needs to be above 8"
	}

	return len(f.Errors) == 0
}
