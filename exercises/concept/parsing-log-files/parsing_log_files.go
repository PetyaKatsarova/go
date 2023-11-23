// package parsinglogfiles

package main

import (
	"fmt"
	"regexp"
)

// import (
//     "bytes"
//     "fmt"
//     "regexp"
// )


func IsValidLine(text string) bool {
	str := `^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`
	isMatching, err := regexp.MatchString(str, text)
	if err != nil {
		panic("upala: error with regex"+str)
	}
	return isMatching
}

func SplitLogLine(text string) []string {
	str := `<[~*=-]>`
	re := regexp.MustCompile(str)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	str := `(?i).*".*[password].*".*`
	count := 0
	for _, v := range lines {
		isMatching, err := regexp.MatchString(str, v)
		fmt.Println(isMatching)
		if err != nil {	panic("upala: error with regex"+str) }
		count++
	}
	return count
}

func main() {
	lines := []string{
		`[INF] passWord`, // contains 'password' but not surrounded by quotation marks
		`"passWord"`,  // count this one
		`[INF] User saw error message "Unexpected Error" on page load.`, // does not contain 'password'
		`[INF] The message "Please reset your password" was ignored by the user`, // count this one 
	}
	CountQuotedPasswords(lines)
}

func RemoveEndOfLineText(text string) string {
	panic("Please implement the RemoveEndOfLineText function")
}

func TagWithUserName(lines []string) []string {
	panic("Please implement the TagWithUserName function")
}

/*
Zero or One Occurrence (?):

The ? symbol is used to indicate that the preceding character or group is optional,
 meaning it may occur zero or one time.Example: colou?r matches both "color" and
  "colour".Zero or More Occurrences (*):
The * symbol means the preceding character or group can appear zero or more times.
Example: ab*c matches "ac", "abc", "abbc", "abbbc", and so on.
One or More Occurrences (+):
The + symbol indicates that the preceding character or group must appear at least
 once (one or more times).Example: ab+c matches "abc", "abbc", "abbbc", and so on,
  but not "ac".
*/

// func main() {

    // match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
    // fmt.Println(match)

    // r, _ := regexp.Compile("p([a-z]+)ch")

    // fmt.Println(r.MatchString("peach"))

    // fmt.Println(r.FindString("peach punch"))

    // fmt.Println("idx:", r.FindStringIndex("peach punch"))

    // fmt.Println(r.FindStringSubmatch("peach punch"))

    // fmt.Println(r.FindStringSubmatchIndex("peach punch"))

    // fmt.Println(r.FindAllString("peach punch pinch", -1))

    // fmt.Println("all:", r.FindAllStringSubmatchIndex(
    //     "peach punch pinch", -1))

    // fmt.Println(r.FindAllString("peach punch pinch", 2))

    // fmt.Println(r.Match([]byte("peach")))

    // r = regexp.MustCompile("p([a-z]+)ch")
    // fmt.Println("regexp:", r)

    // fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

    // in := []byte("a peach")
    // out := r.ReplaceAllFunc(in, bytes.ToUpper)
    // fmt.Println(string(out))
// }
