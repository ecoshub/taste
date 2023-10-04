package utils

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ecoshub/jin"
)

// Error messages
const (
	ErrStringTypeExpectation  = "type expectation failed. expected type: '%s', got type: '%s', path: '%s'"
	ErrStringValueExpectation = "value expectation failed. expected value: '%s', got value: '%s', path: '%s'"
	ErrStringMissingType      = "field must define a type. field: '%s'"
	ErrStringRequiredField    = "field is required. field: '%s'"
)

// expect represents an expected field value.
type expect struct {
	Field      string // The name of the field.
	Value      string // The expected value of the field.
	Type       string // The expected type of the field.
	Required   bool   // Indicates whether the field is required or not.
	IsWildcard bool   // Indicates whether the value of the field is a wildcard.
}

// Validate validates if the expected byte slice matches the given byte slice.
func Validate(expected, got []byte) error {
	// If the expected and given byte slices are equal, return nil.
	if string(expected) == string(got) {
		return nil
	}
	// If the expected byte slice is empty, return nil.
	if len(expected) == 0 {
		return nil
	}
	// Check if the expected byte slice is empty using jin.
	isEmpty, err := jin.IsEmpty(expected)
	if err != nil {
		return err
	}
	if isEmpty {
		return nil
	}
	// Create slices to store the expected and real paths.
	pathExpected := []string{}
	pathReal := []string{}
	pathsReal := make([][]string, 0, 8)
	pathsExpected := make([][]string, 0, 8)
	// Use the tree function to get the expected and real paths.
	err = tree(expected, pathExpected, pathReal, func(pathExpect []string, pathReal []string) (bool, error) {
		if len(pathReal) == 0 {
			return true, nil
		}
		newPathReal := make([]string, len(pathReal))
		newPathExpected := make([]string, len(pathExpect))
		for i := range newPathReal {
			newPathReal[i] = pathReal[i]
			newPathExpected[i] = pathExpect[i]
		}
		pathsExpected = append(pathsExpected, newPathExpected)
		pathsReal = append(pathsReal, newPathReal)
		return true, nil
	})
	if err != nil {
		return err
	}
	// Compare the expected and real values using the paths.
	err = compare(expected, got, pathsExpected, pathsReal)
	if err != nil {
		return err
	}
	// Use jin to walk through the given byte slice and check if all the paths exist.
	jin.Walk(got, func(_ string, _ []byte, path []string) (bool, error) {
		pathString := pathToPathString(path)
		exists := false
		for _, realPath := range pathsReal {
			realPathString := pathToPathString(realPath)
			if pathString == realPathString {
				exists = true
				break
			}
		}
		if !exists {
			err = fmt.Errorf("unexpected path. path: %v", path)
			return false, err
		}
		return true, nil
	})
	if err != nil {
		return err
	}
	// If no errors were found, return nil.
	return nil
}

func pathToPathString(p []string) string {
	return strings.Join(p, ":")
}

// compare function compares two byte arrays, and returns an error if they don't match
// `expect` and `got` represent the expected and actual byte arrays, respectively
// `pathsExpected` and `pathsReal` are slices of slices of strings, where each slice represents a path to a value in a JSON object
func compare(expect, got []byte, pathsExpected, pathsReal [][]string) error {
	// Iterate over the expected paths
	for i := range pathsExpected {
		// Get the current path for the actual JSON object
		realPath := pathsReal[i]
		// Get the current path for the expected JSON object
		expectedPath := pathsExpected[i]
		// Get the last key in the actual path
		realKey := getLastPath(realPath)
		// Get the last key in the expected path
		expectedKey := getLastPath(expectedPath)

		// Get the expected value using the expected path
		expectedValue, err := jin.GetString(expect, expectedPath...)
		if err != nil {
			return fmt.Errorf("fatal parsing error. %v", err)
		}

		// Resolve the expected value, and get the expected type and whether or not it is required
		e, err := resolve(expectedKey, expectedValue, expectedPath)
		if err != nil {
			return err
		}

		// Get the actual value using the actual path
		realValue, err := jin.GetString(got, realPath...)
		// Check if the key exists in the actual object
		exists := err == nil
		if !e.Required {
			if !exists {
				// Skip this iteration if the key is not required and doesn't exist
				continue
			}
		}

		// Get the expected type
		expectedType, err := jin.GetType(expect, expectedPath...)
		if err != nil {
			return fmt.Errorf("fatal parsing error. %v", err)
		}

		// Check if the key exists in the expected object
		if err != nil {
			if jin.ErrEqual(err, jin.ErrCodeKeyNotFound) {
				if e.Required {
					// Return an error if the key is required and doesn't exist in the expected object
					return fmt.Errorf(ErrStringRequiredField, e.Field)
				}
			} else {
				// Return an error if there was an error getting the expected value
				return fmt.Errorf("%s. path: %v", err, realPath)
			}
		}

		// Get the actual type
		realType, err := jin.GetType(got, realPath...)
		if err != nil {
			if jin.ErrEqual(err, jin.ErrCodeKeyNotFound) {
				if e.Required {
					// Return an error if the key is required and doesn't exist in the actual object
					return fmt.Errorf(ErrStringRequiredField, e.Field)
				}
			} else {
				// Return an error if there was an error getting the actual value
				return fmt.Errorf("%s. path: %v", err, realPath)
			}
		}

		// If the actual value is a number, process the type accordingly
		if realType == jin.TypeNumber {
			realType = processNumberType(realValue)
		}

		// If the expected value is a number, process the type accordingly
		if expectedType == jin.TypeNumber {
			expectedType = processNumberType(realValue)
		}

		// Check if the expected and actual keys match
		if realKey != e.Field {
			return fmt.Errorf("fatal parsing error. keys are not same. key: '%v', key: '%v'", realKey, expectedKey)
		}

		if e.Type != "" {
			// If the expected type is not empty, check if the real type matches it
			if realType != e.Type {
				// If the real type does not match the expected type, return an error
				return fmt.Errorf(ErrStringTypeExpectation, e.Type, realType, realPath)
			}
		} else {
			if !e.IsWildcard {
				// If the expected type is empty, check if the type of the value in the real JSON matches the type in the expected JSON
				if expectedType != realType {
					// If the types do not match, return an error
					return fmt.Errorf(ErrStringTypeExpectation, expectedType, realType, realPath)
				}
			}
		}

		if e.IsWildcard {
			continue
		}

		if !(realType == "array" || realType == "object") {
			// If the real type is not an array or object, check if the value matches the expected value
			if realValue != e.Value {
				// If the values do not match, return an error
				return fmt.Errorf(ErrStringValueExpectation, e.Value, realValue, realPath)
			}
		}
	}
	return nil
}

