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

package botUtils

import (
	"github.com/Privisus/botUtils/commands/ball8"
	"github.com/Privisus/botUtils/commands/coin"
	"github.com/Privisus/botUtils/commands/dice"
	"github.com/Privisus/botUtils/commands/latex"
	"github.com/Privisus/botUtils/commands/movies"
	"github.com/Privisus/botUtils/commands/quotes"
	"github.com/Privisus/botUtils/commands/weather"

	"github.com/line/line-bot-sdk-go/linebot"
)

//Caller provides a list of executable command function which simplifies commandHandler's process.
type Caller struct {
	Bot        *linebot.Client
	ReplyToken string
}

//Returns a Caller with Client argument as the pointer of the linebot.Client
func NewCaller(Client *linebot.Client) Caller {
	return Caller{Bot: Client}
}

//Sets a replytoken to ensure proper replying.
//NOTICE: This token must be refreshed everytime it is possible to ensure proper replying.
func (c *Caller) SetReplyToken(ReplyToken string) {
	c.ReplyToken = ReplyToken
}

//Executes Ball8 command.
func (c *Caller) Ball8(Message string) {
	ball8.Execute(Message, c.ReplyToken, c.Bot)
}

//Executes Coin command.
func (c *Caller) Coin() {
	coin.Execute(c.ReplyToken, c.Bot)
}

//Executes Dice command.
func (c *Caller) Dice() {
	dice.Execute(c.ReplyToken, c.Bot)
}

//Executes Latex command.
func (c *Caller) Latex(Eqn string) {
	latex.Execute(Eqn, c.ReplyToken, c.Bot)
}

//Executes Movies command.
func (c *Caller) Movies(lookup, year string) {
	movies.Execute(c.ReplyToken, c.Bot, lookup, year)
}

//Executes Quotes command.
func (c *Caller) Quotes() {
	quotes.Execute(c.ReplyToken, c.Bot)
}

//Executes Weather command.
func (c *Caller) Weather(arg string) {
	weather.Execute(arg, c.ReplyToken, c.Bot)
}
