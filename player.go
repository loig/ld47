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

type player struct {
	x int
	y int
}

func (g *game) movePlayer(move int) {
	dx, dy := 0, 0
	switch move {
	case right:
		dx, dy = 1, 0
	case down:
		dx, dy = 0, 1
	case left:
		dx, dy = -1, 0
	case up:
		dx, dy = 0, -1
	}
	newx := g.player.x + dx
	newy := g.player.y + dy
	if newx < 0 || newx >= g.level.width ||
		newy < 0 || newy >= g.level.height {
		return
	}
	if g.level.field[newy][newx] == wall {
		return
	}
	g.player.x = newx
	g.player.y = newy
}
