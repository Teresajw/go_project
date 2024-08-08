package reflect

import (
	"errors"
	"fmt"
	"reflect"
)

func IterateFields(val any) {
	// 复杂逻辑
	res, err := iterateFields(val)

	//简单逻辑
	if err != nil {
		fmt.Println(err)
		return
	}
	for k, v := range res {
		fmt.Println(k, v)
	}
}

func iterateFields(val any) (map[string]any, error) {
	if val == nil {
		return nil, errors.New("val can't be nil")
	}
	// 获取类型
	tye := reflect.TypeOf(val)
	// 获取值
	value := reflect.ValueOf(val)

	// 循环判断是否是指针
	for tye.Kind() == reflect.Ptr {
		tye = tye.Elem()
		value = value.Elem()
	}

	// 获取字段数量
	NumField := tye.NumField()
	res := make(map[string]any, NumField)
	for i := 0; i < NumField; i++ {
		field := tye.Field(i)
		res[field.Name] = value.Field(i).Interface()
	}
	return res, nil
}

// 获取函数信息,并执行调用
// 考虑可能输入： nil,基本类型，内置类型
func IterateFuncs(val any) (map[string]*FuncInfo, error) {
	if val == nil {
		return nil, errors.New("输入 nil")
	}
	typ := reflect.TypeOf(val)
	//value := reflect.ValueOf(val)

	/*if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
		value = value.Elem()
	}*/

	if typ.Kind() != reflect.Struct && !(typ.Kind() == reflect.Ptr && typ.Elem().Kind() == reflect.Struct) { //判断如果typ是struct或者*struct
		return nil, errors.New("不支持的类型")
	}

	NumMethod := typ.NumMethod() //获取方法数量
	res := make(map[string]*FuncInfo, NumMethod)
	for i := 0; i < NumMethod; i++ {
		method := typ.Method(i)
		mt := method.Type

		numIn := mt.NumIn()
		in := make([]reflect.Type, 0, numIn)
		for j := 0; j < numIn; j++ {
			in = append(in, mt.In(j))
		}

		numOut := mt.NumOut()
		out := make([]reflect.Type, 0, numOut)
		for k := 0; k < numOut; k++ {
			out = append(out, mt.Out(k))
		}

		callRes := method.Func.Call([]reflect.Value{reflect.ValueOf(val)})
		result := make([]any, 0, len(callRes))
		for _, v := range callRes {
			result = append(result, v.Interface())
		}

		res[method.Name] = &FuncInfo{
			Name:   method.Name,
			In:     in,
			Out:    out,
			Result: result,
		}
	}

	return res, nil
}

type FuncInfo struct {
	Name string
	In   []reflect.Type
	Out  []reflect.Type

	// 反射调用得到的结果
	Result []any
}
