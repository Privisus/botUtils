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

//Provides the result of a flipping coin.
package coin

import "botUtils/random"
import "botUtils/replyHandler"
import "github.com/line/line-bot-sdk-go/linebot"

//Provides the result of a flipping coin. It generates a random number between 1 and 2 (inclusive). Will return "heads" if the value is 1, and "tails" if the value is 2, or otherwise, depending on the configuration.
func Execute(ReplyToken string, Client *linebot.Client) {
	n := random.RandomIntBetweenInclusive(1, 2)
	a := ""
	if n == 1 {
		a = "angka"
	}
	if n == 2 {
		a = "garuda"
	}
	replyHandler.ReplyMessage(ReplyToken, Client, []linebot.Message{linebot.NewTextMessage(a)})
}
