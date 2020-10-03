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
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

func (g *game) Draw(screen *ebiten.Image) {

	for y := 0; y < g.level.height; y++ {
		for x := 0; x < g.level.width; x++ {
			switch g.level.field[y][x] {
			case floor:
				ebitenutil.DrawRect(screen, float64(x)*10, float64(y)*10, 10, 10, color.RGBA{0, 255, 0, 255})
			case wall:
				ebitenutil.DrawRect(screen, float64(x)*10, float64(y)*10, 10, 10, color.RGBA{255, 0, 0, 255})
			}
		}
	}

	ebitenutil.DrawRect(screen, float64(g.player.x)*10, float64(g.player.y*10), 10, 10, color.RGBA{0, 0, 255, 255})

}
