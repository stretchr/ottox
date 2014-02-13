package ottox

import (
	"github.com/robertkrimen/otto"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSetOnce(t *testing.T) {

	js := otto.New()

	if didSet, err := SetOnce(js, "myVar", otto.TrueValue()); assert.NoError(t, err) {
		assert.True(t, didSet)
		if val, err := js.Get("myVar"); assert.NoError(t, err) {
			if valBool, err := val.ToBoolean(); assert.NoError(t, err) {
				assert.True(t, valBool)
			}
		}
	}

	if didSet, err := SetOnce(js, "myVar", otto.FalseValue()); assert.NoError(t, err) {
		assert.False(t, didSet)
		if val, err := js.Get("myVar"); assert.NoError(t, err) {
			if valBool, err := val.ToBoolean(); assert.NoError(t, err) {
				assert.True(t, valBool)
			}
		}
	}

}
