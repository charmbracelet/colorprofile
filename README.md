# Color Profile

<p>
    <a href="https://github.com/charmbracelet/colorprofile/releases"><img src="https://img.shields.io/github/release/charmbracelet/colorprofile.svg" alt="Latest Release"></a>
    <a href="https://pkg.go.dev/github.com/charmbracelet/colorprofile?tab=doc"><img src="https://godoc.org/github.com/charmbracelet/colorprofile?status.svg" alt="GoDoc"></a>
    <a href="https://github.com/charmbracelet/colorprofile/actions"><img src="https://github.com/charmbracelet/colorprofile/actions/workflows/build.yml/badge.svg" alt="Build Status"></a>
</p>

A simple, powerful package for detecting terminal color profiles, CSI
sequences, and performing color degradation.

```go
import "github.com/charmbracelet/colorprofile"

// Find the color profile for stdout.
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
fmt.Printf("A cute color we like is: #%x%x%x.\n\n", myCuteColor.R, myCuteColor.G, myCuteColor.B)

// Let's convert it to the detected color profile.
theColorWeNeed := p.Convert(myCuteColor)
fmt.Printf("This terminal needs it to be a %T, at best...\n", theColorWeNeed)
fmt.Printf("...which would be %#v.\n\n", theColorWeNeed)

// Now let's convert it to a color profile that only supports up to 256
// colors.
ansi256Color := colorprofile.ANSI256.Convert(myCuteColor)
fmt.Printf("Apple Terminal would want this color to be: %d (an %T).\n\n", ansi256Color, ansi256Color)

// But really, who has time to convert? Not you? Well, kiddo, here's
// a magical writer that will just auto-convert whatever ANSI you throw at
// it to the appropriate color profile.
myFancyANSI := "\x1b[38;2;107;80;255mCute puppy!!\x1b[m\n"
w := colorprofile.NewWriter(os.Stdout, os.Environ())
w.Printf(myFancyANSI)

// But we're old school. Make the writer only use 4-bit ANSI, 1980s style.
w.Profile = colorprofile.ANSI
w.Printf(myFancyANSI)

// That's too modern. Let's go back to MIT in the 1970s.
w.Profile = colorprofile.NoTTY
w.Printf(myFancyANSI) // not so fancy anymore
```

## Get it

```sh
go get github.com/charmbracelet/colorprofile@latest
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
