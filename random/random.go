/*	botUtils - A set of bot utilities for LINE chat bot.
Copyright (C) 2016 Steven Hans

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

//Provides a random integer generator.
package random

import "math/rand"

var Random *rand.Rand

//Initializes a seed given from the caller.
func Initialize(seed int64) {
	Random = rand.New(rand.NewSource(seed))
}

//Generates a random integer between a and b (inclusive).
func RandomIntBetweenInclusive(a, b int) int {
	return a + Random.Intn(b-a+1)
}

//Generates a random integer between a and b (exclusive).
func RandomIntBetweenExclusive(a, b int) int {
	return RandomIntBetweenInclusive(a-1, b-1)
}
