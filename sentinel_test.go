package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
)

func Example_readTemp() {
	// override tempDev
	tempDev = "testdata/tempDev.txt"
	ts, err := readTemp()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("ts=%0.3f", ts)
	// output: ts=47.774
}

func Example_tempHandler() {
	ts := httptest.NewServer(http.HandlerFunc(tempHandler))
	defer ts.Close()

	// override tempDev
	tempDev = "testdata/tempDev.txt"

	res, err := http.Get(ts.URL)
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(body))
	// output: temperature = 47.8
}
