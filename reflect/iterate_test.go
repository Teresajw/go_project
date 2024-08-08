package reflect

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestIterateFields(t *testing.T) {
	u1 := &User{Name: "Wen"}
	u2 := &u1

	tests := []struct {
		//名字
		name string
		//输入部分
		val any
		//输出部分
		wantRes map[string]any
		wantErr error
	}{
		// TODO: Add test cases.
		{
			name:    "nil",
			val:     nil,
			wantErr: errors.New("val can't be nil"),
		},
		{
			name:    "user",
			val:     User{Name: "Tom"},
			wantErr: nil,
			wantRes: map[string]any{
				"Name": "Tom",
			},
		},
		{
			// 指针类型
			name:    "pointer",
			val:     &User{Name: "Jerry"},
			wantErr: nil,
			wantRes: map[string]any{
				"Name": "Jerry",
			},
		},
		{
			// 多重指针类型
			name:    "multiple pointer",
			val:     u2,
			wantErr: nil,
			wantRes: map[string]any{
				"Name": "Wen",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := iterateFields(tt.val)
			assert.Equal(t, tt.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tt.wantRes, res)
		})
	}
}

type User struct {
	Name string
}

func TestIterateFuncs(t *testing.T) {
	type args struct {
		val any
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]*FuncInfo
		wantErr error
	}{
		// TODO: Add test cases.
		{
			name:    "nil",
			wantErr: errors.New("输入 nil"),
		},
		{
			name:    "basic types",
			args:    args{val: 1},
			wantErr: errors.New("不支持的类型"),
		},
		{
			name: "Order 方法",
			args: args{
				val: Order{
					buyer: 18,
				},
			},
			wantErr: nil,
			want: map[string]*FuncInfo{
				"GetBuyer": {
					Name:   "GetBuyer",
					In:     []reflect.Type{reflect.TypeOf(Order{})},
					Out:    []reflect.Type{reflect.TypeOf(int64(0))},
					Result: []any{int64(18)},
				},
			},
		},
		{
			name: "Order 指针方法",
			args: args{
				val: &OrderV1{
					buyer: 18,
				},
			},
			wantErr: nil,
			want: map[string]*FuncInfo{
				"GetBuyer": {
					Name:   "GetBuyer",
					In:     []reflect.Type{reflect.TypeOf(&OrderV1{})},
					Out:    []reflect.Type{reflect.TypeOf(int64(0))},
					Result: []any{int64(18)},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IterateFuncs(tt.args.val)
			assert.Equal(t, tt.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

type Order struct {
	buyer int64
}

func (o Order) GetBuyer() int64 {
	return o.buyer
}

type OrderV1 struct {
	buyer int64
}

func (o *OrderV1) GetBuyer() int64 {
	return o.buyer
}

type MyInterFace interface {
	Abc()
}

// var _ MyInterFace = abcImpl{}
var _ MyInterFace = &abcImpl{}

type abcImpl struct {
}

func (a *abcImpl) Abc() {
	//TODO implement me
	panic("implement me")
}

type MyService struct {
	GetById func()
}

func Proxy() {
	myService := &MyService{}
	myService.GetById = func() {
		// 发起RPC
		// 解析响应
	}
}
