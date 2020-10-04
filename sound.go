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
	"math"
	"time"

	"github.com/hajimehoshi/ebiten/audio"
	"github.com/hajimehoshi/ebiten/audio/mp3"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type audioManager struct {
	audioContext *audio.Context
	soundPlayer  *audio.Player
	musicPlayer  *audio.Player
}

const (
	talkSound int = iota
	endLevelSound
	moveSound
	dashSound
	missMoveSound
	resetSound
)

// loop the music
func (g *game) updateMusic() {
	if g.audio.musicPlayer != nil {
		if !g.audio.musicPlayer.IsPlaying() {
			g.audio.musicPlayer.Rewind()
			g.audio.musicPlayer.Play()
		}
	} else {
		var error error
		g.audio.musicPlayer, error = audio.NewPlayer(g.audio.audioContext, infiniteMusic)
		if error != nil {
			log.Panic("Audio problem:", error)
		}
    g.audio.musicPlayer.Play()
	}

}

// stop the music
func (g *game) stopMusic() {
	if g.audio.musicPlayer != nil && g.audio.musicPlayer.IsPlaying() {
		error := g.audio.musicPlayer.Close()
		if error != nil {
			log.Panic("Sound problem:", error)
		}
	}
}

// stop the current non-overlaying sound
func (g *game) stopSound() {
	if g.audio.soundPlayer != nil && g.audio.soundPlayer.IsPlaying() {
		error := g.audio.soundPlayer.Close()
		if error != nil {
			log.Panic("Sound problem:", error)
		}
	}
}

// play a sound, telling if it should overlay
// with other sounds or not
func (g *game) playSound(sound int, overlaying bool) {
	var soundBytes []byte
	var error error
	switch sound {
	case dashSound:
		soundBytes = dashSoundBytes
	case moveSound:
		soundBytes = moveSoundBytes
	case endLevelSound:
		soundBytes = endLevelSoundBytes
	case talkSound:
		soundBytes = talkSoundBytes
	case missMoveSound:
		soundBytes = missMoveSoundBytes
	case resetSound:
		soundBytes = resetSoundBytes
	}
	if overlaying {
		soundPlayer, error := audio.NewPlayerFromBytes(g.audio.audioContext, soundBytes)
		if error != nil {
			log.Panic("Sound problem:", error)
		}
		soundPlayer.Play()
	} else {
		if g.audio.soundPlayer != nil && g.audio.soundPlayer.IsPlaying() {
			error = g.audio.soundPlayer.Close()
			if error != nil {
				log.Panic("Sound problem:", error)
			}
		}
		g.audio.soundPlayer, error = audio.NewPlayerFromBytes(g.audio.audioContext, soundBytes)
		if error != nil {
			log.Panic("Sound problem:", error)
		}
		g.audio.soundPlayer.Play()
	}
}

// load all audio assets
func (g *game) initAudio() {
	var error error
	g.audio.audioContext, error = audio.NewContext(44100)
	if error != nil {
		log.Panic("Audio problem:", error)
	}

	// music
	soundFile, error := ebitenutil.OpenFile("sounds/introsong.mp3")
	if error != nil {
		log.Panic("Audio problem:", error)
	}
	sound, error := mp3.Decode(g.audio.audioContext, soundFile)
	if error != nil {
		log.Panic("Audio problem:", error)
	}
	tduration, _ := time.ParseDuration("30s")
	duration := tduration.Seconds()
	bytes := int64(math.Round(duration * 4 * float64(44100)))
	infiniteMusic = audio.NewInfiniteLoop(sound, bytes)

	// sounds
	soundFile, error = ebitenutil.OpenFile("sounds/dashmove.mp3")
	if error != nil {
		log.Panic("Audio problem:", error)
	}
	sound, error = mp3.Decode(g.audio.audioContext, soundFile)
	if error != nil {
		log.Panic("Audio problem:", error)
	}
	dashSoundBytes, error = ioutil.ReadAll(sound)
	if error != nil {
		log.Panic("Audio problem:", error)
	}

	soundFile, error = ebitenutil.OpenFile("sounds/normalmove.mp3")
	if error != nil {
		log.Panic("Audio problem:", error)
	}
	sound, error = mp3.Decode(g.audio.audioContext, soundFile)
	if error != nil {
		log.Panic("Audio problem:", error)
	}
	moveSoundBytes, error = ioutil.ReadAll(sound)
	if error != nil {
		log.Panic("Audio problem:", error)
	}

	soundFile, error = ebitenutil.OpenFile("sounds/endlevel.mp3")
	if error != nil {
		log.Panic("Audio problem:", error)
	}
	sound, error = mp3.Decode(g.audio.audioContext, soundFile)
	if error != nil {
		log.Panic("Audio problem:", error)
	}
	endLevelSoundBytes, error = ioutil.ReadAll(sound)
	if error != nil {
		log.Panic("Audio problem:", error)
	}

	soundFile, error = ebitenutil.OpenFile("sounds/talk.mp3")
	if error != nil {
		log.Panic("Audio problem:", error)
	}
	sound, error = mp3.Decode(g.audio.audioContext, soundFile)
	if error != nil {
		log.Panic("Audio problem:", error)
	}
	talkSoundBytes, error = ioutil.ReadAll(sound)
	if error != nil {
		log.Panic("Audio problem:", error)
	}

	soundFile, error = ebitenutil.OpenFile("sounds/missmove.mp3")
	if error != nil {
		log.Panic("Audio problem:", error)
	}
	sound, error = mp3.Decode(g.audio.audioContext, soundFile)
	if error != nil {
		log.Panic("Audio problem:", error)
	}
	missMoveSoundBytes, error = ioutil.ReadAll(sound)
	if error != nil {
		log.Panic("Audio problem:", error)
	}

	soundFile, error = ebitenutil.OpenFile("sounds/reset.mp3")
	if error != nil {
		log.Panic("Audio problem:", error)
	}
	sound, error = mp3.Decode(g.audio.audioContext, soundFile)
	if error != nil {
		log.Panic("Audio problem:", error)
	}
	resetSoundBytes, error = ioutil.ReadAll(sound)
	if error != nil {
		log.Panic("Audio problem:", error)
	}
}
