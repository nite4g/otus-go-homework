package hw02_unpack_string //nolint:golint,stylecheck

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type test struct {
	input    string
	expected string
	err      error
}

func TestUnpack(t *testing.T) {
	for _, tst := range [...]test{
		{
			input:    "a4bc2d5e",
			expected: "aaaabccddddde",
		},
		{
			input:    "abccd",
			expected: "abccd",
		},
		{
			input:    "3abc",
			expected: "",
			err:      ErrInvalidString,
		},
		{
			input:    "45",
			expected: "",
			err:      ErrInvalidString,
		},
		{
			input:    "aaa10b",
			expected: "",
			err:      ErrInvalidString,
		},
		{
			input:    "",
			expected: "",
		},
		{
			input:    "aaa0b",
			expected: "aab",
		},
		{
			input:    "aaa0",
			expected: "aa",
		},
		{
			input:    "aab4",
			expected: "aabbbb",
		},
		{
			input:    "abcf1cd",
			expected: "abcfcd",
		},
		{
			input:    "ab-2c",
			expected: "ab--c",
		},
	} {
		result, err := Unpack(tst.input)
		require.Equal(t, tst.err, err)
		require.Equal(t, tst.expected, result)
	}
}

func TestUnpackWithEscape(t *testing.T) {

	for _, tst := range [...]test{
		{
			input:    `qwe\4\5`,
			expected: `qwe45`,
		},
		{
			input:    `qwe\45`,
			expected: `qwe44444`,
		},
		{
			input:    `qwe\\5`,
			expected: `qwe\\\\\`,
		},
		{
			input:    `qwe\\\3`,
			expected: `qwe\3`,
		},
		{
			input:    `gfgr\3`,
			expected: `gfgr3`,
		},
		{
			input:    `dfe\n1\`,
			expected: `dfe\n`,
		},
		{
			input:    `f\n2xr`,
			expected: `f\n\nxr`,
		},
		{
			input:    `a4bc2d5e\a`,
			expected: `aaaabccdddddea`,
		},
		{
			input:    `\ 3фыва`,
			expected: `3фыва`,
		},
		{
			input:    `a\30hh`,
			expected: `ahh`,
		},
		{
			input:    `v\n0e`,
			expected: `ve`,
		},
		{
			input:    "\\` 3фыва",
			expected: "",
			err:      ErrInvalidString,
		},
	} {
		result, err := Unpack(tst.input)
		require.Equal(t, tst.err, err)
		require.Equal(t, tst.expected, result)
	}
}
