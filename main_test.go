package main

import (
	"testing"
)

func Test_ipValid(t *testing.T) {

	expectedResult := true
	actualResult := Valid("54.192.0.0")
	failedResult := Valid("123.456.78")

	if actualResult != expectedResult {
		t.Error("FAILED. IS NOT A VALID IP.")
	}
	if failedResult == expectedResult {
		t.Error("FAILED. SHOULD HAVE NOT BEEN VALID IP.")
	}

}

func Test_getGeoIpInfo(t *testing.T) {

	expectedResult := "Amazon.com, Inc."
	actualResult := getIpInfo("54.192.0.0")

	if actualResult.ISP != expectedResult {
		t.Error("FAILED. DIDN'T RETURN EXPECTED ISP. PERHAPS USED WRONG IP ADDRESS")
	}

}

func Test_buildResponse(t *testing.T) {

	expectedResult := getIpInfo("54.192.0.0")
	actualResult := jsonResponse("54.192.0.0")

	if actualResult != expectedResult {
		t.Error("FAILED SHOULD HAVE MATCHING INFO")
	}

}
