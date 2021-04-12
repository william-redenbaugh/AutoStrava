package main

// Import our libraries
import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

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

	for i := 0; i < 24; i++ {
		indexStr := "[" + strconv.Itoa(i) + "]"
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

	// Generating WeatherHour to save.
	var current_weather_hour WeatherHour

	// Getting the date and checking to make sure we are actually parsing real data.
	time, err := jsonparser.GetString(data, "datetime")
	check_err(err)

	// Parsing out the rest of the relevant JSON information. Can be safely assumed that
	// If we have the data coming in, we have the rest of the data.
	tempurature, err := jsonparser.GetFloat(data, "temp")
	humidity, err := jsonparser.GetFloat(data, "humidity")
	feels_like, err := jsonparser.GetFloat(data, "feelslike")

	windspeed, err := jsonparser.GetFloat(data, "windspeed")
	winddirection, err := jsonparser.GetFloat(data, "winddir")
	UVIndex, err := jsonparser.GetFloat(data, "uvindex")
	RainLikelyHood, err := jsonparser.GetFloat(data, "precipprob")
	Rain, err := jsonparser.GetFloat(data, "precip")

	// Saving into data structure.
	current_weather_hour.Time = time
	current_weather_hour.Temp = float32(tempurature)
	current_weather_hour.PerceivedTemp = float32(feels_like)
	current_weather_hour.Humidity = float32(humidity)
	current_weather_hour.WindSpeed = float32(windspeed)
	current_weather_hour.WindDirection = float32(winddirection)
	current_weather_hour.UVIndex = float32(UVIndex)
	current_weather_hour.RainLikelyHood = float32(RainLikelyHood)
	current_weather_hour.Rain = float32(Rain)

	return current_weather_hour
}

func get_weather_session_data(body []byte) WeatherTwoWeek {
	var two_week_weather WeatherTwoWeek

	for i := 0; i < 10; i++ {
		indexStr := "[" + strconv.Itoa(i) + "]"
		value, _, _, err := jsonparser.Get(body, "days", indexStr)
		check_err(err)
		two_week_weather[i] = get_weather_day(value)
	}

	return two_week_weather
}

/*
	@brief Function that get's all the weather data and saves it
*/
func get_weather_data() {
	body := retrieve_weather_file()

	weather_sesion := get_weather_session_data(body)
	fmt.Println(weather_sesion[9].Date)
}
