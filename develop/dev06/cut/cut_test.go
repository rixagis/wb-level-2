package cut

import (
	"reflect"
	"testing"
)

func TestCutUntil(t *testing.T) {
	var testCases = []struct{
		input []string
		from int
		delimeter string
		expected string
	}{
		{
			[]string{"asdf", "qqq", "zzzzz"},
			-1,
			" ",
			"",
		},
		{
			[]string{"asdf", "qqq", "zzzzz"},
			0,
			" ",
			"asdf",
		},
		{
			[]string{"asdf", "qqq", "zzzzz"},
			1,
			" ",
			"asdf qqq",
		},
		{
			[]string{"asdf", "qqq", "zzzzz"},
			5,
			" ",
			"asdf qqq zzzzz",
		},
		{
			[]string{"asdf", "qqq", "zzzzz"},
			1,
			":",
			"asdf:qqq",
		},
	}

	for _, testCase := range testCases {
		var result = cutUntil(testCase.input, testCase.from, testCase.delimeter)
		if result != testCase.expected {
			t.Errorf("testing %v from %d with delimeter %s, expected: %q, got: %q",
				testCase.input,
				testCase.from,
				testCase.delimeter,
				testCase.expected,
				result)
		}
	}
}

func TestCutFrom(t *testing.T) {
	var testCases = []struct{
		input []string
		from int
		delimeter string
		expected string
	}{
		{
			[]string{"asdf", "qqq", "zzzzz"},
			-1,
			" ",
			"",
		},
		{
			[]string{"asdf", "qqq", "zzzzz"},
			0,
			" ",
			"asdf qqq zzzzz",
		},
		{
			[]string{"asdf", "qqq", "zzzzz"},
			1,
			" ",
			"qqq zzzzz",
		},
		{
			[]string{"asdf", "qqq", "zzzzz"},
			5,
			" ",
			"",
		},
		{
			[]string{"asdf", "qqq", "zzzzz"},
			1,
			":",
			"qqq:zzzzz",
		},
	}

	for _, testCase := range testCases {
		var result = cutFrom(testCase.input, testCase.from, testCase.delimeter)
		if result != testCase.expected {
			t.Errorf("testing %v from %d with delimeter %s, expected: %q, got: %q",
				testCase.input,
				testCase.from,
				testCase.delimeter,
				testCase.expected,
				result,
			)
		}
	}
}

func TestCutIndices(t *testing.T) {
	var testCases = []struct{
		input []string
		indices []int
		delimeter string
		expected string
	}{
		{
			[]string{"asdf", "qqq", "zzzzz"},
			[]int{},
			" ",
			"",
		},
		{
			[]string{"asdf", "qqq", "zzzzz"},
			[]int{0, 1, 2},
			" ",
			"asdf qqq zzzzz",
		},
		{
			[]string{"asdf", "qqq", "zzzzz"},
			[]int{1, 2},
			" ",
			"qqq zzzzz",
		},
		{
			[]string{"asdf", "qqq", "zzzzz"},
			[]int{50},
			" ",
			"",
		},
		{
			[]string{"asdf", "qqq", "zzzzz"},
			[]int{1, 2},
			":",
			"qqq:zzzzz",
		},
	}

	for _, testCase := range testCases {
		var result = cutIndices(testCase.input, testCase.indices, testCase.delimeter)
		if result != testCase.expected {
			t.Errorf("testing %v with %v with delimeter %s, expected: %q, got: %q",
				testCase.input,
				testCase.indices,
				testCase.delimeter,
				testCase.expected,
				result,
			)
		}
	}
}


