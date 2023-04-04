package utils

import (
	"errors"
	"strings"

	"github.com/ecoshub/jin"
)

const (
	symbolStart string = "<<"
	symbolEnd   string = ">>"
)

// ProcessBody processes the request body and replaces the symbols with values from the request.
// The symbols are expected to be in the format of "<<name>>" where 'name' is the name of the request parameter.
// If a request parameter is not found or an error occurs while processing the request, the symbol is not replaced.
func ProcessBody(request map[string][]byte, symbol []byte) ([]byte, error) {
	var err error
	done := false
	for !done {
		symbol, done, err = processBodyCore(request, symbol)
		if err != nil {
			return nil, err
		}
	}
	return symbol, nil
}

// processBodyCore is a helper function for ProcessBody that processes a single symbol.
// It searches for the first occurrence of '<<' and the last occurrence of '>>' in the given symbol.
// It then extracts the name of the request parameter from the symbol and tries to replace the symbol with the corresponding value from the request.
// If the request parameter is not found or an error occurs while processing the request, the symbol is not replaced.
// Returns the updated symbol, a boolean flag indicating if the processing is complete, and an error if one occurs.
func processBodyCore(request map[string][]byte, symbol []byte) ([]byte, bool, error) {
	stringBody := string(symbol)
	start := strings.Index(stringBody, symbolStart)
	if start == -1 {
		return symbol, true, nil
	}
	end := strings.Index(stringBody, symbolEnd)
	if end == -1 {
		return symbol, true, nil
	}
	if end < start {
		return nil, false, errors.New("malformed symbol")
	}
	line := stringBody[start+len(symbolStart) : end]
	tokens := strings.Split(line, ".")

	// Get the name of the request parameter
	name := tokens[0]

	// Get the value of the request parameter from the request
	body := request[name]

	var val string
	var err error
	switch len(tokens) {
	case 1:
		// If the symbol has only one token, replace it with the value of the request parameter
		val = string(body)
	default:
		// If the symbol has multiple tokens, try to replace it with the corresponding value from the request using the jin library
		val, err = jin.GetString(body, tokens[1:]...)
		if err != nil {
			// If the error is due to the key not found, skip the replacement and continue processing
			if jin.ErrEqual(err, jin.ErrCodeKeyNotFound) {
				return symbol, true, nil
			}
			// If the error is not due to the key not found, return the error
			return symbol, true, err
		}
	}

	header := stringBody[:start]
	footer := stringBody[end+len(symbolEnd):]

	// Replace the symbol with the value from the request
	newBody := header + val + footer
	symbol = []byte(newBody)

	return symbol, false, nil
}
