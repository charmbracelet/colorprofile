package colorprofile

import (
	"bytes"
	"image/color"
	"io"
	"strconv"

	"github.com/charmbracelet/x/ansi"
)

// NewWriter creates a new color profile writer that downgrades color sequences
// based on the detected color profile.
//
// If environ is nil, it will use os.Environ() to get the environment variables.
//
// It queries the given writer to determine if it supports ANSI escape codes.
// If it does, along with the given environment variables, it will determine
// the appropriate color profile to use for color formatting.
//
// This respects the NO_COLOR, CLICOLOR, and CLICOLOR_FORCE environment variables.
func NewWriter(w io.Writer, environ []string) *Writer {
	return &Writer{
		Forward: w,
		Profile: Detect(w, environ),
	}
}

// Writer represents a color profile writer that writes ANSI sequences to the
// underlying writer.
type Writer struct {
	Forward io.Writer
	Profile Profile
}

// Write writes the given text to the underlying writer.
func (w *Writer) Write(p []byte) (int, error) {
	switch w.Profile {
	case TrueColor:
		return w.Forward.Write(p)
	case NoTTY:
		return io.WriteString(w.Forward, ansi.Strip(string(p)))
	default:
		return w.downsample(p)
	}
}

// downsample downgrades the given text to the appropriate color profile.
func (w *Writer) downsample(p []byte) (int, error) {
	var buf bytes.Buffer
	var state byte

	parser := ansi.GetParser()
	defer ansi.PutParser(parser)

	for len(p) > 0 {
		parser.Reset()
		seq, _, read, newState := ansi.DecodeSequence(p, state, parser)

		switch {
		case ansi.HasCsiPrefix(seq) && parser.Cmd == 'm':
			handleSgr(w, parser, &buf)
		default:
			// If we're not a style SGR sequence, just write the bytes.
			if n, err := buf.Write(seq); err != nil {
				return n, err
			}
		}

		p = p[read:]
		state = newState
	}

	return w.Forward.Write(buf.Bytes())
}

// WriteString writes the given text to the underlying writer.
func (w *Writer) WriteString(s string) (n int, err error) {
	return w.Write([]byte(s))
}

func handleSgr(w *Writer, parser *ansi.Parser, buf *bytes.Buffer) {
	var style ansi.Style
	for i := 0; i < parser.ParamsLen; i++ {
		param := ansi.Param(parser.Params[i])

		switch param := param.Param(); param {
		default:
			// If this is not a color attribute, just append it to the style.
			style = append(style, strconv.Itoa(param))
		case 30, 31, 32, 33, 34, 35, 36, 37: // 8-bit foreground color
			if w.Profile > ANSI {
				continue
			}
			style = style.ForegroundColor(
				w.Profile.Convert(ansi.BasicColor(param - 30))) //nolint:gosec
		case 38: // 16 or 24-bit foreground color
			c := readColor(&i, parser.Params)
			if w.Profile > ANSI {
				continue
			}
			style = style.ForegroundColor(w.Profile.Convert(c))
		case 39: // default foreground color
			if w.Profile > ANSI {
				continue
			}
			style = style.DefaultForegroundColor()
		case 40, 41, 42, 43, 44, 45, 46, 47: // 8-bit background color
			if w.Profile > ANSI {
				continue
			}
			style = style.BackgroundColor(
				w.Profile.Convert(ansi.BasicColor(param - 40))) //nolint:gosec
		case 48: // 16 or 24-bit background color
			c := readColor(&i, parser.Params)
			if w.Profile > ANSI {
				continue
			}
			style = style.BackgroundColor(w.Profile.Convert(c))
		case 49: // default background color
			if w.Profile > ANSI {
				continue
			}
			style = style.DefaultBackgroundColor()
		case 58: // 16 or 24-bit underline color
			c := readColor(&i, parser.Params)
			if w.Profile > ANSI {
				continue
			}
			style = style.UnderlineColor(w.Profile.Convert(c))
		case 59: // default underline color
			if w.Profile > ANSI {
				continue
			}
			style = style.DefaultUnderlineColor()
		case 90, 91, 92, 93, 94, 95, 96, 97: // 8-bit bright foreground color
			if w.Profile > ANSI {
				continue
			}
			style = style.ForegroundColor(
				w.Profile.Convert(ansi.BasicColor(param - 90 + 8))) //nolint:gosec
		case 100, 101, 102, 103, 104, 105, 106, 107: // 8-bit bright background color
			if w.Profile > ANSI {
				continue
			}
			style = style.BackgroundColor(
				w.Profile.Convert(ansi.BasicColor(param - 100 + 8))) //nolint:gosec
		}
	}

	_, _ = buf.WriteString(style.String())
}

func readColor(idxp *int, params []int) (c ansi.Color) {
	i := *idxp
	paramsLen := len(params)
	if i > paramsLen-1 {
		return
	}
	// Note: we accept both main and subparams here
	switch param := ansi.Param(params[i+1]); param.Param() {
	case 2: // RGB
		if i > paramsLen-4 {
			return
		}
		c = color.RGBA{
			R: uint8(ansi.Param(params[i+2]).Param()), //nolint:gosec
			G: uint8(ansi.Param(params[i+3]).Param()), //nolint:gosec
			B: uint8(ansi.Param(params[i+4]).Param()), //nolint:gosec
			A: 0xff,
		}
		*idxp += 4
	case 5: // 256 colors
		if i > paramsLen-2 {
			return
		}
		c = ansi.ExtendedColor(ansi.Param(params[i+2]).Param()) //nolint:gosec
		*idxp += 2
	}
	return
}
