package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/thedevsaddam/gojsonq"
)

//SpeedLog struct to get last speed
type SpeedLog struct {
	Vehicle interface{} `json:"vehicle"`
	Speed   interface{} `json:"speed"`
}

//GetMaxSpeedAllowed get max speed allowed function
func GetMaxSpeedAllowed(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		fail(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	jq := gojsonq.New().File("./data.json")

	res := jq.Find("max_speed_allowed")

	str, err := res.(string)
	if !err {
		fail(w, "oops something went wrong", http.StatusBadRequest)
		return
	}

	data := struct {
		Speed string `json:"speed"`
	}{Speed: str}

	ok(w, data)
}

//GetLastSpeed get last speed by vehicle id
func GetLastSpeed(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		fail(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")
	if id == "" {
		fail(w, "oops something went wrong", http.StatusBadRequest)
		return
	}

	idInt, errr := strconv.Atoi(id)
	if errr != nil {
		fail(w, "oops something went wrong", http.StatusBadRequest)
		return
	}

	jq := gojsonq.New().File("./data.json")
	res := jq.From("speed_log").Where("vehicle", "=", idInt).SortBy("time", "desc").First()

	if res == nil {
		data := struct {
			Message string `json:"message"`
		}{Message: "Invalid id"}
		ok(w, data)
		return

	}

	jsonbody, errr := json.Marshal(res)
	if errr != nil {
		fail(w, "oops something went wrong", http.StatusBadRequest)
		return
	}

	speedLog := SpeedLog{}
	if errr := json.Unmarshal(jsonbody, &speedLog); errr != nil {
		fail(w, "oops something went wrong", http.StatusBadRequest)
		return
	}
	//for performance the data will be storage in int vvalues, they need to be parsed for string, according to the response match
	speedLog.Vehicle = strconv.Itoa(int(speedLog.Vehicle.(float64)))
	speedLog.Speed = strconv.FormatFloat(speedLog.Speed.(float64), 'f', 2, 64)

	ok(w, speedLog)
}

// fail writes a json response with error msg and status header
func fail(w http.ResponseWriter, msg string, status int) {
	w.Header().Set("Content-Type", "application/json")

	data := struct {
		Error string `json:"error"`
	}{Error: msg}

	resp, _ := json.Marshal(data)
	w.WriteHeader(status)

	fmt.Fprintf(w, string(resp))
}

// ok writes data to response with 200 status
func ok(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")

	if s, ok := data.(string); ok {
		fmt.Fprintf(w, s)
		return
	}

	resp, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fail(w, "oops something evil has happened", 500)
		return
	}

	fmt.Fprintf(w, string(resp))
}
