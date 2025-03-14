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
	"github.com/IBM/sarama"
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
		var f = func(start int) []*sarama.ProducerMessage {
			list := make([]string, 0, test.filter)
			realLen := 0
			for s, bytes := range common.NewBuild().BuildSaramaBatch(test.len, test.repeat, start) {
				if len(list) < test.filter {
					list = append(list, s)
				}
				realLen = batch.Push(&bytes, s)
			}
			fmt.Println(fmt.Sprintf("TestSaramaBatch-%d:%+v,realLen=%d,start=%d", i, test, realLen, start))
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
			key, _ := msg.Key.Encode()
			uuidMap[string(key)] = struct{}{}
		}
		var tmpLen int
		for _, msg := range res2 {
			key, _ := msg.Key.Encode()
			if _, ok := uuidMap[string(key)]; ok {
				tmpLen++
			}
		}
		if tmpLen > 0 {
			t.Errorf("again expect 0, but got %d", tmpLen)
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
		var f = func(start int) []*sarama.ProducerMessage {
			list := make([]string, 0, test.filter)
			realLen := 0
			for s, bytes := range common.NewBuild().BuildSaramaBatch(test.len, test.repeat, start) {
				if len(list) < test.filter {
					list = append(list, s)
				}
				realLen = batch.Push(&bytes, s)
			}
			fmt.Println(fmt.Sprintf("TestSaramaBatch-%d:%+v,realLen=%d,start=%d", i, test, realLen, start))
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
			key, _ := msg.Key.Encode()
			uuidMap[string(key)] = struct{}{}
		}
		var tmpLen int
		for _, msg := range res2 {
			key, _ := msg.Key.Encode()
			if _, ok := uuidMap[string(key)]; ok {
				tmpLen++
			}
		}
		if tmpLen > 0 {
			t.Errorf("again expect 0, but got %d", tmpLen)
		}
	}
}
