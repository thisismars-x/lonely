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

package ly

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"lonely/decl"

	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"
)

type window = string
type row = string
type uninit_window lipgloss.Style
type uninit_window_short = lipgloss.Style
type w32 = []float32
type from = []string

type W32 = w32
type From = from

// This function takes a number of smaller windows and embeds them in a single row
//
// Careful:
// Order of args matter, as rows are filled from the left
func MakeRow(list ...window) (result row) {
	initial := ""
	for i := range len(list) {
		result = merge_window(initial, list[i])
		initial = result
	}

	clean_newlines(&result)
	return
}

func merge_window(a, b window) (merged window) {
	a_len, b_len := len(strings.Split(a, "\n")), len(strings.Split(b, "\n"))

	if a_len > b_len {
		for _ = range a_len - b_len {
			b += "\n"
		}
	} else if b_len > a_len {
		for _ = range b_len - a_len {
			a += "\n"
		}
	}

	a_split := strings.Split(a, "\n")
	b_split := strings.Split(b, "\n")

	if len(a_split) != len(b_split) {
		panic("In function merge 'a' and 'b' have different sizes even after padding")
	}

	for i := range len(a_split) {
		merged += a_split[i]
		merged += b_split[i]
		merged += "\n"
	}

	return
}

// This function composes your rows into the terminal screen.
// Use it with your rows from 'MakeRow' to get a bigger composition.
//
// Careful:
// Order of args matter as rows are filled from the topmost
func MergeRow(list ...row) (result row) {
	for i := range len(list) {
		result = fmt.Sprintf("%s\n%s", result, list[i])
	}

	return
}

var SCREEN_WIDTH, SCREEN_HEIGHT int

func TermGetSize() (w, h int) {
	w, h, _ = term.GetSize(int(os.Stdin.Fd()))
	SCREEN_WIDTH, SCREEN_HEIGHT = w, h
	return
}

func clean_newlines(some *string) {
	for len(*some) > 0 && (*some)[len(*some)-1] == '\n' {
		*some = (*some)[:len(*some)-1]
	}
}

var isinit bool

// This function is used to get a bunch of windows that are later
// to be used to make a row
func GetRow(width []float32, height float32) (windows []uninit_window) {
	if !isinit {
		SCREEN_WIDTH, SCREEN_HEIGHT = TermGetSize()
		isinit = true
	}

	n_windows := len(width)

	if height > 1.0 {
		panic("In func GetRow height arg is invalid(over 1.0)")
	}

	if n_windows > 5 {
		panic("You can only have upto 5 windows in a row(violated in func GetRow)")
	}

	windows = make([]uninit_window, 0, n_windows)
	for i := range n_windows {
		width[i] *= 0.99 - (float32(n_windows) / 100)
		windows = append(windows, default_uninit_window(width[i], height))
	}

	return
}

// This function renders some Text to your window
func (self *uninit_window) Text(some string) (final row) {
	final = uninit_window_short(*self).Render(some)
	return
}

func (self *uninit_window) Pad(some string, how int) (final row) {
	width := self.get_width()
	tmp_final := pad([]string{some}, width, how)
	final = uninit_window_short(*self).Render(strings.Join(tmp_final, "\n"))

	return
}

// This function fills a window with a certain color
func (self *uninit_window) Fill(color string) {
	*self = uninit_window(uninit_window_short(*self).Background(lipgloss.Color(color)))
}

// This function is a helper function used by Choices_t
func (self *uninit_window) Choices(some []string, selected, screen_type, how int) (final string) {
	width := self.get_width()
	tmp_some := pad(some, width, how)

	final = ManyChoices(tmp_some, selected, screen_type)
	return
}

