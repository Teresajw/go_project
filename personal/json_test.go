package personal

import (
	"encoding/json"
	"reflect"
	"testing"
)

type Door struct {
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
}

func Test_json(t *testing.T) {
	d := Door{10, 20}

	a, _ := json.Marshal(d)

	t.Log(string(a))

	dt := reflect.TypeOf(d)

	for i := 0; i < dt.NumField(); i++ {
		t.Log()
		t.Logf("字段：%s, 是否可导出：%t, 导出字段名称：%s", dt.Field(i).Name, dt.Field(i).IsExported(), dt.Field(i).Tag.Get("json"))
	}
}
