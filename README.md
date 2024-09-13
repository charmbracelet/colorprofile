# Color Profile

<p>
    <a href="https://github.com/charmbracelet/colorprofile/releases"><img src="https://img.shields.io/github/release/charmbracelet/colorprofile.svg" alt="Latest Release"></a>
    <a href="https://pkg.go.dev/github.com/charmbracelet/colorprofile?tab=doc"><img src="https://godoc.org/github.com/charmbracelet/colorprofile?status.svg" alt="GoDoc"></a>
    <a href="https://github.com/charmbracelet/colorprofile/actions"><img src="https://github.com/charmbracelet/colorprofile/actions/workflows/build.yml/badge.svg" alt="Build Status"></a>
    <a href="https://www.phorm.ai/query?projectId=a0e324b6-b706-4546-b951-6671ea60c13f"><img src="https://stuff.charm.sh/misc/phorm-badge.svg" alt="phorm.ai"></a>
</p>

Color Profile is a Go package for working with terminal color profiles and
color degradation.

## Installation

```sh
go get github.com/charmbracelet/colorprofile@latest
```

## Usage

```go
package main

import (
	"fmt"
	"image/color"
	"os"

	"github.com/charmbracelet/colorprofile"
	"github.com/lucasb-eyer/go-colorful"
)

func printColor(profile colorprofile.Profile, c color.Color) {
	// Print the converted color in the terminal
	c = profile.Convert(c)
	info := fmt.Sprintf("%T(%v)", c, c)
	col, _ := colorful.MakeColor(c)
	fmt.Println("This is a nice color:", col.Hex(), info)
}

func main() {
	// Get the terminal's color profile
	profile := colorprofile.Detect(os.Stdout, os.Environ())

	// Convert 24-bit RGB color to the terminal's color profile.
	// This will return the closest color in the profile's palette
	// if the terminal doesn't support 24-bit color.
	mycolor := color.RGBA{0xff, 0x7b, 0xf5, 0xff} // #ff7bf5
	printColor(profile, mycolor)

	// Use ANSI256 color profile
	printColor(colorprofile.ANSI256, mycolor)
}
```

## Feedback

We’d love to hear your thoughts on this project. Feel free to drop us a note!

- [Twitter](https://twitter.com/charmcli)
- [The Fediverse](https://mastodon.social/@charmcli)
- [Discord](https://charm.sh/chat)

---

Part of [Charm](https://charm.sh).

<a href="https://charm.sh/"><img alt="The Charm logo" src="https://stuff.charm.sh/charm-badge.jpg" width="400"></a>

Charm热爱开源 • Charm loves open source • نحنُ نحب المصادر المفتوحة

