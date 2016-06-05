package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

var (
	tempDev = "/sys/class/thermal/thermal_zone0/temp"
	port    = os.Getenv("PORT")
)

func readTemp() (float32, error) {
	f, err := os.Open(tempDev)
	defer f.Close()
	if err != nil {
		return 0, err
	}

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return 0, err
	}

	tt, err := strconv.Atoi(string(b[:len(b)-1]))
	if err != nil {
		return 0, err
	}

	return float32(tt) / 1000.0, nil
}

func tempHandler(w http.ResponseWriter, r *http.Request) {
	t, err := readTemp()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "temperature = %0.1f", t)
}

func main() {
	log.SetFlags(0)
	log.Println("listening on http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, http.HandlerFunc(tempHandler)))
}
