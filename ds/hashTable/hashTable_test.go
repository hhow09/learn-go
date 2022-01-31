package hashTable

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func toCharStr(i int) rune {
	return rune('A' - 1 + i)
}
func buildTestCase(n int) ([]string, []string) {
	keyRes := make([]string, 0, n)
	valRes := make([]string, 0, n)
	key := ""
	val := ""
	for i := 0; i < n; i++ {
		key = ""
		val = ""
		for j := 0; j < 3; j++ {
			rand.Seed(time.Now().UnixNano())
			key = key + string(toCharStr(rand.Intn(26)))
		}
		for j := 0; j < 10; j++ {
			rand.Seed(time.Now().UnixNano())
			val = val + string(toCharStr(rand.Intn(26)))
		}
		keyRes = append(keyRes, key)
		valRes = append(valRes, val)
	}
	return keyRes, valRes
}

func TestInitHashTable(t *testing.T) {
	ht := InitHashTable()
	assert.Equal(t, ArraySize, len(ht.array), "should hold an array with defined ArraySize.")
}

func TestHash(t *testing.T) {
	inputs := []string{"abcd", "DEF", "ZZ988"}
	expected := []int{2, 4, 6}
	for i, v := range inputs {
		hashed := hash(v)
		assert.Equal(t, expected[i], hashed)
	}
}

func TestHashTable(t *testing.T) {
	keyInputs, valInputs := buildTestCase(5)
	ht := InitHashTable()
	for i, v := range keyInputs {
		ht.Insert(v, valInputs[i])
	}
	for i, v := range keyInputs {
		assert.Equal(t, valInputs[i], ht.Get(v), "should get inserted value")
	}
	assert.Nil(t, ht.Get("cc"), "should not get non-existed value")

	ht.Delete(keyInputs[2])
	assert.Nil(t, ht.Get(keyInputs[2]), "should not get deleted value")
}
