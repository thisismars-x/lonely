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

package key

import (
	"fmt"
	// "github.com/charmbracelet/lipgloss"
	"github.com/inancgumus/screen"
	"golang.org/x/sys/unix"
	"lonely/decl"
	"os"
)

// Clear the screen, and move cursor to top-left position.
// Makes everything ready to render
func ClearScreen() {
	screen.MoveTopLeft()
	screen.Clear()
}

// This function is shorthand for ClearScreen and Println
func Printcl(some string) {
	ClearScreen()
	fmt.Println(some)
}

// This function disables terminal echo and canonical mode
func disable_echo() (*unix.Termios, error) {
	fd := int(os.Stdin.Fd())
	old_state, err := unix.IoctlGetTermios(fd, unix.TCGETS)
	if err != nil {
		return nil, err
	}
	new_state := *old_state
	new_state.Lflag &^= unix.ECHO | unix.ICANON
	new_state.Cc[unix.VMIN] = 1
	new_state.Cc[unix.VTIME] = 0
	if err := unix.IoctlSetTermios(fd, unix.TCSETS, &new_state); err != nil {
		return nil, err
	}
	return old_state, nil
}

// This function restores terminal state
func restore_echo(old_state *unix.Termios) {
	unix.IoctlSetTermios(int(os.Stdin.Fd()), unix.TCSETS, old_state)
}

// This function reads user input at the terminal by suppressing echo.
// Use this function to change your TUI state with input.
func KeyListen() <-chan string {
	out := make(chan string)

	go func() {
		defer close(out)

		old_state, err := disable_echo()
		if err != nil {
			fmt.Println("Error disabling echo:", err)
			return
		}
		defer restore_echo(old_state)

		buf := make([]byte, 3)
		for {
			n, err := os.Stdin.Read(buf)
			if err != nil {
				return
			}

			if n == 1 {
				switch buf[0] {
				case decl.ESC:
					out <- "ESC"
				case decl.SPACE:
					out <- "SPACE"
				case decl.EXCLAMATION:
					out <- "EXCLAMATION"
				case decl.DOUBLE_QUOTE:
					out <- "DOUBLE_QUOTE"
				case decl.HASH:
					out <- "HASH"
				case decl.DOLLAR:
					out <- "DOLLAR"
				case decl.PERCENT:
					out <- "PERCENT"
				case decl.AMPERSAND:
					out <- "AMPERSAND"
				case decl.SINGLE_QUOTE:
					out <- "SINGLE_QUOTE"
				case decl.LEFT_PAREN:
					out <- "LEFT_PAREN"
				case decl.RIGHT_PAREN:
					out <- "RIGHT_PAREN"
				case decl.ASTERISK:
					out <- "ASTERISK"
				case decl.PLUS:
					out <- "PLUS"
				case decl.COMMA:
					out <- "COMMA"
				case decl.MINUS:
					out <- "MINUS"
				case decl.DOT:
					out <- "DOT"
				case decl.SLASH:
					out <- "SLASH"
				case decl.ZERO:
					out <- "ZERO"
				case decl.ONE:
					out <- "ONE"
				case decl.TWO:
					out <- "TWO"
				case decl.THREE:
					out <- "THREE"
				case decl.FOUR:
					out <- "FOUR"
				case decl.FIVE:
					out <- "FIVE"
				case decl.SIX:
					out <- "SIX"
				case decl.SEVEN:
					out <- "SEVEN"
				case decl.EIGHT:
					out <- "EIGHT"
				case decl.NINE:
					out <- "NINE"
				case decl.COLON:
					out <- "COLON"
				case decl.SEMICOLON:
					out <- "SEMICOLON"
				case decl.LESS_THAN:
					out <- "LESS_THAN"
				case decl.EQUAL:
					out <- "EQUAL"
				case decl.GREATER_THAN:
					out <- "GREATER_THAN"
				case decl.QUESTION_MARK:
					out <- "QUESTION_MARK"
				case decl.AT_:
					out <- "AT"
				case decl.A:
					out <- "A"
				case decl.B:
					out <- "B"
				case decl.C:
					out <- "C"
				case decl.D:
					out <- "D"
				case decl.E:
					out <- "E"
				case decl.F:
					out <- "F"
				case decl.G:
					out <- "G"
				case decl.H:
					out <- "H"
				case decl.I:
					out <- "I"
				case decl.J:
					out <- "J"
				case decl.K:
					out <- "K"
				case decl.L:
					out <- "L"
				case decl.M:
					out <- "M"
				case decl.N:
					out <- "N"
				case decl.O:
					out <- "O"
				case decl.P:
					out <- "P"
				case decl.Q:
					out <- "Q"
				case decl.R:
					out <- "R"
				case decl.S:
					out <- "S"
				case decl.T:
					out <- "T"
				case decl.U:
					out <- "U"
				case decl.V:
					out <- "V"
				case decl.W:
					out <- "W"
				case decl.X:
					out <- "X"
				case decl.Y:
					out <- "Y"
				case decl.Z:
					out <- "Z"
				case decl.LEFT_BRACKET:
					out <- "LEFT_BRACKET"
				case decl.BACKSLASH:
					out <- "BACKSLASH"
				case decl.RIGHT_BRACKET:
					out <- "RIGHT_BRACKET"
				case decl.CARET:
					out <- "CARET"
				case decl.UNDERSCORE:
					out <- "UNDERSCORE"
				case decl.BACKTICK:
					out <- "BACKTICK"
				case decl.SMALL_a:
					out <- "a"
				case decl.SMALL_b:
					out <- "b"
				case decl.SMALL_c:
					out <- "c"
				case decl.SMALL_d:
					out <- "d"
				case decl.SMALL_e:
					out <- "e"
				case decl.SMALL_f:
					out <- "f"
				case decl.SMALL_g:
					out <- "g"
				case decl.SMALL_h:
					out <- "h"
				case decl.SMALL_i:
					out <- "i"
				case decl.SMALL_j:
					out <- "j"
				case decl.SMALL_k:
					out <- "k"
				case decl.SMALL_l:
					out <- "l"
				case decl.SMALL_m:
					out <- "m"
				case decl.SMALL_n:
					out <- "n"
				case decl.SMALL_o:
					out <- "o"
				case decl.SMALL_p:
					out <- "p"
				case decl.SMALL_q:
					out <- "q"
				case decl.SMALL_r:
					out <- "r"
				case decl.SMALL_s:
					out <- "s"
				case decl.SMALL_t:
					out <- "t"
				case decl.SMALL_u:
					out <- "u"
				case decl.SMALL_v:
					out <- "v"
				case decl.SMALL_w:
					out <- "w"
				case decl.SMALL_x:
					out <- "x"
				case decl.SMALL_y:
					out <- "y"
				case decl.SMALL_z:
					out <- "z"
				case decl.LEFT_CURLY:
					out <- "LEFT_CURLY"
				case decl.VERTICAL_BAR:
					out <- "VERTICAL_BAR"
				case decl.RIGHT_CURLY:
					out <- "RIGHT_CURLY"
				case decl.TILDE:
					out <- "TILDE"
				}
			} else if n == 3 && buf[0] == decl.ESC && buf[1] == 91 {
				switch buf[2] {
				case decl.UP:
					out <- "UP"
				case decl.DOWN:
					out <- "DOWN"
				case decl.RIGHT:
					out <- "RIGHT"
				case decl.LEFT:
					out <- "LEFT"
				}
			}
		}
	}()

	return out
}
