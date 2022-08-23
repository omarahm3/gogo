package lib

import (
	"fmt"
	"testing"
)

func TestNormalize(t *testing.T) {
	tests := []struct {
		number   string
		expected string
	}{
		{
			number:   "1234567890",
			expected: "1234567890",
		},
		{
			number:   "123 456 7891",
			expected: "1234567891",
		},
		{
			number:   "(123) 456 7892",
			expected: "1234567892",
		},
		{
			number:   "(123) 456-7893",
			expected: "1234567893",
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("should normalize phone number [%s]", test.number), func(t *testing.T) {
			actual := Normalize(test.number)

			if actual != test.expected {
				t.Errorf("Expected %q, got %q\n", test.expected, actual)
			}
		})
	}
}
