package fib

import "testing"

var fibTests = []struct {
	a int
	expected int
}{
	{0, 0 },
	{1, 1 },
	{2, 1 },
	{3, 2 },
	{4, 3 },
	{5, 5 },
	{6, 8 },
	{7, 13 },
	{8, 21 },
	{9, 34 },
	{20, 6765 },
	{42, 267914296 },
	{46, 1836311903 },
	//{50, 12586269025 },
}

func TestFib(t *testing.T) {
	for _, ft := range fibTests {
		if v := Fib(ft.a); v != ft.expected {
			t.Errorf("Fib(%d) returned %d, expected %d", ft.a, v, ft.expected)
		}
	}
}

func TestFibTailRec(t *testing.T) {
	for _, ft := range fibTests {
		if v := FibTailRec(ft.a, 0, 1); v != ft.expected {
		//if v := FibTailRec(ft.a); v != ft.expected {
			t.Errorf("FibTailRec(%d) returned %d, expected %d", ft.a, v, ft.expected)
		}
	}
}

func BenchmarkFib(b *testing.B) {
	fn := Fib
	for i := 0; i < b.N; i++ {
		_ = fn(8)
	}
}

func benchmarkFib(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		Fib(i)
	}
}

func BenchmarkFib1(b *testing.B) { benchmarkFib(1, b) }
func BenchmarkFib2(b *testing.B) { benchmarkFib(2, b) }
func BenchmarkFib3(b *testing.B) { benchmarkFib(3, b) }
func BenchmarkFib4(b *testing.B) { benchmarkFib(4, b) }
func BenchmarkFib20(b *testing.B) { benchmarkFib(20, b) }
func BenchmarkFib40(b *testing.B) { benchmarkFib(42, b) }
