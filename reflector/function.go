package reflector

import (
	"reflect"
	"runtime"
)

// GetFunctionName returns function name in string
func GetFunctionName(funcPtr interface{}) string {
	refVal := reflect.ValueOf(funcPtr)

	if refVal.IsValid() {
		return runtime.FuncForPC(refVal.Pointer()).Name()
	} else {
		return ""
	}
}
