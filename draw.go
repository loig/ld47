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
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

func (g *game) Draw(screen *ebiten.Image) {

	// display level
	for y := 0; y < g.level.height; y++ {
		for x := 0; x < g.level.width; x++ {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(x*tileSize), float64(y*tileSize))
			screen.DrawImage(g.level.field[y][x].image, op)
			if x == g.level.goalx && y == g.level.goaly {
				// display the goal
				screen.DrawImage(goalImage, op)
			}
			if x == g.player.x && y == g.player.y {
				// display the player
				switch g.state {
				case inLevel:
					screen.DrawImage(playerImages[g.player.currentImage], op)
				case levelWon:
					screen.DrawImage(playerWinImages[g.player.currentImage], op)
				}
			}
		}
	}

	// display loop
	for id := 0; id < g.loop.length; id++ {
		s := "....."
		if id < len(g.loop.moves) {
			switch g.loop.moves[id] {
			case right:
				s = "right"
			case up:
				s = "up"
			case left:
				s = "left"
			case down:
				s = "down"
			case dashRight:
				s = "dright"
			case dashUp:
				s = "dup"
			case dashLeft:
				s = "dleft"
			case dashDown:
				s = "ddown"
			}
		}
		ebitenutil.DebugPrintAt(screen, s, 150, 10*id+5)
	}
	ebitenutil.DebugPrintAt(screen, "->", 135, 10*g.loop.currentMoveID+5)
}
