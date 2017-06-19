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

//botUtils provides a quicker method of calling the command and initializing all of the command instead of doing it manually one by one.
package botUtils

import (
	"github.com/line/line-bot-sdk-go/linebot"

	"github.com/Privisus/botUtils/commands/ball8"
	"github.com/Privisus/botUtils/commands/movies"
	"github.com/Privisus/botUtils/commands/quotes"
	"github.com/Privisus/botUtils/commands/weather"
	"github.com/Privisus/botUtils/random"
)

var Call Caller

//Initialize() assigns a linebot.Client pointer to a Bot variable, making it usable for the handler.
//It also initialize other things such as creating a new weather instance, setting a random seed, and reading ball8 choices text file.
//In this way, only this method of initialization is necessary to make the bot work.
func Initialize(Seed int64, botClient *linebot.Client) {
	Call = NewCaller(botClient)
	ball8.Initialize()
	movies.Initialize()
	quotes.Initialize()
	weather.Initialize()
	random.Initialize(Seed)
}

//Execute() returns a Caller, which then provides a list of executable command function.
func Execute() *Caller {
	return &Call
}
