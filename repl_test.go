package main

import (
	"testing"
	"time"

	"github.com/kadlex-web/pokedexcli/internal/pokecache"
)

func TestCleanInput(t *testing.T) {
	/* Need to tests to check the following
	1) Trim leading or trailing whitespace
	2) All words need to be lowercase
	3) all words are stored in a []string (slice of strings)
	*/

	// Begin by creating a []struct to fill with cases
	cases := []struct {
		input    string
		expected []string
	}{
		// define the expected test cases here
		{ // test trailing and leading whitespace
			input:    "    hello world   ",
			expected: []string{"hello", "world"},
		},
		{ //test all caps
			input:    "HELLO TOM!",
			expected: []string{"hello", "tom!"},
		},
		{ // testing multi-spaces between words
			input:    "hello     james",
			expected: []string{"hello", "james"},
		},
	}
	// for every case; use method cleanInput on the input field of the case
	for _, c := range cases {
		actual := cleanInput(c.input)
		// check length of cleaned input slice against what it is expected to be. if they dont match -- fail the test
		if len(actual) != len(c.expected) {
			t.Errorf("length of cleaned input: %d does not match length of expected cleaned input: %d", len(c.input), len(c.expected))
		}
		// for every word in actual (input phrase)
		for i := range actual {
			// get word at position i in the slice
			word := actual[i]
			// get expected word as the element at positon i in the slice
			expectedWord := c.expected[i]
			// if the words don't match fail the test
			if word != expectedWord {
				t.Errorf("%s does not match %s", word, expectedWord)
			}
		}
	}
}

func TestAddGetRemove(t *testing.T) {
	// add keys to the cache and see if they are accessible
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "http://google.com/",
			val: []byte("silly data"),
		},
		{
			key: "http://example.com",
			val: []byte("a fine example"),
		},
	}
	for _, c := range cases {
		cache := pokecache.NewCache(5 * time.Second)
		cache.Add(c.key, c.val)
		val, ok := cache.Get(c.key)
		if !ok {
			t.Errorf("expected to find key")
			return
		}
		if string(val) != string(c.val) {
			t.Errorf("expected to find value")
			return
		}

		cache.Remove(c.key)
		_, ok = cache.Get(c.key)
		if ok {
			t.Errorf("value should be deleted")
		}
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + (5 * time.Millisecond)
	cache := pokecache.NewCache(baseTime)
	cache.Add("https://example.com", []byte("data"))

	_, ok := cache.Get("https://example.com")
	if !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get("https://example.com")
	if ok {
		t.Errorf("expected to not find the key")
		return
	}
}
