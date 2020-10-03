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

type loop struct {
	running   bool
	length    int
	currentID int
	moves     []int
}

const (
	noMove int = iota
	right
	down
	left
	up
)

func (g *game) addToLoop(move int) {
	g.loop.moves = append(g.loop.moves, move)
	g.loop.currentID = len(g.loop.moves)
	if len(g.loop.moves) >= g.loop.length {
		g.loop.running = true
		g.loop.currentID = 0
	}
}