func TestCutLine(t *testing.T) {
	var testCases = []struct{
		name string
		input string
		delimeter string
		fieldNums []int
		until int
		from int
		strict bool
		expected string
		expectedBool bool
	}{
		{
			name: "space delimeter",
			input: "aaa qqq bbb",
			delimeter: " ",
			fieldNums: []int{1, 2},
			until: -1,
			from: -1,
			strict: false,
			expected: "qqq bbb",
			expectedBool: true,
		},
		{
			name: "semicolon delimeter",
			input: "aaa:qqq:bbb",
			delimeter: ":",
			fieldNums: []int{1, 2},
			until: -1,
			from: -1,
			strict: false,
			expected: "qqq:bbb",
			expectedBool: true,
		},
		{
			name: "non-strict semicolon",
			input: "aaa qqq bbb",
			delimeter: ":",
			fieldNums: []int{1, 2},
			until: -1,
			from: -1,
			strict: false,
			expected: "aaa qqq bbb",
			expectedBool: true,
		},
		{
			name: "strict semicolon",
			input: "aaa qqq bbb",
			delimeter: ":",
			fieldNums: []int{1, 2},
			until: -1,
			from: -1,
			strict: true,
			expected: "",
			expectedBool: false,
		},
		{
			name: "from",
			input: "aaa qqq bbb ccc www",
			delimeter: " ",
			fieldNums: []int{1},
			until: -1,
			from: 3,
			strict: false,
			expected: "qqq ccc www",
			expectedBool: true,
		},
		{
			name: "until",
			input: "aaa qqq bbb ccc www",
			delimeter: " ",
			fieldNums: []int{3},
			until: 1,
			from: -1,
			strict: false,
			expected: "aaa qqq ccc",
			expectedBool: true,
		},
		{
			name: "from until",
			input: "aaa qqq bbb ccc www zzz eee",
			delimeter: " ",
			fieldNums: []int{3},
			until: 1,
			from: 5,
			strict: false,
			expected: "aaa qqq ccc zzz eee",
			expectedBool: true,
		},
		{
			name: "out of bounds",
			input: "aaa qqq bbb ccc www zzz eee",
			delimeter: " ",
			fieldNums: []int{15},
			until: -1,
			from: 38,
			strict: false,
			expected: "",
			expectedBool: true,
		},
	}

	for _, testCase := range testCases {
		var result, ok = CutLine(testCase.input,
			testCase.fieldNums,
			testCase.until,
			testCase.from,
			testCase.delimeter,
			testCase.strict,
		)

		if result != testCase.expected || ok != testCase.expectedBool {
			t.Errorf("failed test %q, expected: (%q, %v), got: (%q, %v)",
				testCase.name,
				testCase.expected,
				testCase.expectedBool,
				result,
				ok,
			)
		}
	}
}


func TestParseIndexIntervals(t *testing.T) {
	var testCases = []struct{
		input []string
		expected []int
		expectedError error
	}{
		{
			[]string{"15", "4", "2", "10"},
			[]int{2, 4, 10, 15},
			nil,
		},
		{
			[]string{"15", "4", "-2", "10"},
			nil,
			ErrInvalidFieldRange,
		},
		{
			[]string{},
			nil,
			nil,
		},
		{
			[]string{"15", "4", "asdfasdf", "10"},
			nil,
			ErrInvalidFieldRange,
		},
		{
			[]string{"15", "4", "7-9", "10"},
			[]int{4, 7, 8, 9, 10, 15},
			nil,
		},
		{
			[]string{"9-15", "10-17"},
			[]int{9, 10, 11, 12, 13, 14, 15, 16, 17},
			nil,
		},
	}

	for _, testCase := range testCases {
		result, err := parseIndexIntervals(testCase.input)
		if !reflect.DeepEqual(result, testCase.expected) {
			t.Errorf("testing %v, expected: %v, got: %v", testCase.input, testCase.expected, result)
		}
		if err != testCase.expectedError {
			t.Errorf("testing %v, expected error: %v, got: %v", testCase.input, testCase.expectedError, err)
		}
	}
}


func TestParseFromIntervals(t *testing.T) {
	var testCases = []struct{
		input []string
		expected int
		expectedError error
	}{
		{
			[]string{"15-", "4-", "2-", "10-"},
			2,
			nil,
		},
		{
			[]string{},
			-1,
			nil,
		},
		{
			[]string{"15-", "4-", "-2-", "10-"},
			-1,
			ErrInvalidFieldRange,
		},
		{
			[]string{"15-", "4-", "asdfasdf", "10-"},
			-1,
			ErrInvalidFieldRange,
		},
	}

	for _, testCase := range testCases {
		result, err := parseFromIntervals(testCase.input)
		if !reflect.DeepEqual(result, testCase.expected) {
			t.Errorf("testing %v, expected: %v, got: %v", testCase.input, testCase.expected, result)
		}
		if err != testCase.expectedError {
			t.Errorf("testing %v, expected error: %v, got: %v", testCase.input, testCase.expectedError, err)
		}
	}
}

func TestParseUntilIntervals(t *testing.T) {
	var testCases = []struct{
		input []string
		expected int
		expectedError error
	}{
		{
			[]string{"-15", "-4", "-2", "-10"},
			15,
			nil,
		},
		{
			[]string{},
			-1,
			nil,
		},
		{
			[]string{"-15", "-4", "--2", "-10"},
			-1,
			ErrInvalidFieldRange,
		},
		{
			[]string{"15", "4", "asdfasdf", "10"},
			-1,
			ErrInvalidFieldRange,
		},
	}

	for _, testCase := range testCases {
		result, err := parseUntilIntervals(testCase.input)
		if !reflect.DeepEqual(result, testCase.expected) {
			t.Errorf("testing %v, expected: %v, got: %v", testCase.input, testCase.expected, result)
		}
		if err != testCase.expectedError {
			t.Errorf("testing %v, expected error: %v, got: %v", testCase.input, testCase.expectedError, err)
		}
	}
}

