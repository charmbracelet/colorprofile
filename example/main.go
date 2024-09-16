package main

import (
	"fmt"
	"image/color"
	"os"

	"github.com/charmbracelet/colorprofile"
)

func main() {
	// Detect the color profile for stdout.
	p := colorprofile.Detect(os.Stdout, os.Environ())
	fmt.Printf("Your color profile is what we call '%s'.\n\n", p)

	// Let's talk about the profile.
	fmt.Printf("You know, your colors are quite %s.\n\n", func() string {
		switch p {
		case colorprofile.TrueColor:
			return "fancy"
		case colorprofile.ANSI256:
			return "1990s fancy"
		case colorprofile.ANSI:
			return "normcore"
		case colorprofile.Ascii:
			return "ancient"
		case colorprofile.NoTTY:
			return "naughty!"
		}
		// This should never happen.
		return "...IDK"
	}())

	// Here's a nice color.
	myCuteColor := color.RGBA{0x6b, 0x50, 0xff, 0xff} // #6b50ff
	fmt.Printf("A cute color we like is: %s.\n\n", colorToHex(myCuteColor))

	// Let's convert it to the detected color profile.
	theColorWeNeed := p.Convert(myCuteColor)
	fmt.Printf("This terminal needs that color to be a %T, at best.\n", theColorWeNeed)
	fmt.Printf("In this case that color is %s.\n\n", colorToHex(theColorWeNeed))

	// Now let's convert it to a color profile that only supports up to 256
	// colors.
	ansi256Color := colorprofile.ANSI256.Convert(myCuteColor)
	fmt.Printf("Apple Terminal would want this color to be: %d (an %T).\n\n", ansi256Color, ansi256Color)

	// But really, who has time to convert? Not you? Well, kiddo, here's
	// a magical writer that will just auto-convert whatever ANSI you throw at
	// it to the appropriate color profile.
	myFancyANSI := "\x1b[38;2;107;80;255mCute puppy!!\x1b[m"
	w := colorprofile.NewWriter(os.Stdout, os.Environ())
	w.Println("This terminal:", myFancyANSI)

	// But we're old school. Make the writer only use 4-bit ANSI, 1980s style.
	w.Profile = colorprofile.ANSI
	w.Println("4-bit ANSI:", myFancyANSI)

	// That's way too modern. Let's go back to MIT in the 1970s.
	w.Profile = colorprofile.NoTTY
	w.Println("No TTY :(", myFancyANSI) // less fancy
}

func colorToHex(c color.Color) string {
	r, g, b, _ := c.RGBA()
	return fmt.Sprintf("#%02x%02x%02x", r>>8, g>>8, b>>8)
}
