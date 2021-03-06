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

import (
	"github.com/hajimehoshi/ebiten"
)

// Tile type for defining the field
type tile struct {
	kind  int
	image *ebiten.Image
}

// get a floor tile given some coordinates
func getFloorTile(a, b int) tile {
	if (a+b)%2 == 0 {
		return floorTileA
	}
	return floorTileB
}

// get a box tile given some coordinates
func getBoxTile(a, b int) tile {
	if (a+b)%2 == 0 {
		return boxTileA
	}
	return boxTileB
}

// get a reset tile given some coordinates
func getResetTile(a, b int) tile {
	if (a+b)%2 == 0 {
		return resetTileA
	}
	return resetTileB
}
