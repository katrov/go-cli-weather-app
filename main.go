package main

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"io/ioutil"
	"net/http"
)

type Response struct {
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
	// https://weatherdbi.herokuapp.com/documentation/v1
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

	if len(PrettyPrint(result.Region)) == 2 {
		panic(city + " invalid city! Please repeat again")
	}

	color.Green("Region - " + PrettyPrint(result.Region))
	color.Green("Precip - " + PrettyPrint(result.CurrentConditions.Precip))
	color.Green("Humidity - " + PrettyPrint(result.CurrentConditions.Humidity))
	color.Green("Temp - " + PrettyPrint(result.CurrentConditions.Temp.C) + "°C")
	color.Green("Wind - " + PrettyPrint(result.CurrentConditions.Wind.Km) + "ms")
}

func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}
