package main

import "testing"

func TestCleanInput (t *testing.T) {
	/* Need to tests to check the following
		1) Trim leading or trailing whitespace
		2) All words need to be lowercase
		3) all words are stored in a []string (slice of strings)
	*/
	
	// Begin by creating a []struct to fill with cases
	cases := []struct {
		input string
		expected []string
	}{
	// define the expected test cases here
	{ // test trailing and leading whitespace
		input: "    hello world   ",
		expected: []string{"hello", "world"},
		
	},
	{ //test all caps
		input: "HELLO TOM!",
		expected: []string{"hello", "tom"},
	},
	{ // testing multi-spaces between words
		input: "hello     james",
		expected: []string{"hello", "james"},
	},
	}
}
// for every case; use method cleanInput on the input field of the case
for _, c := range cases {
	actual := cleanInput(c.input)
	// check length of cleaned input slice against what it is expected to be. if they dont match -- fail the test
	if len(actual) != len(expected) {
		t.Errorf(length of cleaned input: %d does not match length of expected cleaned input %d, len(actual), len(expected))
	}
	// for every word in actual (input phrase)
	for i := range actual {
		// get word at position i in the slice
		word := actual[i]
		// get expected word as the element at positon i in the slice
		expectedWord := c.expected[i]
		// if the words don't match fail the test
		if word != expectedWord {
			t.Errorf("%s does not match %s, word, expectedWord)
	}
}

/* Use strings.Fields??
	Could use strings.TrimSpace too? --> Trims all leading and trailing whitespace
 */