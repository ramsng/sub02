package sub02

import "strings"

const Desc2 string = "** Lovely compute **"
const Desc1 string = "** Same word in Lowercase **"

func Lowerconv(s string) string {
	return Desc1 + strings.ToLower(s)
}

func Desc() string {
	return Desc2
}
