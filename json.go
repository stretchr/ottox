package ottox

import (
	"encoding/json"
	"fmt"
	"github.com/robertkrimen/otto"
)

func toJson(o interface{}) string {
	if bytes, err := json.Marshal(o); err == nil {
		return string(bytes)
	}
	return ""
}

func makeMap(runtime *otto.Otto, m map[string]interface{}) (otto.Value, error) {
	return runtime.Run(fmt.Sprintf("eval(%s);", toJson(m)))
}
