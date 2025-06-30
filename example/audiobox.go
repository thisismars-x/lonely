/*
. ----------------------@mail(+392 new)--------
. Date: June 4 - June 30
. From: ***********
.
. Addressed to the public,
.
. Written in lone company, distributed through github.com/thisismars-x/lonely
.
.                    M
. []       {{----}}  O                  7      .
. []       {{ .. }}  N E E D   eeeeee   7    [[.]]
. []       {{    }}  S     O   e    e   7     {}
. []       {{ .. }}  T     G   eeeeee   7     {}
. []       {{____}}  E     2   e        7     {}
. [][@][@]           Rs        eeeeeE   7777  {}
.
.
. With sick passion
. Lethargy and ambition
. Yours
. ***********
.------------------------@(under GPL-license)------
*/

package main

import (
	"fmt"
	"lonely/decl"
	"lonely/key"
	"lonely/ly"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/gopxl/beep"
	"github.com/gopxl/beep/mp3"
	"github.com/gopxl/beep/speaker"
)

// Controls:
// p -> Play
// P -> Pause
// R -> Resume

var speakerInitOnce sync.Once

// Lonely follows a functional paradigm.
// The ideal is to have some global state and a number
// of functions that can access, and mutate it.
// Then reading user input at terminal(through key.KeyListen)
// to know which function to call.
//
// This is the initial screen which is handled by a variable
// 'window: bool' to exit it when user requests to.
func wallpaper() {

	// You get a row and then render some text on that row.
	// Here, {1} is the width-ratio and 0.9 is the height-ratio.
	// You get back as many windows as len({..})
	row := ly.GetRow(ly.W32{1}, 0.9)

	// Functions like Border, Align, PopBorder change the window's
	// qualities. That is why there is no lvalue on its left side.
	// This allows you to share windows better.
	// Also "ASCII", "ascii", "a" are all the same(as long as they begin with a, A)
	// This is a common convention throughout lonely
	row[0].Border("ASCII")

	asset := []string{`
  ______                   __  __                  _______                      
 /      \                 |  \|  \                |       \                     
|  $$$$$$\ __    __   ____| $$ \$$  ______        | $$$$$$$\  ______   __    __ 
| $$__| $$|  \  |  \ /      $$|  \ /      \       | $$__/ $$ /      \ |  \  /  \
| $$    $$| $$  | $$|  $$$$$$$| $$|  $$$$$$\      | $$    $$|  $$$$$$\ \$$\/  $$
| $$$$$$$$| $$  | $$| $$  | $$| $$| $$  | $$      | $$$$$$$\| $$  | $$  >$$  $$ 
| $$  | $$| $$__/ $$| $$__| $$| $$| $$__/ $$      | $$__/ $$| $$__/ $$ /  $$$$\ 
| $$  | $$ \$$    $$ \$$    $$| $$ \$$    $$      | $$    $$ \$$    $$|  $$ \$$\
 \$$   \$$  \$$$$$$   \$$$$$$$ \$$  \$$$$$$        \$$$$$$$   \$$$$$$  \$$   \$$

`}

	// Convert is a function that takes semi-HTML like syntax and pretties it
	// according to the <..> rules. Make is similar to Convert -
	// but Convert is typically useful when you have long text with different
	// style rules.
	caption := ly.Convert(`                            Press <ib fg="#ff0ff0">Space</ib> <b fg="#00ff00">to Start</b>`)
	wallpaperBox := ly.Make(asset[0]+"\n"+caption, 0, "#00ff00", "")

	// Printcl simply Clears the screen and prints the string out
	key.Printcl(row[0].Text(wallpaperBox))
}

// plays, pause, resume are simple functions to play, pause
// and resume your audio-files. Duh! The name!!
func plays(filename string) (*beep.Ctrl, error) {
	args := os.Args
	base := ""

	// Change base directory first arg name(path must have a Music/ Directory)
	if len(args) > 1 {
		base = args[1]
	} else {
		base = os.Getenv("HOME")
	}

	f, err := os.Open(filepath.Join(base, "Music", filename))
	if err != nil {
		return nil, fmt.Errorf("can't open file: %w", err)
	}

	streamer, format, err := mp3.Decode(f)
	if err != nil {
		f.Close()
		return nil, fmt.Errorf("can't decode file: %w", err)
	}

	speakerInitOnce.Do(func() {
		speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	})

	ctrl := &beep.Ctrl{Streamer: streamer, Paused: false}
	speaker.Play(ctrl)

	return ctrl, nil
}

func pause(ctrl *beep.Ctrl) {
	speaker.Lock()
	ctrl.Paused = true
	speaker.Unlock()
}

func resume(ctrl *beep.Ctrl) {
	speaker.Lock()
	ctrl.Paused = false
	speaker.Unlock()
}

