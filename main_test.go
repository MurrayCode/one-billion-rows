package main

import (
	"reflect"
	"testing"
)

func TestReadData(t *testing.T) {
	res := readData("input_test.txt")
	exp := Stations{"Hamburg": Station{min: 0, max: 20, count: 5, sum: 55}, "Istanbul": Station{min: 1, max: 8, count: 4, sum: 14}}
	if !reflect.DeepEqual(exp, res) {
		t.Errorf("Expected %v but got %v", exp, res)
	}
}
