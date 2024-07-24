package main

import (
	"reflect"
	"testing"
)

func TestFindAnagrams(t *testing.T) {
	tests := []struct {
		name     string
		words    []string
		expected map[string][]string
	}{
		{
			name:  "simple case",
			words: []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"},
			expected: map[string][]string{
				"пятак":  {"пятак", "пятка", "тяпка"},
				"листок": {"листок", "слиток", "столик"},
			},
		},
		{
			name:  "case insensitive",
			words: []string{"Пятак", "пятка", "Тяпка", "листок", "Слиток", "столик"},
			expected: map[string][]string{
				"пятак":  {"пятак", "пятка", "тяпка"},
				"листок": {"листок", "слиток", "столик"},
			},
		},
		{
			name:  "single word groups",
			words: []string{"привет", "мир", "рим"},
			expected: map[string][]string{
				"мир": {"мир", "рим"},
			},
		},
		{
			name:     "no anagrams",
			words:    []string{"пятак", "листок", "слово"},
			expected: map[string][]string{},
		},
		{
			name:     "empty input",
			words:    []string{},
			expected: map[string][]string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findAnagrams(&tt.words)
			if !reflect.DeepEqual(*result, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, *result)
			}
		})
	}
}
