/*
ld47, a video game made for Ludum Dare game jam, 47th edition
Copyright (C) 2020  Loïg Jezequel

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see https://www.gnu.org/licenses/
*/
package main

// game structure implementing ebiten interface
type game struct {
	state  int
	level  level
	player player
	loop   loop
	frame  int
}

// list of possible game states
const (
	inLevel int = iota
	levelWon
)

// game initialization
func initGame() *game {
	g := &game{}
	g.initLevel("testlevel")
	return g
}

// game state update
func (g *game) updateState(state int) {
	g.state = state
	g.frame = 0
}
