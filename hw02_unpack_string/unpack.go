package hw02_unpack_string //nolint:golint,stylecheck

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(comp_str string) (string, error) { // 60 max length in the 2020 !?

	var sbResult strings.Builder
	var sbEscSeq strings.Builder

	comp_str = strings.ReplaceAll(comp_str, " ", "")  // remove all white spaces
	comp_str = strings.ReplaceAll(comp_str, "\t", "") // remove all white spaces
	for i := 1; i < len(comp_str); i++ {
		if unicode.IsDigit(rune(comp_str[0])) {
			// statement will cath fist digit
			return "", ErrInvalidString
		}

		if i == len(comp_str)-1 && string(comp_str[len(comp_str)-1]) == `\` {
			continue
		}

		current := comp_str[i-1]
		next := comp_str[i]
		if sbEscSeq.Len() > 0 && unicode.IsDigit(rune(next)) { // switch not an issue
			var num int
			num, _ = strconv.Atoi(string(next))
			if num == 0 {
				continue
			}
			sbResult.WriteString(strings.Repeat(sbEscSeq.String(), num))
			sbEscSeq.Reset()
			continue
		} else if sbEscSeq.Len() > 0 && !unicode.IsDigit(rune(next)) {
			sbResult.WriteString(sbEscSeq.String())
			sbEscSeq.Reset()
			continue
		} else if unicode.IsDigit(rune(current)) && !unicode.IsDigit(rune(next)) {
			if i == len(comp_str)-1 {
				sbResult.WriteByte(next)
			}
			continue
		} else if string(current) == `\` {
			// create sequences of chars
			if string(next) == `n` { // switch not the best choise here
				sbEscSeq.WriteByte(current)
				sbEscSeq.WriteByte(next)
			} else if string(next) == `\` || unicode.IsDigit(rune(next)) {
				sbEscSeq.WriteByte(next)
			} else if string(next) == "`" {
				return "", ErrInvalidString
			} else {
				sbResult.WriteByte(current)
				if i == len(comp_str)-1 {
					sbResult.WriteByte(next)
				}
			}
			if i == len(comp_str)-1 {
				sbResult.WriteString(sbEscSeq.String())
			}
		} else if !unicode.IsDigit(rune(current)) && unicode.IsDigit(rune(next)) {
			var num int
			num, _ = strconv.Atoi(string(next))
			if num == 0 {
				continue
			}
			sbResult.WriteString(strings.Repeat(string(current), num))
		} else {
			sbResult.WriteByte(current)
			if i == len(comp_str)-1 {
				sbResult.WriteByte(next)
			}
		}

		if unicode.IsDigit(rune(current)) &&
			unicode.IsDigit(rune(next)) && sbEscSeq.Len() == 0 {
			return "", ErrInvalidString
		}
	}

	return sbResult.String(), nil
}
