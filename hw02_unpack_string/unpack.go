package hw02_unpack_string //nolint:golint,stylecheck

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func repeatOrDel(seq strings.Builder, char rune) (string, error) {
	var resStr string
	var num int
	num, err := strconv.Atoi(string(char))

	if err != nil {
		return "", ErrInvalidString
	}

	if num == 0 {
		return "", nil
	}

	resStr = strings.Repeat(seq.String(), num)

	return resStr, nil
}

func normString(st string) (string, error) {
	if len(st) == 0 {
		return "", nil
	}

	if unicode.IsDigit(rune(st[0])) {
		return "", ErrInvalidString
	}

	if string(st[len(st)-1]) == `\` {
		return "", ErrInvalidString
	}

	return st, nil
}

func Unpack(compStr string) (string, error) { //nolint:gocognit,funlen
	var wasDigit bool // just to control pair digit in the raw
	var charSeq strings.Builder
	var resultStr strings.Builder

	type states string
	const (
		defaultState states = "default"
		slashedState states = "slashed"
	)

	state := defaultState
	compStr, nErr := normString(compStr)

	if nErr != nil {
		return "", nErr
	}

	for _, char := range compStr {
		switch state {
		case defaultState:
			if unicode.IsDigit(char) {
				if wasDigit {
					return "", ErrInvalidString
				}
				tStr, err := repeatOrDel(charSeq, char)
				if err != nil {
					return "", ErrInvalidString
				}
				resultStr.WriteString(tStr)
				charSeq.Reset()
				wasDigit = true
				continue
			}

			if string(char) == `\` {
				state = slashedState
				if charSeq.Len() > 0 {
					resultStr.WriteString(charSeq.String())
				}
				charSeq.Reset()
			} else {
				if charSeq.Len() != 0 {
					resultStr.WriteString(charSeq.String())
					charSeq.Reset()
				}
				charSeq.WriteRune(char)
			}
		case slashedState:
			if string(char) != `\` && !unicode.IsDigit(char) {
				return "", ErrInvalidString
			}
			charSeq.WriteRune(char)

			state = defaultState
		default:
			state = defaultState
		}

		wasDigit = false
	}

	if charSeq.Len() != 0 {
		resultStr.WriteString(charSeq.String())
	}

	return resultStr.String(), nil
}
