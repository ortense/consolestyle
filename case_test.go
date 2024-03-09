package consolestyle_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	input    string
	act      func(string) string
	expected string
}

func run(test *testing.T, cases map[string]testcase) {
	for name, tc := range cases {
		test.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.act(tc.input), tc.expected)
		})
	}
}
