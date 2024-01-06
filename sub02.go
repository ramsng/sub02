package sub02

import "strings"

const Desc string = "** Lovely compute **"
const Desc1 string = "** Same word in Lowercase **"

func Lowerconv(s string) string {
	return "Desc1" + strings.ToLower(s)
}
