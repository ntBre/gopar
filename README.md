# gopar

gopar is a very simple parser generator written in Go

## Usage

Example input is a series of regular expressions followed on the same
lines by what you want them to be called. Running gopar as a
standalone program requires passing the name of the input file as the
first command line argument and will generate a file called parse.go
that contains the function ParseText(), which takes a []byte and
returns a struct holding the fields defined in the input file and the
text that was matched by the corresponding regular expressions.
Alternatively, the parser can be generated from within other Go code
by calling the WriteParser() function, which takes the names of the
input and output files as arguments.

As of now, the parser only parses one instance of matching text and
anything after that will be lumped into the last field of the
output. Consequently, the caller is responsible for splitting the text
between groups of matches.

Since Go does not support negative lookbehinds, the regular expression
input is adapted slightly to support ignoring characters after a
forward-slash (/).  Any characters after a "/" in a regular expression
will also be matched but not included in the saved field.

## Example

The included test.in is an example input file for parsing a LaTeX
bibliography entry from test.parse. 