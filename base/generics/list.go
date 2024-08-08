package generics

type List[T any] interface {
	Add(idx int, t T)
	Append(t T)
}

func UseList() {
	var l List[int]
	l.Append(1)
}
