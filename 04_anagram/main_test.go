package main

import (
	"strings"
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
	{
		[]string{"abc", "bca", "cab", "aaa", "aa", "a"},
		[][]string{
			{"abc", "bca", "cab"},
			{"aaa"},
			{"aa"},
			{"a"},
		},
	},
	{
		[]string{"abc", "bca", "cab", "abc", "bca", "cab"},
		[][]string{
			{"abc", "bca", "cab", "abc", "bca", "cab"},
		},
	},
	{
		[]string{"", "", "a", "aa", "a", "aa"},
		[][]string{
			{"", ""},
			{"a", "a"},
			{"aa", "aa"},
		},
	},
	{
		[]string{""},
		[][]string{
			{""},
		},
	},
	{
		[]string{"a"},
		[][]string{
			{"a"},
		},
	},
	{
		[]string{},
		[][]string{},
	},
	{
		[]string{"aaa", "aab", "aac"},
		[][]string{
			{"aaa"},
			{"aab"},
			{"aac"},
		},
	},
	{
		[]string{"abc", "bcd", "cde"},
		[][]string{
			{"abc"},
			{"bcd"},
			{"cde"},
		},
	},
	{
		[]string{"abccba", "bccb", "cc"},
		[][]string{
			{"abccba"},
			{"bccb"},
			{"cc"},
		},
	},
}

func equal(output, expected [][]string) bool {
	if len(output) != len(expected) {
		return false
	}

	check := map[string]bool{}

	for _, v := range output {
		check[strings.Join(v, "")] = true
	}

	for _, v := range expected {
		key := strings.Join(v, "")

		if !check[key] {
			return false
		} else {
			check[key] = false
		}
	}

	return true
}

func TestGroupAnagram(t *testing.T) {
	for _, s := range samples {
		if !equal(groupAnagram(s.sample), s.expected) {
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
