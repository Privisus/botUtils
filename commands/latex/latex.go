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

//Provides a quick way into parsing a message into the chat into a latex equation (generally math equation).
package latex

import (
	"botUtils/replyHandler"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/line/line-bot-sdk-go/linebot"
)

var base_url string = "https://latex.codecogs.com/png.latex?\\inline&space;\\dpi{300}&space;\\bg_white&space;"
var APP_BASE_URL string = os.Getenv("APP_BASE_URL")

//Parses an incoming message into a latex equation; creates a temporary file that is hosted in /downloaded/ and send that file to the user.
func Execute(eqn string, ReplyToken string, Client *linebot.Client) {
	r, err := http.Get(base_url + eqn)
	if err != nil {
		log.Println("Failed to parse equation")
	} else {
		eq, err := ioutil.TempFile("downloaded", "eqn")
		if err != nil {
			log.Println("Failed to create a temporary file")
			return
		}

		content, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println("Failed parsing to content")
		} else {
			eq.Write(content)

			path, err := filepath.Abs(eq.Name())
			if err != nil {
				log.Println("Error getting absolute path")
			}

			eq.Close()

			os.Rename(path, path+".png")

			//defer os.Remove(path + ".png")

			replyHandler.ReplyMessage(ReplyToken, Client, []linebot.Message{linebot.NewImageMessage(APP_BASE_URL+eq.Name()+".png", APP_BASE_URL+eq.Name()+".png")})
		}
	}
	defer r.Body.Close()
}
