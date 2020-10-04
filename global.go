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
	"golang.org/x/image/font"
)

// size of the screen
const (
	screenWidth  = 384
	screenHeight = 272
	tilex        = 24
	menux        = 4
	tiley        = 17
	tileyOffset  = -8
)

// duration of a loop step in frames
var stepDuration = 25

// number of frames between end of level and next level
var endLevelDuration = 120

// title screen
var titleImage *ebiten.Image

const (
	subtitleText = "Stuck in a loop"
	creditText   = "A game made in 48h for Ludum Dare 47"
	infoText     = "Press enter"
)

// in level info
const (
	totalNumLevel  = 10
	levelInfoText1 = "Benchmark: C.U.B.-"
	levelInfoText2 = " ::: Progress: "
)

// infos on tiles
var (
	tilesImage *ebiten.Image
	tileSize   = 16
)

// All types of tiles
var wallTile = tile{kind: wall}
var floorTileA = tile{kind: floor}
var floorTileB = tile{kind: floor}
var nothingTile = tile{kind: nothing}

// Over the floor objects
var goalImage *ebiten.Image

// infos on player
const (
	numPlayerImages    = 2
	numPlayerWinImages = 6
)

var playerImages []*ebiten.Image
var playerFrames []int = []int{40, 10}
var playerWinImages []*ebiten.Image
var playerWinFrames []int = []int{10, 10, 10, 10, 10, 10}

// menu
const (
	maxLoopLength = 8
)

var menuLeftPartsUp []*ebiten.Image
var menuLeftPartDown *ebiten.Image
var menuRightPartUp *ebiten.Image
var menuRightPartDown *ebiten.Image
var menuCenterPartUp *ebiten.Image
var menuCenterPartDown *ebiten.Image
var menuNotification *ebiten.Image
var menuMoveInfos []*ebiten.Image
var menuEmptySpot *ebiten.Image

var selectedLeftUp *ebiten.Image
var selectedLeftDown *ebiten.Image
var selectedCenterUp *ebiten.Image
var selectedCenterDown *ebiten.Image
var selectedRightUp *ebiten.Image
var selectedRightDown *ebiten.Image

var menuMoveImages []*ebiten.Image

// talking
var displayFont font.Face

const talkFrames = 180
