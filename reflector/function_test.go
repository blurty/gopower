package reflector

import "testing"

func TestGetFunctionName(t *testing.T) {
	testCases := []struct{
		Input interface{}
		Output string
	} {
		{
			Input:TestGetFunctionName,
			Output:"github.com/blurty/gopower/reflector.TestGetFunctionName",
		},
		{
			Input:nil,
			Output:"",
		},
		{
			Input:[]int{1},
			Output:"",
		},
	}
	for _, testCase := range testCases {
		got := GetFunctionName(testCase.Input)
		if got != testCase.Output {
			t.Errorf("GetFunctionName(%v) got %v, want%v", testCase.Input, got, testCase.Output)
		}
	}
}
