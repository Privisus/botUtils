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

//Provides a summary of a movie based on http://omdbapi.com/
package movies

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"botUtils/replyHandler"

	"github.com/line/line-bot-sdk-go/linebot"
)

var baseURL_format string = "http://www.omdbapi.com/?t=%s&y=%s&plot=short&r=json"

//Handles the JSON decoding in Execute(...).
type Movie struct {
	Response   string
	Error      string
	Title      string
	Year       string
	Rated      string
	Released   string
	Runtime    string
	Genre      string
	Plot       string
	Metascore  string
	ImdbRating string
}

//Returns a new instance of Movie.
func NewMovie() Movie {
	return Movie{}
}

/*Performs a movie lookup in http://omdbapi.com/, decoding the reponse to the Movie struct and displaying the results to the user.

(what is that "judul, tahun, rated" means): Judul -> title, Tahun -> year, Rated -> rated, Tanggal rilis -> Released and so on.
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

	if mv.Response == "True" {
		reply = fmt.Sprintf("Judul: %s\nTahun: %s\nRated: %s\nTanggal rilis: %s\nDurasi: %s\nGenre: %s\nPlot: %s\nMetascore: %s\nIMDB rating: %s", mv.Title, mv.Year, mv.Rated, mv.Released, mv.Runtime, mv.Genre, mv.Plot, mv.Metascore, mv.ImdbRating)
	} else {
		//"Cannot find the movie. Try another keyword".
		reply = "Film tidak ditemukan. Coba kata yang lain."
	}

	replyHandler.ReplyMessage(ReplyToken, Client, []linebot.Message{linebot.NewTextMessage(reply)})
}
