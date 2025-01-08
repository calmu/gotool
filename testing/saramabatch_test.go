// Package gotool
//
// ----------------develop info----------------
//
//	@Author xunmuhuang@rastar.com
//	@DateTime 2025-1-8 14:54
//
// --------------------------------------------
package testing

import (
	"fmt"
	"github.com/calmu/gotool/saramabatch"
	"github.com/calmu/gotool/testing/common"
	"testing"
)

func TestSaramaBatch(t *testing.T) {
	tests := []struct {
		len, repeat, filter int
	}{
		{len: 100, repeat: 0, filter: 0},
		{len: 120, repeat: 30, filter: 0},
		{len: 200, repeat: 3, filter: 10},
		{len: 300, repeat: 50, filter: 20},
	}

	for i, test := range tests {
		batch := saramabatch.NewBatchMsgList(test.len)
		list := make([]string, 0, test.filter)
		realLen := 0
		for s, bytes := range common.NewBuild().BuildSaramaBatch(test.len, test.repeat) {
			if len(list) < test.filter {
				list = append(list, s)
			}
			realLen = batch.Push(&bytes, s)
		}
		fmt.Println(fmt.Sprintf("TestSaramaBatch-%d:%+v,realLen=%d", i, test, realLen))
		batch.FilterMulti(list)
		if res := batch.GetClean(); len(res) != realLen-test.filter {
			t.Errorf("expect %d, but got %d", realLen-test.filter, len(res))
		}
	}
}

func TestSaramaBatch2(t *testing.T) {
	tests := []struct {
		len, repeat, filter int
	}{
		{len: 100, repeat: 0, filter: 0},
		{len: 120, repeat: 30, filter: 0},
		{len: 200, repeat: 3, filter: 10},
		{len: 300, repeat: 50, filter: 20},
	}

	for i, test := range tests {
		batch := saramabatch.NewBatchMsgList(test.len)
		list := make([]string, 0, test.filter)
		realLen := 0
		for s, bytes := range common.NewBuild().BuildSaramaBatch(test.len, test.repeat) {
			if len(list) < test.filter {
				list = append(list, s)
			}
			realLen = batch.Push(&bytes, s)
		}
		fmt.Println(fmt.Sprintf("TestSaramaBatch2-%d:%+v,realLen=%d", i, test, realLen))
		for _, s := range list {
			batch.Filter(s)
		}
		if res := batch.GetClean(); len(res) != realLen-test.filter {
			t.Errorf("expect %d, but got %d", realLen-test.filter, len(res))
		}
	}
}
