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

// state of the player
type player struct {
	x            int
	y            int
	currentImage int
	currentFrame int
}

// reset player animation
func (g *game) resetPlayerAnimation() {
	g.player.currentImage = 0
	g.player.currentFrame = 0
}

// reset player position and animation
func (g *game) resetPlayer() {
	g.player.x = g.level.startx
	g.player.y = g.level.starty
	g.resetPlayerAnimation()
}

// update player animation
func (g *game) updatePlayerAnimation() {
	g.player.currentFrame++

	if g.state == inLevel {
		if g.player.currentFrame >= playerFrames[g.player.currentImage] {
			g.player.currentFrame = 0
			g.player.currentImage = (g.player.currentImage + 1) % numPlayerImages
		}
	}

	if g.state == levelWon {
		if g.player.currentImage+1 < numPlayerWinImages {
			if g.player.currentFrame >= playerWinFrames[g.player.currentImage] {
				g.player.currentFrame = 0
				g.player.currentImage = g.player.currentImage + 1
			}
		}
	}
}

// move the player if possible
func (g *game) movePlayer(move int) {
	dx, dy := 0, 0
	switch move {
	case right, dashRight:
		dx, dy = 1, 0
	case down, dashDown:
		dx, dy = 0, 1
	case left, dashLeft:
		dx, dy = -1, 0
	case up, dashUp:
		dx, dy = 0, -1
	}
	newx := g.player.x + dx
	newy := g.player.y + dy
	if newx < 0 || newx >= g.level.width ||
		newy < 0 || newy >= g.level.height {
		g.playSound(missMoveSound, true)
		return
	}
	if g.level.field[newy][newx].kind == wall {
		g.playSound(missMoveSound, true)
		return
	}
	if move == dashRight || move == dashDown ||
		move == dashLeft || move == dashUp {
		for newx+dx >= 0 && newx+dx < g.level.width &&
			newy+dy >= 0 && newy+dy < g.level.height &&
			g.level.field[newy+dy][newx+dx].kind != wall {
			newx += dx
			newy += dy
		}
		g.playSound(dashSound, true)
	} else {
		g.playSound(moveSound, true)
	}
	g.player.x = newx
	g.player.y = newy
}
