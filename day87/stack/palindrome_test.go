package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPalindrome(t *testing.T) {

	assert.True(t, IsPalindrome(121))
	assert.False(t, IsPalindrome(810), "810 Is marked as palindrome")
}
