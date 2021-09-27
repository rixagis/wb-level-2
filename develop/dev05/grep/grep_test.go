package grep

import (
	"strings"
	"testing"
)

func TestGrep(t *testing.T) {
	testCases := []struct{
		name string
		input []string
		target string
		params Parameters
		expected string
	}{
		{
			"basic",
			[]string{"aaa", "aaaa", "bbb"},
			"aaa",
			Parameters{

			},
			"aaa\naaaa\n",
		},
		{
			"invert",
			[]string{"aaa", "aaaa", "bbb"},
			"aaa",
			Parameters{
				Invert: true,
			},
			"bbb\n",
		},
		{
			"not found",
			[]string{"aaa", "aaaa", "bbb"},
			"xxx",
			Parameters{

			},
			"",
		},
		{
			"case insensitive off",
			[]string{"aaa", "AAA", "bbb"},
			"aaa",
			Parameters{

			},
			"aaa\n",
		},
		{
			"case insensitive on",
			[]string{"aaa", "AAA", "bbb"},
			"aaa",
			Parameters{
				IgnoreCase: true,
			},
			"aaa\nAAA\n",
		},
		{
			"line numbers",
			[]string{"aaa", "aaaa", "bbb"},
			"aaa",
			Parameters{
				LineNum: true,
			},
			"1:aaa\n2:aaaa\n",
		},
		{
			"fixed off",
			[]string{"aaa", "asdfasdf", "bbb", "[a-z]"},
			"[a-z]",
			Parameters{

			},
			"aaa\nasdfasdf\nbbb\n[a-z]\n",
		},
		{
			"fixed on",
			[]string{"aaa", "asdfasdf", "bbb", "[a-z]"},
			"[a-z]",
			Parameters{
				Fixed: true,
			},
			"[a-z]\n",
		},
		{
			"fixed ignore case",
			[]string{"aaa", "asdfasdf", "bbb", "[a-z]", "[A-Z]"},
			"[a-z]",
			Parameters{
				Fixed: true,
				IgnoreCase: true,
			},
			"[a-z]\n[A-Z]\n",
		},
		{
			"count",
			[]string{"aaa", "aaaa", "bbb"},
			"aaa",
			Parameters{
				Count: true,
			},
			"2\n",
		},
		{
			"context",
			[]string{"aaa", "aaaa", "bbb", "ccc"},
			"bbb",
			Parameters{
				After: 2,
				Before: 2,
			},
			"aaa\naaaa\nbbb\nccc\n",
		},
	}


	for _, testCase := range testCases {
		builder := strings.Builder{}
		Grep(testCase.input, testCase.target, testCase.params, &builder)
		result := builder.String()
		if result != testCase.expected {
			t.Errorf("testing %s, expected: %s, got: %s", testCase.name, testCase.expected, result)
		}
	}
}