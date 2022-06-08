package linq

import "sort"

type Collection[T any] []T

type Predicate[T any] func(obj T) bool

type Action[T any] func(obj T)

type Selector[T1 any, T2 any] func(s T1) T2

type Comparer[T any] func(obj1, obj2 T) int

func (e Collection[T]) FirstOrDefault(p Predicate[T]) T {

	n := len(e)

	for i := 0; i < n; i++ {
		if p(e[i]) {
			return e[i]
		}
	}

	var result T
	return result
}

func (e Collection[T]) FirstIndex(p Predicate[T]) int {

	n := len(e)

	for i := 0; i < n; i++ {
		if p(e[i]) {
			return i
		}
	}

	return -1
}

func (e Collection[T]) LastIndex(p Predicate[T]) int {

	for i := len(e) - 1; i >= 0; i-- {
		if p(e[i]) {
			return i
		}
	}

	return -1
}

func (e Collection[T]) NextIndex(start int, p Predicate[T]) int {

	n := len(e)

	if start < 0 {
		start = 0
	}

	for i := start; i < n; i++ {
		if p(e[i]) {
			return i
		}
	}

	return -1
}

func (e Collection[T]) All(p Predicate[T]) bool {

	n := len(e)

	for i := 0; i < n; i++ {

		if !p(e[i]) {
			return false
		}
	}

	return true
}

func (e Collection[T]) Any(p Predicate[T]) bool {

	n := len(e)

	for i := 0; i < n; i++ {
		if p(e[i]) {
			return true
		}
	}

	return false
}

func (e Collection[T]) Where(p Predicate[T]) Collection[T] {

	n := len(e)
	result := make(Collection[T], 0, n/2)

	for i := 0; i < n; i++ {
		if p(e[i]) {
			result = append(result, e[i])
		}
	}

	return result
}

func (e Collection[T]) Take(n int) Collection[T] {

	if n >= len(e) {
		return e
	}

	if n >= 0 {
		return e[0:n]
	}

	i := len(e) + n

	if i <= 0 {
		return e
	}

	return e[i:]
}

func Select[T1 any, T2 any](e Collection[T1], s Selector[T1, T2]) Collection[T2] {

	result := make(Collection[T2], len(e))

	for i := 0; i < len(e); i++ {
		result[i] = s(e[i])
	}

	return result
}

func (e Collection[T]) ForEach(a Action[T]) {

	n := len(e)

	for i := 0; i < n; i++ {
		a(e[i])
	}
}

func (e Collection[T]) Order(c Comparer[T]) Collection[T] {

	sort.Slice(e, func(i, j int) bool { return c(e[i], e[j]) < 0 })
	return e
}

func (e Collection[T]) OrderDesc(c Comparer[T]) Collection[T] {

	sort.Slice(e, func(i, j int) bool { return c(e[i], e[j]) >= 0 })
	return e
}

func (e Collection[T]) RemoveRange(start, count int) Collection[T] {

	n := len(e)

	if start >= n || count <= 0 || start+count <= 0 {
		return e
	}

	if start < 0 {
		count += start
		start = 0
	}

	if start == 0 {

		if count >= n {
			count = n
		}

		return e[count:]
	}

	if start+count >= n {
		return e[0:start]
	}

	finish := start + count

	for i := start; i < finish; i++ {
		e[i] = e[i+count]
	}

	return e[0 : n-count]
}
