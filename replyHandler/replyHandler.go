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

//Provides a simple reply handler to the user.
package replyHandler

import (
	"log"

	"github.com/line/line-bot-sdk-go/linebot"
)

//Replies a message to the user with provided message(s).
func ReplyMessage(ReplyToken string, Client *linebot.Client, Messages []linebot.Message) {
	if _, err := Client.ReplyMessage(ReplyToken, Messages...).Do(); err != nil {
		log.Print("Error replying message: ", nil)
	}
}

//Replies a message to the user with provided image(s).
func ReplyImgMessage(ReplyToken string, Client *linebot.Client, Message *linebot.ImageMessage) {
	if _, err := Client.ReplyMessage(ReplyToken, Message).Do(); err != nil {
		log.Print("Error replying message: ", err)
	}
}
