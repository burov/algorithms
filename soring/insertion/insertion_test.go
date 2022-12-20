package insertion

import (
	"reflect"
	"testing"
)

func TestAscending_returnsProperResults(t *testing.T) {
	tests := []struct {
		input          []int
		expectedResult []int
	}{
		{
			input:          []int{},
			expectedResult: []int{},
		},
		{
			input:          []int{5, 4, 3, 2, 1},
			expectedResult: []int{1, 2, 3, 4, 5},
		},
		{
			input:          []int{5, 2, 4, 6, 1, 3},
			expectedResult: []int{1, 2, 3, 4, 5, 6},
		},
	}

	for i, tt := range tests {
		Ascending(tt.input)

		result := tt.input
		if !reflect.DeepEqual(tt.expectedResult, result) {
			t.Errorf("Test %d, unexpected result, expected %+v, got %+v", i, tt.expectedResult, result)
			t.Fail()
		}

	}

}

func TestDescending_returnsProperResults(t *testing.T) {
	tests := []struct {
		input          []int
		expectedResult []int
	}{
		{
			input:          []int{},
			expectedResult: []int{},
		},
		{
			input:          []int{1, 2, 3, 4, 5},
			expectedResult: []int{5, 4, 3, 2, 1},
		},
		{
			input:          []int{5, 2, 4, 6, 1, 3},
			expectedResult: []int{6, 5, 4, 3, 2, 1},
		},
	}

	for i, tt := range tests {
		Descending(tt.input)

		result := tt.input
		if !reflect.DeepEqual(tt.expectedResult, result) {
			t.Errorf("Test %d, unexpected result, expected %+v, got %+v", i, tt.expectedResult, result)
			t.Fail()
		}

	}

}
