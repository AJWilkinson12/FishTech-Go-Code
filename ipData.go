package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func jsonResponse(ip string) ipInfoStruct {

	ipInfoStruct := getIpInfo(ip)

	return ipInfoStruct

}

func getIpInfo(ip string) ipInfoStruct {

	var ipInfoStruct ipInfoStruct
	if Valid(ip) {
		response, err := http.Get("https://json.geoiplookup.io/" + ip)

		if err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}

		resultData, err := ioutil.ReadAll(response.Body)

		if err != nil {
			log.Fatal(err)
		}

		json.Unmarshal(resultData, &ipInfoStruct)
	}
	return ipInfoStruct
}
