# Color Profile

<p>
    <a href="https://github.com/charmbracelet/colorprofile/releases"><img src="https://img.shields.io/github/release/charmbracelet/colorprofile.svg" alt="Latest Release"></a>
    <a href="https://pkg.go.dev/github.com/charmbracelet/colorprofile?tab=doc"><img src="https://godoc.org/github.com/charmbracelet/colorprofile?status.svg" alt="GoDoc"></a>
    <a href="https://github.com/charmbracelet/colorprofile/actions"><img src="https://github.com/charmbracelet/colorprofile/actions/workflows/build.yml/badge.svg" alt="Build Status"></a>
</p>

A simple, powerful package for detecting terminal color profiles, CSI
sequences, and performing color degradation.

## Example

Import the lib.
```go
import "github.com/charmbracelet/colorprofile"
```

Detect the color profile.
```go
p := colorprofile.Detect(os.Stdout, os.Environ())
fmt.Printf("Your color profile is what we call '%s'.", p)
```

Let's talk about the profile.
```go
fmt.Printf("You know, your colors are quite %s.", func() string {
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
    return "...IDK" // this should never happen
}())
```

Downsample a color to the detected profile, when necessary.
```go
c := color.RGBA{0x6b, 0x50, 0xff, 0xff} // #6b50ff
convertedColor := p.Convert(c)
```

Convert it to ANSI256.
```go
ansi256Color := colorprofile.ANSI256.Convert(c)
```

Magically downsample any ANSI to the detected profile, when necessary.
```go
fancyANSI := "\x1b[38;2;107;80;255mCute puppy!!\x1b[m"
w := colorprofile.NewWriter(os.Stdout, os.Environ())
w.Printf(fancyANSI)
```

Magically downsample to 4-bit ANSI.
```go
w.Profile = colorprofile.ANSI
w.Printf(myFancyANSI)
```

Strip ANSI altogether.
```go
w.Profile = colorprofile.NoTTY
w.Printf(myFancyANSI) // not so fancy anymore
```

## Get it

```sh
go get "github.com/charmbracelet/colorprofile@latest"
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
