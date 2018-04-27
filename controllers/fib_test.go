package controllers

import (
	"reflect"
	"testing"
)

var testData = []struct {
	num      int
	expected int
}{
	{-1, 0},
	{0, 0},
	{3, 3},
	{4, 5},
	{20, 10946},
	{42, 433494437},
}

func Examplefibonacci(t *testing.T) {
	for _, f := range testData {
		if val := fibonacci(f.num); val != f.expected {
			t.Errorf("fibonacci(%d) returned %d, expected %d", f.num, val, f.expected)
		}
	}
}

func Benchmarkfibonacci(b *testing.B) {
	f := fibonacci
	for i := 0; i < b.N; i++ {
		_ = f(10)
	}
}

var testArrayData = []struct {
	num      int
	expected []int
}{
	{5, []int{0, 1, 1, 2, 3}},
	{0, []int{0}},
	{-1, []int{0}},
	{9, []int{0, 1, 1, 2, 3, 5, 8, 13, 21}},
}

func ExampleGetFibSequence(t *testing.T) {
	for j, f := range testArrayData {
		for i := range GetFibSequence(f.num) {
			if ok := reflect.DeepEqual(i, f.expected[j]); !ok {
				t.Errorf("GetFibSequence(%d) returned %v, expected %v", f.num, i, f.expected[j])
			}
		}
	}
}

func BenchmarkGetFibSequence(b *testing.B) {
	f := GetFibSequence
	for i := 0; i < b.N; i++ {
		_ = f(10)
	}
}
