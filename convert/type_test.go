package convert

import (
	"reflect"
	"testing"
)

var (
	int8Val = int8(1)
	int16Val = int16(1)
	int32Val = int32(1)
	int64Val = int64(1)
	intVal = int(1)
	uint8Val = uint8(1)
	uint16Val = uint16(1)
	uint32Val = uint32(1)
	uint64Val = uint64(1)
	stringVal = "1"
	boolVal = true
	float32Val = float32(1)
	float64val = float64(1)

	stringTrueVal = "true"
)

func TestToInt(t *testing.T) {
	testCases := []struct{
		Input interface{}
		Output int
		WantErr bool	// 是否期待返回错误
	} {
		{
			Input:nil,
			WantErr:true,
		},
		{
			Input:[]int{intVal},
			WantErr:true,
		},
		{
			Input:map[int]interface{}{intVal:true},
			WantErr:true,
		},
		{
			Input:intVal,
			Output:1,
		},
		{
			Input:int8Val,
			Output:1,
		},
		{
			Input:int16Val,
			Output:1,
		},
		{
			Input:int32Val,
			Output:1,
		},
		{
			Input:int64Val,
			Output:1,
		},
		{
			Input:uint8Val,
			Output:1,
		},
		{
			Input:uint16Val,
			Output:1,
		},
		{
			Input:uint32Val,
			Output:1,
		},
		{
			Input:uint64Val,
			Output:1,
		},
		{
			Input:&intVal,
			Output:1,
		},
		{
			Input:&int8Val,
			Output:1,
		},
		{
			Input:&int16Val,
			Output:1,
		},
		{
			Input:&int32Val,
			Output:1,
		},
		{
			Input:&int64Val,
			Output:1,
		},
		{
			Input:&uint8Val,
			Output:1,
		},
		{
			Input:&uint16Val,
			Output:1,
		},
		{
			Input:&uint32Val,
			Output:1,
		},
		{
			Input:&uint64Val,
			Output:1,
		},
		{
			Input:boolVal,
			Output:1,
		},
		{
			Input:&boolVal,
			Output:1,
		},
		{
			Input:stringVal,
			Output:1,
		},
		{
			Input:&stringVal,
			Output:1,
		},
		{
			Input:float32Val,
			Output:1,
		},
		{
			Input:&float32Val,
			Output:1,
		},
		{
			Input:float64val,
			Output:1,
		},
		{
			Input:&float64val,
			Output:1,
		},
	}
	for _, testCase := range testCases {
		gotVal, err := ToInt(testCase.Input)
		if err != nil && !testCase.WantErr{
			t.Errorf("ToInt(%+v) got error:%v", reflect.TypeOf(testCase.Input), err)
			continue
		} else if err == nil && testCase.WantErr {
			t.Errorf("ToInt(%+v) want error, got no error", reflect.TypeOf(testCase.Input))
			continue
		}
		if gotVal != testCase.Output {
			t.Fatalf("ToInt(%+v) want %d, got %d", reflect.TypeOf(testCase.Input), testCase.Output, gotVal)
			continue
		}
		t.Logf("ToInt(%+v) got correct result %d,%v", reflect.TypeOf(testCase.Input), gotVal, err)
	}
}

func TestToString(t *testing.T) {
	testCases := []struct{
		Input interface{}
		Output string
		WantErr bool
	} {
		{
			Input:nil,
			WantErr:true,
		},
		{
			Input:[]string{"a"},
			WantErr:true,
		},
		{
			Input:map[string]bool {"a":true},
			WantErr:true,
		},
		{
			Input:intVal,
			Output:stringVal,
		},
		{
			Input:&intVal,
			Output:stringVal,
		},
		{
			Input:uint8Val,
			Output:stringVal,
		},
		{
			Input:&uint8Val,
			Output:stringVal,
		},
		{
			Input:stringVal,
			Output:stringVal,
		},
		{
			Input:float64val,
			Output:stringVal,
		},
		{
			Input:&float64val,
			Output:stringVal,
		},
		{
			Input:boolVal,
			Output:"true",
		},
		{
			Input:&boolVal,
			Output:"true",
		},
	}
	for _, testCase := range testCases {
		gotVal, err := ToString(testCase.Input)
		if err != nil && !testCase.WantErr{
			t.Errorf("ToString(%+v) got error:%v", reflect.TypeOf(testCase.Input), err)
			continue
		} else if err == nil && testCase.WantErr {
			t.Errorf("ToString(%+v) want error, got no error", reflect.TypeOf(testCase.Input))
			continue
		}
		if gotVal != testCase.Output {
			t.Fatalf("ToString(%+v) want %v, got %v", reflect.TypeOf(testCase.Input), testCase.Output, gotVal)
			continue
		}
		t.Logf("ToString(%+v) got correct result %v,%v", reflect.TypeOf(testCase.Input), gotVal, err)
	}
}

