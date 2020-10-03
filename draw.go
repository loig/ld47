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
	"fmt"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

func (g *game) Draw(screen *ebiten.Image) {

	// draw floor
	for y := 0; y < tiley; y++ {
		for x := 0; x < tilex; x++ {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(x*tileSize), float64(y*tileSize+tileyOffset))
			if (x+y)%2 == 0 {
				screen.DrawImage(floorTileA.image, op)
			} else {
				screen.DrawImage(floorTileB.image, op)
			}
		}
	}

	// display level
	levelxOffset := ((tilex - menux) - g.level.width) / 2
	levelyOffset := (tiley - g.level.height) / 2
	for y := 0; y < g.level.height; y++ {
		for x := 0; x < g.level.width; x++ {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64((x+levelxOffset)*tileSize), float64((y+levelyOffset)*tileSize+tileyOffset))
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
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(screenWidth-menux*tileSize+tileSize/2), float64(2*id*tileSize))
		if id < len(g.loop.moves) {
			screen.DrawImage(menuLeftPartsUp[id], op)
		} else {
			screen.DrawImage(menuEmptySpot, op)
		}
		if id == g.loop.currentMoveID {
			screen.DrawImage(selectedLeftUp, op)
		}
		op.GeoM.Translate(0, 16)
		screen.DrawImage(menuLeftPartDown, op)
		if id == g.loop.currentMoveID {
			screen.DrawImage(selectedLeftDown, op)
		}
		op.GeoM.Translate(16, 0)
		screen.DrawImage(menuCenterPartDown, op)
		if id == g.loop.currentMoveID {
			screen.DrawImage(selectedCenterDown, op)
		}
		op.GeoM.Translate(0, -16)
		screen.DrawImage(menuCenterPartUp, op)
		if id == g.loop.currentMoveID {
			screen.DrawImage(selectedCenterUp, op)
		}
		op.GeoM.Translate(16, 0)
		screen.DrawImage(menuRightPartUp, op)
		if id == g.loop.currentMoveID {
			screen.DrawImage(selectedRightUp, op)
		}
		op.GeoM.Translate(0, 16)
		screen.DrawImage(menuRightPartDown, op)
		if id == g.loop.currentMoveID {
			screen.DrawImage(selectedRightDown, op)
		}

		if id < len(g.loop.moves) {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(screenWidth-menux*tileSize+tileSize*2), float64(2*id*tileSize+tileSize/2))
			screen.DrawImage(menuMoveImages[g.loop.moves[id]], op)
		}
	}

	ebitenutil.DebugPrint(screen, fmt.Sprint(ebiten.CurrentTPS(), ", ", ebiten.CurrentFPS()))
}
