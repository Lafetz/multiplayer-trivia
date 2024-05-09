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

var rxEmail = regexp.MustCompile(`^[a-zA-Z-0-9\\_\\-\\.]+@[a-zA-Z\\_\\-\\.]+$`)
var rxUsername = regexp.MustCompile(`^[a-zA-Z-0-9\\_\\-\\.]`)

func (f *SignupUser) Valid() bool {

	f.Errors = make(map[string]string)

	if strings.TrimSpace(f.Username) == "" {
		f.Errors["username"] = "username is required."
	}
	if utf8.RuneCountInString(f.Username) > 10 || utf8.RuneCountInString(f.Username) < 3 {
		f.Errors["username"] = "username length must be between 3 and 10"
	}
	if !rxUsername.MatchString(f.Username) {
		f.Errors["username"] = "username be characters must be letters"
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
func (f *SignupUser) AddError(errTarget string, errMsg string) {
	f.Errors[errTarget] = errMsg
}
func (f *SigninUser) Valid() bool {

	f.Errors = make(map[string]string)

	if strings.TrimSpace(f.Email) == "" {
		f.Errors["email"] = "email is required"
	}
	if strings.TrimSpace(f.Password) == "" {
		f.Errors["password"] = "password is required"
	}

	if utf8.RuneCountInString(f.Password) < 8 {
		f.Errors["password"] = "password is too short."
	}

	return len(f.Errors) == 0
}
