package convert

import (
	"github.com/pkg/errors"
	"reflect"
	"strconv"
)

func toInt(value reflect.Value) (int, error) {
	switch f := value; f.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16,reflect.Int32,reflect.Int64:
		return int(f.Int()), nil
	case reflect.Uint,reflect.Uint8, reflect.Uint16,reflect.Uint32,reflect.Uint64:
		return int(f.Uint()), nil
	case reflect.String:
		s := f.String()
		result, err := strconv.ParseInt(s, 10, 64)
		return int(result), err
	case reflect.Float32,reflect.Float64:
		fv := f.Float()
		return int(fv), nil
	case reflect.Bool:
		b := f.Bool()
		if b {
			return 1, nil
		} else {
			return 0, nil
		}
	case reflect.Ptr:
		return toInt(f.Elem())
	default:
		return 0, errors.New("data type not supported, type:" + f.Kind().String())
	}
}

// ToInt try to convert interface{} to int
// ToInt accepts data like: int, *int, uint8, *uint8, string, bool, *string, *bool, float32, *float32 etc
func ToInt(data interface{}) (int, error) {
	if data == nil {
		return 0, errors.New("source data is nil")
	}
	return toInt(reflect.ValueOf(data))
}
