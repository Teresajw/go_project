package types

type List interface {
	Add(idx int, val any) error
	Append(val any) error
	Delete(idx int) (any, error)
}
