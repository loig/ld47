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

// state of the move loop
type loop struct {
	running       bool
	length        int
	currentMoveID int
	nextMoveID    int
	moves         []int
}

// possible moves
const (
	noMove int = iota
	right
	down
	left
	up
	dashRight
	dashDown
	dashLeft
	dashUp
	numMove
)

// initialize the loop
func (g *game) initLoop() {
	g.loop = loop{
		running:       false,
		length:        g.level.loopLength,
		currentMoveID: 0,
		nextMoveID:    0,
		moves:         make([]int, 0),
	}
}

// add a move to the loop and start it if full
func (g *game) addToLoop(move int) {
	g.loop.moves = append(g.loop.moves, move)
	g.loop.nextMoveID = len(g.loop.moves)
	g.loop.currentMoveID = g.loop.nextMoveID
	if len(g.loop.moves) >= g.loop.length {
		g.loop.running = true
		g.loop.currentMoveID = len(g.loop.moves) - 1
		g.loop.nextMoveID = 0
		g.frame = 0
		if g.level.number == 1 && g.talk.nextTalk == 3 {
			g.playSound(talkSound, false)
		}
	}
}

// execute next move in the loop
func (g *game) runLoop() {
	g.loop.currentMoveID = g.loop.nextMoveID
	g.movePlayer(g.loop.moves[g.loop.nextMoveID])
	g.loop.nextMoveID = (g.loop.nextMoveID + 1) % g.loop.length
}
