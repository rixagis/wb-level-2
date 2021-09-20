package unpacker

import (
	"testing"
)

func TestAtoi(t *testing.T) {
	var (
		intResult int
		isDigit bool
	)

	for i, c := range "0123456789" {
		intResult, isDigit = atoi(c)
		if !isDigit {
			t.Errorf("isDigit result of atoi(%s) is expected to be true, got false", string(c))
		}
		if intResult != i {
			t.Errorf("int result of atoi(%s) is expected to be %d, got %d", string(c), i, intResult)
		}
	}

	for _, c := range "abcdefghijklmnopqrstuvwxyz!@#$%^&*()_+|{}[]<>'\"/\\`" {
		intResult, isDigit = atoi(c)
		if isDigit {
			t.Errorf("isDigit result of atoi(%s) is expected to be false, got true", string(c))
		}
		if intResult != -1 {
			t.Errorf("int result of atoi(%s) is expected to be -1, got %d", string(c), intResult)
		}
	}
}


func TestMultiplyRune(t *testing.T) {
	var testCases = []struct{
		name string
		inputRune rune
		inputInt int
		expected string
	}{
		{
			"single char",
			'a',
			1,
			"a",
		},
		{
			"no chars",
			'a',
			0,
			"",
		},
		{
			"several characters",
			'a',
			15,
			"aaaaaaaaaaaaaaa",
		},
		{
			"incorrect input",
			'a',
			-18,
			"",
		},
	}

	var result string
	for _, testCase := range testCases {
		result = multiplyRune(testCase.inputRune, testCase.inputInt)
		if result != testCase.expected {
			t.Errorf("failed testCase %q: multiply(%s, %d) is expected to be %q, got: %q",
				testCase.name,
				string(testCase.inputRune),
				testCase.inputInt,
				testCase.expected,
				result)
		}
	}
}



func TestProcessChar(t *testing.T) {
	var testCases = []struct{
		name string
		state state
		input rune
		expectedNextStateID int
		expectedNextStateChar rune
		expectedNextStateNumber int
		expectedOutput string
	}{
		{
			"digit on start",
			&stateStart{},
			'4',
			stateErrorID,
			0,
			0,
			"",
		},
		{
			"char on start",
			&stateStart{},
			'a',
			stateCharID,
			'a',
			0,
			"",
		},
		{
			"backslash on start",
			&stateStart{},
			'\\',
			stateBackslashID,
			0,
			0,
			"",
		},
		{
			"digit on backslash",
			&stateBackslash{},
			'4',
			stateCharID,
			'4',
			0,
			"",
		},
		{
			"char on backslash",
			&stateBackslash{},
			'a',
			stateErrorID,
			0,
			0,
			"",
		},
		{
			"backslash on backslash",
			&stateBackslash{},
			'\\',
			stateCharID,
			'\\',
			0,
			"",
		},
		{
			"digit on char",
			&stateChar{stateData: stateData{currentChar: 'a'}},
			'4',
			stateNumberID,
			'a',
			4,
			"",
		},
		{
			"char on char",
			&stateChar{stateData: stateData{currentChar: 'a'}},
			'b',
			stateCharID,
			'b',
			0,
			"a",
		},
		{
			"backslash on char",
			&stateChar{stateData: stateData{currentChar: 'a'}},
			'\\',
			stateBackslashID,
			'a',
			0,
			"a",
		},
		{
			"digit on digit",
			&stateNumber{stateData: stateData{currentChar: 'a', currentNumber: 4}},
			'4',
			stateNumberID,
			'a',
			44,
			"",
		},
		{
			"char on digit",
			&stateNumber{stateData: stateData{currentChar: 'a', currentNumber: 4}},
			'b',
			stateCharID,
			'b',
			0,
			"aaaa",
		},
		{
			"backslash on digit",
			&stateNumber{stateData: stateData{currentChar: 'a', currentNumber: 4}},
			'\\',
			stateBackslashID,
			0,
			0,
			"aaaa",
		},
		{
			"finish on digit",
			&stateNumber{stateData: stateData{currentChar: 'a', currentNumber: 4}},
			0,
			stateFinishID,
			0,
			0,
			"aaaa",
		},
		{
			"finish on char",
			&stateChar{stateData: stateData{currentChar: 'a'}},
			0,
			stateFinishID,
			0,
			0,
			"a",
		},
		{
			"finish on backslash",
			&stateBackslash{},
			0,
			stateErrorID,
			0,
			0,
			"",
		},
		{
			"char on finish",
			&stateFinish{},
			'a',
			stateErrorID,
			0,
			0,
			"",
		},
		{
			"char on error",
			&stateError{},
			'a',
			stateErrorID,
			0,
			0,
			"",
		},
	}

	var unpacker = NewUnpacker()
	var output string

	if unpacker.currentStateID() != stateStartID {
		t.Errorf("inital state should be Start state (%d), got: %d", stateStartID, unpacker.currentStateID())
	}

	for _, testCase := range testCases {
		unpacker.currentState = testCase.state
		testCase.state.setContext(unpacker)
		output = unpacker.processChar(testCase.input)
		if unpacker.currentStateID() != testCase.expectedNextStateID {
			t.Errorf("failed case %q: expected next state: %v, got: %v", testCase.name, testCase.expectedNextStateID, testCase.state.id())
		}
		if testCase.expectedOutput != output {
			t.Errorf("faled case %q: expected output: %q, got: %q", testCase.name, testCase.expectedOutput, output)
		}
	}
}



func TestUnpack(t *testing.T) {
	var testCases = []struct{
		input string
		expected string
		expectedError bool
	}{
		{
			input: "a4bc2d5e",
			expected: "aaaabccddddde",
			expectedError: false,
		},
		{
			input: "abcd",
			expected: "abcd",
			expectedError: false,
		},
		{
			input: "45",
			expected: "",
			expectedError: true,
		},
		{
			input: "",
			expected: "",
			expectedError: false,
		},
		{
			input: `qwe\4\5`,
			expected: `qwe45`,
			expectedError: false,
		},
		{
			input: `qwe\45`,
			expected: `qwe44444`,
			expectedError: false,
		},
		{
			input: `qwe\\5`,
			expected: `qwe\\\\\`,
			expectedError: false,
		},
		{
			input: `\`,
			expected: "",
			expectedError: true,
		},
	}

	var unpacker *Unpacker
	var output string
	var err error
	for _, testCase := range testCases {
		unpacker = NewUnpacker()
		output, err = unpacker.Unpack(testCase.input)
		if testCase.expected != output {
			t.Errorf("testing %q, expected: %q, got: %q", testCase.input, testCase.expected, output)
		}
		if testCase.expectedError {
			if err == nil {
				t.Errorf("testing %q, expected error, got none", testCase.input)
			}
		} else {
			if err != nil {
				t.Errorf("testing %q, expected no error, got: %q", testCase.input, err)
			}
		}
	}
}