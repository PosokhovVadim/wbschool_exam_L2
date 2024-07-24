package main

import (
	"reflect"
	"testing"
)

func TestCut(t *testing.T) {
	tests := []struct {
		name      string
		input     []string
		fields    []int
		delimiter string
		separated bool
		expected  []string
	}{
		{
			name:      "select first and third fields with tab delimiter",
			input:     []string{"a\tb\tc", "d\te\tf", "g\th\ti"},
			fields:    []int{1, 3},
			delimiter: "\t",
			separated: false,
			expected:  []string{"a\tc", "d\tf", "g\ti"},
		},
		{
			name:      "select second field with comma delimiter",
			input:     []string{"a,b,c", "d,e,f", "g,h,i"},
			fields:    []int{2},
			delimiter: ",",
			separated: false,
			expected:  []string{"b", "e", "h"},
		},
		{
			name:      "select second field with comma delimiter and separated flag",
			input:     []string{"abc", "d,e,f", "ghij"},
			fields:    []int{2},
			delimiter: ",",
			separated: true,
			expected:  []string{"e"},
		},
		{
			name:      "select fourth field which does not exist",
			input:     []string{"a\tb\tc", "d\te\tf", "g\th\ti"},
			fields:    []int{4},
			delimiter: "\t",
			separated: false,
			expected:  []string{"", "", ""},
		},
		{
			name:      "select second field with comma delimiter from mixed content",
			input:     []string{"a,b,c", "d e f", "g,h,i"},
			fields:    []int{2},
			delimiter: ",",
			separated: false,
			expected:  []string{"b", "", "h"},
		},
		{
			name:      "select all fields with space delimiter",
			input:     []string{"a b c", "d e f", "g h i"},
			fields:    []int{1, 2, 3},
			delimiter: " ",
			separated: false,
			expected:  []string{"a b c", "d e f", "g h i"},
		},
		{
			name:      "select only lines with delimiter",
			input:     []string{"no delimiter", "has,delimiter", "another,one"},
			fields:    []int{2},
			delimiter: ",",
			separated: true,
			expected:  []string{"delimiter", "one"},
		},
		{
			name:      "select first field with space delimiter",
			input:     []string{"one two three", "four five six", "seven eight nine"},
			fields:    []int{1},
			delimiter: " ",
			separated: false,
			expected:  []string{"one", "four", "seven"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			f = tt.fields
			d = tt.delimiter
			s = tt.separated

			result := cut(tt.input)

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("expected %q, got %q", tt.expected, result)
			}
		})
	}
}