// This function is used to create a list of choices to allow the user
// to chose from them.
//
// Note:
// selected is 1-indexed
// screen_type maybe any of decl.SIMPLE_CHOICES, decl.RED_DOUBLE_ARROW, and many more
// (read at lonely/decl)
// how maybe any of decl.LEFT_, decl.RIGHT_, decl.TOP, decl,BOTTOM
func (self *uninit_window) Choices_t(some []string, selected, screen_type, how int) (final row) {
	if selected > len(some) {
		selected = len(some)
	}

	if selected < 0 {
		selected = 1
	}

	switch screen_type {

	case decl.DOUBLE_ARROW_CHOICES, decl.PURPLE_DOUBLE_ARROW, decl.ORANGE_DOUBLE_ARROW, decl.RED_DOUBLE_ARROW:
		some[selected-1] = fmt.Sprintf(">> %s  ", some[selected-1])

	case decl.ARROW_CHOICES, decl.PURPLE_ARROW, decl.RED_ARROW, decl.ORANGE_ARROW:
		some[selected-1] = fmt.Sprintf("-> %s  ", some[selected-1])

	}

	final = self.Text(self.Choices(some, selected, screen_type, how))
	return
}

// This function is used to set Border-style on your window.
// Unlike many other 'self' functions, this function does not return anything
// but mutates it permanently, this is so that you can share
// your windows better.
// Note: ASCII borders work much better for bigger windows(other styles do not
// have any problems- ASCII fits better)
func (self *uninit_window) Border(some string) {
	switch string(some[0]) {
	case "d", "D":
		*self = uninit_window(uninit_window_short(*self).Border(lipgloss.DoubleBorder()))

	case "n", "N":
		*self = uninit_window(uninit_window_short(*self).Border(lipgloss.NormalBorder()))

	case "a", "A":
		*self = uninit_window(uninit_window_short(*self).Border(lipgloss.ASCIIBorder()))

	case "r", "R":
		*self = uninit_window(uninit_window_short(*self).Border(lipgloss.RoundedBorder()))

	default:
		panic("Invalid args to uninit_window's border method")

	}
}

// Similar to BorderColor returns nothing
func (self *uninit_window) BorderColor(some string) {
	*self = uninit_window(
		uninit_window_short(*self).
			BorderForeground(lipgloss.Color(some)))
}

// This function uses VERTICAL-HORIZONTAL align string
// Example:
//
// window_example.align("cB")
// This sets vertical align to center and align to bottom
//
// window_example.align("xb")
// This sets only horizontal align
func (self *uninit_window) Align(some string) {
	if len(some) > 0 {
		switch some[0] {
		case 'c', 'C':
			*self = uninit_window(uninit_window_short(*self).
				AlignVertical(lipgloss.Center))

		case 'l', 'L':
			*self = uninit_window(uninit_window_short(*self).
				AlignVertical(lipgloss.Left))

		case 'r', 'R':
			*self = uninit_window(uninit_window_short(*self).
				AlignVertical(lipgloss.Right))

		case 'b', 'B':
			*self = uninit_window(uninit_window_short(*self).
				AlignVertical(lipgloss.Bottom))

		case 't', 'T', 'u', 'U':
			*self = uninit_window(uninit_window_short(*self).
				AlignVertical(lipgloss.Top))
		}

		if len(some) > 1 {
			switch some[1] {
			case 'c', 'C':
				*self = uninit_window(uninit_window_short(*self).
					AlignHorizontal(lipgloss.Center))

			case 'l', 'L':
				*self = uninit_window(uninit_window_short(*self).
					AlignHorizontal(lipgloss.Left))

			case 'r', 'R':
				*self = uninit_window(uninit_window_short(*self).
					AlignHorizontal(lipgloss.Right))

			case 'b', 'B':
				*self = uninit_window(uninit_window_short(*self).
					AlignHorizontal(lipgloss.Bottom))

			case 't', 'T', 'u', 'U':
				*self = uninit_window(uninit_window_short(*self).
					AlignHorizontal(lipgloss.Top))
			}
		}
	}
}

func (self *uninit_window) get_width() (some int) {
	tmp_self := uninit_window_short(*self)
	some = tmp_self.GetWidth()
	return
}

