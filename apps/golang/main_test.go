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

func TestColorize(t *testing.T) {
	const art = "hello"

	tests := []struct {
		name    string
		color   string
		want    string
		wantErr bool
	}{
		{"empty leaves art untouched", "", art, false},
		{"red", "red", "\x1b[31mhello\x1b[0m", false},
		{"green", "green", "\x1b[32mhello\x1b[0m", false},
		{"blue", "blue", "\x1b[34mhello\x1b[0m", false},
		{"case insensitive", "ReD", "\x1b[31mhello\x1b[0m", false},
		{"surrounding whitespace", "  blue  ", "\x1b[34mhello\x1b[0m", false},
		{"unknown color", "purple", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := colorize(art, tt.color)
			if (err != nil) != tt.wantErr {
				t.Fatalf("colorize(%q, %q) err = %v, wantErr %v", art, tt.color, err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("colorize(%q, %q) = %q, want %q", art, tt.color, got, tt.want)
			}
		})
	}
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
