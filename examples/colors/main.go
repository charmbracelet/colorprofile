package main

import (
	"fmt"

	"github.com/charmbracelet/x/ansi"
)

func printBlock(c ansi.Color, fg ansi.Color) {
	r, g, b, _ := c.RGBA()
	block := ansi.NewStyle().BackgroundColor(c).ForegroundColor(fg)
	fmt.Print(block.Styled(fmt.Sprintf(" #%02X%02X%02X ", r>>8, g>>8, b>>8)))
}

func main() {
	header := ansi.NewStyle().Bold()
	fmt.Println(header.Styled("Basic ANSI colors"))
	for i := range 16 {
		c := ansi.BasicColor(i)
		fg := ansi.Black
		if i == 0 {
			fg = ansi.White
		}
		printBlock(c, fg)
		if i == 7 {
			fmt.Println()
		}
	}

	fmt.Println()
	fmt.Println()

	fmt.Println(header.Styled("256 ANSI colors"))
	for i := 16; i < 232; i++ {
		c := ansi.IndexedColor(i)
		fg := ansi.Black
		if i >= 232 {
			fg = ansi.White
		}
		printBlock(c, fg)
		if (i-15)%6 == 0 {
			fmt.Println()
		}
	}

	fmt.Println()
	fmt.Println()

	fmt.Println(header.Styled("256 ANSI grayscale colors"))
	for i := 232; i < 256; i++ {
		c := ansi.IndexedColor(i)
		fg := ansi.White
		printBlock(c, fg)
		if (i-231)%6 == 0 {
			fmt.Println()
		}
	}

	fmt.Println()
}
