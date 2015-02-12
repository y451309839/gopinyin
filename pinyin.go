package gopinyin

import "regexp"

var pyLength = len(pyValue)
var spLength = len(spValue)
var hzRegexp = regexp.MustCompile("^[\u4e00-\u9fa5]$")

func Code(s string) int {
	gbkString := UTF8ToGBK(s)
	var i1, i2 int
	i1 = int(gbkString[0])
	i2 = int(gbkString[1])
	return i1*256 + i2 - 65536
}

func ConvertPY(s string) string {
	pyString := ""
	var str string
	var code int

	for _, rune := range s {
		str = string(rune)
		if hzRegexp.MatchString(str) { //chinese

			code = Code(str)

			if code > 0 && code < 160 {
				pyString += str
			} else {

				if v, ok := pyTable[code]; ok { //map by table
					pyString += v
				} else {

					for i := (pyLength - 1); i >= 0; i-- {

						if pyValue[i] <= code {
							pyString += pyName[i]
							break
						}
					}
				}
			}
		} else { //other
			pyString += str
		}
	}
	return pyString
}

func ConvertSP(s string) string {
	spString := ""
	var str string
	var code int

	for _, rune := range s {
		str = string(rune)
		if hzRegexp.MatchString(str) {
			code = Code(str)

			if code > 0 && code < 160 {
				spString += str
			} else {

				if v, ok := spTable[code]; ok {
					spString += v
				} else {

					for i := (spLength - 1); i >= 0; i-- {
						if spValue[i] <= code {
							spString += spName[i]
							break
						}
					}
				}
			}
		} else {
			spString += str
		}
	}
	return spString
}
