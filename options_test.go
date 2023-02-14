package log

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_getOptions(t *testing.T) {
	options := getOptions()
	assert.Equal(t, options.Level, "info")
}
