package formatter

import (
	"testing"
)

func TestTruncate(t *testing.T) {
	tests := []struct {
		name   string
		text   string
		length uint
		want   string
	}{
		{name: "raw", text: "onetwo", length: 6, want: "onetwo"},
		{name: "truncated", text: "onetwothree", length: 6, want: "one..."},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Truncate(tt.text, tt.length); got != tt.want {
				t.Errorf(`Truncate("%s", %d) = "%s", want "%s"`, tt.text, tt.length, got, tt.want)
			}
		})
	}
}
