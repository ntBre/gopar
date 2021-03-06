package main

import (
	"reflect"
	"testing"
	"fmt"
	"io/ioutil"
	"strings"
)

// func TestReadInput(t *testing.T) {
// 	got := ReadInput("test.in")
// 	want := "@\narticle Type\n{\nMP2 Key"
// 	if got != want {
// 		t.Errorf("got %q, wanted %q", got, want)
// 	}
// }

func TestIsSpace(t *testing.T) {
	t.Run("tab", func(t *testing.T) {
		got := IsSpace('\t')
		want := true
		if got != want {
			t.Errorf("wrong")
		}
	})
	t.Run("newline", func(t *testing.T) {
		got := IsSpace('\n')
		want := true
		if got != want {
			t.Errorf("wrong")
		}
	})
	t.Run("space", func(t *testing.T) {
		got := IsSpace(' ')
		want := true
		if got != want {
			t.Errorf("wrong")
		}
	})
}

func TestParseInput(t *testing.T) {
	got1, got2, _ := ParseInputString("@\narticle Type\n{\nMP2 Key\n,")
	want1, want2 := []string{"@", "article", "{", "MP2", ","},
		[]string{"", "Type", "", "Key", ""}
	if !reflect.DeepEqual(got1, want1) || !reflect.DeepEqual(got2, want2) {
		t.Errorf("Something wrong, gots: %q, %q; wants: %q, %q",
			got1, got2, want1, want2)
	}
}

func TestMakeGo(t *testing.T) {
	// t.Errorf("got:\n%s", MakeGo(ParseInputString("@\narticle Type\n{\nMP2 Key")))
}

func TestParseText(t *testing.T) {
	text, err := ioutil.ReadFile("test.parse")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(text), "\n")
	for i, line := range lines {
		lines[i] = strings.TrimSpace(line)
	}
	text = []byte(strings.Join(lines, "\n"))
	got := ParseText(text)
	fmt.Printf("%#v\n", got)
}
