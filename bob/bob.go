// This package handles communication with Bob
package bob

import "strings"

// Hey returns Bobs reply to any remark
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)
	switch {
	case remark == "":
		return "Fine. Be that way!"
	case strings.HasSuffix(remark, "?"):
		if strings.ToUpper(remark) == remark && remark != strings.ToLower(remark) {
			return "Calm down, I know what I'm doing!"
		}
		return "Sure."
	case remark == strings.ToUpper(remark) && remark != strings.ToLower(remark):
		return "Whoa, chill out!"
	default:
		return "Whatever."
	}
	return ""
}
