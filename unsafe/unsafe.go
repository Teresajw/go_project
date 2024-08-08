package unsafe

import (
	"errors"
	"reflect"
	"unsafe"
)

type FieldAccessor interface {
	Field(field string) (int, error)
	SetField(field string, val int) error
}

type UnsafeAccessor struct {
	fields     map[string]FieldMeta
	entityAddr unsafe.Pointer
}

type FieldMeta struct {
	typ    reflect.Type
	offset uintptr
}

func NewUnsafeAccessor(entity any) (*UnsafeAccessor, error) {
	if entity == nil {
		return nil, errors.New("invalid entity")
	}
	val := reflect.ValueOf(entity)
	typ := reflect.TypeOf(entity)

	val.UnsafeAddr()
	if typ.Kind() != reflect.Ptr || typ.Elem().Kind() != reflect.Struct {
		return nil, errors.New("invalid entity")
	}
	fields := make(map[string]FieldMeta, typ.Elem().NumField())
	elemType := typ.Elem()
	for i := 0; i < elemType.NumField(); i++ {
		fd := elemType.Field(i)
		fields[fd.Name] = FieldMeta{offset: fd.Offset}
	}
	return &UnsafeAccessor{entityAddr: val.UnsafePointer(), fields: fields}, nil
}

func (u *UnsafeAccessor) Field(field string) (int, error) {
	fdMeta, ok := u.fields[field]
	if !ok {
		return 0, errors.New("field not found")
	}
	prt := unsafe.Pointer(uintptr(u.entityAddr) + fdMeta.offset)
	if prt == nil {
		return 0, errors.New("invalid pointer")
	}
	res := *(*int)(prt)
	return res, nil
}