func TestParseIntervals(t *testing.T) {
	var testCases = []struct{
		name string
		input string
		expectedIndices []int
		expectedFrom int
		expectedUntil int
		expectedError error
	}{
		{
			name: "indices",
			input: "15,3,24,1,5",
			expectedIndices: []int{1,3,5,15,24},
			expectedFrom: -1,
			expectedUntil: -1,
			expectedError: nil,
		},
		{
			name: "indices repeating",
			input: "15,3,24,1,5,1,3,15,15",
			expectedIndices: []int{1,3,5,15,24},
			expectedFrom: -1,
			expectedUntil: -1,
			expectedError: nil,
		},
		{
			name: "from",
			input: "15-,17-,2-",
			expectedIndices: nil,
			expectedFrom: 2,
			expectedUntil: -1,
			expectedError: nil,
		},
		{
			name: "until",
			input: "-15,-17,-2",
			expectedIndices: nil,
			expectedFrom: -1,
			expectedUntil: 17,
			expectedError: nil,
		},
		{
			name: "from until",
			input: "-15,17-,-2,22-",
			expectedIndices: nil,
			expectedFrom: 17,
			expectedUntil: 15,
			expectedError: nil,
		},
		{
			name: "from until indices",
			input: "-10,17-,13,-2,22-,14",
			expectedIndices: []int{13, 14},
			expectedFrom: 17,
			expectedUntil: 10,
			expectedError: nil,
		},
		{
			name: "normalization",
			input: "-5,3,4,8,10-,9,15",
			expectedIndices: []int{8, 9},
			expectedFrom: 10,
			expectedUntil: 5,
			expectedError: nil,
		},
		{
			name: "intersection",
			input: "-10,3,4,8,5-,9,15",
			expectedIndices: nil,
			expectedFrom: 0,
			expectedUntil: -1,
			expectedError: nil,
		},
		{
			name: "range",
			input: "-5,3,4,8,30-,9,12-15",
			expectedIndices: []int{8, 9, 12, 13, 14, 15},
			expectedFrom: 30,
			expectedUntil: 5,
			expectedError: nil,
		},
		{
			name: "invalid indices",
			input: "-10,3,4,asdf,5-,9,15",
			expectedIndices: nil,
			expectedFrom: -1,
			expectedUntil: -1,
			expectedError: ErrInvalidFieldRange,
		},
		{
			name: "invalid range",
			input: "-10,3,4,5,5-,9-4,15",
			expectedIndices: nil,
			expectedFrom: -1,
			expectedUntil: -1,
			expectedError: ErrInvalidFieldRange,
		},
		{
			name: "invalid range 2",
			input: "-10,3,4,5,5-,3-4-5,15",
			expectedIndices: nil,
			expectedFrom: -1,
			expectedUntil: -1,
			expectedError: ErrInvalidFieldRange,
		},
		{
			name: "invalid range 3",
			input: "-10,3,4,5,5-,3-asdf5,15",
			expectedIndices: nil,
			expectedFrom: -1,
			expectedUntil: -1,
			expectedError: ErrInvalidFieldRange,
		},
		{
			name: "invalid until",
			input: "-10asd,3,4,5,5-,4-9,15",
			expectedIndices: nil,
			expectedFrom: -1,
			expectedUntil: -1,
			expectedError: ErrInvalidFieldRange,
		},
		{
			name: "invalid from",
			input: "-10,3,4,5,5asdf-,4-9,15",
			expectedIndices: nil,
			expectedFrom: -1,
			expectedUntil: -1,
			expectedError: ErrInvalidFieldRange,
		},
	}

	for _, testCase := range testCases {
		indices, until, from, err := ParseIntervals(testCase.input)
		if !reflect.DeepEqual(indices, testCase.expectedIndices) {
			t.Errorf("failed test %q, expected indices: %v, got: %v", testCase.name, testCase.expectedIndices, indices)
		}

		if from != testCase.expectedFrom {
			t.Errorf("failed test %q, expected from: %d, got: %d", testCase.name, testCase.expectedFrom, from)
		}

		if until != testCase.expectedUntil {
			t.Errorf("failed test %q, expected until: %d, got: %d", testCase.name, testCase.expectedUntil, until)
		}

		if err != testCase.expectedError {
			t.Errorf("failed test %q, expected error: %v, got: %v", testCase.name, testCase.expectedError, err)
		}
		
	}
}