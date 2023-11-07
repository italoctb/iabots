package helper

import (
	"fmt"
	"regexp"
	"strings"

	"go.mau.fi/whatsmeow/types"
)

func ParseJID(arg string) (types.JID, bool) {
	if arg[0] == '+' {
		arg = arg[1:]
	}

	tel := strings.Split(arg, "@")[0]
	if len(tel) < 12 {
		return types.NewJID(arg, types.DefaultUserServer), false
	}

	if !strings.ContainsRune(arg, '@') {
		fmt.Printf("Invalid JID: %s", arg)
		return types.NewJID(arg, types.DefaultUserServer), true
	} else {
		recipient, err := types.ParseJID(arg)
		if err != nil {
			fmt.Printf("Invalid JID: %s", arg)
			return recipient, false
		} else if recipient.User == "" {
			fmt.Printf("Invalid JID: %s", arg)
			return recipient, false
		}
		return recipient, true
	}
}

func ExtractPhoneFromVcard(vcard string) string {
	pattern := `TEL;.*:\+?(\d[\d\- ]+\d)`

	// compile regex pattern
	regex := regexp.MustCompile(pattern)

	// find all matches in vCard string
	matches := regex.FindAllStringSubmatch(vcard, -1)

	// extract phone numbers from matches
	numbers := []string{}
	for _, match := range matches {
		//remove all non-numeric characters
		reg, err := regexp.Compile("[^0-9]+")
		if err != nil {
			fmt.Println(err)
		}
		match[1] = reg.ReplaceAllString(match[1], "")
		numbers = append(numbers, match[1])
	}

	return strings.Join(numbers, ",")

}

func ExtractLinkFromFile(link string) string {
	// has http
	if !strings.Contains(link, "http") {
		return link
	}

	link = strings.ReplaceAll(link, "<", "")
	link = strings.ReplaceAll(link, ">", "")
	return strings.Split(link, "|")[0]
}
