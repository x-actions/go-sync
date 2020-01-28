package utils

import (
	"testing"

	"gotest.tools/assert"
)


func TestUtils(t *testing.T) {
	assert.Equal(t, Md5sum("webhooks"), "C10F40999B74C408263F790B30E70EFE", "Md5sum method is mistake.")
}
