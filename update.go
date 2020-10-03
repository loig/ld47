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
	"github.com/hajimehoshi/ebiten/inpututil"
)

func (g *game) Update(screen *ebiten.Image) error {

	g.frame = (g.frame + 1) % stepDuration

	switch g.state {

	case inLevel:
		if !g.loop.running {
			move := noMove
			if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
				move = right
			} else if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
				move = down
			} else if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
				move = left
			} else if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
				move = up
			}
			if move != noMove {
				g.movePlayer(move)
				g.addToLoop(move)
			}
		} else {
			if g.frame == 0 {
				g.runLoop()
			}
		}
		if g.levelFinished() {
			g.state = levelWon
		}

	case levelWon:
	}

	return nil
}
