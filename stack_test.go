package ds

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	st := NewDefaultStack()
	st.Push(10)
	assert.Equal(t, 10, st.Top())
}
