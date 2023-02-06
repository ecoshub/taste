package taste

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ecoshub/jin"
)

var (
	errStringTypeExpectation  string = "type expectation failed. expected type: '%s', got type: '%s', path: '%s'"
	errStringValueExpectation string = "value expectation failed. expected value: '%s', got value: '%s', path: '%s'"
	errStringMissingType      string = "field must define a type. filed: '%s'"
	errStringRequiredField    string = "field is required. field: '%s'"
)

// expect parsed expectation
type expect struct {
	field      string
	value      string
	jsonType   string
	required   bool
	isWildcard bool
}

// validate validate json value with expectations
func validate(expect, got []byte) error {
	// if no expectation defined but there is something to validate, return error
	if len(expect) == 0 && len(got) != 0 {
		return fmt.Errorf("no expectation specified but got response body. body: '%s'", got)
	}

	if len(got) == 0 {
		if len(expect) == 0 {
			return nil
		}
		return fmt.Errorf("expected something but got nothing (got nil body)")
	}

	// if its same return with no error
	if string(expect) == string(got) {
		return nil
	}

	// paths definitions
	pathExpect := []string{}
	pathReal := []string{}
	pathsReal := make([][]string, 0, 8)
	pathsExpected := make([][]string, 0, 8)

	// walk through all paths anca resolve real paths vs validation schemes
	err := tree(expect, pathExpect, pathReal, func(pathExpect []string, pathReal []string) (bool, error) {
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

	// this means it is not a JSON
	if len(pathExpect) == 0 && len(pathsReal) == 0 {
		return fmt.Errorf("unexpected result. got: '%s', expected: '%s'", got, expect)
	}

	err = compare(expect, got, pathsExpected, pathsReal)
	if err != nil {
		return err
	}
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
	return nil
}

func pathToPathString(p []string) string {
	return strings.Join(p, ":")
}

func compare(expect, got []byte, pathsExpected, pathsReal [][]string) error {
	for i := range pathsExpected {
		realPath := pathsReal[i]
		expectedPath := pathsExpected[i]
		realKey := getLastPath(realPath)
		expectedKey := getLastPath(expectedPath)

		expectedValue, err := jin.GetString(expect, expectedPath...)
		if err != nil {
			return fmt.Errorf("fatal parsing error. %v", err)
		}

		e, err := resolve(expectedKey, expectedValue, expectedPath)
		if err != nil {
			return err
		}

		realValue, err := jin.GetString(got, realPath...)
		exists := err == nil
		if !e.required {
			if !exists {
				continue
			}
		}

		_type, err := jin.GetType(expect, expectedPath...)
		if err != nil {
			return fmt.Errorf("fatal parsing error. %v", err)
		}

		if err != nil {
			if jin.ErrEqual(err, jin.ErrCodeKeyNotFound) {
				if e.required {
					return fmt.Errorf(errStringRequiredField, e.field)
				}
			} else {
				return fmt.Errorf("%s. path: %v", err, realPath)
			}
		}

		realType, err := jin.GetType(got, realPath...)
		if err != nil {
			if jin.ErrEqual(err, jin.ErrCodeKeyNotFound) {
				if e.required {
					return fmt.Errorf(errStringRequiredField, e.field)
				}
			} else {
				return fmt.Errorf("%s. path: %v", err, realPath)
			}
		}

		if realType == jin.TypeNumber {
			realType = processNumberType(realValue)
		}

		if _type == jin.TypeNumber {
			_type = processNumberType(realValue)
		}

		if realKey != e.field {
			return fmt.Errorf("fatal parsing error. keys are not same. key: '%v', key: '%v'", realKey, expectedKey)
		}

		if e.jsonType != "" {
			if realType != e.jsonType {
				return fmt.Errorf(errStringTypeExpectation, e.jsonType, realType, realPath)
			}
		} else {
			if _type != realType {
				return fmt.Errorf(errStringTypeExpectation, _type, realType, realPath)
			}
		}

		if !(realType == "array" || realType == "object") {
			if !e.isWildcard {
				if realValue != e.value {
					return fmt.Errorf(errStringValueExpectation, e.value, realValue, realPath)
				}
			}
		}
	}
	return nil
}

func processNumberType(value string) string {
	_, err := strconv.ParseInt(value, 10, 64)
	if err == nil {
		return "int"
	}
	_, err = strconv.ParseFloat(value, 64)
	if err == nil {
		return "float"
	}
	return value
}

func resolve(key, value string, path []string) (*expect, error) {
	e := &expect{}
	// strip '*' prefix if exists, and set the required field
	e.required = true
	if strings.HasPrefix(key, "*") {
		key = key[1:]
		e.required = false
	}

	tokens := strings.Split(key, "|")
	switch len(tokens) {
	case 0:
		return nil, fmt.Errorf(errStringMissingType, path)
	case 1:
		e.field = tokens[0]
	default:
		e.field = tokens[0]
		e.jsonType = tokens[1]
	}
	if value == "*" {
		e.isWildcard = true
	}
	e.value = value
	return e, nil
}

func tree(body []byte, pathExpect []string, pathReal []string, f func(pathExpect []string, pathReal []string) (bool, error)) error {
	keepRunning, err := f(pathExpect, pathReal)
	if err != nil {
		return err
	}

	if !keepRunning {
		return nil
	}

	if len(body) == 0 {
		return nil
	}

	t, err := jin.GetType(body)
	if err != nil {
		return err
	}

	switch t {
	case jin.TypeObject:
		err = jin.IterateKeyValue(body, func(kb, vb []byte) (bool, error) {
			key := string(kb)
			tok := strings.Split(key, "|")
			pathExpect = append(pathExpect, key)
			key = strings.TrimPrefix(tok[0], "*")
			pathReal = append(pathReal, key)
			err = tree(vb, pathExpect, pathReal, f)
			if err != nil {
				return false, err
			}
			pathExpect = pathExpect[:len(pathExpect)-1]
			pathReal = pathReal[:len(pathReal)-1]
			return true, nil
		})
		if err != nil {
			return err
		}
	case jin.TypeArray:
		index := 0
		err = jin.IterateArray(body, func(val []byte) (bool, error) {
			indexString := strconv.Itoa(index)
			pathExpect = append(pathExpect, indexString)
			pathReal = append(pathReal, indexString)
			err = tree(val, pathExpect, pathReal, f)
			if err != nil {
				return false, err
			}
			pathExpect = pathExpect[:len(pathExpect)-1]
			pathReal = pathReal[:len(pathReal)-1]
			index++
			return true, nil
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func getLastPath(path []string) string {
	if len(path) == 1 {
		return path[0]
	}
	return path[len(path)-1]
}
