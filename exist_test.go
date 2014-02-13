package ottox

import (
	"github.com/robertkrimen/otto"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExist(t *testing.T) {

	js := otto.New()

	assert.False(t, Exist(js, "myVar"))

	js.Set("myVar", otto.TrueValue())
	assert.True(t, Exist(js, "myVar"), "true should exist")

	js.Set("myVar", otto.FalseValue())
	assert.True(t, Exist(js, "myVar"), "false still exists")

	js.Set("myVar", otto.UndefinedValue())
	assert.False(t, Exist(js, "myVar"), "undefined doesn't exist")

}
