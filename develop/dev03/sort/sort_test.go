package sort

import (
	"testing"
)


func TestSortLexicographical(t *testing.T) {
	var testCases = []struct{
		input []string
		reversed bool
		expected []string
	}{
		{
			[]string{"10", "1000", "15", "2", "102"},
			false,
			[]string{"10", "1000", "102", "15", "2"},
		},
		{
			[]string{"ccc", "fff", "bbb", "aaa", "hhh"},
			false,
			[]string{"aaa", "bbb", "ccc", "fff", "hhh"},
		},
		{
			[]string{"a", "1", "2", "", " "},
			false,
			[]string{"", " ", "1", "2", "a"},
		},
		{
			[]string{"10", "1000", "15", "2", "102"},
			true,
			[]string{"2", "15", "102", "1000", "10"},
		},
		{
			[]string{"ccc", "fff", "bbb", "aaa", "hhh"},
			true,
			[]string{"hhh", "fff", "ccc", "bbb", "aaa"},
		},
		{
			[]string{"a", "1", "2", "", " "},
			true,
			[]string{"a", "2", "1", " ", ""},
		},
	}

	var result []string
	for _, testCase := range testCases {
		result = nil
		result = append(result, testCase.input...)
		sortByWhole(result, lessLexicographical, testCase.reversed, false)
		if !slicesEqual(result, testCase.expected) {
			t.Errorf("sorting %v with reversed = %v, expected: %v, got: %v", testCase.input, testCase.reversed, testCase.expected, result)
		}
	}
}


func TestGetNumberPart(t *testing.T) {
	var testCases = []struct{
		input string
		expected float64
	}{
		{
			"",
			0,
		},
		{
			"1",
			1,
		},
		{
			"13574",
			13574,
		},
		{
			"0.15",
			0.15,
		},
		{
			"45.89",
			45.89,
		},
		{
			"-5",
			-5,
		},
		{
			"asdf",
			0,
		},
		{
			"1asdf",
			1,
		},
		{
			"13574asdf",
			13574,
		},
		{
			"0.15asdf",
			0.15,
		},
		{
			"45.89asdf",
			45.89,
		},
		{
			"-5asdf",
			-5,
		},
		{
			"-asdf",
			0,
		},
		{
			"asdf",
			0,
		},
		{
			"-15.7.5",
			-15.7,
		},
		{
			"-",
			0,
		},
	}

	var result float64 = 0
	for _, testCase := range testCases {
		result = getNumberPart(testCase.input)
		if float64(result) != testCase.expected {
			t.Errorf("tested: %q, expected: %f, got: %f", testCase.input, testCase.expected, result)
		}
	}
}



func TestLessNumerical(t *testing.T) {
	var testCases = []struct{
		a string
		b string
		expectedLess bool
	}{
		{
			"15",
			"145",
			true,
		},
		{
			"156",
			"145",
			false,
		},
		{
			"1.5",
			"145",
			true,
		},
		{
			"15.0",
			"-1.45",
			false,
		},
		{
			"15asf",
			"145fdfdfd",
			true,
		},
		{
			"156asfa",
			"145fsasefa",
			false,
		},
		{
			"1.5.asdfas",
			"145sdsss",
			true,
		},
		{
			"15.0 asd ",
			"-1.45f as",
			false,
		},
		{
			"-15.0 asd ",
			"",
			true,
		},
		{
			"15",
			"  as",
			false,
		},
		{
			"115asdf",
			"115asdg",
			true,
		},
		{
			"-11.5b",
			"-11.5a",
			false,
		},
	}

	var result = false
	var arr = []string{"", ""}
	for _, testCase := range testCases {
		arr[0] = testCase.a
		arr[1] = testCase.b
		result = lessNumerical(arr, 0, 1, false)
		if result != testCase.expectedLess {
			t.Errorf("compared %q and %q, expected: %v, got: %v", testCase.a, testCase.b, testCase.expectedLess, result)
		}
	}
}

