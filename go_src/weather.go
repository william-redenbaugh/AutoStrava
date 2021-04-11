package main

// Import our libraries
import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/buger/jsonparser"
)

type WeatherHour struct {
	Time          string
	Temp          float32
	Humidity      float32
	PerceivedTemp float32
	WindSpeed     float32
	WindDirection float32
	UVIndex       float32
}

func write_weather_file() []byte {
	resp, err := http.Get("https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline/4512%20Middle%20Park%20Dr%2C%20San%20Jose%20CA?unitGroup=us&key=WFKDKA26EJQE8RHVEWLP7UV9K")
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	file_handler, err := os.Create("latest_weather.json")

	if err != nil {
		log.Fatal(err)
	}

	// Writes JSON data as character array to file handler. Closes file descriptor
	file_handler.Write(body)
	file_handler.Close()

	// Return contents of JSON to be used later in program
	return body
}

func retrieve_weather_file() {

}

/*!
@brief Function that get's all the weather data and saves it
*/
func get_weather_data() {
	body := write_weather_file()
	days, err := jsonparser.GetFloat(body, "days", "[0]", "datetimeEpoch")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(days)
}
