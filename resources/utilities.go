package resources

import (
	"fmt"
	"hash/fnv"
	"math/rand"
	"regexp"
	"time"

	bot "github.com/yanzay/tbot"
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

//GetName extrats the username of the user who send the message
func GetName(msg *bot.Message) string {
	if msg.From.Username != "" {
		return msg.From.Username
	}

	return fmt.Sprintf("%s %s", msg.From.FirstName, msg.From.LastName)
}

//GetHash create a hash value from values
func GetHash(elements ...interface{}) string {
	return hash(fmt.Sprint(elements...))
}

func hash(s string) string {
	h := fnv.New32a()
	h.Write([]byte(s))
	return string(h.Sum32())
}
