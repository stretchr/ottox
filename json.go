package ottox

import (
	"encoding/json"
	"fmt"
	"github.com/robertkrimen/otto"
)

// toJson marshals the specified object into a JSON string.
func toJson(o interface{}) string {
	if bytes, err := json.Marshal(o); err == nil {
		return string(bytes)
	}
	return ""
}

// makeMap makes a map inside the specified runtime, and returns the
// otto.Value containing the object, or an error if something went wrong.
func makeMap(runtime *otto.Otto, m map[string]interface{}) (otto.Value, error) {
	return runtime.Run(fmt.Sprintf("eval(%s);", toJson(m)))
}
