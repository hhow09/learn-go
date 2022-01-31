package trie

//https://youtu.be/nL7BHR5vJDc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func getTestCases() []string {
	testCases := []string{"aragorn", "argon", "aragorg", "eragorn", "oregorn", "oreo", "orc"}
	return testCases
}

func TestTrie(t *testing.T) {
	trie := InitTrie()
	tc := getTestCases()
	for _, v := range tc {
		trie.Insert(v)
	}
	for _, v := range tc {
		assert.True(t, trie.Search(v))
	}
	assert.False(t, trie.Search("aragornm"))
	assert.False(t, trie.Search("abcabc"))
}
