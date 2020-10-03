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
	"log"

	"github.com/hajimehoshi/ebiten"
)

func main() {
	g := &game{
		level:  testLevel,
		player: player{1, 1},
		loop:   loop{false, 5, 0, 0, make([]int, 0)},
	}

	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Ludum Dare 47")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
