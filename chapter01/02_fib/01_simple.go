package fib

func Fib(x int) int {
	if x == 0 {
		return 0
	} else if x <= 2 {
		return 1
	} else {
		return Fib(x - 2) + Fib(x - 1)
	}
}

func FibTailRec(x, prev, cur int) int {
	if x == 0 {
		return prev
	} else {
		return FibTailRec(x - 1, cur, prev + cur)
	}
}

//func FibTailRec(x int) int {
//	type FibTailRecLoop = func(int, int, int) int
//	var loop FibTailRecLoop
//	loop = func(x, prev, cur int) int {
//		if x == 0 {
//			return prev
//		} else {
//			return loop(x - 1, cur, prev + cur)
//		}
//	}
//	return loop(x, 0, 1)
//}
