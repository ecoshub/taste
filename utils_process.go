package taste

import (
	"strings"

	"github.com/ecoshub/jin"
)

const (
	symbolStart string = "<<"
	symbolEnd   string = ">>"
)

// inPlaceStoredValues put in place stored values
func inPlaceStoredValues(request map[string][]byte, symbol []byte) ([]byte, error) {
	var err error
	done := false
	for !done {
		symbol, done, err = inPlaceStoredValuesCore(request, symbol)
		if err != nil {
			return nil, err
		}
	}
	return symbol, nil
}

// inPlaceStoredValuesCore put in place stored values core function
func inPlaceStoredValuesCore(request map[string][]byte, symbol []byte) ([]byte, bool, error) {
	stringBody := string(symbol)
	start := strings.Index(stringBody, symbolStart)
	if start == -1 {
		return symbol, true, nil
	}
	end := strings.Index(stringBody, symbolEnd)
	if end == -1 {
		return symbol, true, nil
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
				return symbol, true, nil
			}
			return symbol, true, err
		}
	}
	newBody := stringBody[:start] + val + stringBody[end+len(symbolEnd):]
	symbol = []byte(newBody)
	return symbol, false, nil
}
