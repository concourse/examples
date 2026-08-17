package main

import (
	"reflect"
	"sort"
	"strings"
	"testing"
)

func TestLookup(t *testing.T) {
	tests := []struct {
		name string
		word string
		want bool
	}{
		{"exact match", "concourse", true},
		{"uppercase", "PLANE", true},
		{"mixed case", "PiLoT", true},
		{"surrounding whitespace", "  plane  ", true},
		{"unknown word", "nope", false},
		{"empty string", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			art, ok := lookup(tt.word)
			if ok != tt.want {
				t.Fatalf("lookup(%q) ok = %v, want %v", tt.word, ok, tt.want)
			}
			if !ok {
				if art != "" {
					t.Errorf("lookup(%q) returned art %q for a miss, want empty", tt.word, art)
				}
				return
			}
			if art == "" {
				t.Errorf("lookup(%q) returned empty art for a hit", tt.word)
			}
			if strings.HasPrefix(art, "\n") || strings.HasSuffix(art, "\n") {
				t.Errorf("lookup(%q) art should be trimmed of leading/trailing newlines", tt.word)
			}
		})
	}
}

func TestWithBorder(t *testing.T) {
	t.Run("single line", func(t *testing.T) {
		got := withBorder("abc")
		want := strings.Join([]string{
			"+-----+",
			"| abc |",
			"+-----+",
		}, "\n")
		if got != want {
			t.Errorf("withBorder(\"abc\") =\n%s\nwant\n%s", got, want)
		}
	})

	t.Run("ragged lines are padded to the widest", func(t *testing.T) {
		got := withBorder("a\nbbbb\ncc")
		want := strings.Join([]string{
			"+------+",
			"| a    |",
			"| bbbb |",
			"| cc   |",
			"+------+",
		}, "\n")
		if got != want {
			t.Errorf("withBorder ragged =\n%s\nwant\n%s", got, want)
		}
	})

	t.Run("width is measured in runes not bytes", func(t *testing.T) {
		// ⣿ is 3 bytes but one column; the border must line up on rune count.
		got := withBorder("⣿⣿")
		want := strings.Join([]string{
			"+----+",
			"| ⣿⣿ |",
			"+----+",
		}, "\n")
		if got != want {
			t.Errorf("withBorder(\"⣿⣿\") =\n%s\nwant\n%s", got, want)
		}
	})

	t.Run("every row has equal display width", func(t *testing.T) {
		bordered := withBorder(lookupOrFatal(t, "pilot"))
		lines := strings.Split(bordered, "\n")
		width := len([]rune(lines[0]))
		for i, line := range lines {
			if n := len([]rune(line)); n != width {
				t.Errorf("line %d has %d runes, want %d: %q", i, n, width, line)
			}
		}
		if !strings.HasPrefix(lines[0], "+") || !strings.HasSuffix(lines[0], "+") {
			t.Errorf("top border = %q, want +...+", lines[0])
		}
	})
}

func lookupOrFatal(t *testing.T, word string) string {
	t.Helper()
	art, ok := lookup(word)
	if !ok {
		t.Fatalf("lookup(%q) failed", word)
	}
	return art
}

func TestNames(t *testing.T) {
	got := names()

	if len(got) != len(images) {
		t.Errorf("names() returned %d entries, want %d", len(got), len(images))
	}

	if !sort.StringsAreSorted(got) {
		t.Errorf("names() is not sorted: %v", got)
	}

	want := []string{"concourse", "pilot", "plane"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("names() = %v, want %v", got, want)
	}

	// Every name should resolve via lookup.
	for _, n := range got {
		if _, ok := lookup(n); !ok {
			t.Errorf("names() returned %q but lookup(%q) failed", n, n)
		}
	}
}
