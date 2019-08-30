package sinebow

import (
	"fmt"
	"math"
	"strings"
	"unicode/utf8"
)

// ANSI escape code breakdown example:
//
// set fg color      bold     reset
//  ┌───┴───┐          │     ┌──┴──┐
//  \x1b[38;2;%d;%d;%d;1;7m%c\x1b[0m
//             └──┬──┘   │
//              r g b  invert

const (
	foreground = "\x1b[38;2;%d;%d;%d;1m%c\x1b[0m"
	background = "\x1b[38;2;%d;%d;%d;1;7m%c\x1b[0m"
)

// Rainbro sexifies a given string rune-by-rune using 3
// out-of-phase sine waves to create a rainbro effect.
func Rainbro(s string) string {
	return rainbro(s, foreground)
}

// RainbroBG inverts the colors to make the bg sexy
func RainbroBG(s string) string {
	return rainbro(s, background)
}

// (ﾉ☉ヮ⚆)ﾉ ⌒*:･ﾟ✧ https://krazydad.com/tutorials/makecolors.php
func rainbro(input, format string) string {
	var (
		sb = strings.Builder{}
		freq = (2*math.Pi)/float64(utf8.RuneCountInString(input))
	)

	for i, c := range input {
		_, _ = sb.WriteString(fmt.Sprintf(
			format,
			int(math.Round(math.Sin(freq*float64(i))*127+128)),
			int(math.Round(math.Sin(freq*float64(i)+(2*math.Pi)/3)*127+128)),
			int(math.Round(math.Sin(freq*float64(i)+(4*math.Pi)/3)*127+128)),
			c,
		))
	}

	return sb.String()
}