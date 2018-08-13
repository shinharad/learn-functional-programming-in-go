package fib

func Memoize(mf Memoized) Memoized {
	cache := make(map[int]int)
	return func(key int) int {
		if val, found := cache[key]; found {
			return val
		}
		temp := mf(key)
		cache[key] = temp
		return temp
	}
}

type Memoized func(int) int
var fibMem Memoized
func FibMemoized(n int) int {
	fibMem = Memoize(func(x int) int {
		if x == 0 {
			return 0
		} else if x <= 2 {
			return 1
		} else {
			return fibMem(x - 2) + fibMem(x - 1)
		}
	})
	return fibMem(n)
}


