package main

import (
	"detector-demo/api"
	"net/http"
)

func main() {

	http.HandleFunc("/max-speed-allowed", api.GetMaxSpeedAllowed)
	http.HandleFunc("/last-speed", api.GetLastSpeed)

	http.ListenAndServe(":8080", nil)
}
