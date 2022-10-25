package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"time"
)

type StatusCuaca struct {
	Status Cuaca
}

type Cuaca struct {
	Water int
	Wind  int
}

type HasilCuaca struct {
	Water       int
	Wind        int
	StatusWater string
	StatusWind  string
}

func UpdateWeather() {
	max := 100
	min := 1

	for {
		fmt.Println("Running Update Cuaca setiap 15 detik")
		rand.Seed(time.Now().UnixNano())
		statusCuaca := StatusCuaca{}
		statusCuaca.Status.Water = rand.Intn(max-min) + min
		statusCuaca.Status.Wind = rand.Intn(max-min) + min

		dataCuaca, err := json.Marshal(statusCuaca)
		if err != nil {
			fmt.Println(err)
		}

		errs := ioutil.WriteFile("data.json", dataCuaca, 0644)
		if errs != nil {
			fmt.Println(errs)
		}
		time.Sleep(15 * time.Second)

	}
}