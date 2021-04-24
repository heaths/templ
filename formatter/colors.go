package formatter

import (
	"os"

	"github.com/mattn/go-isatty"
	"github.com/mgutz/ansi"
)

var (
	green  = ansi.ColorFunc("green")
	red    = ansi.ColorFunc("red")
	yellow = ansi.ColorFunc("yellow")
)

type Colors struct {
	enabled bool
}

func SystemColors() Colors {
	return Colors{
		enabled: isTerminal(os.Stdout),
	}
}

func (c *Colors) Green(s string) string {
	if c.enabled {
		return green(s)
	}

	return s
}

func (c *Colors) Red(s string) string {
	if c.enabled {
		return red(s)
	}

	return s
}

func (c *Colors) Yellow(s string) string {
	if c.enabled {
		return yellow(s)
	}

	return s
}

func isTerminal(f *os.File) bool {
	return isatty.IsTerminal(f.Fd()) || isatty.IsCygwinTerminal(f.Fd())
}
