package main

import (
	"reflect"
	"testing"
)

var samples = []struct {
	sample   []string
	expected [][]string
}{
	{
		[]string{"kita", "atik", "tika", "aku", "kia", "makan", "kua"},
		[][]string{
			{"makan"},
			{"kita", "atik", "tika"},
			{"aku", "kua"},
			{"kia"},
		},
	},
}

func TestGroupAnagram(t *testing.T) {
	for _, s := range samples {
		if reflect.DeepEqual(groupAnagram(s.sample), s.expected) {
			t.Errorf(`got "%s" from "%s", expected "%s"`, groupAnagram(s.sample), s.sample, s.expected)
		}
	}
}

func BenchmarkGroupAnagram(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, s := range samples {
			groupAnagram(s.sample)
		}
	}
}
