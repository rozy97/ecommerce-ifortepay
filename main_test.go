package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_(t *testing.T) {
	t.Run("run example test", func(t *testing.T) {
		num := 2 + 2
		assert.Equal(t, 6, num)
	})

}
