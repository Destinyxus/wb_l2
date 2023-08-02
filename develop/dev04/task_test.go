package main

import (
	"reflect"
	"testing"
)

func TestSearchAn(t *testing.T) {
	input := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"}
	expectedOutput := map[string][]string{
		"столик": []string{"листок", "слиток", "столик"},
		"тяпка":  []string{"пятак", "пятка", "тяпка"},
	}

	result := searchAn(input)

	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, result)
	}
}