func (self *uninit_window) get_height() (some int) {
	tmp_self := uninit_window_short(*self)
	some = tmp_self.GetHeight()
	return
}

func pad(some []string, width, how int) (final []string) {
	width = int(float32(width) * .96)
	pad_left := 0
	pad_right := 0

	for i := range len(some) {
		switch how {

		case decl.CENTER:
			pad_right = (len(some[i]) - width) / 2
			pad_left = pad_right

		case decl.RIGHT:
			fallthrough

		case decl.LEFT:
			pad_left = 0
			pad_right = len(some[i]) - width

		}

		if how == decl.RIGHT {
			pad_left, pad_right = pad_right, pad_left
		}

		if pad_right < 0 {
			pad_right *= -1
		}

		if pad_left < 0 {
			pad_left *= -1
		}

		a := ""
		for _ = range pad_left {
			a += " "
		}

		a += some[i]

		for _ = range pad_right {
			a += " "
		}

		// Trailing " " are removed and not displayed by terminals
		some[i] = strings.ReplaceAll(a, " ", "\u00A0")
	}

	final = some
	return
}

// This function removes all border of your window
func (self *uninit_window) HideBorder() {
	*self = uninit_window(uninit_window_short(*self).Border(lipgloss.HiddenBorder()))
}

// This function is used to remove either top or bottom border.
// You would use this function when you get a row with many windows
// and you desire some window has a bigger height than it
// has been reserved. Notice all windows in a row have same height
// (that's what rows are!) but may have different widths.
//
// You can use this function in such cases(and other cases too) to make
// a window look more height-y than it really is
func (self *uninit_window) PopBorder(some string) {
	style := uninit_window_short(*self)
	border, _, _, _, _ := style.GetBorder()

	removeTop := false
	removeBottom := false

	if len(some) > 0 {
		switch some[0] {
		case 'u', 'U', 't', 'T':
			removeTop = true
		case 'b', 'B', 'l', 'L':
			removeBottom = true
		}
	}

	if len(some) > 1 {
		switch some[1] {
		case 'u', 'U', 't', 'T':
			removeTop = true
		case 'b', 'B':
			removeBottom = true
		}
	}

	if removeTop {
		style = style.BorderTop(false)
		border.Top = ""
		border.TopLeft = "|"
		border.TopRight = "|"
	}

	if removeBottom {
		style = style.BorderBottom(false)
		border.Bottom = ""
		border.BottomLeft = "|"
		border.BottomRight = "|"
	}

	style = style.Border(border)
	*self = uninit_window(style)
}

func default_uninit_window(w_ratio, h_ratio float32) (some uninit_window) {
	some = uninit_window(
		lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			Width(int(float32(SCREEN_WIDTH) * w_ratio)).
			Height(int(float32(SCREEN_HEIGHT) * h_ratio)).
			AlignHorizontal(lipgloss.Center).
			AlignVertical(lipgloss.Center).
			BorderForeground(lipgloss.Color(decl.Green)))

	return
}

