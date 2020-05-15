package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
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

func ParseInputString(input string) (names, tokens, trims []string) {
	var (
		inName    bool = true
		inToken   bool = false
		inBetween bool = false
		name      string
		token     string
		trim      string
	)

	for i := 0; i < len(input); i++ {
		c := input[i]
		switch {
		case c != ' ' && c != '\n' && inBetween:
			trim += string(c)
		case (c == ' ' || c == '\n') && inBetween:
			inBetween = false
			names = append(names, name)
			name = ""
			inName = false
			inToken = true
		case i == len(input)-1 && inToken:
			token += string(c)
			tokens = append(tokens, token)
			trims = append(trims, trim)
			trim = ""
		case i == len(input)-1 && inName:
			name += string(c)
			names = append(names, name)
			tokens = append(tokens, token)
			trims = append(trims, trim)
			trim = ""
		case c == '/' && inName:
			inBetween = true
		case c == ' ' && inName:
			names = append(names, name)
			name = ""
			inName = false
			inToken = true
		case c == '\n' && inName:
			names = append(names, name)
			name = ""
			trims = append(trims, trim)
			trim = ""
			tokens = append(tokens, token)
			token = ""
		case c == '\n' && !inName:
			inName = true
			trims = append(trims, trim)
			trim = ""
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

func MakeGo(names, tokens, trims []string) (lines string) {
	// could also read imports from input file
	imports := MakeImports("regexp", "strings")
	lines += fmt.Sprintf("package main\nimport (\n%s)\n", imports)
	lines += fmt.Sprintf("func ParseText(text []byte) interface{} {\n")
	// read input file
	lines += "type Regexp struct {\nExpr *regexp.Regexp\nTokenize bool\nTrim string\n}\n"
	lines += "regexes := []Regexp{\n"
	tokenString := ""
	goodTokens := make([]string, 0)
	for i, _ := range names {
		tokenize := "false"
		if tokens[i] != "" {
			tokenString += fmt.Sprintf("%s interface{}\n", tokens[i])
			tokenize = "true"
			goodTokens = append(goodTokens, tokens[i])
		}
		lines += fmt.Sprintf("Regexp{regexp.MustCompile(`%s`),"+
			"%s, %q},\n", names[i], tokenize, trims[i])
	}
	// close regexes
	lines += "}\n"
	lines += "type Tokens struct {\n" + tokenString
	// close struct
	lines += "}\n"
	// match expressions and build tokens
	// open for loop
	lines += fmt.Sprintln("tokenSlice := make([]string, 0)")
	lines += fmt.Sprintf("%s", "for _, regex := range regexes {\n")
	lines += fmt.Sprintln("var (\nendices []int\nmatch []byte\n)")
	lines += fmt.Sprintln("if regex.Trim != \"\" {")
	lines += fmt.Sprintln("endices = regexp.MustCompile(regex.Trim)." +
		"FindStringIndex(string(text))")
	lines += fmt.Sprintln("}")
	lines += fmt.Sprintln("if regex.Expr.Match(text) {")
	lines += fmt.Sprintln("matchIndices := regex.Expr.FindStringIndex(string(text))")
	lines += fmt.Sprintln("if regex.Trim != \"\" {")
	lines += fmt.Sprintln("match = text[matchIndices[0]:endices[0]]")
	lines += fmt.Sprintln("text = text[endices[1]:]")
	lines += fmt.Sprintln("} else {")
	lines += fmt.Sprintln("match = text[matchIndices[0]:matchIndices[1]]")
	lines += fmt.Sprintln("text = text[matchIndices[1]:]")
	lines += fmt.Sprintln("}")
	lines += fmt.Sprintln("if regex.Tokenize {")
	lines += fmt.Sprintln("tokenSlice = append(tokenSlice, strings.ReplaceAll(string(match), " +
		"\"\\n\", \" \"))")
	lines += fmt.Sprintln("}")
	// close if
	lines += fmt.Sprintf("}\n")
	// close for loop
	lines += fmt.Sprintf("}\n")
	lines += fmt.Sprintln("t := Tokens{}")
	for i, _ := range goodTokens {
		lines += fmt.Sprintf("t.%s = tokenSlice[%d]\n", goodTokens[i], i)
	}
	lines += fmt.Sprintf("return t")
	// close main
	lines += fmt.Sprintf("}\n")
	return
}

func WriteGo(names, tokens, trims []string, w io.Writer) {
	lines := MakeGo(names, tokens, trims)
	w.Write([]byte(lines))
}

func main() {
	names, tokens, trims := ParseInputString(ReadInput("test.in"))
	f, _ := os.Create("test.go")
	WriteGo(names, tokens, trims, f)
}
