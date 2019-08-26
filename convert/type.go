package convert

import (
	"fmt"
	"github.com/pkg/errors"
	"reflect"
	"strconv"
)

var (
	ErrSourceDataNil = errors.New("source data is nil")
	ErrTypeNotSupported = errors.New("data type not supported")
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
		return 0, ErrTypeNotSupported
	}
}

// ToInt try to convert interface{} to int
// ToInt accepts data like: int, *int, uint8, *uint8, string, bool, *string, *bool, float32, *float32 etc
func ToInt(data interface{}) (int, error) {
	if data == nil {
		return 0, ErrSourceDataNil
	}
	return toInt(reflect.ValueOf(data))
}

//	Must ToInt version
func MustToInt(data interface{}) int {
	result, err := ToInt(data)
	if err != nil {
		panic("converter: MustToInt(" + fmt.Sprintf("%v", data) + "): " + err.Error())
	}

	return result
}

func toString(value reflect.Value) (string, error) {
	switch f := value; f.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16,reflect.Int32,reflect.Int64:
		return strconv.FormatInt(f.Int(), 10), nil
	case reflect.Uint,reflect.Uint8, reflect.Uint16,reflect.Uint32,reflect.Uint64:
		return strconv.FormatUint(f.Uint(), 10), nil
	case reflect.String:
				return f.String(), nil
	case reflect.Float32,reflect.Float64:
		fv := f.Float()
		return strconv.FormatFloat(fv, 'f', -1, 64), nil
	case reflect.Bool:
		b := f.Bool()
		return strconv.FormatBool(b), nil
	case reflect.Ptr:
		return toString(f.Elem())
	default:
		return "", ErrTypeNotSupported
	}
}

// ToString try to convert interface{} to string
// ToString accepts data like: set of int, set of uint, set of float, bool, string  and pointer of these etc
func ToString(data interface{}) (string, error) {
	if data == nil {
		return "", ErrSourceDataNil
	}

	return toString(reflect.ValueOf(data))
}

//	Must ToString version
func MustToString(data interface{}) string {
	result, err := ToString(data)
	if err != nil {
		panic("converter: MustToString(" + fmt.Sprintf("%v", data) + "): " + err.Error())

	}
	return result
}

func toBool(value reflect.Value) (bool, error) {
	switch f := value; f.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16,reflect.Int32,reflect.Int64:
		return f.Int() != 0, nil
	case reflect.Uint,reflect.Uint8, reflect.Uint16,reflect.Uint32,reflect.Uint64:
		return f.Uint() != 0, nil
	case reflect.String:
		s := f.String()
		return strconv.ParseBool(s)
	case reflect.Bool:
		return f.Bool(), nil
	case reflect.Ptr:
		return toBool(f.Elem())
	default:
		return false, ErrTypeNotSupported
	}
}

// ToString try to convert interface{} to string
// ToString accepts data like: set of int, set of uint, bool, string and pointer of these etc
func ToBool(data interface{}) (bool, error) {
	if data == nil {
		return false, ErrSourceDataNil
	}

	return toBool(reflect.ValueOf(data))
}

//	Must ToBool version
func MustToBool(srcData interface{}) bool {
	result, err := ToBool(srcData)
	if err != nil {
		panic("converter: MustToBool(" + fmt.Sprintf("%v", srcData) + "): " + err.Error())
	}

	return result
}

func toFloat64(value reflect.Value) (float64, error) {
	switch f := value; f.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16,reflect.Int32,reflect.Int64:
		return float64(f.Int()), nil
	case reflect.Uint,reflect.Uint8, reflect.Uint16,reflect.Uint32,reflect.Uint64:
		return float64(f.Uint()), nil
	case reflect.String:
		s := f.String()
		return strconv.ParseFloat(s, 64)
	case reflect.Float32,reflect.Float64:
		return  f.Float(), nil
	case reflect.Bool:
		b := f.Bool()
		if b {
			return 1, nil
		} else {
			return 0, nil
		}
	case reflect.Ptr:
		return toFloat64(f.Elem())
	default:
		return 0, ErrTypeNotSupported
	}
}

// ToString try to convert interface{} to string
// ToString accepts data like: set of int, set of uint, bool, string and pointer of these etc
func ToFloat64(data interface{}) (float64, error) {
	if data == nil {
		return 0, ErrSourceDataNil
	}

	return toFloat64(reflect.ValueOf(data))
}

//	Must ToFloat64 version
func MustToFloat64(srcData interface{}) float64 {
	result, err := ToFloat64(srcData)
	if err != nil {
		panic("converter: MustToFloat64(" + fmt.Sprintf("%v", srcData) + "): " + err.Error())
	}

	return result
}

//	Convert interface to interface list, this function only works on slice
func ToInterfaceSlice(srcData interface{}) ([]interface{}, error) {
	if srcData == nil {
		return nil, ErrSourceDataNil
	}

	if reflect.TypeOf(srcData).Kind() != reflect.Slice {
		return nil, ErrTypeNotSupported
	}

	//	Create src and dst slice
	srcSlice := reflect.ValueOf(srcData)
	dstSlice := reflect.MakeSlice(reflect.TypeOf([]interface{}{}), 0, srcSlice.Len())

	//	Copy src to dst
	for idx := 0; idx < srcSlice.Len(); idx++ {
		dstSlice = reflect.Append(dstSlice, srcSlice.Index(idx))
	}

	return dstSlice.Interface().([]interface{}), nil
}

//	MustToInterfaceSlice version is like ToInterfaceSlice but panic if convert failed
func MustToInterfaceSlice(srcData interface{}) []interface{} {
	result, err := ToInterfaceSlice(srcData)
	if err != nil {
		panic("converter: MustToInterfaceSlice(" + fmt.Sprintf("%v", srcData) + "): " + err.Error())
	}

	return result
}
