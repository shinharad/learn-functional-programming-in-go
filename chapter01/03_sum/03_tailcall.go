package fib

func SumTailCall(vs []int, acc int) int {
	if len(vs) == 0 {
		return acc
	}
	return SumTailCall(vs[1:], vs[0] + acc)
}
