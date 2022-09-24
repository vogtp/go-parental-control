package user

import (
	"os/user"
	"strings"
)

func RemoveHostPart(usr string) string {
	idx := strings.Index(usr, "\\")
	if idx > 0 {
		usr = usr[idx+1:]
	}
	return usr
}

func Current() string {
	usr, err := user.Current()
	if err != nil {
		return "UNKNOWN"
	}

	return RemoveHostPart(usr.Username)
}
