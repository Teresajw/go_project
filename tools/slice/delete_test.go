package slice

import (
	"github.com/Teresajw/go_project/tools/internal/errs"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDelete(t *testing.T) {
	testCases := []struct {
		name      string
		slice     []int
		delIndex  int
		wantSlice []int
		wantVal   int
		wantErr   error
	}{
		{
			name:      "delete 1",
			slice:     []int{1, 2, 3},
			delIndex:  1,
			wantSlice: []int{1, 3},
			wantVal:   2,
			wantErr:   nil,
		},
		{
			name:      "index middle",
			slice:     []int{123, 124, 125},
			delIndex:  1,
			wantSlice: []int{123, 125},
			wantVal:   124,
			wantErr:   nil,
		},
		{
			name:     "index out of range",
			slice:    []int{123, 100},
			delIndex: 5,
			wantErr:  errs.NewErrIndexOutOfRange(2, 5),
		},
		{
			name:     "index less than 0",
			slice:    []int{123, 100},
			delIndex: -1,
			wantErr:  errs.NewErrIndexOutOfRange(2, -1),
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			slice, val, err := Delete(tt.slice, tt.delIndex)
			assert.Equal(t, tt.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tt.wantSlice, slice)
			assert.Equal(t, tt.wantVal, val)
		})
	}
}

func TestFilterDelete(t *testing.T) {
	testCases := []struct {
		name            string
		src             []int
		deleteCondition func(idx int, src int) bool

		wantRes []int
	}{
		{
			name: "空切片",
			src:  []int{},
			deleteCondition: func(idx int, src int) bool {
				return false
			},

			wantRes: []int{},
		},
		{
			name: "不删除元素",
			src:  []int{0, 1, 2, 3, 4, 5, 6, 7},
			deleteCondition: func(idx int, src int) bool {
				return false
			},

			wantRes: []int{0, 1, 2, 3, 4, 5, 6, 7},
		},
		{
			name: "删除首位元素",
			src:  []int{0, 1, 2, 3, 4, 5, 6},
			deleteCondition: func(idx int, src int) bool {
				return idx == 0
			},

			wantRes: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name: "删除前面两个元素",
			src:  []int{0, 1, 2, 3, 4, 5, 6, 7},
			deleteCondition: func(idx int, src int) bool {
				return idx == 0 || idx == 1
			},

			wantRes: []int{2, 3, 4, 5, 6, 7},
		},
		{
			name: "删除中间单个元素",
			src:  []int{0, 1, 2, 3, 4, 5, 6, 7},
			deleteCondition: func(idx int, src int) bool {
				return idx == 3
			},

			wantRes: []int{0, 1, 2, 4, 5, 6, 7},
		},
		{
			name: "删除中间多个不连续元素",
			src:  []int{0, 1, 2, 3, 4, 5, 6, 7},
			deleteCondition: func(idx int, src int) bool {
				return idx == 2 || idx == 4
			},

			wantRes: []int{0, 1, 3, 5, 6, 7},
		},
		{
			name: "删除中间多个连续元素",
			src:  []int{0, 1, 2, 3, 4, 5, 6, 7},
			deleteCondition: func(idx int, src int) bool {
				return idx == 3 || idx == 4
			},

			wantRes: []int{0, 1, 2, 5, 6, 7},
		},
		{
			name: "删除中间多个元素，第一部分为一个元素，第二部分为连续元素",
			src:  []int{0, 1, 2, 3, 4, 5, 6, 7},
			deleteCondition: func(idx int, src int) bool {
				return idx == 2 || idx == 4 || idx == 5
			},

			wantRes: []int{0, 1, 3, 6, 7},
		},
		{
			name: "删除中间多个元素，第一部分为连续元素，第二部分为一个元素",
			src:  []int{0, 1, 2, 3, 4, 5, 6, 7},
			deleteCondition: func(idx int, src int) bool {
				return idx == 2 || idx == 3 || idx == 5
			},

			wantRes: []int{0, 1, 4, 6, 7},
		},
		{
			name: "删除后面两个元素",
			src:  []int{0, 1, 2, 3, 4, 5, 6, 7},
			deleteCondition: func(idx int, src int) bool {
				return idx == 6 || idx == 7
			},

			wantRes: []int{0, 1, 2, 3, 4, 5},
		},
		{
			name: "删除末尾元素",
			src:  []int{0, 1, 2, 3, 4, 5, 6, 7},
			deleteCondition: func(idx int, src int) bool {
				return idx == 7
			},

			wantRes: []int{0, 1, 2, 3, 4, 5, 6},
		},
		{
			name: "删除所有元素",
			src:  []int{0, 1, 2, 3, 4, 5, 6, 7},
			deleteCondition: func(idx int, src int) bool {
				return true
			},

			wantRes: []int{},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Log(tc.name)
			t.Log(tc.src)
			res := FilterDelete(tc.src, tc.deleteCondition)
			t.Log(res)
			assert.Equal(t, tc.wantRes, res)
		})
	}
}
