package parser

import (
	"fmt"
	"strings"
)


func main() {
	fmt.Println("Hello, world")
}

type Msisdn struct {
	CDC string
	SN string
	CI string
	MNO string
}

func MsisdnParse(msisdnString string) Msisdn {
	msisdnString = sanitize(msisdnString)
	
}


/*
	Sanitizes the received MSISDN string by removing whitespace characters, 00 and + prefix and remove all ( and )
*/
func sanitize(msisdnString string) string {
	// remove trailing 00
	var sanitizedString string;
	sanitizedString = strings.TrimPrefix(msisdnString, "00")
	// remove trailing +
	sanitizedString = strings.TrimPrefix(sanitizedString, "+")

	// remove ( and )
	sanitizedString = strings.Replace(sanitizedString, "(", "", -1)
	sanitizedString = strings.Replace(sanitizedString, ")", "", -1)

	// remove whitespace characters
	sanitizedString = strings.Replace(sanitizedString, " ", "", -1)
	sanitizedString = strings.Replace(sanitizedString, "	", "", -1)

	return sanitizedString;
}
