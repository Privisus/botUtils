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

//Provides a qod of the specified category (http://quotes.rest/qod/categories to see the list of available public categories).
package quotes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"botUtils/replyHandler"

	"github.com/line/line-bot-sdk-go/linebot"
)

var BaseURL, Category string = "http://quotes.rest/qod.json?category=", "love"

var Quotes QuoteHandler

type QuoteHandler struct {
	URL      string
	Category string
}

//Construct a new QuoteHandler from specified URL and Category.
func NewQuote(BaseURL, Category string) QuoteHandler {
	return QuoteHandler{
		URL:      BaseURL,
		Category: Category,
	}
}

type QuoteJSON struct {
	Contents *QuoteInfo
}

//Construct a new QuoteJSON to handle JSON decoding.
func NewQuoteJSON() QuoteJSON {
	return QuoteJSON{}
}

type QuoteInfo struct {
	Quotes []*QuoteContent
}

type QuoteContent struct {
	Quote  string
	Author string
}

//Initializes Quotes.
func Initialize() {
	Quotes = NewQuote(BaseURL+Category, Category)
}

//Sends the quote to the user.
func Execute(ReplyToken string, Client *linebot.Client) {
	r, err := http.Get(Quotes.URL)
	if err != nil {
		log.Println("Error: cannot get quotes")
	}
	defer r.Body.Close()

	qc := NewQuoteJSON()
	json.NewDecoder(r.Body).Decode(&qc)

	replyHandler.ReplyMessage(ReplyToken, Client, []linebot.Message{linebot.NewTextMessage(fmt.Sprintf("%s - %s", qc.Contents.Quotes[0].Quote, qc.Contents.Quotes[0].Author))})
}
