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

/*
	@brief Contains all of our weather information
*/
type WeatherHour struct {
	Time           string
	Temp           float32
	Humidity       float32
	PerceivedTemp  float32
	WindSpeed      float32
	WindDirection  float32
	UVIndex        float32
	RainLikelyHood float32
	Rain           float32
}

/*
	@brief Contains the current date, and a list of all the weather over the course of the day
*/
type WeatherDay struct {
	Date             string
	WeatherHourArray [24]WeatherHour
}

/*
	@brief Generated Type that is essentially a container for all the weather data that can be contained
*/
type WeatherTwoWeek [14]WeatherDay

/*
	@brief Check and flags errors.
*/
func check_err(err error) {
	if err != nil {
		log.Fatal(err)
	}
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

/*
	@brief Gets the array data from the weather file
*/
func retrieve_weather_file() []byte {
	dat, err := ioutil.ReadFile("latest_weather.json")
	if err != nil {
		return nil
	}

	return dat
}

/*
	@brief Parses out the day weather data and saves.
*/
func get_weather_day(data []byte) WeatherDay {
	var current_weather_day WeatherDay

	// Parse out the current date.
	date, err := jsonparser.GetString(data, "datetime")
	current_weather_day.Date = date
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 1; i++ {
		indexStr := "[" + string(i) + "]"
		hour, _, _, err := jsonparser.Get(data, "hours", indexStr)
		check_err(err)
		current_weather_day.WeatherHourArray[i] = get_weather_hour(hour)

	}

	return current_weather_day
}

/*
	@brief Get the current weather data by hour.
*/
func get_weather_hour(data []byte) WeatherHour {
	var current_weather_hour WeatherHour
	time, err := jsonparser.GetString(data, "datetime")

	// The current time of the hour
	current_weather_hour.Time = time
	check_err(err)

	return current_weather_hour
}

/*
	@brief Function that get's all the weather data and saves it
*/
func get_weather_data() {
	body := retrieve_weather_file()
	value, _, _, err := jsonparser.Get(body, "days", "[0]")

	todaysWeather := get_weather_day(value)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(todaysWeather.Date)
}
