package token

import "fmt"

// Type alias for a string
type Type string

const (
	// Unrecognize token or character
	Ilegal Type = "ILEGAL"

	// End of file
	EOF Type = "EOF"

	// Literals
	String Type = "STRING"
	Number Type = "Number"

	// Structural tokens
	LeftBrace    Type = "{"
	RightBrace   Type = "}"
	LeftBracket  Type = "["
	RightBracket Type = "]"
	Comma        Type = ","
	Colon        Type = ":"
	Whitespace   Type = "WHITESPACE"

	// Comments
	LineComment  Type = "//"
	BlockComment Type = "/*"

	// Values
	True  Type = "TRUE"
	False Type = "FALSE"
	Null  Type = "NULL"

	// Commands
	Post   Type = "POST"
	Get    Type = "GET"
	Header Type = "HEADER"
	Http   Type = "HTTP"
)

type Token struct {
	Type    Type
	Literal string
	Line    int
	Start   int
	End     int
}

var validIdentifiers = map[string]Type{
	"POST":   Post,
	"GET":    Get,
	"header": Header,
	"HTTP":   Http,
}

func LookupIdentifier(identifier string) (Type, error) {
	if token, ok := validIdentifiers[identifier]; ok {
		return token, nil
	}
	return "", fmt.Errorf("error: expected a valid identifier, found: %s", identifier)
}
