package main

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"io/ioutil"
	"net/http"
)

type Response struct {
	Status            string `json:"status"`
	Region            string `json:"region"`
	CurrentConditions struct {
		Precip   string `json:"precip"`
		Humidity string `json:"humidity"`
		Temp     struct {
			C int `json:"c"`
		}
		Wind struct {
			Km int `json:"km"`
		}
	}
}

func main() {
	var city string
	var result Response
	// https://weatherdbi.herokuapp.com/documentation/v1 documentation for weatherdbi
	const WeatherPublicApi = "https://weatherdbi.herokuapp.com/data/weather/"

	color.Green("Weather Shell app")
	color.Yellow("Enter city")
	color.Blue("+++++++++++++++++")

	fmt.Scan(&city)

	resp, err := http.Get(WeatherPublicApi + city)
	if err != nil {
		panic("No response from request!")
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	json.Unmarshal(body, &result)

	if result.Status == "fail" {
		panic(city + " invalid city! Please repeat again")
	}

	color.Green("Region - " + result.Region)
	color.Green("Precip - " + result.CurrentConditions.Precip)
	color.Green("Humidity - " + result.CurrentConditions.Humidity)
	//color.Green("Temp - " + strconv.Itoa(result.CurrentConditions.Temp.C) + "°C")
	color.Green(fmt.Sprintf("Temp -  %d°C", result.CurrentConditions.Temp.C))
	color.Green(fmt.Sprintf("Wind -  %dms", result.CurrentConditions.Wind.Km))
}
