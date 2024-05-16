package arrays

func Reduce[A, B any](collection []A, accumulator func(B, A) B, initial B) B {
	res := initial
	for _, x := range collection {
		res = accumulator(res, x)
	}

	return res
}

func Find[A any](items []A, predicate func(A) bool) (value A, found bool) {
	for _, v := range items {
		if predicate(v) {
			return v, true
		}
	}
	return
}