// This function provides many load screen options. See at lonely/decl
func LoadScreen(curr, final, screen_type int) (some string) {
	curr = (curr * 20 / final)
	final = 20

	switch screen_type {

	case decl.SIMPLE:
		some = load_screen_simple(curr, final)

	case decl.SIMPLE_BLOCKS:
		some = load_screen_simple_blocks(curr, final)

	case decl.SIMPLE_EQUAL:
		some = load_screen_simple_equal(curr, final)

	case decl.AT:
		some = load_screen_at(curr, final)

	case decl.RED_BLOCKS:
		some = fmt.Sprintf("\033[31m%s\033[0m", load_screen_simple_blocks(curr, final))

	case decl.YELLOW_BLOCKS:
		some = fmt.Sprintf("\033[33m%s\033[0m", load_screen_simple_blocks(curr, final))

	case decl.ORANGE_BLOCKS:
		some = fmt.Sprintf("\033[34m%s\033[0m", load_screen_simple_blocks(curr, final))

	case decl.GREEN_BLOCKS:
		some = fmt.Sprintf("\033[32m%s\033[0m", load_screen_simple_blocks(curr, final))

	case decl.BLUE_BLOCKS:
		some = fmt.Sprintf("\033[94m%s\033[0m", load_screen_simple_blocks(curr, final))

	case decl.RED_EQUAL:
		some = fmt.Sprintf("\033[31m%s\033[0m", load_screen_simple_equal(curr, final))

	case decl.YELLOW_EQUAL:
		some = fmt.Sprintf("\033[33m%s\033[0m", load_screen_simple_equal(curr, final))

	case decl.ORANGE_EQUAL:
		some = fmt.Sprintf("\033[34m%s\033[0m", load_screen_simple_equal(curr, final))

	case decl.GREEN_EQUAL:
		some = fmt.Sprintf("\033[32m%s\033[0m", load_screen_simple_equal(curr, final))

	case decl.BLUE_EQUAL:
		some = fmt.Sprintf("\033[94m%s\033[0m", load_screen_simple_equal(curr, final))

	case decl.RED_AT:
		some = fmt.Sprintf("\033[31m%s\033[0m", load_screen_at(curr, final))

	case decl.YELLOW_AT:
		some = fmt.Sprintf("\033[33m%s\033[0m", load_screen_at(curr, final))

	case decl.ORANGE_AT:
		some = fmt.Sprintf("\033[34m%s\033[0m", load_screen_at(curr, final))

	case decl.GREEN_AT:
		some = fmt.Sprintf("\033[32m%s\033[0m", load_screen_at(curr, final))

	case decl.BLUE_AT:
		some = fmt.Sprintf("\033[94m%s\033[0m", load_screen_at(curr, final))

	case decl.RED_SIMPLE:
		some = fmt.Sprintf("\033[31m%s\033[0m", load_screen_simple(curr, final))

	case decl.YELLOW_SIMPLE:
		some = fmt.Sprintf("\033[33m%s\033[0m", load_screen_simple(curr, final))

	case decl.ORANGE_SIMPLE:
		some = fmt.Sprintf("\033[34m%s\033[0m", load_screen_simple(curr, final))

	case decl.GREEN_SIMPLE:
		some = fmt.Sprintf("\033[32m%s\033[0m", load_screen_simple(curr, final))

	case decl.BLUE_SIMPLE:
		some = fmt.Sprintf("\033[94m%s\033[0m", load_screen_simple(curr, final))

	}

	return
}

func load_screen_simple_equal(curr, final int) (some string) {
	some += "[ "
	for _ = range curr {
		some += "="
	}

	if curr != final {
		some += ">"
	}

	for _ = range final - curr {
		some += " "
	}

	some += " ]"

	return
}

func load_screen_simple_blocks(curr, final int) (some string) {
	some += "[ "
	for _ = range curr {
		some += "■"
	}

	for _ = range final - curr {
		some += " "
	}

	some += " ]"

	return
}

func load_screen_simple(curr, final int) (some string) {
	some += "[ "
	for _ = range curr {
		some += "*"
	}

	for _ = range final - curr {
		some += " "
	}

	some += " ]"
	return
}

func load_screen_at(curr, final int) (some string) {
	some += "[ "
	for _ = range curr {
		some += "@"
	}

	for _ = range final - curr {
		some += " "
	}

	some += " ]"
	return
}

