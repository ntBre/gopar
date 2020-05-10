package main

import (
	"fmt"
	"io"
	"io/ioutil"
)

func ReadInput(filename string) (lines string) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

func IsSpace(c byte) bool {
	spaces := "\n\t "
	for i, _ := range spaces {
		if c == spaces[i] {
			return true
		}
	}
	return false
}

func ParseInputString(input string) (names, tokens []string) {
	var (
		inName  bool = true
		inToken bool = false
		name    string
		token   string
	)

	for i := 0; i < len(input); i++ {
		c := input[i]
		switch {
		case i == len(input)-1:
			token += string(c)
			tokens = append(tokens, token)
			token = ""
		case c == ' ' && inName:
			names = append(names, name)
			name = ""
			inName = false
			inToken = true
		case c == '\n' && inName:
			names = append(names, name)
			name = ""
			tokens = append(tokens, token)
			token = ""
		case c == '\n' && !inName:
			inName = true
			tokens = append(tokens, token)
			token = ""
			inToken = false
		case c != ' ' && inName:
			name += string(c)
		case c != ' ' && inToken:
			token += string(c)
		}
	}
	return
}

func MakeImports(imports ...string) (line string) {
	for _, s := range imports {
		line += "\"" + s + "\"\n"
	}
	return
}

func MakeGo(names, tokens []string) (lines string) {
	// could also read imports from input file
	imports := MakeImports("regexp", "io/ioutil")
	lines += fmt.Sprintf("package main\nimport (\n%s)\n", imports)
	lines += fmt.Sprintf("func main() {\n")
	// TODO read input file
	// build expressions
	lines += "regexes := []*regexp.Regexp{\n"
	for i, _ := range names {
		lines += fmt.Sprintf("regexp.MustCompile(`%s`),\n", names[i])
	}
	// close regexes
	lines += "}\n"
	// TODO match expressions and build tokens
	// close main
	lines += fmt.Sprintf("}\n")
	return
}

func WriteGo(names, tokens []string, w io.Writer) {
	lines := MakeGo(names, tokens)
	w.Write([]byte(lines))
}

// The first element is a regular expression, optionally inside quotes
// quotes are not allowed right now
// the second thing is where it gets stored
// names -> regexp for matching
// tokens -> if regexp.Match(names) {thing.Token =

// what should the parser return? right now im working on the code that will write the parser
