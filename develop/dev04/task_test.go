package anagrams

import (
	"testing"
	"reflect"
)

func TestAreAnagrams(t *testing.T) {
	var testCases = []struct{
		input1 string
		input2 string
		expected bool
	}{
		{
			"пятка",
			"тяпка",
			true,
		},
		{
			"пятка",
			"пятак",
			true,
		},
		{
			"пятак",
			"пятка",
			true,
		},
		{
			"asdfasdf",
			"fdasfdsa",
			true,
		},
		{
			"пятка",
			"тяпк",
			false,
		},
		{
			"пятка",
			"тяпкк",
			false,
		},
		{
			"aaaab",
			"aaabb",
			false,
		},
		{
			"",
			"",
			true,
		},
		{
			"",
			"aaabb",
			false,
		},
	}

	for _, testCase := range testCases {
		var result = areAnagrams(testCase.input1, testCase.input2)
		if result != testCase.expected {
			t.Errorf("testing %q and %q, expected: %v, got: %v", testCase.input1, testCase.input2, testCase.expected, result)
		}
	}
}


func TestMakeAnagramMap(t *testing.T) {
	var testCases = []struct{
		input []string
		expected map[string][]string
	}{
		{
			[]string{"тяпка", "пятка", "пятак"},
			map[string][]string{
				"тяпка": {"тяпка", "пятка", "пятак"},
			},
		},
		{
			[]string{"тяпка", "пятка", "пятак", "asdfasdf"},
			map[string][]string{
				"тяпка": {"тяпка", "пятка", "пятак"},
			},
		},
		{
			[]string{"пятка", "тяпка", "пятак", "asdf"},
			map[string][]string{
				"пятка": {"пятка", "тяпка", "пятак"},
			},
		},
		{
			[]string{"Пятка", "ТЯПКА", "пяТак", "asdf"},
			map[string][]string{
				"пятка": {"пятка", "тяпка", "пятак"},
			},
		},
		{
			[]string{"Пятка", "ТЯПКА", "пяТак", "asdf", "FSDA"},
			map[string][]string{
				"пятка": {"пятка", "тяпка", "пятак"},
				"asdf":  {"asdf", "fsda"},
			},
		},
		{
			[]string{"Пятка", "asdf"},
			map[string][]string{
			},
		},
		{
			[]string{},
			map[string][]string{
			},
		},
	}

	for _, testCase := range testCases {
		var result = MakeAnagramMap(testCase.input)
		if !reflect.DeepEqual(testCase.expected, result) {
			t.Errorf("testing %v, expected: %v, got: %v", testCase.input, testCase.expected, result)
		}
	}
}

