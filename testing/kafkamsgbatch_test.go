// Package testing
//
// ----------------develop info----------------
//
//	@Author xunmuhuang@rastar.com
//	@DateTime 2025-1-8 17:26
//
// --------------------------------------------
package testing

import (
	"fmt"
	"github.com/calmu/gotool/kafkamsgbatch"
	"github.com/calmu/gotool/testing/common"
	"testing"
)

func TestKafkaMsgBatch(t *testing.T) {
	tests := []struct {
		len, repeat, filter int
	}{
		{len: 100, repeat: 0, filter: 0},
		{len: 120, repeat: 30, filter: 0},
		{len: 200, repeat: 3, filter: 10},
		{len: 300, repeat: 50, filter: 20},
	}

	for i, test := range tests {
		batch := kafkamsgbatch.NewSliceMsgBatch(test.len)
		var f = func(start int) []*kafkamsgbatch.KafkaMsg {
			list := make([]string, 0, test.filter)
			realLen := 0
			for s, bytes := range common.NewBuild().BuildKafkaMsgBatch(test.len, test.repeat, start) {
				if len(list) < test.filter {
					list = append(list, s)
				}
				realLen = batch.Push(bytes, s)
			}
			fmt.Println(fmt.Sprintf("TestKafkaMsgBatch-%d:%+v,realLen=%d,start=%d", i, test, realLen, start))
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
			uuidMap[msg.Key] = struct{}{}
		}
		var tmpLen int
		for _, msg := range res2 {
			if _, ok := uuidMap[msg.Key]; ok {
				tmpLen++
			}
		}
		if tmpLen > 0 {
			t.Errorf("again expect 0, but got %d", tmpLen)
		}
	}
}

func TestKafkaMsgBatch2(t *testing.T) {
	tests := []struct {
		len, repeat, filter int
	}{
		{len: 100, repeat: 0, filter: 0},
		{len: 120, repeat: 30, filter: 0},
		{len: 200, repeat: 3, filter: 10},
		{len: 300, repeat: 50, filter: 20},
	}

	for i, test := range tests {
		batch := kafkamsgbatch.NewSliceMsgBatch(test.len)
		var f = func(start int) []*kafkamsgbatch.KafkaMsg {
			list := make([]string, 0, test.filter)
			realLen := 0
			for s, bytes := range common.NewBuild().BuildKafkaMsgBatch(test.len, test.repeat, start) {
				if len(list) < test.filter {
					list = append(list, s)
				}
				realLen = batch.Push(bytes, s)
			}
			fmt.Println(fmt.Sprintf("TestKafkaMsgBatch2-%d:%+v,realLen=%d,start=%d", i, test, realLen, start))
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
			uuidMap[msg.Key] = struct{}{}
		}
		var tmpLen int
		for _, msg := range res2 {
			if _, ok := uuidMap[msg.Key]; ok {
				tmpLen++
			}
		}
		if tmpLen > 0 {
			t.Errorf("again expect 0, but got %d", tmpLen)
		}
	}
}
