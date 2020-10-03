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

import "strconv"

var speaker1 = "???"
var cubNum = 7
var speaker2BaseName = "C.U.B-"
var speaker2 = speaker2BaseName + strconv.Itoa(cubNum)

var talks = [][]sentence{
	[]sentence{
		sentence{&speaker1, []string{"Welcome to life C.U.B-" + strconv.Itoa(cubNum) + "!"}},
		sentence{&speaker2, []string{"..."}},
		sentence{&speaker1, []string{"You are a freshly generated Cyber-", "netic Unit Benchmark. Hopefully you", "will perform better than your pre-", "decessors."}},
		sentence{&speaker2, []string{"..."}},
		sentence{&speaker1, []string{"You do not look very loquacious."}},
		sentence{&speaker2, []string{"..."}},
		sentence{&speaker1, []string{"I guess we go on then. Press enter."}},
	},
	[]sentence{
		sentence{&speaker1, []string{"Just use the arrow keys to move and", "reach the orange circled tile.", "You can use backspace to restart.", "Press enter when you are ready."}},
	},
	[]sentence{
		sentence{&speaker1, []string{"Oups! Did I forgot to mention that", "you have a move limit after wich", "you will loop forever? Well, sorry.", "Now press enter and enjoy looping."}},
	},
	[]sentence{
		sentence{&speaker1, []string{"You can use space combined with", "some arrow key to dash.", "Press enter when you are ready."}},
	},
	[]sentence{},
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
	g.talk = talk{
		dialog:    talks[0],
		talkState: 1,
		nextTalk:  1,
	}
}

// go to next talk
func (g *game) updateTalks() {
	nextTalk := g.talk.nextTalk
	g.talk = talk{
		dialog:    talks[nextTalk],
		talkState: 1,
		nextTalk:  (nextTalk + 1) % len(talks),
	}
}

const talkFrames = 120
