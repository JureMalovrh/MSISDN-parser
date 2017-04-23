package parser

import (	
	"testing"
	//"fmt"
)

func TestFindMNOProviders(t *testing.T) {	
	number := "38641"
	mno, exists := "Mobitel", true
	actualMno, actualExists := findMNOProviders(number)

	if actualMno != mno {
		t.Errorf("Expected the MNO provider of %s to be %s but instead got %s!", number, mno, actualMno)
	}
	if actualExists != exists {
		t.Errorf("Expected the exists param of %s to be %b but instead got %b!", number, exists, actualExists)
	}

	number = "141"
	mno, exists = "", false
	actualMno, actualExists = findMNOProviders(number)
	if actualMno != mno {
		t.Errorf("Expected the MNO provider of %s to be %s but instead got %s!", number, mno, actualMno)
	}
	if actualExists != exists {
		t.Errorf("Expected the exists param of %s to be %b but instead got %b!", number, exists, actualExists)
	}

	number = ""
	mno, exists = "", false
	actualMno, actualExists = findMNOProviders(number)
	if actualMno != mno {
		t.Errorf("Expected the MNO provider of %s to be %s but instead got %s!", number, mno, actualMno)
	}
	if actualExists != exists {
		t.Errorf("Expected the exists param of %s to be %b but instead got %b!", number, exists, actualExists)
	}
}


func TestFindCountry(t *testing.T) {	
	number := "386"
	cc, exists := "SI", true
	actualCC, actualExists := findCountry(number)

	if actualCC != cc {
		t.Errorf("Expected the country code of %s to be %s but instead got %s!", number, cc, actualCC)
	}
	if actualExists != exists {
		t.Errorf("Expected the exists param of %s to be %b but instead got %b!", number, exists, actualExists)
	}

	number = "141"
	cc, exists = "", false
	actualCC, actualExists = findCountry(number)
	if actualCC != cc {
		t.Errorf("Expected the country code of %s to be %s but instead got %s!", number, cc, actualCC)
	}
	if actualExists != exists {
		t.Errorf("Expected the exists param of %s to be %b but instead got %b!", number, exists, actualExists)
	}

	number = ""
	cc, exists = "", false
	actualCC, actualExists = findCountry(number)
	if actualCC != cc {
		t.Errorf("Expected the country code of %s to be %s but instead got %s!", number, cc, actualCC)
	}
	if actualExists != exists {
		t.Errorf("Expected the exists param of %s to be %b but instead got %b!", number, exists, actualExists)
	}
}

func TestGetMNOProvider(t *testing.T) {
	number := "38641123123"
	mno, sn := "Mobitel", "123123"
	actualMno, actualSn := GetMNOProvider(number)

	if(actualMno != mno) {
		t.Errorf("Expected the mobile network operator of %s to be %s but instead got %s!", number, mno, actualMno)
	}

	if(actualSn != sn) {
		t.Errorf("Expected the subscriber number of %s to be %s but instead got %s!", number, sn, actualSn)
	}

	number = "123123123123"
	mno, sn = "Unknown provider", ""
	actualMno, actualSn = GetMNOProvider(number)

	if(actualMno != mno) {
		t.Errorf("Expected the mobile network operator of %s to be %s but instead got %s!", number, mno, actualMno)
	}

	if(actualSn != sn) {
		t.Errorf("Expected the subscriber number of %s to be %s but instead got %s!", number, sn, actualSn)
	}

	number = "1202123450"
	mno, sn = "Unknown provider", ""
	actualMno, actualSn = GetMNOProvider(number)

	if(actualMno != mno) {
		t.Errorf("Expected the mobile network operator of %s to be %s but instead got %s!", number, mno, actualMno)
	}

	if(actualSn != sn) {
		t.Errorf("Expected the subscriber number of %s to be %s but instead got %s!", number, sn, actualSn)
	}	

	number = "37497111111"
	mno, sn ="Karabakh Telecom", "111111"
	actualMno, actualSn = GetMNOProvider(number)

	if(actualMno != mno) {
		t.Errorf("Expected the mobile network operator of %s to be %s but instead got %s!", number, mno, actualMno)
	}

	if(actualSn != sn) {
		t.Errorf("Expected the subscriber number of %s to be %s but instead got %s!", number, sn, actualSn)
	}
	number = "37497"
	mno, sn ="Karabakh Telecom", ""
	actualMno, actualSn = GetMNOProvider(number)

	if(actualMno != mno) {
		t.Errorf("Expected the mobile network operator of %s to be %s but instead got %s!", number, mno, actualMno)
	}

	if(actualSn != sn) {
		t.Errorf("Expected the subscriber number of %s to be %s but instead got %s!", number, sn, actualSn)
	}
}

func TestGetCountryFromCountryCode(t *testing.T) {
	number := "38641"
	cc, cnc := "SI", "386"
	actualCC, actualCNC := GetCountryFromCountryCode(number)

	if(actualCC != cc) {
		t.Errorf("Expected the country code of %s to be %s but instead got %s!", number, cc, actualCC)
	}

	if(actualCNC != cnc) {
		t.Errorf("Expected the country number code of %s to be %s but instead got %s!", number, cnc, actualCNC)
	}

	number = "38"
	cc, cnc = "", ""
	actualCC, actualCNC = GetCountryFromCountryCode(number)

	if(actualCC != cc) {
		t.Errorf("Expected the country code of %s to be %s but instead got %s!", number, cc, actualCC)
	}

	if(actualCNC != cnc) {
		t.Errorf("Expected the country number code of %s to be %s but instead got %s!", number, cnc, actualCNC)
	}

	number = "1"
	cc, cnc = "", ""
	actualCC, actualCNC = GetCountryFromCountryCode(number)

	if(actualCC != cc) {
		t.Errorf("Expected the country code of %s to be %s but instead got %s!", number, cc, actualCC)
	}

	if(actualCNC != cnc) {
		t.Errorf("Expected the country number code of %s to be %s but instead got %s!", number, cnc, actualCNC)
	}	

	number = "1202"
	cc, cnc = "US", "1202"
	actualCC, actualCNC = GetCountryFromCountryCode(number)

	if(actualCC != cc) {
		t.Errorf("Expected the country code of %s to be %s but instead got %s!", number, cc, actualCC)
	}

	if(actualCNC != cnc) {
		t.Errorf("Expected the country number code of %s to be %s but instead got %s!", number, cnc, actualCNC)
	}
}