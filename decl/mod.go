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

package decl

const (
	Black      = "#000000"
	White      = "#ffffff"
	Red        = "#ff0000"
	Green      = "#00ff00"
	Blue       = "#0000ff"
	Yellow     = "#ffff00"
	Cyan       = "#00ffff"
	Magenta    = "#ff00ff"
	Gray       = "#808080"
	LightGray  = "#d3d3d3"
	DarkGray   = "#404040"
	Orange     = "#ffa500"
	Purple     = "#800080"
	Violet     = "#8a2be2"
	Pink       = "#ff69b4"
	Teal       = "#008080"
	Indigo     = "#4b0082"
	Brown      = "#a52a2a"
	Maroon     = "#800000"
	Lime       = "#bfff00"
	Olive      = "#808000"
	Navy       = "#000080"
	Coral      = "#ff7f50"
	Gold       = "#ffd700"
	Salmon     = "#fa8072"
	Crimson    = "#dc143c"
	Khaki      = "#f0e68c"
	Aqua       = "#00ffff"
	Mint       = "#98ff98"
	Silver     = "#c0c0c0"
	Beige      = "#f5f5dc"
	Turquoise  = "#40e0d0"
	Chartreuse = "#7fff00"
	Plum       = "#dda0dd"
	Orchid     = "#da70d6"
	Tomato     = "#ff6347"
	Azure      = "#f0ffff"
	Lavender   = "#e6e6fa"
	None       = ""

	// RIGHT_ is not written RIGHT because it is redeclared in
	// keyboard captures(same with LEFT)
	CENTER = iota + 69
	LEFT_
	RIGHT_
	TOP
	BOTTOM

	// Load screen patterns
	SIMPLE
	SIMPLE_BLOCKS
	SIMPLE_EQUAL
	AT

	//... <COLOR>_<BLOCKTYPE> generate colored string
	//... using ANSII codes
	RED_BLOCKS
	YELLOW_BLOCKS
	GREEN_BLOCKS
	BLUE_BLOCKS
	ORANGE_BLOCKS

	RED_EQUAL
	YELLOW_EQUAL
	GREEN_EQUAL
	BLUE_EQUAL
	ORANGE_EQUAL

	RED_AT
	YELLOW_AT
	GREEN_AT
	BLUE_AT
	ORANGE_AT

	RED_SIMPLE
	YELLOW_SIMPLE
	GREEN_SIMPLE
	BLUE_SIMPLE
	ORANGE_SIMPLE

	// Chose from a list of options
	SIMPLE_CHOICES
	ARROW_CHOICES
	DOUBLE_ARROW_CHOICES
	HORIZONTAL_CHOICES

	// The declarations for load_screen(), i.e ORANGE_SIMPLE, GREEN_SIMPLE, RED_SIMPLE
	// are recycled for SIMPLE_CHOICES
	ORANGE_ARROW
	RED_ARROW
	PURPLE_ARROW

	ORANGE_DOUBLE_ARROW
	RED_DOUBLE_ARROW
	PURPLE_DOUBLE_ARROW

	GREEN_HORIZONTAL_CHOICES
	RED_HORIZONTAL_CHOICES
	ORANGE_HORIZONTAL_CHOICES
)

// These declarations are for keyboard captures
// Note: Although many key captures are provided it is suggested
// to use only a core subset of them. For a better experience, do not
// use [, ], {, }, !, ~, as controls although they work
const (
	// Control keys
	ESC   = 27
	SPACE = 32

	// Arrow keys (used in ANSI escape sequences with ESC)
	UP    = 65
	DOWN  = 66
	RIGHT = 67
	LEFT  = 68

	// Digits
	ZERO  = 48
	ONE   = 49
	TWO   = 50
	THREE = 51
	FOUR  = 52
	FIVE  = 53
	SIX   = 54
	SEVEN = 55
	EIGHT = 56
	NINE  = 57

	// Uppercase letters
	A = 65
	B = 66
	C = 67
	D = 68
	E = 69
	F = 70
	G = 71
	H = 72
	I = 73
	J = 74
	K = 75
	L = 76
	M = 77
	N = 78
	O = 79
	P = 80
	Q = 81
	R = 82
	S = 83
	T = 84
	U = 85
	V = 86
	W = 87
	X = 88
	Y = 89
	Z = 90

	// Lowercase letters
	SMALL_a = 97
	SMALL_b = 98
	SMALL_c = 99
	SMALL_d = 100
	SMALL_e = 101
	SMALL_f = 102
	SMALL_g = 103
	SMALL_h = 104
	SMALL_i = 105
	SMALL_j = 106
	SMALL_k = 107
	SMALL_l = 108
	SMALL_m = 109
	SMALL_n = 110
	SMALL_o = 111
	SMALL_p = 112
	SMALL_q = 113
	SMALL_r = 114
	SMALL_s = 115
	SMALL_t = 116
	SMALL_u = 117
	SMALL_v = 118
	SMALL_w = 119
	SMALL_x = 120
	SMALL_y = 121
	SMALL_z = 122

	// Punctuation and symbols
	EXCLAMATION   = 33  // !
	DOUBLE_QUOTE  = 34  // "
	HASH          = 35  // #
	DOLLAR        = 36  // $
	PERCENT       = 37  // %
	AMPERSAND     = 38  // &
	SINGLE_QUOTE  = 39  // '
	LEFT_PAREN    = 40  // (
	RIGHT_PAREN   = 41  // )
	ASTERISK      = 42  // *
	PLUS          = 43  // +
	COMMA         = 44  // ,
	MINUS         = 45  // -
	DOT           = 46  // .
	SLASH         = 47  // /
	COLON         = 58  // :
	SEMICOLON     = 59  // ;
	LESS_THAN     = 60  // <
	EQUAL         = 61  // =
	GREATER_THAN  = 62  // >
	QUESTION_MARK = 63  // ?
	AT_           = 64  // @
	LEFT_BRACKET  = 91  // [
	BACKSLASH     = 92  // \
	RIGHT_BRACKET = 93  // ]
	CARET         = 94  // ^
	UNDERSCORE    = 95  // _
	BACKTICK      = 96  // `
	LEFT_CURLY    = 123 // {
	VERTICAL_BAR  = 124 // |
	RIGHT_CURLY   = 125 // }
	TILDE         = 126 // ~
)

const (
	Italic    = 0x01
	Bold      = 0x02
	Underline = 0x04
	Strike    = 0x08
)