// This function provides choice-lists. See screen_type at lonely/decl
func ManyChoices(list []string, selected, screen_type int) (some string) {

	if len(list) < selected {
		selected = len(list)
	}

	if selected < 1 {
		selected = 1
	}

	switch screen_type {

	case decl.SIMPLE_CHOICES:
		some = many_choices_simple(list, selected)

	case decl.ARROW_CHOICES:
		some = many_choices_arrow(list, selected)

	case decl.DOUBLE_ARROW_CHOICES:
		some = many_choices_dbl_arrow(list, selected)

	case decl.ORANGE_SIMPLE:
		some = many_choices_simple_orange(list, selected)

	case decl.GREEN_SIMPLE:
		some = many_choices_simple_green(list, selected)

	case decl.RED_SIMPLE:
		some = many_choices_simple_red(list, selected)

	case decl.ORANGE_ARROW:
		some = many_choices_arrow_orange(list, selected)

	case decl.RED_ARROW:
		some = many_choices_arrow_red(list, selected)

	case decl.PURPLE_ARROW:
		some = many_choices_arrow_purple(list, selected)

	case decl.ORANGE_DOUBLE_ARROW:
		some = many_choices_dbl_arrow_orange(list, selected)

	case decl.RED_DOUBLE_ARROW:
		some = many_choices_dbl_arrow_red(list, selected)

	case decl.PURPLE_DOUBLE_ARROW:
		some = many_choices_dbl_arrow_purple(list, selected)

	case decl.HORIZONTAL_CHOICES:
		some = many_choices_horizontal(list, selected)

	case decl.GREEN_HORIZONTAL_CHOICES:
		some = many_choices_horizontal_green(list, selected)

	case decl.RED_HORIZONTAL_CHOICES:
		some = many_choices_horizontal_red(list, selected)

	case decl.ORANGE_HORIZONTAL_CHOICES:
		some = many_choices_horizontal_orange(list, selected)

	}

	return
}

func many_choices_horizontal(list []string, selected int) (some string) {
	for i := 0; i < len(list); i++ {
		if i == selected-1 {
			box := lipgloss.NewStyle().
				Foreground(lipgloss.Color(decl.Black)).
				Background(lipgloss.Color(decl.White)).
				Render(list[i])

			some += box + " "
		} else {
			some += list[i] + " "
		}
	}

	return
}

func many_choices_horizontal_green(list []string, selected int) (some string) {
	for i := 0; i < len(list); i++ {
		if i == selected-1 {
			box := lipgloss.NewStyle().
				Foreground(lipgloss.Color(decl.Black)).
				Background(lipgloss.Color(decl.Green)).
				Bold(true).
				Render(list[i])

			some += box + " "
		} else {
			some += list[i] + " "
		}
	}

	return
}

func many_choices_horizontal_red(list []string, selected int) (some string) {
	for i := 0; i < len(list); i++ {
		if i == selected-1 {
			box := lipgloss.NewStyle().
				Foreground(lipgloss.Color(decl.Black)).
				Background(lipgloss.Color(decl.Red)).
				Bold(true).
				Render(list[i])

			some += box + " "
		} else {
			some += list[i] + " "
		}
	}

	return
}

func many_choices_horizontal_orange(list []string, selected int) (some string) {
	for i := 0; i < len(list); i++ {
		if i == selected-1 {
			box := lipgloss.NewStyle().
				Foreground(lipgloss.Color(decl.Black)).
				Background(lipgloss.Color(decl.Orange)).
				Bold(true).
				Render(list[i])

			some += box + " "
		} else {
			some += list[i] + " "
		}
	}

	return
}

func many_choices_dbl_arrow(list []string, selected int) (some string) {
	for i := 0; i < len(list); i++ {
		if i == selected-1 {
			box := lipgloss.NewStyle().
				Foreground(lipgloss.Color(decl.Green)).
				Render(list[i])
			some += box
			some += "\n"
		} else {
			some += list[i] + "\n"
		}
	}

	return
}

