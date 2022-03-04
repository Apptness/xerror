package xerror

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsError(t *testing.T) {
	assert.Equal(t, IsError(errors.New("boop")), true, "should be true")
	assert.Equal(t, IsError(nil), false, "should be false")
}
