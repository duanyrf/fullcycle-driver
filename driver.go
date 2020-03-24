package main

import (
	"encoding/json"
	"fullcycle_driver/entity"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"os"
)

// This function load the list of drivers from a JSON file.
func LoadDrivers(file string) []byte {
	jsonFile, err := os.Open(file)
	if err != nil {
		panic(err.Error())
	}

	defer jsonFile.Close()

	data, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		panic(err.Error())
	}

	return data
}

// This function send the list of all registered drivers to client page.
func ListDrivers(w http.ResponseWriter, r *http.Request) {
	drivers := LoadDrivers("drivers.json")
	w.Write([]byte(drivers))
}

// This function send the select driver (by id) to the cliente page.
func GetDriverByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	data := LoadDrivers("drivers.json")

	var drivers entity.Drivers
	json.Unmarshal(data, &drivers)

	for _, v := range drivers.Drivers {
		if v.Uuid == vars["id"] {
			driver, _ := json.Marshal(v)
			w.Write([]byte(driver))
			break
		}
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/drivers", ListDrivers)
	r.HandleFunc("/drivers/{id}", GetDriverByID)
	http.ListenAndServe(":8081", r)
}
