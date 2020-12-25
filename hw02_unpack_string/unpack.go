package hw02_unpack_string //nolint:golint,stylecheck

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func repeatOrDel(seq *strings.Builder, char rune) (string, error) {
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
	st = strings.ReplaceAll(st, " ", "")  // remove all white spaces
	st = strings.ReplaceAll(st, "\t", "") // remove all white spaces

	if len(st) == 0 {
		return "", nil
	}

	if unicode.IsDigit(rune(st[0])) {
		return "", ErrInvalidString
	}

	return st, nil
}

var state string = `default`
var wasDigit bool // just to control pair digit in the raw

func Unpack(compStr string) (string, error) { //nolint:gocognit
	var charSeq strings.Builder
	var resultStr strings.Builder

	compStr, nErr := normString(compStr)
	if nErr != nil {
		return "", nErr
	}

	for _, char := range compStr {
		switch state {
		case `default`:
			if unicode.IsDigit(char) {
				if wasDigit {
					return "", ErrInvalidString
				}
				tStr, err := repeatOrDel(&charSeq, char)
				if err != nil {
					return "", ErrInvalidString
				}
				resultStr.WriteString(tStr)
				charSeq.Reset()
				wasDigit = true
				continue
			}

			if string(char) == `\` {
				state = `slashed`
				// resultStr.WriteString(charSeq.String())
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
		case `slashed`:
			if string(char) == "`" {
				return "", ErrInvalidString
			}
			if string(char) == `n` {
				charSeq.WriteString(`\`)
				charSeq.WriteRune(char)
			} else {
				charSeq.WriteRune(char)
			}
			state = `default`
		}
		wasDigit = false
	}
	if charSeq.Len() != 0 {
		resultStr.WriteString(charSeq.String())
	}

	return resultStr.String(), nil
}
