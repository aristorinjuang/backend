package main

import "testing"

var samples = []struct {
	sample   string
	expected string
}{
	{"Hello World!", "Hello World!"},
	{"Hello World! ()", "Hello World! ()"},
	{"Hello World! (x)", "x"},
	{"Hello (World!)", "World!"},
	{"Hello (W)orld!", "W"},
	{"Hello()World!", "Hello()World!"},
	{"Hello()World(!)", "Hello()World(!)"},
	{"Hello()World(! Hi)", "Hello()World(! Hi)"},
	{"Hello(Hi)World(! Hi)", "Hi"},
	{"()()()()()", "()()()()()"},
	{"()(x)()()()", "()(x)()()()"},
	{"(x)()()()()", "x"},
	{"x", "x"},
	{"", ""},
	{"(", "("},
}

func TestFindFirstStringInBracket(t *testing.T) {
	for _, s := range samples {
		if findFirstStringInBracket(s.sample) != s.expected {
			t.Errorf(`got "%s" from "%s", expected "%s"`, findFirstStringInBracket(s.sample), s.sample, s.expected)
		}
	}
}

func BenchmarkFindFirstStringInBracket(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, s := range samples {
			findFirstStringInBracket(s.sample)
		}
	}
}
