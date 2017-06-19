/*	botUtils - A set of bot utilities for LINE chat bot.
Copyright (C) 2017 Steven Hans

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as
published by the Free Software Foundation, either version 3 of the
License, or (at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

//Provides a quick weather lookup.
package weather

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/Privisus/botUtils/replyHandler"

	"github.com/line/line-bot-sdk-go/linebot"
)

//TODO extend to 168h instead of 48h
var weather_token string = os.Getenv("DARKSKY_TOKEN")
var lat_longitude string = os.Getenv("LATLONGITUDE")

var w Weather

type Weather struct {
	UpdateURL    string
	Token        string
	LatLongItude string

	Summary        string
	Temp           float32
	Temp_FeelsLike float32
	Humidity       float32 //don't forget make percentage
	CloudCover     float32 //same
	WindSpeed      float32
}

//Creates a new Weather instance with provided token and latitude longitude.
func NewWeather(token string, latlongitude string) Weather {
	return Weather{
		Token:        token,
		LatLongItude: latlongitude,
		UpdateURL:    "https://api.darksky.net/forecast/" + token + "/" + latlongitude + "?exclude=currently,minutely,daily,flags&lang=id&units=si",
	}
}

//Updates the weather condition.
func (w *Weather) Update(hourAdvance int) {
	r, err := http.Get(w.UpdateURL)
	if err != nil {
		log.Println("cannot get update URL")
		return
	}
	defer r.Body.Close()

	weather_content := NewWeatherContent()
	json.NewDecoder(r.Body).Decode(&weather_content)

	for _, data := range weather_content.Hourly.Data {
		w_Date := time.Unix(int64(data.Time), 0)
		w_Day, w_Hour := w_Date.Day(), w_Date.Hour()

		duration, err := time.ParseDuration(strconv.Itoa(hourAdvance) + "h")
		if err != nil {
			return
		}

		Day, Hour := time.Now().Add(duration).Day(), time.Now().Hour()

		if w_Day == Day && w_Hour == Hour {
			w.Summary = data.Summary
			w.Temp = data.Temperature
			w.Temp_FeelsLike = data.ApparentTemperature
			w.Humidity = data.Humidity
			w.CloudCover = data.CloudCover
			w.WindSpeed = data.WindSpeed
			break
		}
	}
}

type WeatherContent struct {
	Hourly *WeatherInfo
}

func NewWeatherContent() WeatherContent {
	return WeatherContent{}
}

type WeatherInfo struct {
	Summary string
	Data    []*WeatherData
}

type WeatherData struct {
	Time                int32
	Summary             string
	Temperature         float32
	ApparentTemperature float32
	Humidity            float32
	CloudCover          float32
	WindSpeed           float32
}

//Initializes thw initial Weather struct.
func Initialize() {
	w = NewWeather(weather_token, lat_longitude)
}

//Sends the summary of the forecast to the user.
func Execute(arg string, ReplyToken string, Client *linebot.Client) {
	advancedHour, err := strconv.Atoi(arg)
	if err != nil {
		advancedHour = 0
	}
	switch arg {
	case "":
		advancedHour = 0 //Basically means "now".
	case "nanti": //"Later"
		advancedHour = 1
	case "nanti-nanti": //"Laterrrrr"
		advancedHour = 2
	case "nantinya": //"next time"
		advancedHour = 3
	case "selanjutnya": //"next timeeeee"
		advancedHour = 6
	case "besok": //"tomorrow"
		advancedHour = 24
	case "lusa": //"the day after tomorrow"
		advancedHour = 48
	case "depan": //"12 hours from now"
		advancedHour = 12
	}
	w.Update(advancedHour)

	text_display := ""
	if advancedHour == 0 {
		text_display = "SEKARANG"
	} else {
		text_display = fmt.Sprintf("%d JAM DARI SEKARANG", advancedHour)
	}
	replyHandler.ReplyMessage(ReplyToken, Client, []linebot.Message{linebot.NewTextMessage(fmt.Sprintf("PERKIRAAN CUACA %s:\n%s\n----------\nTemperatur: %.2fC\nTerasa seperti: %.2fC\nKelembapan: %.2f%%\nHamparan awan: %.2f%%\nKelajuan angin: %.2fm/s\n----------\n(SUMBER: DarkSky.net)", text_display, w.Summary, w.Temp, w.Temp_FeelsLike, w.Humidity*100, w.CloudCover*100, w.WindSpeed))})
}
