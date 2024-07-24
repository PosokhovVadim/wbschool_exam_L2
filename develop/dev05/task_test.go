package main

import (
	"reflect"
	"testing"
)

func TestGrepProcess(t *testing.T) {
	tests := []struct {
		name     string
		lines    []string
		pattern  string
		flags    map[string]interface{}
		expected []string
	}{
		{
			name:    "simple match",
			lines:   []string{"hello world", "foo bar", "hello again"},
			pattern: "hello",
			flags:   map[string]interface{}{},
			expected: []string{
				"hello world",
				"hello again",
			},
		},
		{
			name:    "ignore case",
			lines:   []string{"Hello world", "foo bar", "HELLO again"},
			pattern: "hello",
			flags:   map[string]interface{}{"i": true},
			expected: []string{
				"Hello world",
				"HELLO again",
			},
		},
		{
			name:    "invert match",
			lines:   []string{"hello world", "foo bar", "hello again"},
			pattern: "hello",
			flags:   map[string]interface{}{"v": true},
			expected: []string{
				"foo bar",
			},
		},
		{
			name:    "fixed string match",
			lines:   []string{"hello world", "foo bar", "hello again"},
			pattern: "foo",
			flags:   map[string]interface{}{"F": true},
			expected: []string{
				"foo bar",
			},
		},
		{
			name:    "line number",
			lines:   []string{"hello world", "foo bar", "hello again"},
			pattern: "hello",
			flags:   map[string]interface{}{"n": true},
			expected: []string{
				"1: hello world",
				"3: hello again",
			},
		},
		{
			name:    "context lines",
			lines:   []string{"hello world", "foo bar", "baz qux", "hello again", "another line"},
			pattern: "baz",
			flags:   map[string]interface{}{"C": 1},
			expected: []string{
				"foo bar",
				"baz qux",
				"hello again",
			},
		},
		{
			name:    "before lines",
			lines:   []string{"hello world", "foo bar", "baz qux", "hello again", "another line"},
			pattern: "hello",
			flags:   map[string]interface{}{"B": 1},
			expected: []string{
				"hello world",
				"baz qux",
				"hello again",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for key, value := range tt.flags {
				switch key {
				case "A":
					a = value.(int)
				case "B":
					b = value.(int)
				case "C":
					ctx = value.(int)
				case "c":
					count = value.(bool)
				case "i":
					i = value.(bool)
				case "v":
					v = value.(bool)
				case "F":
					f = value.(bool)
				case "n":
					n = value.(bool)
				}
			}

			result := grepProcess(tt.lines, tt.pattern)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
			resetFlags()
		})
	}
}

func resetFlags() {
	a = 0
	b = 0
	ctx = 0
	count = false
	i = false
	v = false
	f = false
	n = false
}
