package wget

import (
	"log"
	"net/url"
	"regexp"
	"testing"
)

func TestFixLinks(t *testing.T){
	srcRE, err := regexp.Compile(`src=\"(.*?)\"`)
	if err != nil {
		log.Panicf("regex is incorrect: %s", err)
	}
	hrefRE, err := regexp.Compile(`href=\"(.*?)\"`)
	if err != nil {
		log.Panicf("regex is incorrect: %s", err)
	}

	googleURL, err := url.Parse("http://google.com")
	if err != nil {
		log.Panicf("url is incorrect: %s", err)
	}
	gnuURL, err := url.Parse("https://www.gnu.org/software/wget")
	if err != nil {
		log.Panicf("url is incorrect: %s", err)
	}
	var testCases = []struct{
		input string
		base *url.URL
		re *regexp.Regexp
		expected string
	}{
		{
			`<script async="" src="/w/load.php?lang=en&amp;modules=startup&amp;only=scripts&amp;raw=1&amp;skin=vector"></script>`,
			googleURL,
			srcRE,
			`<script async="" src="http://google.com/w/load.php?lang=en&amp;modules=startup&amp;only=scripts&amp;raw=1&amp;skin=vector"></script>`,
		},
		{
			`<link rel="icon" type="image/png" href="/graphics/gnu-head-mini.png" />`,
			gnuURL,
			hrefRE,
			`<link rel="icon" type="image/png" href="https://www.gnu.org/graphics/gnu-head-mini.png" />`,
		},
	}

	
	
	for _, testCase := range testCases {
		result := fixLinks(testCase.input, testCase.base, testCase.re)
		if result != testCase.expected {
			t.Errorf("expected: %s, got: %s\n", testCase.expected, result)
		}
	}
}


func TestMakeFileName(t *testing.T) {
	var testCases = []struct{
		input string
		expected string
		expectedError bool
	}{
		{
			"http://examgle.org",
			"index.html",
			false,
		},
		{
			"http://examgle.org/",
			"index.html",
			false,
		},
		{
			"http://examgle.org/asdf",
			"asdf.html",
			false,
		},
		{
			"http://examgle.org/asdf.jpg",
			"asdf.jpg",
			false,
		},
		{
			":asdfasdf",
			"",
			true,
		},
	}

	for _, testCase := range testCases {
		result, err := MakeFileName(testCase.input)
		if result != testCase.expected {
			t.Errorf("expected: %s, got: %s", testCase.expected, result)
		}
		if err != nil && !testCase.expectedError {
			t.Errorf("expected no error, got %s", err)
		}
		if err == nil && testCase.expectedError {
			t.Errorf("expected error, got nil")
		}
	}
}