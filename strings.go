package phpfunc

import "strings"

func AddSlashes(str string) string {
	replace := map[string]string{`\`: `\\`, `"`: `\"`, `'`: `\'`, "`": "\\`"}
	for k, v := range replace {
		str = strings.Replace(str, k, v, -1)
	}

	return str
}

func SubStr(str string, start int, length int) string {
	if start < 0 {
		start = 0
	}

	if length == 0 {
		length = len(str)
	}

	strArr := strings.Split(str, "")
	strArrLength := len(strArr)
	i := 0
	str = ""
	for i = start; i < start+length; i++ {
		if i >= strArrLength {
			break
		}
		str += strArr[i]
	}
	return str
}

func StrPos(haystack string, needle string, offset int) int {
	if offset > 0 {
		haystack = SubStr(haystack, offset, len(haystack))
	}

	return strings.Index(haystack, needle) + offset
}

func Explode(arg string, str string, maxlimit int) []string {
	if maxlimit <= 0 {
		return strings.Split(str, arg)
	}
	return strings.SplitN(str, arg, maxlimit)
}

func Implode(arg string, array []string) string {
	if len(array) == 0 {
		return ""
	}

	return strings.Join(array, arg)
}

func Trim(str string, cutset string) string {
	return strings.Trim(str, " \n\r\t"+cutset)
}

func LTrim(str string) string {
	return strings.TrimLeft(str, " \n\r\t")
}

func RTrim(str string) string {
	return strings.TrimRight(str, " \n\r\t")
}
