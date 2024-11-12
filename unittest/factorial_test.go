//
// Golang Workshop 2024
//

package main

import "testing"

func TestFactorialA(t *testing.T) {
	got := FactorialA(10)
	want := 3628800
	if got != want {
		t.Errorf("Factorial(10 = %d; want %d", got, want)
	}
}

func TestFactorialB(t *testing.T) {
	got := FactorialB(10)
	want := 3628800
	if got != want {
		t.Errorf("Factorial(10 = %d; want %d", got, want)
	}
}

func TestFactorialC(t *testing.T) {
	cache := make(map[int]int, 100)
	want := 3628800
	got := FactorialC(cache, 10)
	if got != want {
		t.Errorf("Factorial(10 = %d; want %d", got, want)
	}
	got = FactorialC(cache, 10)
	if got != want {
		t.Errorf("Factorial(10 = %d; want %d", got, want)
	}
}

/*func TestRaceCondition(t *testing.T) {
	got := RaceCondition(5)
	want := 10
	if got != want {
		t.Errorf("RaceCondition(5 = %d; want %d", got, want)
	}
}*/

var result int

func BenchmarkFactorialA(b *testing.B) {
	var r int
	for i := 0; i < b.N; i++ {
		r = FactorialA(100)
	}
	result = r
}

func BenchmarkFactorialB(b *testing.B) {
	var r int
	for i := 0; i < b.N; i++ {
		r = FactorialB(100)
	}
	result = r
}

func BenchmarkFactorialC(b *testing.B) {
	cache := make(map[int]int, 100)
	var r int
	for i := 0; i < b.N; i++ {
		r = FactorialC(cache, 100)
	}
	result = r
}