// processNumberType takes in a string value and returns the string representation of its type.
func processNumberType(value string) string {
	_, err := strconv.Atoi(value)
	if err == nil {
		return "int"
	}

	_, err = strconv.ParseFloat(value, 64)
	if err == nil {
		return "float"
	}

	return value
}

// resolve function takes in a key, value, and path as input arguments and returns a pointer to expect struct and an error
func resolve(key, value string, path []string) (*expect, error) {
	// create a new instance of expect struct
	e := &expect{}

	// expected field always required
	e.Required = true

	// unless field starts with an asterisk (*)
	if strings.HasPrefix(key, "*") {
		key = key[1:]
		if key != "*" {
			e.Required = false
		}
	}

	// split the key string by '|' separator and handle different cases
	tokens := strings.Split(key, "|")
	switch len(tokens) {
	case 0:
		return nil, fmt.Errorf(ErrStringMissingType, path)
	case 1:
		// if there is only one token, set the Field field of the expect struct to that token
		e.Field = tokens[0]
	default:
		// if there are two tokens, set the Field field to the first token and the Type field to the second token
		e.Field = tokens[0]
		e.Type = tokens[1]
	}

	// check if the value is a wildcard and set the IsWildcard field accordingly
	if value == "*" {
		e.IsWildcard = true
	}
	e.Value = value

	// return the expect struct and nil error if there are no issues
	return e, nil
}

// tree function recursively traverses the provided JSON body, applying the provided function 'f' to each element
// while keeping track of the expected and actual paths.
// If 'f' returns false, the recursion stops and the function returns nil.
func tree(body []byte, pathExpect []string, pathReal []string, f func(pathExpect []string, pathReal []string) (bool, error)) error {
	// Call the provided function 'f' with the current expected and actual paths and check if it returns false.
	keepRunning, err := f(pathExpect, pathReal)
	if err != nil {
		return err
	}

	if !keepRunning {
		return nil
	}

	// If the JSON body is empty, there is nothing to traverse.
	if len(body) == 0 {
		return nil
	}

	// Determine the type of the current JSON element.
	t, err := jin.GetType(body)
	if err != nil {
		return err
	}

	switch t {
	case jin.TypeObject:
		// If the current element is a JSON object, iterate over its key-value pairs.
		err = jin.IterateKeyValue(body, func(kb, vb []byte) (bool, error) {
			key := string(kb)
			tok := strings.Split(key, "|")
			pathExpect = append(pathExpect, key)
			// Add the current key to the actual path, stripping the '*' prefix if it exists.
			key = strings.TrimPrefix(tok[0], "*")
			pathReal = append(pathReal, key)
			err = tree(vb, pathExpect, pathReal, f)
			if err != nil {
				return false, err
			}
			// Remove the current key from both paths after finishing recursion.
			pathExpect = popLastPath(pathExpect)
			pathReal = popLastPath(pathReal)
			return true, nil
		})
		if err != nil {
			return err
		}
	case jin.TypeArray:
		// If the current element is a JSON array, iterate over its values.
		index := 0
		err = jin.IterateArray(body, func(val []byte) (bool, error) {
			indexString := strconv.Itoa(index)
			pathExpect = append(pathExpect, indexString)
			pathReal = append(pathReal, indexString)
			err = tree(val, pathExpect, pathReal, f)
			if err != nil {
				return false, err
			}
			// Remove the current index from both paths after finishing recursion.
			pathExpect = popLastPath(pathExpect)
			pathReal = popLastPath(pathReal)
			index++
			return true, nil
		})
		if err != nil {
			return err
		}
	}
	return nil
}

// popLastPath takes a slice of strings representing a path and returns a new slice with the last element removed.
// If the original slice is empty or has only one element, an empty slice is returned. Otherwise, a new slice with the
// last element removed is returned. This function does not modify the original slice.
func popLastPath(path []string) []string {
	switch len(path) {
	case 0:
		return path
	case 1:
		return []string{}
	default:
		return path[:len(path)-1]
	}
}

// getLastPath returns the last element of the provided path.
func getLastPath(path []string) string {
	if len(path) == 1 {
		return path[0]
	}
	return path[len(path)-1]
}
