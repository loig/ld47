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
	"io/ioutil"
	"log"
	"path/filepath"
	"strconv"
	"strings"
)

// description of one level
type level struct {
	width      int
	height     int
	startx     int
	starty     int
	goalx      int
	goaly      int
	loopLength int
	field      [][]int
	nextLevel  string
}

// field tile types
const (
	floor int = iota
	wall
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
	g.initLoop()
}

// init level from file
func (g *game) initLevel(levelName string) {
	file, error := ioutil.ReadFile(filepath.Join("levels/", levelName))
	if error != nil {
		log.Panic("Cannot read level ", levelName, ": ", error)
	}
	log.Print(string(file))

	lines := strings.Split(string(file), "\n")
	if len(lines) < 5 {
		log.Panic("Cannot read level", levelName, ": not enough lines in file")
	}
	log.Print("Num lines: ", len(lines))

	// get width
	width, error := strconv.Atoi(strings.TrimPrefix(lines[0], "width="))
	if error != nil {
		log.Panic("Cannot read level", levelName, ": width is not a correct integer")
	}
	log.Print("Width: ", width)

	// get height
	height, error := strconv.Atoi(strings.TrimPrefix(lines[1], "height="))
	if error != nil {
		log.Panic("Cannot read level", levelName, ": height is not a correct integer")
	}
	log.Print("Height: ", height)

	// get loop length
	loopLength, error := strconv.Atoi(strings.TrimPrefix(lines[2], "loop="))
	if error != nil {
		log.Panic("Cannot read level", levelName, ": loop is not a correct integer")
	}
	log.Print("Loop: ", loopLength)

	// get next level name
	nextLevel := strings.TrimPrefix(lines[3], "next=")
	log.Print("Next level: ", nextLevel)

	// get field
	var startx, starty, goalx, goaly int
	field := make([][]int, height)
	if len(lines) < 4+height {
		log.Panic("Cannot read level", levelName, ": number of lines in file does not correspond to level height")
	}
	for line := 4; line < height+4; line++ {
		if len(lines[line]) < width {
			log.Panic("Cannot read level", levelName, ": number of characters per line in file does not correspond to level width")
		}
		log.Print(lines[line])
		fieldLine, isStart, isGoal, tmpStartx, tmpGoalx := getLevelLine(lines[line], width)
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

	g.level = level{
		width:      width,
		height:     height,
		startx:     startx,
		starty:     starty,
		goalx:      goalx,
		goaly:      goaly,
		loopLength: loopLength,
		field:      field,
		nextLevel:  nextLevel,
	}
	g.initPlayer()
	g.initLoop()
}

// read a string describing one line of a level
func getLevelLine(line string, width int) (levelLine []int, isStart, isGoal bool, startx, goalx int) {
	levelLine = make([]int, width)
	for column := 0; column < width; column++ {
		currentField := nothing
		switch line[column] {
		case '#':
			currentField = wall
		case '.':
			currentField = floor
		case 's':
			currentField = floor
			isStart = true
			startx = column
		case 'g':
			currentField = floor
			isGoal = true
			goalx = column
		default:
			log.Panic("Cannot read level: unrecognized character in field description: ", line[column])
		}
		levelLine[column] = currentField
	}
	return levelLine, isStart, isGoal, startx, goalx
}
