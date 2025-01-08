// Package gotool
//
// ----------------develop info----------------
//
//	@Author xunmuhuang@rastar.com
//	@DateTime 2025-1-7 17:42
//
// --------------------------------------------
package testing

import (
	"fmt"
	"github.com/calmu/gotool/normalbatch"
	"github.com/calmu/gotool/testing/common"
	"testing"
)

func TestByte(t *testing.T) {
	tests := []struct {
		len, repeat, filter int
	}{
		{len: 100, repeat: 0, filter: 0},
		{len: 120, repeat: 30, filter: 0},
		{len: 200, repeat: 3, filter: 10},
		{len: 300, repeat: 50, filter: 20},
	}

	for i, test := range tests {
		batch := normalbatch.NewSliceByteBatch(test.len)
		list := make([]string, 0, test.filter)
		realLen := 0
		for s, bytes := range common.NewBuild().BuildByte(test.len, test.repeat) {
			if len(list) < test.filter {
				list = append(list, s)
			}
			realLen = batch.Push(bytes, s)
		}
		fmt.Println(fmt.Sprintf("TestByte-%d:%+v,realLen=%d", i, test, realLen))
		batch.FilterMulti(list)
		if res := batch.GetClean(); len(res) != realLen-test.filter {
			t.Errorf("expect %d, but got %d", realLen-test.filter, len(res))
		}
	}
}

func TestByte2(t *testing.T) {
	tests := []struct {
		len, repeat, filter int
	}{
		{len: 100, repeat: 0, filter: 0},
		{len: 120, repeat: 30, filter: 0},
		{len: 200, repeat: 3, filter: 10},
		{len: 300, repeat: 50, filter: 20},
	}

	for i, test := range tests {
		batch := normalbatch.NewSliceByteBatch(test.len)
		list := make([]string, 0, test.filter)
		realLen := 0
		for s, bytes := range common.NewBuild().BuildByte(test.len, test.repeat) {
			if len(list) < test.filter {
				list = append(list, s)
			}
			realLen = batch.Push(bytes, s)
		}
		fmt.Println(fmt.Sprintf("TestByte2-%d:%+v,realLen=%d", i, test, realLen))
		for _, s := range list {
			batch.Filter(s)
		}
		if res := batch.GetClean(); len(res) != realLen-test.filter {
			t.Errorf("expect %d, but got %d", realLen-test.filter, len(res))
		}
	}
}

func TestString(t *testing.T) {
	tests := []struct {
		len, repeat, filter int
	}{
		{len: 100, repeat: 0, filter: 0},
		{len: 120, repeat: 30, filter: 0},
		{len: 200, repeat: 3, filter: 10},
		{len: 300, repeat: 50, filter: 20},
	}

	for i, test := range tests {
		batch := normalbatch.NewSliceStringBatch(test.len)
		list := make([]string, 0, test.filter)
		realLen := 0
		for s, bytes := range common.NewBuild().BuildString(test.len, test.repeat) {
			if len(list) < test.filter {
				list = append(list, s)
			}
			realLen = batch.Push(bytes, s)
		}
		fmt.Println(fmt.Sprintf("TestString-%d:%+v,realLen=%d", i, test, realLen))
		batch.FilterMulti(list)
		if res := batch.GetClean(); len(res) != realLen-test.filter {
			t.Errorf("expect %d, but got %d", realLen-test.filter, len(res))
		}
	}
}

func TestString2(t *testing.T) {
	tests := []struct {
		len, repeat, filter int
	}{
		{len: 100, repeat: 0, filter: 0},
		{len: 120, repeat: 30, filter: 0},
		{len: 200, repeat: 3, filter: 10},
		{len: 300, repeat: 50, filter: 20},
	}

	for i, test := range tests {
		batch := normalbatch.NewSliceStringBatch(test.len)
		list := make([]string, 0, test.filter)
		realLen := 0
		for s, bytes := range common.NewBuild().BuildString(test.len, test.repeat) {
			if len(list) < test.filter {
				list = append(list, s)
			}
			realLen = batch.Push(bytes, s)
		}
		fmt.Println(fmt.Sprintf("TestString2-%d:%+v,realLen=%d", i, test, realLen))
		for _, s := range list {
			batch.Filter(s)
		}
		if res := batch.GetClean(); len(res) != realLen-test.filter {
			t.Errorf("expect %d, but got %d", realLen-test.filter, len(res))
		}
	}
}

func TestInterface(t *testing.T) {
	tests := []struct {
		len, repeat, filter int
	}{
		{len: 100, repeat: 0, filter: 0},
		{len: 120, repeat: 30, filter: 0},
		{len: 200, repeat: 3, filter: 10},
		{len: 300, repeat: 50, filter: 20},
	}

	for i, test := range tests {
		batch := normalbatch.NewSliceInterfaceBatch(test.len)
		list := make([]string, 0, test.filter)
		realLen := 0
		for s, bytes := range common.NewBuild().BuildInterface(test.len, test.repeat) {
			if len(list) < test.filter {
				list = append(list, s)
			}
			realLen = batch.Push(bytes, s)
		}
		fmt.Println(fmt.Sprintf("TestInterface-%d:%+v,realLen=%d", i, test, realLen))
		batch.FilterMulti(list)
		if res := batch.GetClean(); len(res) != realLen-test.filter {
			t.Errorf("expect %d, but got %d", realLen-test.filter, len(res))
		}
	}
}

func TestInterface2(t *testing.T) {
	tests := []struct {
		len, repeat, filter int
	}{
		{len: 100, repeat: 0, filter: 0},
		{len: 120, repeat: 30, filter: 0},
		{len: 200, repeat: 3, filter: 10},
		{len: 300, repeat: 50, filter: 20},
	}

	for i, test := range tests {
		batch := normalbatch.NewSliceInterfaceBatch(test.len)
		list := make([]string, 0, test.filter)
		realLen := 0
		for s, bytes := range common.NewBuild().BuildInterface(test.len, test.repeat) {
			if len(list) < test.filter {
				list = append(list, s)
			}
			realLen = batch.Push(bytes, s)
		}
		fmt.Println(fmt.Sprintf("TestInterface2-%d:%+v,realLen=%d", i, test, realLen))
		for _, s := range list {
			batch.Filter(s)
		}
		if res := batch.GetClean(); len(res) != realLen-test.filter {
			t.Errorf("expect %d, but got %d", realLen-test.filter, len(res))
		}
	}
}
