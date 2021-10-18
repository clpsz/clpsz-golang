package pkg

import "regexp"

func MatchRe(pattern, str string) bool {
	re := regexp.MustCompile(pattern)

	return re.MatchString(str)
}

// GetReGroupN
// caller should ensure that
// 1. if pattern matches, group index really exists
func GetReGroupN(pattern, str string, index int) string {
	re := regexp.MustCompile(pattern)

	if re.MatchString(str) {
		match := re.FindStringSubmatch(str)
		return match[index]
	} else {
		return ""
	}
}

// GetReGroupN2
// caller should ensure that
// 1. if pattern matches, group i1, i2 really exist
func GetReGroupN2(pattern, str string, i1, i2 int) (string, string) {
	re := regexp.MustCompile(pattern)

	if re.MatchString(str) {
		match := re.FindStringSubmatch(str)
		return match[i1], match[i2]
	} else {
		return "", ""
	}
}

func GetReGroupMany(pattern, str string, indices []int) []string {
	res := make([]string, 0)
	re := regexp.MustCompile(pattern)

	if re.MatchString(str) {
		match := re.FindStringSubmatch(str)
		for _, i := range indices {
			res = append(res, match[i])
		}
	}

	return res
}
