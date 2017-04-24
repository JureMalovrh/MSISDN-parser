package parser

import (	
	"testing"
	//"fmt"
)

func TestSanitize(t *testing.T) {
	// trailing zeros
	number := "0038640400400"
	sanitized := "38640400400"
	actualSanitized := sanitize(number)

	if actualSanitized != sanitized {
		t.Errorf("Expected the sanitized string of %s to be %s but instead got %s!", number, sanitized, actualSanitized)
	}	
	// trailing +
	number = "+38640400400"
	sanitized = "38640400400"
	actualSanitized = sanitize(number)

	if actualSanitized != sanitized {
		t.Errorf("Expected the sanitized string of %s to be %s but instead got %s!", number, sanitized, actualSanitized)
	}
	// ()
	number = "+(386)40400400"
	sanitized = "38640400400"
	actualSanitized = sanitize(number)

	if actualSanitized != sanitized {
		t.Errorf("Expected the sanitized string of %s to be %s but instead got %s!", number, sanitized, actualSanitized)
	}

	// -
	number = "+38640-400-400"
	sanitized = "38640400400"
	actualSanitized = sanitize(number)

	if actualSanitized != sanitized {
		t.Errorf("Expected the sanitized string of %s to be %s but instead got %s!", number, sanitized, actualSanitized)
	}

	// " " and "	"
	number = "+38640 400	400"
	sanitized = "38640400400"
	actualSanitized = sanitize(number)

	if actualSanitized != sanitized {
		t.Errorf("Expected the sanitized string of %s to be %s but instead got %s!", number, sanitized, actualSanitized)
	}

	//all together
	number = "000(386) 40 - 400	-400"
	sanitized = "38640400400"
	actualSanitized = sanitize(number)

	if actualSanitized != sanitized {
		t.Errorf("Expected the sanitized string of %s to be %s but instead got %s!", number, sanitized, actualSanitized)
	}

}

func TestParseMsisdn(t *testing.T) {
	// normal number (all data)

	number := "0038640400400"
	parsed := Msisdn{CDC: "386", SN: "400400", CI: "SI", MNO: "Si.mobil"}
	actualParsed := ParseMsisdn(number)

	if actualParsed.CDC != parsed.CDC {
		t.Errorf("Expected the parsed string MSISDN, CDC of %s to be %s but instead got %s!", number, parsed.CDC, actualParsed.CDC)
	}	
	
	if actualParsed.SN != parsed.SN {
		t.Errorf("Expected the parsed string MSISDN, SN of %s to be %s but instead got %s!", number, parsed.SN, actualParsed.SN)
	}	

	if actualParsed.CI != parsed.CI {
		t.Errorf("Expected the parsed string MSISDN, CI of %s to be %s but instead got %s!", number, parsed.CI, actualParsed.CI)
	}	

	if actualParsed.MNO != parsed.MNO {
		t.Errorf("Expected the parsed string MSISDN, MNO of %s to be %s but instead got %s!", number, parsed.MNO, actualParsed.MNO)
	}	
	// too long number
	number = "0038640400400123123123"
	parsed = Msisdn{ERR: "Number too long"}
	actualParsed = ParseMsisdn(number)

	if actualParsed.ERR != parsed.ERR {
		t.Errorf("Expected the parsed string MSISDN, ERR of %s to be %s but instead got %s!", number, parsed.ERR, actualParsed.ERR)
	}	
	// too short number
	number = "0038640"
	parsed = Msisdn{ERR: "Number too short"}
	actualParsed = ParseMsisdn(number)

	if actualParsed.ERR != parsed.ERR {
		t.Errorf("Expected the parsed string MSISDN, ERR of %s to be %s but instead got %s!", number, parsed.ERR, actualParsed.ERR)
	}
	// country code not recognized
	number = "1999123123123"
	parsed = Msisdn{ERR: "Country code not recognized"}
	actualParsed = ParseMsisdn(number)

	if actualParsed.ERR != parsed.ERR {
		t.Errorf("Expected the parsed string MSISDN, ERR of %s to be %s but instead got %s!", number, parsed.ERR, actualParsed.ERR)
	}		
	// number does not contains only numbers
	number = "199a9123123123a"
	parsed = Msisdn{ERR: "Number does not contain only numbers"}
	actualParsed = ParseMsisdn(number)

	if actualParsed.ERR != parsed.ERR {
		t.Errorf("Expected the parsed string MSISDN, ERR of %s to be %s but instead got %s!", number, parsed.ERR, actualParsed.ERR)
	}
	
	// MNO unknown -> SN should be everything after CC
	number = "1212700123123"
	parsed = Msisdn{CDC: "1212", SN: "700123123", CI: "US", MNO: "Unknown provider"}
	actualParsed = ParseMsisdn(number)

	if actualParsed.CDC != parsed.CDC {
		t.Errorf("Expected the parsed string MSISDN, CDC of %s to be %s but instead got %s!", number, parsed.CDC, actualParsed.CDC)
	}	
	
	if actualParsed.SN != parsed.SN {
		t.Errorf("Expected the parsed string MSISDN, SN of %s to be %s but instead got %s!", number, parsed.SN, actualParsed.SN)
	}	

	if actualParsed.CI != parsed.CI {
		t.Errorf("Expected the parsed string MSISDN, CI of %s to be %s but instead got %s!", number, parsed.CI, actualParsed.CI)
	}	

	if actualParsed.MNO != parsed.MNO {
		t.Errorf("Expected the parsed string MSISDN, MNO of %s to be %s but instead got %s!", number, parsed.MNO, actualParsed.MNO)
	}		

	number = "973663123123"
	parsed = Msisdn{CDC: "973", SN: "123123", CI: "BA", MNO: "Zain"}
	actualParsed = ParseMsisdn(number)

	if actualParsed.CDC != parsed.CDC {
		t.Errorf("Expected the parsed string MSISDN, CDC of %s to be %s but instead got %s!", number, parsed.CDC, actualParsed.CDC)
	}	
	
	if actualParsed.SN != parsed.SN {
		t.Errorf("Expected the parsed string MSISDN, SN of %s to be %s but instead got %s!", number, parsed.SN, actualParsed.SN)
	}	

	if actualParsed.CI != parsed.CI {
		t.Errorf("Expected the parsed string MSISDN, CI of %s to be %s but instead got %s!", number, parsed.CI, actualParsed.CI)
	}	

	if actualParsed.MNO != parsed.MNO {
		t.Errorf("Expected the parsed string MSISDN, MNO of %s to be %s but instead got %s!", number, parsed.MNO, actualParsed.MNO)
	}	

}

func TestReturnString(t *testing.T) {
	parsed := Msisdn{CDC: "973", SN: "123123", CI: "BA", MNO: "Zain"}
	returnedString := "cdc:973, sn:123123, ci:BA, mno:Zain, err: "
	actualReturnedString := ReturnString(parsed)
	if returnedString != actualReturnedString {
		t.Errorf("Expected the string to be %s but instead got %s!",  returnedString, actualReturnedString)
	}	
}