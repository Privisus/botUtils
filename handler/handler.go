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

//Provides a simple way to handle a potential command from a message sent by a user.
package handler

import (
	"botUtils/commandHandler"

	"strings"
)

//Handles a command by passing it to a commandHandler with the arguments from Slice(..), a replytoken, and a client. The last two is given by the creator. The client will either be replying with a response message, or none (because of an error/nonexistent command).
func HandleCommand(Command string) {
	args := InsensitiveSlice(Command)

	if commandHandler.StartsWith("=", args) {
		commandHandler.ExecuteCommand(SensitiveSlice(Command))
	} else {
		commandHandler.ExecuteCommand(args)
	}
}

//Splits the message into a potential command arguments that is separated based on whitespace. This also makes the command case-insensitive.
func InsensitiveSlice(Command string) []string {
	return strings.Split(strings.ToLower(Command), " ")
}

//Splits the message into a potential command arguments that is separated based on whitespace. The case of the command remains unchanged.
func SensitiveSlice(Command string) []string {
	return strings.Split(Command, " ")
}
