package formatter

import (
	"testing"

	"github.com/mgutz/ansi"
)

type colorFunc func(s string) string
type colorFuncFunc func(c *Colors) colorFunc

func TestColors(t *testing.T) {
	tests := []struct {
		name    string
		f       colorFuncFunc
		enabled bool
		want    string
	}{
		{
			name:    "green_disabled",
			f:       func(c *Colors) colorFunc { return c.Green },
			enabled: false,
			want:    "text",
		},
		{
			name:    "green_enabled",
			f:       func(c *Colors) colorFunc { return c.Green },
			enabled: true,
			want:    ansi.Color("text", "green"),
		},
		{
			name:    "red_disabled",
			f:       func(c *Colors) colorFunc { return c.Red },
			enabled: false,
			want:    "text",
		},
		{
			name:    "red_enabled",
			f:       func(c *Colors) colorFunc { return c.Red },
			enabled: true,
			want:    ansi.Color("text", "red"),
		},
		{
			name:    "yellow_disabled",
			f:       func(c *Colors) colorFunc { return c.Yellow },
			enabled: false,
			want:    "text",
		},
		{
			name:    "yellow_enabled",
			f:       func(c *Colors) colorFunc { return c.Yellow },
			enabled: true,
			want:    ansi.Color("text", "yellow"),
		},
	}

	for _, tt := range tests {
		c := Colors{
			enabled: tt.enabled,
		}

		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f(&c)("text"); got != tt.want {
				t.Errorf(`got "%s", want "%s"`, got, tt.want)
			}
		})
	}
}