func slicesEqual(a, b []string) bool {
	var len1 = len(a)
	var len2 = len(b)
	if len1 != len2 {
		return false
	}
	for i := 0; i < len1; i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestSortNumerical(t *testing.T) {
	var testCases = []struct{
		input []string
		reversed bool
		expected []string
	}{
		{
			[]string{"115", "325", "41", "-10", "85.0"},
			false,
			[]string{"-10", "41", "85.0", "115", "325"},
		},
		{
			[]string{"115asdf", "325asdf", "41asdf", "-10asdf", "85.0asdf"},
			false,
			[]string{"-10asdf", "41asdf", "85.0asdf", "115asdf", "325asdf"},
		},
		{
			[]string{"115", "325", "41", "f", "a"},
			false,
			[]string{"a", "f", "41", "115", "325"},
		},
		{
			[]string{"115", "41b", "325", "41aa"},
			false,
			[]string{"41aa", "41b", "115", "325"},
		},
		{
			[]string{"115", "325", "41", "-10", "85.0"},
			true,
			[]string{"325", "115", "85.0", "41", "-10"},
		},
		{
			[]string{"115asdf", "325asdf", "41asdf", "-10asdf", "85.0asdf"},
			true,
			[]string{"325asdf", "115asdf", "85.0asdf", "41asdf", "-10asdf"},
		},
		{
			[]string{"115", "325", "41", "f", "a"},
			true,
			[]string{"325", "115", "41", "f", "a"},
		},
		{
			[]string{"115", "41b", "325", "41aa"},
			true,
			[]string{"325", "115", "41b", "41aa"},
		},
	}

	var result []string
	for _, testCase := range testCases {
		result = nil
		result = append(result, testCase.input...)
		sortByWhole(result, lessNumerical, testCase.reversed, false)
		if !slicesEqual(result, testCase.expected) {
			t.Errorf("sorting %v with reversed = %v, expected: %v, got: %v", testCase.input, testCase.reversed, testCase.expected, result)
		}
	}
}

func TestGetUniques(t *testing.T) {
	var testCases = []struct{
		input []string
		expected []string
	}{
		{
			[]string{"a", "b", "c"},
			[]string{"a", "b", "c"},
		},
		{
			[]string{"a", "a", "b", "b", "c", "c", "c"},
			[]string{"a", "b", "c"},
		},
		{
			[]string{},
			[]string{},
		},
	}

	var result []string
	for _, testCase := range testCases {
		result = nil
		result = append(result, getUniques(testCase.input)...)
		if !slicesEqual(result, testCase.expected) {
			t.Errorf("testing %v, expected: %v, got: %v", testCase.input, testCase.expected, result)
		}
	}
}

func TestSortByField(t * testing.T) {
	const (
		lex = iota
		num
	)
	var testCases = []struct{
		input []string
		field int
		sortBy int
		reversed bool
		expected []string
	}{
		{
			[]string{"bbb 100", "zzz 50", "aaa 300.0"},
			1,
			lex,
			false,
			[]string{"aaa 300.0", "bbb 100", "zzz 50"},
		},
		{
			[]string{"bbb 100", "zzz 50", "aaa 300.0"},
			2,
			lex,
			false,
			[]string{"bbb 100", "aaa 300.0", "zzz 50"},
		},
		{
			[]string{"bbb 100", "zzz 50", "aaa 300.0"},
			1,
			num,
			false,
			[]string{"aaa 300.0", "bbb 100", "zzz 50"},
		},
		{
			[]string{"bbb 100", "zzz 50", "aaa 300.0"},
			2,
			num,
			false,
			[]string{"zzz 50", "bbb 100", "aaa 300.0"},
		},
		{
			[]string{"100", "zzz 50", "aaa 300.0"},
			2,
			num,
			false,
			[]string{"100", "zzz 50", "aaa 300.0"},
		},
		{
			[]string{"bbb 100", "zzz 50", "300.0"},
			2,
			num,
			false,
			[]string{"300.0", "zzz 50", "bbb 100"},
		},
		{
			[]string{"bbb 100", "", ""},
			2,
			num,
			false,
			[]string{"", "", "bbb 100"},
		},
		{
			[]string{"bbb 100", "zzz 50", "aaa 300.0"},
			1,
			lex,
			true,
			[]string{"zzz 50", "bbb 100", "aaa 300.0"},
		},
		{
			[]string{"bbb 100", "zzz 50", "aaa 300.0"},
			2,
			lex,
			true,
			[]string{"zzz 50", "aaa 300.0", "bbb 100"},
		},
		{
			[]string{"bbb 100", "zzz 50", "aaa 300.0"},
			1,
			num,
			true,
			[]string{"zzz 50", "bbb 100", "aaa 300.0"},
		},
		{
			[]string{"bbb 100", "zzz 50", "aaa 300.0"},
			2,
			num,
			true,
			[]string{"aaa 300.0", "bbb 100", "zzz 50"},
		},
	}

	var result []string
	for _, testCase := range testCases {
		var lessFunc func([]string, int, int, bool) bool
		switch testCase.sortBy {
		case num:
			lessFunc = lessNumerical
		default:
			lessFunc = lessLexicographical
		}

		result = nil
		result = append(result, testCase.input...)
		sortByField(result, testCase.field, lessFunc, testCase.reversed)
		if !slicesEqual(result, testCase.expected) {
			t.Errorf("testing %v by field %d (type %d), expected: %v, got: %v",
			testCase.input,
			testCase.field,
			testCase.sortBy,
			testCase.expected,
			result)
		}
	}
}

func TestMonthToInt(t *testing.T) {
	var testCases = []struct{
		input string
		expected int
	}{
		{
			"jan",
			1,
		},
		{
			"FEB",
			2,
		},
		{
			"march",
			3,
		},
		{
			"ApRiL",
			4,
		},
		{
			"MAY",
			5,
		},
		{
			"jun",
			6,
		},
		{
			"july",
			7,
		},
		{
			"AUGUST",
			8,
		},
		{
			"SEP",
			9,
		},
		{
			"oct",
			10,
		},
		{
			"NOVEMBER",
			11,
		},
		{
			"DEC",
			12,
		},
		{
			"asdfasdfasdf",
			-1,
		},
	}

	var result int
	for _, testCase := range testCases {
		result = monthToInt(testCase.input)
		if testCase.expected != result {
			t.Errorf("testing %q, expected: %d, got: %d", testCase.input, testCase.expected, result)
		}
	}
}

func TestSortMonths(t *testing.T) {
	var testCases = []struct{
		input []string
		reversed bool
		expected []string
	}{
		{
			[]string{"dec", "feb", "jan"},
			false,
			[]string{"jan", "feb", "dec"},
		},
		{
			[]string{"dec", "asdfasdf", "jan"},
			false,
			[]string{"asdfasdf", "jan", "dec"},
		},
		{
			[]string{"dec", "feb", "jan"},
			true,
			[]string{"dec", "feb", "jan"},
		},
		{
			[]string{"dec", "asdfasdf", "jan"},
			true,
			[]string{"dec", "jan" , "asdfasdf"},
		},
	}

	var result []string
	for _, testCase := range testCases {
		result = nil
		result = append(result, testCase.input...)
		sortByWhole(result, lessMonth, testCase.reversed, false)
		if !slicesEqual(result, testCase.expected) {
			t.Errorf("testing %v with reversed=%v, expected: %v, got: %v",
			testCase.input,
			testCase.reversed,
			testCase.expected,
			result)
		}
	}
}


func TestSortSuffix(t *testing.T) {
	var testCases = []struct{
		input []string
		reversed bool
		expected []string
	}{
		{
			[]string{"150M", "149k", "150k","asdf", "-15k", "150"},
			false,
			[]string{"asdf", "150", "-15k", "149k", "150k", "150M"},
		},
		{
			[]string{"150M", "149k", "150k", "asdf", "-15k", "150"},
			true,
			[]string{"150M", "150k", "149k", "-15k", "150", "asdf"},
		},
		{
			[]string{"-150.0k", "-101", "15k", "15.0.0", "14.0", "asdf"},
			false,
			[]string{"-101", "asdf", "14.0", "15.0.0", "-150.0k", "15k"},
		},
		{
			[]string{"-150.0k", "-101", "15k", "15.0.0", "-101", "14.0", "asdf"},
			true,
			[]string{"15k", "-150.0k", "15.0.0", "14.0", "asdf", "-101", "-101"},
		},
	}

	var result []string
	for _, testCase := range testCases {
		result = nil
		result = append(result, testCase.input...)
		sortByWhole(result, lessSuffix, testCase.reversed, false)
		if !slicesEqual(result, testCase.expected) {
			t.Errorf("testing %v with reversed=%v, expected: %v, got: %v",
			testCase.input,
			testCase.reversed,
			testCase.expected,
			result)
		}
	}
}

func TestIgnoreTrailingWhitespace(t *testing.T) {
	var testCases = []struct{
		input []string
		less func([]string, int, int, bool) bool
		reversed bool
		expected []string
	}{
		{
			[]string{"ccc", "aaa", "aaa ", "bbb ", "bbb"},
			lessLexicographical,
			false,
			[]string{"aaa", "aaa ", "bbb ", "bbb", "ccc"},
		},
		{
			[]string{"ccc", "bbb ", "bbb", "aaa", "aaa "},
			lessLexicographical,
			true,
			[]string{"ccc", "bbb ", "bbb", "aaa", "aaa "},
		},
		{
			[]string{"115 ", "115", "41", "41 ", "85.0"},
			lessNumerical,
			false,
			[]string{"41", "41 ", "85.0", "115 ", "115"},
		},
		{
			[]string{"115 ", "115", "41", "41 ", "85.0"},
			lessNumerical,
			true,
			[]string{"115 ", "115", "85.0", "41", "41 "},
		},
		{
			[]string{"DEC ", "apr", "APR ", "dec", "september"},
			lessMonth,
			false,
			[]string{"apr", "APR ", "september", "DEC ", "dec"},
		},
		{
			[]string{"DEC ", "apr", "APR ", "dec", "september"},
			lessMonth,
			true,
			[]string{"DEC ", "dec", "september", "apr", "APR ", },
		},
		{
			[]string{"115k ", "150M", "150M ", "115k", "300"},
			lessSuffix,
			false,
			[]string{"300", "115k ", "115k", "150M", "150M "},
		},
		{
			[]string{"115k ", "150M", "150M ", "115k", "300"},
			lessSuffix,
			true,
			[]string{"150M", "150M ", "115k ", "115k", "300"},
		},

	}

	var result []string
	for _, testCase := range testCases {
		result = nil
		result = append(result, testCase.input...)
		sortByWhole(result, testCase.less, testCase.reversed, true)
		if !slicesEqual(result, testCase.expected) {
			t.Errorf("testing %v with reversed=%v, expected: %v, got: %v",
			testCase.input,
			testCase.reversed,
			testCase.expected,
			result)
		}
	}
}

func TestSort(t *testing.T) {
	var testCases = []struct{
		name string
		input []string
		key int
		numerical, month, human, reversed, unique, ignoreTrailSpaces bool
		expected []string
	}{
		{
			name: "general lexicographical",
			input: []string{"ccc", "fff", "bbb", "aaa", "hhh", "aaa"},
			key: -1,
			numerical: false, month: false, human: false,
			reversed: false, unique: false, ignoreTrailSpaces: false,
			expected: []string{"aaa", "aaa", "bbb", "ccc", "fff", "hhh"},
		},
		{
			name: "lex reversed",
			input: []string{"ccc", "fff", "bbb", "aaa", "hhh", "aaa"},
			key: -1,
			numerical: false, month: false, human: false,
			reversed: true, unique: false, ignoreTrailSpaces: false,
			expected: []string{"hhh", "fff", "ccc", "bbb", "aaa", "aaa"},
		},
		{
			name: "lex unique",
			input: []string{"ccc", "fff", "bbb", "aaa", "hhh", "aaa"},
			key: -1,
			numerical: false, month: false, human: false,
			reversed: false, unique: true, ignoreTrailSpaces: false,
			expected: []string{"aaa", "bbb", "ccc", "fff", "hhh"},
		},
		{
			name: "lex unique reversed",
			input: []string{"ccc", "fff", "bbb", "aaa", "hhh", "aaa"},
			key: -1,
			numerical: false, month: false, human: false,
			reversed: true, unique: true, ignoreTrailSpaces: false,
			expected: []string{"hhh", "fff", "ccc", "bbb", "aaa"},
		},
		{
		name: "numerical",
		input: []string{"150", "10.5", "-5", "40"},
		key: -1,
		numerical: true, month: false, human: false,
		reversed: false, unique: false, ignoreTrailSpaces: false,
		expected: []string{"-5", "10.5", "40", "150"},
		},
		{
			name: "month",
			input: []string{"sep", "APR", "DECEMBER", "january"},
			key: -1,
			numerical: false, month: true, human: false,
			reversed: false, unique: false, ignoreTrailSpaces: false,
			expected: []string{"january", "APR", "sep", "DECEMBER"},
		},
		{
			name: "suffix",
			input: []string{"150M", "149k", "150k","asdf", "-15k", "150"},
			key: -1,
			numerical: false, month: false, human: true,
			reversed: false, unique: false, ignoreTrailSpaces: false,
			expected: []string{"asdf", "150", "-15k", "149k", "150k", "150M"},
		},
		{
			name: "fields",
			input: []string{"bbb 100", "zzz 50", "aaa 300.0"},
			key: 2,
			numerical: true, month: false, human: false,
			reversed: false, unique: false, ignoreTrailSpaces: false,
			expected: []string{"zzz 50", "bbb 100", "aaa 300.0"},
		},

	}

	var result []string
	for _, testCase := range testCases {
		result = Sort(result,
			testCase.key, 
			testCase.numerical,
			testCase.month,
			testCase.human,
			testCase.reversed,
			testCase.unique,
			testCase.ignoreTrailSpaces)
		if slicesEqual(result, testCase.expected) {
			t.Errorf("failed test %q: expected: %v, got: %v", testCase.name, testCase.expected, result)
		}
	}
}

func TestCheck(t *testing.T) {
	var testCases = []struct{
		name string
		input []string
		key int
		numerical, month, human, reversed, ignoreTrailSpaces bool
		expected int
	}{
		{
			name: "lexicographical unsorted",
			input: []string{"ccc", "fff", "bbb", "aaa", "hhh", "aaa"},
			key: -1,
			numerical: false, month: false, human: false,
			reversed: false, ignoreTrailSpaces: false,
			expected: 3,
		},
		{
			name: "lex reversed sorted",
			input: []string{"hhh", "fff", "ccc", "bbb", "aaa", "aaa"},
			key: -1,
			numerical: false, month: false, human: false,
			reversed: true, ignoreTrailSpaces: false,
			expected: -1,
		},
		{
			name: "lex reversed unsorted",
			input: []string{"ccc", "fff", "bbb", "aaa", "hhh", "aaa"},
			key: -1,
			numerical: false, month: false, human: false,
			reversed: true, ignoreTrailSpaces: false,
			expected: 2,
		},
		{
			name: "numerical sorted",
			input: []string{"-5", "10.5", "40", "150"},
			key: -1,
			numerical: true, month: false, human: false,
			reversed: false, ignoreTrailSpaces: false,
			expected: -1,
		},
		{
			name: "numerical unsorted",
			input: []string{"150", "10.5", "-5", "40"},
			key: -1,
			numerical: true, month: false, human: false,
			reversed: false, ignoreTrailSpaces: false,
			expected: 2,
		},
		{
			name: "month sorted",
			input: []string{"january", "APR", "sep", "DECEMBER"},
			key: -1,
			numerical: false, month: true, human: false,
			reversed: false, ignoreTrailSpaces: false,
			expected: -1,
		},
		{
			name: "month unsorted",
			input: []string{"sep", "APR", "DECEMBER", "january"},
			key: -1,
			numerical: false, month: true, human: false,
			reversed: false, ignoreTrailSpaces: false,
			expected: 2,
		},
		{
			name: "suffix sorted",
			input: []string{"asdf", "150", "-15k", "149k", "150k", "150M"},
			key: -1,
			numerical: false, month: false, human: true,
			reversed: false, ignoreTrailSpaces: false,
			expected: -1,
		},
		{
			name: "suffix unsorted",
			input: []string{"150M", "149k", "150k","asdf", "-15k", "150"},
			key: -1,
			numerical: false, month: false, human: true,
			reversed: false, ignoreTrailSpaces: false,
			expected: 2,
		},
		{
			name: "fields sorted",
			input: []string{"zzz 50", "bbb 100", "aaa 300.0"},
			key: 2,
			numerical: true, month: false, human: false,
			reversed: false, ignoreTrailSpaces: false,
			expected: -1,
		},
		{
			name: "fields unsorted",
			input: []string{"bbb 100", "aaa 300.0", "zzz 50"},
			key: 2,
			numerical: true, month: false, human: false,
			reversed: false, ignoreTrailSpaces: false,
			expected: 3,
		},

	}

	var result int
	for _, testCase := range testCases {
		result = Check(testCase.input,
			testCase.key, 
			testCase.numerical,
			testCase.month,
			testCase.human,
			testCase.reversed,
			testCase.ignoreTrailSpaces)
		if result != testCase.expected {
			t.Errorf("failed test %q: expected: %v, got: %v", testCase.name, testCase.expected, result)
		}
	}
}