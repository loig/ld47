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
	"strconv"
)

var speaker1 = "???"

const initCubNum = 3

var cubNum = initCubNum

const speaker2BaseName = "C.U.B Mk."

var speaker2 string

var talks = [][]sentence{
	[]sentence{
		sentence{&speaker1, nil},
		sentence{&speaker1, []string{"You are a freshly generated C.U.B:", "Cybernetic Unit Benchmark."}},
		sentence{&speaker1, []string{"Hopefully you will perform better", "than your predecessors."}},
		sentence{&speaker2, []string{"..."}},
		sentence{&speaker1, []string{"You do not look very loquacious."}},
		sentence{&speaker2, []string{"..."}},
		sentence{&speaker1, []string{"I guess we go on then. Press enter."}},
	},
	[]sentence{
		sentence{&speaker1, []string{"Use the arrow keys to reach the", "orange-bordered tile. Your moves", "are recorded on the right. In case", "of emergency, backspace will let", "you restart. Press enter now."}},
	},
	[]sentence{
		sentence{&speaker1, []string{"Oups! Did I forgot to mention that", "you have a move memory limit after", "which you will loop on the recorded", "move sequence forever? Well, sorry.", "Now press enter and enjoy looping."}},
	},
	[]sentence{
		sentence{&speaker1, []string{"You can use space combined with", "some arrow key to dash.", "Press enter when you are ready."}},
	},
	[]sentence{
		sentence{&speaker1, []string{"Do you see this new element on the", "field? This is a box. You can push", "it if you do not dash. Press enter."}},
	},
	[]sentence{
		sentence{&speaker1, []string{"Do you see this new element on the", "field? If you stop on it it will", "free your move memory. Press enter."}},
	},
	[]sentence{
		sentence{&speaker1, []string{"Wow, you did it. Impressive."}},
		sentence{&speaker1, []string{"You were maybe even faster than me."}},
		sentence{&speaker2, []string{"..."}},
		sentence{&speaker2, []string{"..."}},
		sentence{&speaker2, []string{"Wait... what!?"}},
		sentence{&speaker1, []string{"It's your turn to handle it now."}},
		sentence{&speaker2, []string{"Ok, then."}},
		sentence{&speaker1, []string{"One last time: press enter."}},
	},
}

// type for describing talk
type talk struct {
	dialog    []sentence
	talkState int
	nextTalk  int
}

// type for one sentence
type sentence struct {
	speaker *string
	text    []string
}

// init the talks
func (g *game) initTalks() {
	if cubNum != initCubNum {
		speaker1 = speaker2
	}
	speaker2 = speaker2BaseName + strconv.Itoa(cubNum)
	talks[0][0].text = []string{"Welcome to life " + speaker2BaseName + strconv.Itoa(cubNum) + "!"}
	g.talk = talk{
		dialog:    talks[0],
		talkState: 1,
		nextTalk:  1,
	}
	cubNum++
}

// go to next talk
func (g *game) updateTalks() {
	nextTalk := g.talk.nextTalk
	g.talk = talk{
		dialog:    talks[nextTalk],
		talkState: 1,
		nextTalk:  nextTalk + 1,
	}
}

// update the current talk state
func (g *game) updateCurrentTalk() {
	if (g.talk.dialog[g.talk.talkState].speaker == &speaker1) ||
		(g.talk.nextTalk >= len(talks) && g.talk.talkState >= 4) {
		g.playSound(talkSound, false)
	} else {
		g.stopSound()
	}
	g.talk.talkState++
	if cubNum == initCubNum+1 &&
		g.talk.nextTalk >= len(talks) &&
		g.talk.talkState == 6 {
		speaker1 = speaker2BaseName + strconv.Itoa(initCubNum-1)
	}
}
