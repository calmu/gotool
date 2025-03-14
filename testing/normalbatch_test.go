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
	"encoding/json"
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
		var f = func(start int) [][]byte {
			list := make([]string, 0, test.filter)
			realLen := 0
			for s, bytes := range common.NewBuild().BuildByte(test.len, test.repeat, start) {
				if len(list) < test.filter {
					list = append(list, s)
				}
				realLen = batch.Push(bytes, s)
			}
			fmt.Println(fmt.Sprintf("TestByte-%d:%+v,realLen=%d,start=%d", i, test, realLen, start))
			batch.FilterMulti(list)
			res := batch.GetClean()
			if len(res) != realLen-test.filter {
				t.Errorf("expect %d, but got %d", realLen-test.filter, len(res))
			}
			return res
		}
		res1 := f(0)
		res2 := f(test.len)

		uuidMap := make(map[string]struct{}, len(res1))
		for _, msg := range res1 {
			var tmpMap map[string]interface{}
			_ = json.Unmarshal(msg, &tmpMap)
			uuidMap[tmpMap["id"].(string)] = struct{}{}
		}
		var tmpLen int
		for _, msg := range res2 {
			var tmpMap map[string]interface{}
			_ = json.Unmarshal(msg, &tmpMap)
			if _, ok := uuidMap[tmpMap["id"].(string)]; ok {
				tmpLen++
			}
		}
		if tmpLen > 0 {
			t.Errorf("again expect 0, but got %d", tmpLen)
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
		var f = func(start int) [][]byte {
			list := make([]string, 0, test.filter)
			realLen := 0
			for s, bytes := range common.NewBuild().BuildByte(test.len, test.repeat, start) {
				if len(list) < test.filter {
					list = append(list, s)
				}
				realLen = batch.Push(bytes, s)
			}
			fmt.Println(fmt.Sprintf("TestByte2-%d:%+v,realLen=%d,start=%d", i, test, realLen, start))
			for _, s := range list {
				batch.Filter(s)
			}
			res := batch.GetClean()
			if len(res) != realLen-test.filter {
				t.Errorf("expect %d, but got %d", realLen-test.filter, len(res))
			}
			return res
		}
		res1 := f(0)
		res2 := f(test.len)

		uuidMap := make(map[string]struct{}, len(res1))
		for _, msg := range res1 {
			var tmpMap map[string]interface{}
			_ = json.Unmarshal(msg, &tmpMap)
			uuidMap[tmpMap["id"].(string)] = struct{}{}
		}
		var tmpLen int
		for _, msg := range res2 {
			var tmpMap map[string]interface{}
			_ = json.Unmarshal(msg, &tmpMap)
			if _, ok := uuidMap[tmpMap["id"].(string)]; ok {
				tmpLen++
			}
		}
		if tmpLen > 0 {
			t.Errorf("again expect 0, but got %d", tmpLen)
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
		var f = func(start int) []string {
			list := make([]string, 0, test.filter)
			realLen := 0
			for s, bytes := range common.NewBuild().BuildString(test.len, test.repeat, start) {
				if len(list) < test.filter {
					list = append(list, s)
				}
				realLen = batch.Push(bytes, s)
			}
			fmt.Println(fmt.Sprintf("TestString-%d:%+v,realLen=%d,start=%d", i, test, realLen, start))
			batch.FilterMulti(list)
			res := batch.GetClean()
			if len(res) != realLen-test.filter {
				t.Errorf("expect %d, but got %d", realLen-test.filter, len(res))
			}
			return res
		}
		res1 := f(0)
		res2 := f(test.len)

		uuidMap := make(map[string]struct{}, len(res1))
		for _, msg := range res1 {
			var tmpMap map[string]interface{}
			_ = json.Unmarshal([]byte(msg), &tmpMap)
			uuidMap[tmpMap["id"].(string)] = struct{}{}
		}
		var tmpLen int
		for _, msg := range res2 {
			var tmpMap map[string]interface{}
			_ = json.Unmarshal([]byte(msg), &tmpMap)
			if _, ok := uuidMap[tmpMap["id"].(string)]; ok {
				tmpLen++
			}
		}
		if tmpLen > 0 {
			t.Errorf("again expect 0, but got %d", tmpLen)
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
		var f = func(start int) []string {
			list := make([]string, 0, test.filter)
			realLen := 0
			for s, bytes := range common.NewBuild().BuildString(test.len, test.repeat, start) {
				if len(list) < test.filter {
					list = append(list, s)
				}
				realLen = batch.Push(bytes, s)
			}
			fmt.Println(fmt.Sprintf("TestString2-%d:%+v,realLen=%d,start=%d", i, test, realLen, start))
			for _, s := range list {
				batch.Filter(s)
			}
			res := batch.GetClean()
			if len(res) != realLen-test.filter {
				t.Errorf("expect %d, but got %d", realLen-test.filter, len(res))
			}
			return res
		}
		res1 := f(0)
		res2 := f(test.len)

		uuidMap := make(map[string]struct{}, len(res1))
		for _, msg := range res1 {
			var tmpMap map[string]interface{}
			_ = json.Unmarshal([]byte(msg), &tmpMap)
			uuidMap[tmpMap["id"].(string)] = struct{}{}
		}
		var tmpLen int
		for _, msg := range res2 {
			var tmpMap map[string]interface{}
			_ = json.Unmarshal([]byte(msg), &tmpMap)
			if _, ok := uuidMap[tmpMap["id"].(string)]; ok {
				tmpLen++
			}
		}
		if tmpLen > 0 {
			t.Errorf("again expect 0, but got %d", tmpLen)
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
		var f = func(start int) []interface{} {
			list := make([]string, 0, test.filter)
			realLen := 0
			for s, bytes := range common.NewBuild().BuildInterface(test.len, test.repeat, start) {
				if len(list) < test.filter {
					list = append(list, s)
				}
				realLen = batch.Push(bytes, s)
			}
			fmt.Println(fmt.Sprintf("TestInterface-%d:%+v,realLen=%d,start=%d", i, test, realLen, start))
			batch.FilterMulti(list)
			res := batch.GetClean()
			if len(res) != realLen-test.filter {
				t.Errorf("expect %d, but got %d", realLen-test.filter, len(res))
			}
			return res
		}
		res1 := f(0)
		res2 := f(test.len)

		uuidMap := make(map[string]struct{}, len(res1))
		for _, msg := range res1 {
			var tmpMap map[string]interface{}
			if tmpVal, ok := msg.([]uint8); ok {
				_ = json.Unmarshal(tmpVal, &tmpMap)
			} else if _, ok = msg.(map[string]interface{}); ok {
				tmpMap = msg.(map[string]interface{})
			} else {
				_ = json.Unmarshal([]byte(msg.(string)), &tmpMap)
			}
			uuidMap[tmpMap["id"].(string)] = struct{}{}
		}
		var tmpLen int
		for _, msg := range res2 {
			var tmpMap map[string]interface{}
			if tmpVal, ok := msg.([]uint8); ok {
				_ = json.Unmarshal(tmpVal, &tmpMap)
			} else if _, ok = msg.(map[string]interface{}); ok {
				tmpMap = msg.(map[string]interface{})
			} else {
				_ = json.Unmarshal([]byte(msg.(string)), &tmpMap)
			}
			if _, ok := uuidMap[tmpMap["id"].(string)]; ok {
				tmpLen++
			}
		}
		if tmpLen > 0 {
			t.Errorf("again expect 0, but got %d", tmpLen)
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
		var f = func(start int) []interface{} {
			list := make([]string, 0, test.filter)
			realLen := 0
			for s, bytes := range common.NewBuild().BuildInterface(test.len, test.repeat, start) {
				if len(list) < test.filter {
					list = append(list, s)
				}
				realLen = batch.Push(bytes, s)
			}
			fmt.Println(fmt.Sprintf("TestInterface2-%d:%+v,realLen=%d,start=%d", i, test, realLen, start))
			for _, s := range list {
				batch.Filter(s)
			}
			res := batch.GetClean()
			if len(res) != realLen-test.filter {
				t.Errorf("expect %d, but got %d", realLen-test.filter, len(res))
			}
			return res
		}
		res1 := f(0)
		res2 := f(test.len)

		uuidMap := make(map[string]struct{}, len(res1))
		for _, msg := range res1 {
			var tmpMap map[string]interface{}
			if tmpVal, ok := msg.([]uint8); ok {
				_ = json.Unmarshal(tmpVal, &tmpMap)
			} else if _, ok = msg.(map[string]interface{}); ok {
				tmpMap = msg.(map[string]interface{})
			} else {
				_ = json.Unmarshal([]byte(msg.(string)), &tmpMap)
			}
			uuidMap[tmpMap["id"].(string)] = struct{}{}
		}
		var tmpLen int
		for _, msg := range res2 {
			var tmpMap map[string]interface{}
			if tmpVal, ok := msg.([]uint8); ok {
				_ = json.Unmarshal(tmpVal, &tmpMap)
			} else if _, ok = msg.(map[string]interface{}); ok {
				tmpMap = msg.(map[string]interface{})
			} else {
				_ = json.Unmarshal([]byte(msg.(string)), &tmpMap)
			}
			if _, ok := uuidMap[tmpMap["id"].(string)]; ok {
				tmpLen++
			}
		}
		if tmpLen > 0 {
			t.Errorf("again expect 0, but got %d", tmpLen)
		}
	}
}
