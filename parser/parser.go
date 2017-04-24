package parser

import (
	"fmt"
	"strings"
	"regexp"
)


type Msisdn struct {
	CDC string
	SN string
	CI string
	MNO string
	ERR string
}

func ParseMsisdn(msisdnString string) Msisdn {
	msisdnString = sanitize(msisdnString)
	if(len(msisdnString) < 8) {
		return Msisdn{ERR: "Number too short"}
	}

	if(len(msisdnString) > 15) {
		return Msisdn{ERR: "Number too long"}
	}

	var re = regexp.MustCompile(`^\d+$`)

	var testNumbers = re.MatchString(msisdnString)
	if(!testNumbers) {
		return Msisdn{ERR: "Number does not contain only numbers"}
	}


	country, countryCode := GetCountryFromCountryCode(msisdnString)

	if(country == "" || countryCode == "") {
		return Msisdn{ERR: "Country code not recognized"}
	}
	mnoProvider, subscriberNumber := GetMNOProvider(msisdnString)

	// if subscriber number is empty string -> no MNO has been found, so we return subscriber number only without CNC
	if subscriberNumber == "" {
		subscriberNumber = msisdnString[len(countryCode):]
	}

	return Msisdn{CDC: countryCode, SN: subscriberNumber, CI: country, MNO: mnoProvider}
}


/*
	Sanitizes the received MSISDN string by removing whitespace characters, trailing 0 and + prefix and remove all (, ) and -
*/
func sanitize(msisdnString string) string {
	// remove trailing 00
	var sanitizedString string;
	sanitizedString = strings.TrimLeft(msisdnString, "0")
	// remove trailing +
	sanitizedString = strings.TrimPrefix(sanitizedString, "+")

	// remove (, ) and -
	sanitizedString = strings.Replace(sanitizedString, "(", "", -1)
	sanitizedString = strings.Replace(sanitizedString, ")", "", -1)
	sanitizedString = strings.Replace(sanitizedString, "-", "", -1)

	// remove whitespace characters
	sanitizedString = strings.Replace(sanitizedString, " ", "", -1)
	sanitizedString = strings.Replace(sanitizedString, "	", "", -1)

	return sanitizedString;
}

func ReturnString(msisdn Msisdn) string {
	if(msisdn.ERR != "") {
		return fmt.Sprintf("cdc: %s, sn: %s, ci: %s, mno: %s, err: %s", msisdn.CDC, msisdn.SN, msisdn.CI, msisdn.MNO, msisdn.ERR)
	} else {
		return fmt.Sprintf("cdc: %s, sn: %s, ci: %s, mno: %s", msisdn.CDC, msisdn.SN, msisdn.CI, msisdn.MNO)
	}
} 

