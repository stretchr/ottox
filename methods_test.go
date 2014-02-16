package ottox

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestGenerateMethodName(t *testing.T) {

	m := generateMethodName("prefix")
	assert.True(t, strings.HasPrefix(m, "prefix"))
	assert.NotEqual(t, "prefix", m, "Should be more than just the prefix")

	m = generateMethodName()
	assert.True(t, strings.HasPrefix(m, "ottoxMethod"))
	assert.NotEqual(t, "ottoxMethod", m, "Should be more than just the prefix")

}