func many_choices_dbl_arrow_red(list []string, selected int) (some string) {
	for i := 0; i < len(list); i++ {
		if i == selected-1 {
			box := lipgloss.NewStyle().
				Bold(true).
				Foreground(lipgloss.Color(decl.Red)).
				Render(list[i])
			some += box
			some += "\n"
		} else {
			some += list[i] + "\n"
		}
	}

	return
}
func many_choices_dbl_arrow_orange(list []string, selected int) (some string) {
	for i := 0; i < len(list); i++ {
		if i == selected-1 {
			box := lipgloss.NewStyle().
				Foreground(lipgloss.Color(decl.Orange)).
				Render(">> ", list[i])
			some += box
			some += "\n"
		} else {
			some += list[i] + "\n"
		}
	}

	return
}
func many_choices_dbl_arrow_purple(list []string, selected int) (some string) {
	for i := 0; i < len(list); i++ {
		if i == selected-1 {
			box := lipgloss.NewStyle().
				Foreground(lipgloss.Color(decl.Magenta)).
				Render(list[i])
			some += box
			some += "\n"
		} else {
			some += list[i] + "\n"
		}
	}

	return
}

func many_choices_arrow(list []string, selected int) (some string) {
	for i := 0; i < len(list); i++ {
		if i == selected-1 {
			box := lipgloss.NewStyle().
				Foreground(lipgloss.Color(decl.Green)).
				Render(list[i])
			some += box
			some += "\n"
		} else {
			some += list[i] + "\n"
		}
	}

	return
}

func many_choices_arrow_red(list []string, selected int) (some string) {
	for i := 0; i < len(list); i++ {
		if i == selected-1 {
			box := lipgloss.NewStyle().
				Bold(true).
				Foreground(lipgloss.Color(decl.Red)).
				Render(list[i])
			some += box
			some += "\n"
		} else {
			some += list[i] + "\n"
		}
	}

	return
}
func many_choices_arrow_orange(list []string, selected int) (some string) {
	for i := 0; i < len(list); i++ {
		if i == selected-1 {
			box := lipgloss.NewStyle().
				Foreground(lipgloss.Color(decl.Orange)).
				Render(list[i])
			some += box
			some += "\n"
		} else {
			some += list[i] + "\n"
		}
	}

	return
}
func many_choices_arrow_purple(list []string, selected int) (some string) {
	for i := 0; i < len(list); i++ {
		if i == selected-1 {
			box := lipgloss.NewStyle().
				Foreground(lipgloss.Color(decl.Magenta)).
				Render(list[i])
			some += box
			some += "\n"
		} else {
			some += list[i] + "\n"
		}
	}

	return
}

func many_choices_simple(list []string, selected int) (some string) {
	for i := 0; i < len(list); i++ {
		if i == selected-1 {
			box := lipgloss.NewStyle().
				Background(lipgloss.Color(decl.White)).
				Foreground(lipgloss.Color(decl.Black)).
				Render(list[i])
			some += box
			some += "\n"
		} else {
			some += list[i] + "\n"
		}
	}

	return
}

func many_choices_simple_orange(list []string, selected int) (some string) {
	for i := 0; i < len(list); i++ {
		if i == selected-1 {
			box := lipgloss.NewStyle().
				Background(lipgloss.Color(decl.Orange)).
				Foreground(lipgloss.Color(decl.Black)).
				Bold(true).
				Render(list[i])
			some += box
			some += "\n"
		} else {
			some += list[i] + "\n"
		}
	}

	return
}

func many_choices_simple_green(list []string, selected int) (some string) {
	for i := 0; i < len(list); i++ {
		if i == selected-1 {
			box := lipgloss.NewStyle().
				Background(lipgloss.Color(decl.Green)).
				Foreground(lipgloss.Color(decl.Black)).
				Bold(true).
				Render(list[i])
			some += box
			some += "\n"
		} else {
			some += list[i] + "\n"
		}
	}

	return
}

func many_choices_simple_red(list []string, selected int) (some string) {
	for i := 0; i < len(list); i++ {
		if i == selected-1 {
			box := lipgloss.NewStyle().
				Background(lipgloss.Color(decl.Red)).
				Foreground(lipgloss.Color(decl.Black)).
				Bold(true).
				Render(list[i])
			some += box
			some += "\n"
		} else {
			some += list[i] + "\n"
		}
	}

	return
}

