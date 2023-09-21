package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDateToInt(t *testing.T) {
	testCases := []struct {
		input    []string
		expected []int
	}{
		{[]string{"2022-01-01"}, []int{6}},
		{[]string{"2022-01-01", "2022-01-02"}, []int{6, 0}},
		{[]string{"2022-01-25", "2022-01-02", "2022-01-03"}, []int{2, 0, 1}},
	}

	for _, k := range testCases {
		actual, err := DateToInt(k.input...)
		assert.Nil(t, err)
		assert.Equal(t, len(k.expected), len(actual))
		for i := range actual {
			assert.Equal(t, k.expected[i], actual[i])
		}
	}
}

func TestIsParamNull(t *testing.T) {
	testCases := []struct {
		input    []string
		expected bool
	}{
		{[]string{"a", "b", "c"}, false},
		{[]string{"a", "", "c"}, true},
		{[]string{"", "", ""}, true},
		{[]string{""}, true},
		{[]string{}, false},
	}

	for _, k := range testCases {
		actual := IsParamNull(k.input...)
		assert.Equal(t, k.expected, actual)
	}
}
