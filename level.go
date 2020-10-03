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

// description of one level
type level struct {
	width  int
	height int
	startx int
	starty int
	goalx  int
	goaly  int
	field  [][]int
}

// field tile types
const (
	floor int = iota
	wall
	nothing
)

// check if level is finished
func (g *game) levelFinished() bool {
	return g.player.x == g.level.goalx && g.player.y == g.level.goaly
}

// test level
var testLevel = level{
	width:  6,
	height: 4,
	startx: 1,
	starty: 1,
	goalx:  4,
	goaly:  2,
	field: [][]int{
		[]int{wall, wall, wall, wall, wall, wall},
		[]int{wall, floor, floor, floor, wall, wall},
		[]int{wall, floor, floor, floor, floor, wall},
		[]int{wall, wall, wall, wall, wall, wall},
	},
}
