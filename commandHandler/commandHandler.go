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

//Provides a gateway for the handler into handling a command (message) sent by an user.
package commandHandler

import (
	"github.com/Privisus/botUtils"

	"regexp"
	"strings"
	"time"
)

var cooldown_time, time_t int64 = 3, 0
var re *regexp.Regexp = regexp.MustCompile("(.*)\\s+\\(+(\\d+)\\)+")

//ExecuteCommand finds a matching pattern for a potential command in a non-regex way and pick a suitable module to pass it to.
func ExecuteCommand(args []string) {
	if !cooldown() {
		if StartsWith("!", args) {
			if EndsWith("?", args) {
				botUtils.Execute().Ball8(strings.Join(args, ""))
				return
			} else {
				switch args[0][1:] {
				case "roll", "dadu", "dice":
					botUtils.Execute().Dice()
				case "coin":
					botUtils.Execute().Coin()
				case "weather", "cuaca":
					if len(args) >= 2 {
						botUtils.Execute().Weather(args[1])
					} else {
						botUtils.Execute().Weather("")
					}
				case "boxoffice", "cinema", "film", "theater", "movielookup", "movie":
					substr := re.FindStringSubmatch(strings.Join(args[1:], " "))

					film_name := strings.Replace(strings.Join(args[1:], " "), " ", "+", -1)
					year := ""

					if substr != nil {
						film_name = strings.Replace(substr[1], " ", "+", -1)
						year = substr[2]
					}

					botUtils.Execute().Movies(film_name, year)
				case "quote", "qod", "quoteoftheday":
					botUtils.Execute().Quotes()
				default:
					undoCooldown()
					return
				}
			}
		} else if StartsWith("=", args) {
			botUtils.Execute().Latex(strings.Join(args, "")[1:])
			return
		} else {
			undoCooldown()
		}
	}
}

//Adds a specified cooldown time (depending on the configuration) before the user can use any command again.
func cooldown() bool {
	if time.Now().Unix()-time_t > 0 {
		time_t = time.Now().Unix() + cooldown_time
		return false
	}
	return true
}

//Undo the cooldown.
func undoCooldown() {
	time_t -= cooldown_time
}

//Checks if the string is equal.
func checkEqual(match, n string) bool {
	return match == n
}

//Checks if the very last character is equal to the provided match character.
func EndsWith(match string, args []string) bool {
	return checkEqual(match, args[len(args)-1][len(args[len(args)-1])-1:])
}

//Checks if the very first character is equal to the provided match character.
func StartsWith(match string, args []string) bool {
	return checkEqual(match, args[0][:1])
}
