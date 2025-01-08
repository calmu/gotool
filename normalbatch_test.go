// Package gotool
//
// ----------------develop info----------------
//
//	@Author xunmuhuang@rastar.com
//	@DateTime 2025-1-7 17:42
//
// --------------------------------------------
package gotool

import (
	"fmt"
	"github.com/calmu/gotool/normalbatch"
	"math/rand"
	"testing"
	"time"
)

func buildByte(len, repeat int) map[string][]byte {
	m := make(map[string][]byte, len)
	realLen := len - repeat
	for i := 1; i <= realLen; i++ {
		if repeat > 0 {
			m[fmt.Sprintf("a%d", repeat)] = []byte(fmt.Sprintf(`{"id":"a%d","name":"a%db%d"}`, repeat, repeat, repeat))
		}
		m[fmt.Sprintf("a%d", i)] = []byte(fmt.Sprintf(`{"id":"a%d","name":"a%db%d"}`, i, i, i))
		repeat--
	}
	return m
}

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
		for s, bytes := range buildByte(test.len, test.repeat) {
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
		/*{len: 100, repeat: 0, filter: 0},
		{len: 120, repeat: 30, filter: 0},
		{len: 200, repeat: 3, filter: 10},
		{len: 300, repeat: 50, filter: 20},*/
		{len: 10, repeat: 5, filter: 2},
	}

	for i, test := range tests {
		batch := normalbatch.NewSliceByteBatch(test.len)
		list := make([]string, 0, test.filter)
		realLen := 0
		for s, bytes := range buildByte(test.len, test.repeat) {
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

func buildString(len, repeat int) map[string]string {
	m := make(map[string]string, len)
	realLen := len - repeat
	for i := 1; i <= realLen; i++ {
		if repeat > 0 {
			m[fmt.Sprintf("a%d", repeat)] = fmt.Sprintf(`{"id":"a%d","name":"a%db%d"}`, repeat, repeat, repeat)
		}
		m[fmt.Sprintf("a%d", i)] = fmt.Sprintf(`{"id":"a%d","name":"a%db%d"}`, i, i, i)
		repeat--
	}
	return m
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
		for s, bytes := range buildString(test.len, test.repeat) {
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
		for s, bytes := range buildString(test.len, test.repeat) {
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

func buildInterface(len, repeat int) map[string]interface{} {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	index := r.Intn(3)
	m := make(map[string]interface{}, len)
	switch index {
	case 0:
		for k, v := range buildByte(len, repeat) {
			m[k] = v
		}
		return m
	case 1:
		for k, v := range buildString(len, repeat) {
			m[k] = v
		}
		return m
	}
	realLen := len - repeat
	for i := 1; i <= realLen; i++ {
		if repeat > 0 {
			m[fmt.Sprintf("a%d", repeat)] = map[string]interface{}{"id": fmt.Sprintf("a%d", repeat), "name": fmt.Sprintf("a%db%d", repeat, repeat)}
		}
		m[fmt.Sprintf("a%d", i)] = map[string]interface{}{"id": fmt.Sprintf("a%d", i), "name": fmt.Sprintf("a%db%d", i, i)}
		repeat--
	}
	return m
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
		for s, bytes := range buildInterface(test.len, test.repeat) {
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
		for s, bytes := range buildInterface(test.len, test.repeat) {
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
