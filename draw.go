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
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
)

func (g *game) Draw(screen *ebiten.Image) {

	// draw floor
	for y := 0; y < tiley; y++ {
		for x := 0; x < tilex; x++ {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(x*tileSize), float64(y*tileSize+tileyOffset))
			if (x+y)%2 == 0 {
				screen.DrawImage(getFloorTile(y, x).image, op)
			} else {
				screen.DrawImage(getFloorTile(y, x).image, op)
			}
		}
	}

	if g.state == inLevel || g.state == levelWon {
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
					if g.level.field[y][x].kind == box {
						screen.DrawImage(goalBoxImage, op)
					} else {
						screen.DrawImage(goalImage, op)
					}
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

		// display info
		s := fmt.Sprint(levelInfoText1, cubNum-1, levelInfoText2)
		if g.state == inLevel {
			s = fmt.Sprint(s, g.level.number-1)
		} else {
			s = fmt.Sprint(s, g.level.number)
		}
		s = fmt.Sprint(s, "/", totalNumLevel)
		text.Draw(screen, s, displayFont, 9, 268, color.RGBA{45, 47, 74, 255})

		if g.state == inLevel {
			text.Draw(screen, "Backslash to restart level", displayFont, 9, 10, color.RGBA{45, 47, 74, 255})
		}
	}

	if g.state == intro || g.state == gameWon ||
		(g.state == inLevel && g.level.number == 1 && g.talk.nextTalk == 2) ||
		(g.state == inLevel && g.level.number == 1 && g.talk.nextTalk == 3 && g.loop.running) ||
		(g.state == inLevel && g.level.number == 2 && g.talk.nextTalk == 4) ||
		(g.state == inLevel && g.level.number == 10 && g.talk.nextTalk == 5) {
		textyPos := 16
		dy := tileSize/2 - 2
		dx := tileSize - 3
		xloop := 22
		for pos := 0; pos < g.talk.talkState; pos++ {
			boxxPos := 6
			textxPos := boxxPos + 8
			if g.talk.dialog[pos].speaker == &speaker2 {
				boxxPos = 88
				textxPos = boxxPos + 8
			}
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(boxxPos), float64(textyPos-tileSize*3/2))
			for j := 0; j < len(g.talk.dialog[pos].text)+1; j++ {
				for i := 0; i < xloop; i++ {
					screen.DrawImage(menuEmptySpot, op)
					op.GeoM.Translate(0, float64(dy))
					screen.DrawImage(menuEmptySpot, op)
					op.GeoM.Translate(float64(dx), float64(-dy))
				}
				op.GeoM.Translate(float64(-xloop*dx), 1.6*float64(dy))
			}

			speaker := *(g.talk.dialog[pos].speaker) + ":"
			text.Draw(screen, speaker, displayFont, textxPos, textyPos, color.RGBA{249, 65, 9, 255})
			textyPos += 10
			for _, line := range g.talk.dialog[pos].text {
				text.Draw(screen, line, displayFont, textxPos, textyPos, color.RGBA{169, 49, 13, 255})
				textyPos += 10
			}

			textyPos += 14
		}
	}

	if g.state == titlescreen {
		op := &ebiten.DrawImageOptions{}
		screen.DrawImage(titleImage, op)
		text.Draw(screen, infoText, displayFont, 148, 220, color.RGBA{249, 65, 9, 255})
		text.Draw(screen, creditText, displayFont, 90, 11, color.RGBA{45, 47, 74, 255})
	}

	// debug
	//ebitenutil.DebugPrint(screen, fmt.Sprint(ebiten.CurrentTPS(), ", ", ebiten.CurrentFPS()))
}
