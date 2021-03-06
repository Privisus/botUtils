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

package handler

import (
	"fmt"
)

func ExampleHandleCommand() {
	//...
	switch message := event.Message.(type) {
	case *linebot.TextMessage:
		handler.HandleCommand(message.Text, replyToken, bot)
	}
	// Output: none (nothing)
}

func ExampleSlice() {
	fmt.Println(Slice("aaA bBb CcC"))
	fmt.Println(Slice("!RoLL A Dice"))
	// Output: [aaa bbb ccc]
	// [!roll a dice]
}
