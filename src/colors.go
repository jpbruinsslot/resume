package main

import "github.com/mgutz/ansi"

// Color(s, "red")            // red
// Color(s, "red+b")          // red bold
// Color(s, "red+B")          // red blinking
// Color(s, "red+u")          // red underline
// Color(s, "red+bh")         // red bold bright
// Color(s, "red:white")      // red on white
// Color(s, "red+b:white+h")  // red bold on white bright
// Color(s, "red+B:white+h")  // red blink on white bright
// Color(s, "off")            // turn off ansi codes

// "foregroundColor+attributes:backgroundColor+attributes"

// Colors
//
// 	black
// 	red
// 	green
// 	yellow
// 	blue
// 	magenta
// 	cyan
// 	white
//
// Attributes
//
// 	b = bold foreground
// 	B = Blink foreground
// 	u = underline foreground
// 	i = inverse
// 	h = high intensity (bright) foreground, background

type Colors struct {
	Black        string
	Red          string
	Green        string
	Yellow       string
	Blue         string
	Magenta      string
	Cyan         string
	White        string
	LightBlack   string
	LightRed     string
	LightGreen   string
	LightYellow  string
	LightBlue    string
	LightMagenta string
	LightCyan    string
	LightWhite   string
	Bold         string
	Underline    string
	Blink        string
	Inverse      string
	Bright       string
	Reset        string
}

func GetANSIColors() Colors {
	return Colors{
		Black:        ansi.Black,
		Red:          ansi.Red,
		Green:        ansi.Green,
		Yellow:       ansi.Yellow,
		Blue:         ansi.Blue,
		Magenta:      ansi.Magenta,
		Cyan:         ansi.Cyan,
		White:        ansi.White,
		LightBlack:   ansi.LightBlack,
		LightRed:     ansi.LightRed,
		LightGreen:   ansi.LightGreen,
		LightYellow:  ansi.LightYellow,
		LightBlue:    ansi.LightBlue,
		LightMagenta: ansi.LightMagenta,
		LightCyan:    ansi.LightCyan,
		LightWhite:   ansi.LightWhite,
		Bold:         ansi.ColorCode("off+b"),
		Underline:    ansi.ColorCode("off+u"),
		Blink:        ansi.ColorCode("off+B"),
		Inverse:      ansi.ColorCode("off+i"),
		Bright:       ansi.ColorCode("off+h"),
		Reset:        ansi.ColorCode("reset"),
	}
}
