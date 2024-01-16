package hash_maps

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckArrayResult(t *testing.T) {
	assert.ElementsMatch(t, RunTwoNums([]int{2, 3, 4}, 6), []int{0, 2})
	assert.ElementsMatch(t, RunTwoNums([]int{3, 3}, 6), []int{0, 1})
	assert.ElementsMatch(t, RunTwoNums([]int{3, 2, 7, 4}, 7), []int{0, 3})

}
