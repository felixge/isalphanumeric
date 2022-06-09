package isalphanumeric

import "regexp"

var matcher = regexp.MustCompile("^[a-zA-Z0-9]+$")

func IsAlphaNumericRegex(s string) bool {
	return matcher.MatchString(s)
}

func IsAlphaNumericLoop(s string) bool {
	for i := 0; i < len(s); i++ {
		if !isAlphaNumeric(s[i]) {
			return false
		}
	}
	return true
}

func isAlphaNumeric(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
}

var lookupTable [128]byte

func init() {
	for i := 0; i < len(lookupTable); i++ {
		if c := byte(i); !isAlphaNumeric(c) {
			lookupTable[i] = 1
		}
	}
}

func IsAlphaNumericSIMD(s string) bool
