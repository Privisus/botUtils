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

//Provides a "yes/no" type reply when executed.
package ball8

import (
	"botUtils/replyHandler"
	"crypto/sha1"
	"encoding/hex"
	"hash"
	"io/ioutil"
	"log"
	"strings"

	"github.com/line/line-bot-sdk-go/linebot"
)

var SHA_HASH hash.Hash

var choices []string

/*
Provides an agree/disagree type of reply. This reply must be consistent if and only if the message is the exact same (ignoring the case) as before.

First, the user sends the bot a ball8 command with the message, then it gets evaluated by the commandHandler and eventually go to this Execute function.

Second, this Execute function works as follows: It creates a brand new sha1 hash to write into, then it writes the message into the hash, then finally it creates the SHA1 hash based on the data written.

Finally, it sums up every int value of every character in the sha1 hash string, then doing a modulo operation of length of the choices to get to the answer.
*/
func Execute(Message string, ReplyToken string, Client *linebot.Client) {
	SHA_HASH = sha1.New()

	SHA_HASH.Write([]byte(Message))
	message_hash := hex.EncodeToString(SHA_HASH.Sum(nil))
	modulo_num := mod(message_hash)

	replyHandler.ReplyMessage(ReplyToken, Client, []linebot.Message{linebot.NewTextMessage(choices[modulo_num%len(choices)])})
}

//Initializes the ball8 module; reading the choices file. usually this file is located in vendor/botUtils/commands/ball8/ball8.txt, but I think this hardcoding can be a little bit volatile.
func Initialize() {
	file, err := ioutil.ReadFile("vendor/botUtils/commands/ball8/ball8.txt")
	if err != nil {
		log.Println("Error opening ball8.txt")
	} else {
		choices = strings.Split(string(file), "\n")
	}
}

//Despite the name of the function, it does not perform a modulo operation; it sums every int character of the text and returns the result of that sum in int.
func mod(text string) int {
	sum := 0
	for _, n := range text {
		sum += int(n)
	}
	return sum
}
