package utils

import (
	"strings"

	"github.com/ecoshub/jin"
)

const (
	symbolStart string = "<<"
	symbolEnd   string = ">>"
)

func ProcessBody(request map[string][]byte, symbol []byte) ([]byte, error) {
	stringBody := string(symbol)
	start := strings.Index(stringBody, symbolStart)
	if start == -1 {
		return symbol, nil
	}
	end := strings.Index(stringBody, symbolEnd)
	if end == -1 {
		return symbol, nil
	}
	line := stringBody[start+len(symbolStart) : end]
	tokens := strings.Split(line, ".")

	name := tokens[0]
	body := request[name]

	var val string
	var err error
	switch len(tokens) {
	case 1:
		val = string(body)
	default:
		val, err = jin.GetString(body, tokens[1:]...)
		if err != nil {
			if jin.ErrEqual(err, jin.ErrCodeKeyNotFound) {
				return symbol, nil
			}
			return symbol, err
		}
	}
	newBody := stringBody[:start] + val + stringBody[end+len(symbolEnd):]
	symbol = []byte(newBody)
	return symbol, nil
}
