package pkg

import "regexp"

func MatchRe(pattern, str string) bool {
	re := regexp.MustCompile(pattern)

	return re.MatchString(str)
}

func GetReGroupN(pattern, str string, index int) string {
	re := regexp.MustCompile(pattern)

	if re.MatchString(str) {
		match := re.FindStringSubmatch(str)
		return match[index]
	} else {
		return ""
	}
}
