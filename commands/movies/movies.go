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

//Provides a summary of a movie based on themoviedb.org
package movies

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Privisus/botUtils/replyHandler"

	"github.com/line/line-bot-sdk-go/linebot"
)

var movie_token string = os.Getenv("MOVIE_TOKEN")
var poster_baseURL string
var poster_size string
var poster_link string

var movie_conf string = "https://api.themoviedb.org/3/configuration?api_key=" + movie_token
var baseURL_format string = "http://api.themoviedb.org/3/search/movie?api_key=" + movie_token + "&language=en-US&page=1&include_adult=true&query=%s&year=%s"

type Images struct {
	Images *ImageData
}

type ImageData struct {
	Secure_Base_Url string
}

//Handles the JSON decoding in Execute(...).
type Movie struct {
	Total_Results int
	Results       []*MovieData
}

type MovieData struct {
	Title        string
	Release_Date string
	Overview     string
	Vote_Average float32
	Adult        bool
	Poster_Path  string
}

//Returns a new instance of Movie.
func NewMovie() Movie {
	return Movie{}
}

//Returns a new instance of Images.
func NewImages() Images {
	return Images{}
}

/*Performs a movie lookup in themoviedb.org, decoding the reponse to the Movie struct and displaying the results to the user.
 */
func Execute(ReplyToken string, Client *linebot.Client, lookup string, year string) {
	r, err := http.Get(fmt.Sprintf(baseURL_format, lookup, year))
	if err != nil {
		log.Print("Error movie lookup")
	}
	defer r.Body.Close()

	mv := NewMovie()

	json.NewDecoder(r.Body).Decode(&mv)

	reply := ""

	if mv.Total_Results != 0 {
		reply = fmt.Sprintf("Judul: %s\nTanggal rilis: %s\nRating: %.2f\nAdult: %t\nPlot: %s", mv.Results[0].Title, mv.Results[0].Release_Date, mv.Results[0].Vote_Average, mv.Results[0].Adult, mv.Results[0].Overview)
	} else {
		//"Cannot find the movie. Try another keyword".
		reply = "Film tidak ditemukan. Coba kata yang lain."
	}

	log.Print("Poster link is: ", poster_link+mv.Results[0].Poster_Path)

	replyHandler.ReplyMessage(ReplyToken, Client, []linebot.Message{linebot.NewImageMessage(poster_link+mv.Results[0].Poster_Path, poster_link+mv.Results[0].Poster_Path), linebot.NewTextMessage(reply)})
}

func Initialize() {
	r, err := http.Get(movie_conf)
	if err != nil {
		log.Print("Error initializing movie")
	}
	defer r.Body.Close()

	conf := NewImages()

	json.NewDecoder(r.Body).Decode(&conf)

	poster_baseURL = conf.Images.Secure_Base_Url
	poster_link = poster_baseURL + "original"
}
