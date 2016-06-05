package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func readTemp() (string, error) {
	f, err := os.Open("/sys/class/thermal/thermal_zone0/temp")
	defer f.Close()
	if err != nil {
		return "", err
	}

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func main() {
	ts, err := readTemp()
	if err != nil {
		fmt.Println("err")
		os.Exit(-1)
	}
	fmt.Printf("ts=%q", ts)
}
