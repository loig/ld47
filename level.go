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
	"path/filepath"
	"strconv"
	"strings"

	"github.com/hajimehoshi/ebiten/ebitenutil"
)

// description of one level
type level struct {
	width         int
	height        int
	startx        int
	starty        int
	goalx         int
	goaly         int
	loopLength    int
	field         [][]tile
	originalfield [][]tile
	nextLevel     string
	number        int
}

// field tile types
const (
	floor int = iota
	wall
	box
	resetFloor
	nothing
)

// check if level is finished
func (g *game) levelFinished() bool {
	return g.player.x == g.level.goalx && g.player.y == g.level.goaly
}

// reset level
func (g *game) resetLevel() {
	g.player.x = g.level.startx
	g.player.y = g.level.starty
	g.copyField()
	g.initLoop()
}

// copy a field in a new array
func (g *game) copyField() {
	f := make([][]tile, g.level.height)
	for y := 0; y < g.level.height; y++ {
		f[y] = make([]tile, g.level.width)
		for x := 0; x < g.level.width; x++ {
			f[y][x] = g.level.originalfield[y][x]
		}
	}
	g.level.field = f
}

// init level from file
func (g *game) initLevel(levelName string) {
	file, error := ebitenutil.OpenFile(filepath.Join("levels/", levelName))
	if error != nil {
		log.Panic("Cannot read level ", levelName, ": ", error)
	}

	bytes := make([]byte, 1) // maybe make this size larger
	n, error := file.Read(bytes)
	if n == 0 && error != nil {
		log.Panic("Cannot read level ", levelName, ": ", error)
	}
	for n == len(bytes) {
		bytes2 := make([]byte, len(bytes))
		n, _ = file.Read(bytes2)
		n += len(bytes)
		bytes = append(bytes, bytes2...)
	}

	lines := strings.Split(string(bytes), "\n")
	if len(lines) < 5 {
		log.Panic("Cannot read level", levelName, ": not enough lines in file")
	}

	// get width
	width, error := strconv.Atoi(strings.TrimPrefix(lines[0], "width="))
	if error != nil {
		log.Panic("Cannot read level", levelName, ": width is not a correct integer")
	}

	// get height
	height, error := strconv.Atoi(strings.TrimPrefix(lines[1], "height="))
	if error != nil {
		log.Panic("Cannot read level", levelName, ": height is not a correct integer")
	}

	// get loop length
	loopLength, error := strconv.Atoi(strings.TrimPrefix(lines[2], "loop="))
	if error != nil {
		log.Panic("Cannot read level", levelName, ": loop is not a correct integer")
	}

	// get next level name
	nextLevel := strings.TrimPrefix(lines[3], "next=")

	// get field
	var startx, starty, goalx, goaly int = 1, 1, width - 2, height - 2
	field := make([][]tile, height)
	if len(lines) < 4+height {
		log.Panic("Cannot read level", levelName, ": number of lines in file does not correspond to level height")
	}
	for line := 4; line < height+4; line++ {
		if len(lines[line]) < width {
			log.Panic("Cannot read level", levelName, ": number of characters per line in file does not correspond to level width")
		}
		fieldLine, isStart, isGoal, tmpStartx, tmpGoalx := getLevelLine(lines[line], width, height, line-4)
		field[line-4] = fieldLine
		if isStart {
			startx = tmpStartx
			starty = line - 4
		}
		if isGoal {
			goalx = tmpGoalx
			goaly = line - 4
		}
	}

	levelNum := g.level.number + 1
	g.level = level{
		width:         width,
		height:        height,
		startx:        startx,
		starty:        starty,
		goalx:         goalx,
		goaly:         goaly,
		loopLength:    loopLength,
		originalfield: field,
		nextLevel:     nextLevel,
		number:        levelNum,
	}
	g.copyField()
	g.updateState(inLevel)
	g.resetPlayer()
	g.initLoop()
}

// read a string describing one line of a level
func getLevelLine(line string, width, height int, lineNum int) (levelLine []tile, isStart, isGoal bool, startx, goalx int) {
	levelLine = make([]tile, width)
	colOffset := ((tilex - menux) - width) / 2
	lineOffset := (tiley - height) / 2
	for column := 0; column < width; column++ {
		currentField := nothingTile
		switch line[column] {
		case '#':
			currentField = wallTile
		case '.':
			currentField = getFloorTile(lineNum+lineOffset, column+colOffset)
		case 's':
			currentField = getFloorTile(lineNum+lineOffset, column+colOffset)
			isStart = true
			startx = column
		case 'g':
			currentField = getFloorTile(lineNum+lineOffset, column+colOffset)
			isGoal = true
			goalx = column
		case 'b':
			currentField = getBoxTile(lineNum+lineOffset, column+colOffset)
		case 'r':
			currentField = getResetTile(lineNum+lineOffset, column+colOffset)
		default:
			log.Panic("Cannot read level: unrecognized character in field description: ", line[column])
		}
		levelLine[column] = currentField
	}
	return levelLine, isStart, isGoal, startx, goalx
}
