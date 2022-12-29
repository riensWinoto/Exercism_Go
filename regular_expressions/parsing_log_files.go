package main

import (
	"fmt"
	"os"
	"regexp"
)

var (
	isValidlog   string   = "[ERR] A good error here"
	splitLogText string   = "section 1<*>section 2<~~~>section 3"
	countPass    []string = []string{
		`[INF] passWord`,
		`"passWord"`,
		`[INF] User saw error message "Unexpected Error" on page load.`,
		`[INF] The message "Please reset your password" was ignored by the user`,
	}
	endOfLineText   string   = "[INF] end-of-line23033 Network Failure end-of-line27"
	tagUsernameText []string = []string{
		"[WRN] User James123 has exceeded storage space.",
		"[WRN] Host down. User   Michelle4 lost connection.",
		"[INF] Users can login again after 23:00.",
		"[DBG] We need to check that user names are at least 6 chars long.",
	}
)

func IsValidLine(text string) bool {
	regx, err := regexp.Compile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)]`)
	if err != nil {
		fmt.Println("IsValidLine Error")
		os.Exit(1)
	}
	return regx.MatchString(text)
}

func SplitLogLine(text string) []string {
	regx, err := regexp.Compile(`<[*~=-]*>`)
	if err != nil {
		fmt.Println("SplitLogLine Error")
		os.Exit(1)
	}
	return regx.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	var counter int = 0
	regx, err := regexp.Compile(`(?i)".*password.*"`)
	if err != nil {
		fmt.Println("CountQuotedPasswords Error")
		os.Exit(1)
	}
	for _, val := range lines {
		if regx.MatchString(val) == true {
			counter += 1
		}
	}
	return counter
}

func RemoveEndOfLineText(text string) string {
	regx, err := regexp.Compile(`end-of-line\d+`)
	if err != nil {
		fmt.Println("CountQuotedPasswords Error")
		os.Exit(1)
	}
	return regx.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	var tagRet []string = []string{}
	regx, err := regexp.Compile(`User\s+(\w+)`)
	if err != nil {
		fmt.Println("TagWithUserName Error")
		os.Exit(1)
	}
	for _, val := range lines {
		founds := regx.FindStringSubmatch(val)
		if founds != nil {
			val = fmt.Sprintf("[USR] %s %s", founds[1], val)
		}
		tagRet = append(tagRet, val)
	}
	return tagRet
}

func main() {
	fmt.Println(IsValidLine(isValidlog))
	fmt.Println(SplitLogLine(splitLogText))
	fmt.Println(CountQuotedPasswords(countPass))
	fmt.Println(RemoveEndOfLineText(endOfLineText))
	fmt.Println(TagWithUserName(tagUsernameText))
}
