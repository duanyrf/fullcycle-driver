package main

import (
	"encoding/json"
	"fullcycle_driver/entity"
	"strings"
	"testing"
)

func DriversTest(t *testing.T) {
	result := LoadDrivers("drivers.json")
	var drivers entity.Drivers
	json.Unmarshal(result, &drivers)

	for _, d := range drivers.Drivers {
		if !(strings.Contains(d.Name, "Duany")) {
			t.Errorf("Test fail: %s", result)
		}
	}
}
