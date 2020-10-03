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

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

func init() {
	img, _, error := ebitenutil.NewImageFromFile("images/tiles.png", ebiten.FilterDefault)
	if error != nil {
		log.Panic(error)
	}
	tilesImage = img

	wallTile.image = tilesImage.SubImage(image.Rect(0, 0, 16, 24)).(*ebiten.Image)
	floorTileA.image = tilesImage.SubImage(image.Rect(16, 0, 32, 24)).(*ebiten.Image)
	floorTileB.image = tilesImage.SubImage(image.Rect(32, 0, 48, 24)).(*ebiten.Image)
	nothingTile.image = tilesImage.SubImage(image.Rect(0, 0, 16, 16)).(*ebiten.Image)

	goalImage = tilesImage.SubImage(image.Rect(48, 0, 64, 24)).(*ebiten.Image)

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

}

func main() {
	g := initGame()

	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Ludum Dare 47")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