// This function is used to make prettier strings
//
// Example:
// Make("hello", Bold|Underline, "#ff0000", "")
func Make(some string, flag int, fg, bg string) (final string) {
	final = lipgloss.NewStyle().
		Bold(flag&decl.Bold != 0).
		Italic(flag&decl.Italic != 0).
		Strikethrough(flag&decl.Strike != 0).
		Underline(flag&decl.Underline != 0).
		Foreground(lipgloss.Color(fg)).
		Background(lipgloss.Color(bg)).
		Render(some)

	return
}

// This function uses HTML-style strings to make prettier strings.
// It is similar to Make but is useful when u need to embed many styles
// within the same string.
//
// Example:
// `
// This text is <ui> Underlined + italic </ui>
// Every <uisb> tag must be closed without mistake </uisb>
// You can also have <fg="#00ff00"> Different color text </fg>
// You can <i fg="#0f0f0f" bg="#ff00f0"> use any combinations of tags </i>
// Always remember the order <b> First u,i,s,b then fg then bg </b>
// You may <i fg="#ff00f0"> omit any tag </i>
// `
//
// Caution:
// Always follow this order:
// First any combination of uisb(any can be ommitted)
// Then fg=COLOR, then bg=COLOR
// If you only use colors, you may omit </fg> and use </>
func Convert(input string) string {
	var builder strings.Builder

	re := regexp.MustCompile(`<([ubis]*)\s*(fg="(#[0-9a-fA-F]{6})")?\s*(bg="(#[0-9a-fA-F]{6})")?\s*>`)

	for {
		loc := re.FindStringSubmatchIndex(input)
		if loc == nil {
			builder.WriteString(input)
			break
		}

		builder.WriteString(input[:loc[0]])

		styleStr := ""
		if loc[2] != -1 && loc[3] != -1 {
			styleStr = input[loc[2]:loc[3]]
		}

		fg := ""
		bg := ""

		if loc[6] != -1 && loc[7] != -1 {
			fg = input[loc[6]:loc[7]]
		}
		if loc[10] != -1 && loc[11] != -1 {
			bg = input[loc[10]:loc[11]]
		}

		input = input[loc[1]:]

		endTag := "</" + styleStr + ">"
		closeIndex := strings.Index(input, endTag)
		if closeIndex == -1 {
			break
		}
		content := input[:closeIndex]
		input = input[closeIndex+len(endTag):]

		style := lipgloss.NewStyle()
		for _, ch := range styleStr {
			switch ch {
			case 'b':
				style = style.Bold(true)
			case 'i':
				style = style.Italic(true)
			case 'u':
				style = style.Underline(true)
			case 's':
				style = style.Strikethrough(true)
			}
		}
		if fg != "" {
			style = style.Foreground(lipgloss.Color(fg))
		}
		if bg != "" {
			style = style.Background(lipgloss.Color(bg))
		}

		builder.WriteString(style.Render(content))
	}

	return builder.String()
}

// This function is used to pad two strings to same width
func Pad(this string, to string, how int) (some string) {
	switch how {
	case decl.LEFT_:
		some = this + strings.Repeat(" ", len(to)-len(this))

	case decl.RIGHT_:
		some = pad_rt(this, to)

	case decl.CENTER:
		some = pad_center(this, to)
	}
	return
}

func pad_rt(this, to string) (some string) {
	diff := len(to) - len(this)
	if diff <= 0 {
		return this
	}
	padding := strings.Repeat(" ", diff)
	return padding + this
}

func pad_center(this, to string) (some string) {
	diff := len(to) - len(this)
	if diff <= 0 {
		return this
	}
	left := diff / 2
	right := diff - left
	return strings.Repeat(" ", left) + this + strings.Repeat(" ", right)
}