var ltime int

// This TUI has only two simple screens - welcome and audio
// This screen provides an interface for the user to
// select and play an audio file
func load_audios(which *int, shouldPlay bool, playing int) *beep.Ctrl {
	starts := time.Now()

	args := os.Args
	base := ""
	if len(args) > 1 {
		base = args[1]
	} else {
		base = os.Getenv("HOME")
	}

	dir := filepath.Join(base, "Music")
	files, err := os.ReadDir(dir)
	if err != nil {
		panic("Could not read files")
	}

	audio_list := []string{}
	for _, entry := range files {
		if !entry.IsDir() {
			audio_list = append(audio_list, entry.Name())
		}
	}

	if *which > len(audio_list) {
		*which = len(audio_list)
	}
	if *which < 1 {
		*which = 1
	}

	// As before, getting two rows
	row := ly.GetRow(ly.W32{1}, 0.65)
	row2 := ly.GetRow(ly.W32{1}, 0.2)

	// As noted before, "ascii", "a", "ASS-ID" are all same
	row[0].Border("ascii")
	row2[0].Border("a/A")

	// This is a trick you will probably play multiple times
	// using Lonely.
	// In lonely you work with rows, which are composed of windows.
	// Windows are horizontally stacked to create rows.
	// Rows and vertically stacked to create the screen.
	//
	// When you PopBorder("l") - "l", "lower", "L"
	// you remove window(row[0])'s lower border
	// PopBorder("u") removes window(row2[0])'s upper border
	// Then you stack them up giving an illusion of a
	// more height-y window.
	// You would generally do this in a case where you have
	// multiple windows in a row where one window needs to
	// be a different height than that of other windows in same
	// row.
	row[0].PopBorder("l")
	row2[0].PopBorder("u")

	// This sets the alignment of the text rendered within
	// the window.
	// Order of alignment -> VERTICAL-HORIZONTAL
	// "rr" places the text on right most both vertically and horizontally
	// "rb" places rightmost vertically and bottom horizontally
	//
	// Note:
	// You can omit "rr" and use "r" if you are fine with default
	// horizontal alignment
	row2[0].Align("rr")

	if ltime == 0 {
		ltime = int(time.Since(starts).Microseconds())
	}

	// Choices_t is used to create a choice list from which you are free to chose
	// from. There are many more styles you can use in lonely/decl,
	// like SIMPLE, DOUBLE_ARROW, RED_SIMPLE, etc.
	s1 := row[0].Choices_t(audio_list, *which, decl.GREEN_SIMPLE, decl.CENTER)

	caption := fmt.Sprintf("Loaded audio file(s) from %s ", dir)
	caption += "\n" + ly.Pad(fmt.Sprintf("in < %d microsecs ", ltime), caption, decl.RIGHT_)

	// You use Text to render text within the window
	s2 := row2[0].Text(ly.Make(caption, 0, "#0fffff", ""))

	key.Printcl(ly.MergeRow(s1, s2))

	if shouldPlay {
		audio := strings.TrimSpace(strings.ReplaceAll(audio_list[*which-1], "\u00A0", " "))
		ctrl, err := plays(audio)
		if err != nil {
			panic(fmt.Sprintf("Could not open %s: %s", audio, err))
		}
		return ctrl
	}

	return nil
}

func main() {

	// This function runs a goroutine to give you what input
	// is provided in string-form.
	// <Esc> becomes "ESC"
	// <Space> becomes "SPACE"
	//
	// for a full list, read lonely/decl
	inKey := key.KeyListen()

	// Some global state to decide what function to call
	// on different user input.
	welcome, playing, paused := true, false, false
	which, this := 1, 0

	var audio *beep.Ctrl

	// Draw the wallpaper to the screen on start
	wallpaper()

	for {

		// Check user input. Consult global state.
		// If you should, call the function that renders
		// to the screen. That is the simplicity of Lonely
		select {
		case keyVal, ok := <-inKey:
			if !ok {
				return
			}

			switch keyVal {

			case "SPACE":
				if welcome {
					welcome = false
					load_audios(&which, false, -1)
				}

			case "UP":
				if !welcome {
					which--
					load_audios(&which, false, -1)
				}

			case "DOWN":
				if !welcome {
					which++
					load_audios(&which, false, -1)
				}

			case "p":
				if !welcome && !playing {
					playing = true
					this = which
					audio = load_audios(&which, true, this)
				}

			case "P":
				if !welcome && playing {
					playing = false
					paused = true
					pause(audio)
				}

			case "R":
				if !welcome && paused {
					playing = true
					paused = false
					resume(audio)
				}

			case "ESC":
				key.ClearScreen()
				return
			}
		}
	}
}