func TestToBool(t *testing.T) {
	testCases := []struct{
		Input interface{}
		Output bool
		WantErr bool
	} {
		{
			Input:nil,
			WantErr:true,
		},
		{
			Input:[]bool{true},
			WantErr:true,
		},
		{
			Input:map[bool]bool {true:true},
			WantErr:true,
		},
		{
			Input:int8Val,
			Output:true,
		},
		{
			Input:&int8Val,
			Output:true,
		},
		{
			Input:uint8Val,
			Output:true,
		},
		{
			Input:&uint8Val,
			Output:true,
		},
		{
			Input:stringTrueVal,
			Output:true,
		},
		{
			Input:&stringTrueVal,
			Output:true,
		},
		{
			Input:float64val,
			WantErr:true,
		},
		{
			Input:&float64val,
			WantErr:true,
		},
		{
			Input:boolVal,
			Output:true,
		},
		{
			Input:&boolVal,
			Output:true,
		},
	}
	for _, testCase := range testCases {
		gotVal, err := ToBool(testCase.Input)
		if err != nil && !testCase.WantErr{
			t.Errorf("ToBool(%+v) got error:%v", reflect.TypeOf(testCase.Input), err)
			continue
		} else if err == nil && testCase.WantErr {
			t.Errorf("ToBool(%+v) want error, got no error", reflect.TypeOf(testCase.Input))
			continue
		}
		if gotVal != testCase.Output {
			t.Fatalf("ToBool(%+v) want %v, got %v", reflect.TypeOf(testCase.Input), testCase.Output, gotVal)
			continue
		}
		t.Logf("ToBool(%+v) got correct result %v,%v", reflect.TypeOf(testCase.Input), gotVal, err)
	}
}


func TestToFloat64(t *testing.T) {
	testCases := []struct{
		Input interface{}
		Output float64
		WantErr bool	// 是否期待返回错误
	} {
		{
			Input:nil,
			WantErr:true,
		},
		{
			Input:[]int{intVal},
			WantErr:true,
		},
		{
			Input:map[int]interface{}{intVal:true},
			WantErr:true,
		},
		{
			Input:intVal,
			Output:1,
		},
		{
			Input:int8Val,
			Output:1,
		},
		{
			Input:int16Val,
			Output:1,
		},
		{
			Input:int32Val,
			Output:1,
		},
		{
			Input:int64Val,
			Output:1,
		},
		{
			Input:uint8Val,
			Output:1,
		},
		{
			Input:uint16Val,
			Output:1,
		},
		{
			Input:uint32Val,
			Output:1,
		},
		{
			Input:uint64Val,
			Output:1,
		},
		{
			Input:&intVal,
			Output:1,
		},
		{
			Input:&int8Val,
			Output:1,
		},
		{
			Input:&int16Val,
			Output:1,
		},
		{
			Input:&int32Val,
			Output:1,
		},
		{
			Input:&int64Val,
			Output:1,
		},
		{
			Input:&uint8Val,
			Output:1,
		},
		{
			Input:&uint16Val,
			Output:1,
		},
		{
			Input:&uint32Val,
			Output:1,
		},
		{
			Input:&uint64Val,
			Output:1,
		},
		{
			Input:boolVal,
			Output:1,
		},
		{
			Input:&boolVal,
			Output:1,
		},
		{
			Input:stringVal,
			Output:1,
		},
		{
			Input:&stringVal,
			Output:1,
		},
		{
			Input:float32Val,
			Output:1,
		},
		{
			Input:&float32Val,
			Output:1,
		},
		{
			Input:float64val,
			Output:1,
		},
		{
			Input:&float64val,
			Output:1,
		},
	}
	for _, testCase := range testCases {
		gotVal, err := ToFloat64(testCase.Input)
		if err != nil && !testCase.WantErr{
			t.Errorf("ToFloat64(%+v) got error:%v", reflect.TypeOf(testCase.Input), err)
			continue
		} else if err == nil && testCase.WantErr {
			t.Errorf("ToFloat64(%+v) want error, got no error", reflect.TypeOf(testCase.Input))
			continue
		}
		if gotVal != testCase.Output {
			t.Fatalf("ToFloat64(%+v) want %v, got %v", reflect.TypeOf(testCase.Input), testCase.Output, gotVal)
			continue
		}
		t.Logf("ToFloat64(%+v) got correct result %v,%v", reflect.TypeOf(testCase.Input), gotVal, err)
	}
}