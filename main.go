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
	"image"
	"log"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"golang.org/x/image/font"
)

func init() {
	img, _, error := ebitenutil.NewImageFromFile("images/tiles.png", ebiten.FilterDefault)
	if error != nil {
		log.Panic(error)
	}
	tilesImage = img

	// init tiles images
	wallTile.image = tilesImage.SubImage(image.Rect(0, 0, 16, 24)).(*ebiten.Image)
	floorTileA.image = tilesImage.SubImage(image.Rect(16, 0, 32, 24)).(*ebiten.Image)
	floorTileB.image = tilesImage.SubImage(image.Rect(32, 0, 48, 24)).(*ebiten.Image)
	nothingTile.image = tilesImage.SubImage(image.Rect(0, 0, 16, 16)).(*ebiten.Image)

	// init goal image
	goalImage = tilesImage.SubImage(image.Rect(48, 0, 64, 24)).(*ebiten.Image)

	// init player images
	playerImages = make([]*ebiten.Image, numPlayerImages)
	for pos := 0; pos < numPlayerImages; pos++ {
		playerImages[pos] = tilesImage.SubImage(image.Rect(pos*16, 24, (pos+1)*16, 48)).(*ebiten.Image)
	}

	playerWinImages = make([]*ebiten.Image, numPlayerWinImages)
	for pos := 0; pos < numPlayerImages; pos++ {
		playerWinImages[pos] = tilesImage.SubImage(image.Rect(pos*16, 24, (pos+1)*16, 48)).(*ebiten.Image)
	}
	for pos := numPlayerImages; pos < numPlayerWinImages; pos++ {
		playerWinImages[pos] = tilesImage.SubImage(image.Rect((pos-numPlayerImages)*16, 48, ((pos-numPlayerImages)+1)*16, 72)).(*ebiten.Image)
	}

	// init menu images
	menuLeftPartsUp = make([]*ebiten.Image, maxLoopLength)
	for numPart := 0; numPart < maxLoopLength; numPart++ {
		var sx, sy int = 16 * (numPart % 4), 72 + (numPart/4)*24
		menuLeftPartsUp[numPart] = tilesImage.SubImage(image.Rect(sx, sy, sx+16, sy+24)).(*ebiten.Image)
	}
	menuLeftPartDown = tilesImage.SubImage(image.Rect(32, 120, 48, 144)).(*ebiten.Image)
	menuRightPartUp = tilesImage.SubImage(image.Rect(16, 120, 32, 144)).(*ebiten.Image)
	menuRightPartDown = tilesImage.SubImage(image.Rect(0, 144, 16, 168)).(*ebiten.Image)
	menuCenterPartUp = tilesImage.SubImage(image.Rect(0, 120, 16, 144)).(*ebiten.Image)
	menuCenterPartDown = tilesImage.SubImage(image.Rect(48, 120, 64, 144)).(*ebiten.Image)

	selectedLeftDown = tilesImage.SubImage(image.Rect(16, 144, 32, 168)).(*ebiten.Image)
	selectedLeftUp = tilesImage.SubImage(image.Rect(32, 144, 48, 168)).(*ebiten.Image)
	selectedCenterDown = tilesImage.SubImage(image.Rect(48, 144, 64, 168)).(*ebiten.Image)
	selectedCenterUp = tilesImage.SubImage(image.Rect(0, 168, 16, 192)).(*ebiten.Image)
	selectedRightDown = tilesImage.SubImage(image.Rect(16, 168, 32, 192)).(*ebiten.Image)
	selectedRightUp = tilesImage.SubImage(image.Rect(32, 168, 48, 192)).(*ebiten.Image)

	menuEmptySpot = tilesImage.SubImage(image.Rect(48, 168, 64, 192)).(*ebiten.Image)

	menuMoveImages = make([]*ebiten.Image, numMove)
	for move := 1; move < numMove; move++ {
		var sx, sy int = ((move - 1) % 4) * 16, 192 + ((move-1)/4)*24
		menuMoveImages[move] = tilesImage.SubImage(image.Rect(sx, sy, sx+16, sy+24)).(*ebiten.Image)
	}

	// font for dialogs
	ttfont, err := truetype.Parse(fonts.ArcadeN_ttf)
	if err != nil {
		panic(err)
	}
	displayFont = truetype.NewFace(ttfont, &truetype.Options{
		Size:    8,
		DPI:     72,
		Hinting: font.HintingFull,
	})

}

func main() {
	g := initGame()

	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Ludum Dare 47")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
