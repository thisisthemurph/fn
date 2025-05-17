package fn

// Contains returns true if the value is in the slice, false otherwise.
func Contains[T comparable](s []T, value T) bool {
	for _, v := range s {
		if v == value {
			return true
		}
	}
	return false
}

func First[T any](in []T, compareFn func(T) bool) (T, bool) {
	var zero T
	for _, v := range in {
		if compareFn(v) {
			return v, true
		}
	}
	return zero, false
}

// Map applies the function f to each element in the slice and returns a new slice with the results.
// If the input slice is nil, Map returns nil.
func Map[TIn any, TOut any](in []TIn, f func(TIn) TOut) []TOut {
	if in == nil {
		return nil
	}

	out := make([]TOut, len(in))
	for i, v := range in {
		out[i] = f(v)
	}
	return out
}
