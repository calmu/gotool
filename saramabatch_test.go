// Package gotool
//
// ----------------develop info----------------
//
//	@Author xunmuhuang@rastar.com
//	@DateTime 2025-1-8 14:54
//
// --------------------------------------------
package gotool

import (
	"encoding/json"
	"fmt"
	"github.com/IBM/sarama"
	"github.com/calmu/gotool/saramabatch"
	"math/rand"
	"testing"
	"time"
)

func buildSaramaBatch(len, repeat int) map[string]sarama.ProducerMessage {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	index := r.Intn(3)
	m := make(map[string]sarama.ProducerMessage, len)
	switch index {
	case 0:
		for k, v := range buildByte(len, repeat) {
			m[k] = sarama.ProducerMessage{
				Topic:     "test_topic",
				Key:       sarama.StringEncoder(k),
				Value:     sarama.ByteEncoder(v),
				Partition: int32(index),
			}
		}
		return m
	case 1:
		for k, v := range buildString(len, repeat) {
			m[k] = sarama.ProducerMessage{
				Topic:     "test_topic",
				Key:       sarama.StringEncoder(k),
				Value:     sarama.ByteEncoder(v),
				Partition: int32(index),
			}
		}
		return m
	}
	realLen := len - repeat
	for i := 1; i <= realLen; i++ {
		if repeat > 0 {
			k := fmt.Sprintf("a%d", repeat)
			info := map[string]interface{}{"id": k, "name": fmt.Sprintf("a%db%d", repeat, repeat)}
			v, _ := json.Marshal(info)
			m[fmt.Sprintf("a%d", repeat)] = sarama.ProducerMessage{
				Topic:     "test_topic",
				Key:       sarama.StringEncoder(k),
				Value:     sarama.ByteEncoder(v),
				Partition: int32(index),
			}
		}
		k := fmt.Sprintf("a%d", i)
		info := map[string]interface{}{"id": k, "name": fmt.Sprintf("a%db%d", i, i)}
		v, _ := json.Marshal(info)
		m[fmt.Sprintf("a%d", i)] = sarama.ProducerMessage{
			Topic:     "test_topic",
			Key:       sarama.StringEncoder(k),
			Value:     sarama.ByteEncoder(v),
			Partition: int32(index),
		}
		repeat--
	}
	return m
}

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
		for s, bytes := range buildSaramaBatch(test.len, test.repeat) {
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
		for s, bytes := range buildSaramaBatch(test.len, test.repeat) {
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
