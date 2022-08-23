package lib

import (
	"regexp"
)

func Normalize(number string) string {
	reg := regexp.MustCompile("[^0-9]")
	return reg.ReplaceAllString(number, "")
}
