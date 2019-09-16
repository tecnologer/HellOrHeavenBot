package resources

import (
	"math/rand"
	"regexp"
	"time"
)

var rgAtSignPrefix *regexp.Regexp

func init() {
	rgAtSignPrefix, _ = regexp.Compile("^@*")
}

//LeftTrimAtSign removes the "at signs" (@) on the left
func LeftTrimAtSign(text string) string {
	return string(rgAtSignPrefix.ReplaceAll([]byte(text), []byte{}))
}

//GetRandomIntFromRange returns a random integer between the number provided
func GetRandomIntFromRange(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}
