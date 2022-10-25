package main

import (
	"assignment3/services"
	"encoding/json"
	"html/template"
	"io/ioutil"
	"net/http"
)

func main() {
	go services.UpdateWeather()
	http.HandleFunc("/", getCuaca)
	http.ListenAndServe(":9000", nil)
}

func getCuaca(w http.ResponseWriter, r *http.Request) {
	dataCuaca, err := ioutil.ReadFile("data.json")
	if err != nil {
		writeJsonResponse(w, http.StatusNotFound, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	statusCuaca := services.StatusCuaca{}
	errUn := json.Unmarshal(dataCuaca, &statusCuaca)
	if errUn != nil {
		writeJsonResponse(w, http.StatusNotFound, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	water := statusCuaca.Status.Water
	wind := statusCuaca.Status.Wind
	var statusWater string
	var statusWind string

	if water <= 5 {
		statusWater = "Aman"
	} else if water >= 6 && water <= 8 {
		statusWater = "Siaga"
	} else {
		statusWater = "Bahaya"
	}

	if wind <= 6 {
		statusWind = "Aman"
	} else if wind >= 7 && wind <= 15 {
		statusWind = "Siaga"
	} else {
		statusWind = "Bahaya"
	}

	resultWeather := services.HasilCuaca{}
	resultWeather.Water = water
	resultWeather.Wind = wind
	resultWeather.StatusWater = statusWater
	resultWeather.StatusWind = statusWind

	tpl, errTmpl := template.ParseFiles("weather.html")
	if errTmpl != nil {
		writeJsonResponse(w, http.StatusNotFound, map[string]interface{}{
			"error": errTmpl.Error(),
		})
		return
	}
	tpl.Execute(w, resultWeather)

}

func writeJsonResponse(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}