// Package common
//
// ----------------develop info----------------
//
//	@Author xunmuhuang@rastar.com
//	@DateTime 2025-1-8 17:43
//
// --------------------------------------------
package common

import (
	"encoding/json"
	"fmt"
	"github.com/IBM/sarama"
	"github.com/calmu/gotool/kafkamsgbatch"
	"math/rand"
	"time"
)

type Build struct {
}

func NewBuild() *Build {
	return &Build{}
}

func (a *Build) BuildString(len, repeat int) map[string]string {
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

func (a *Build) BuildByte(len, repeat int) map[string][]byte {
	m := make(map[string][]byte, len)
	for s, s2 := range a.BuildString(len, repeat) {
		m[s] = []byte(s2)
	}
	return m
}

func (a *Build) BuildKafkaMsgBatch(len, repeat int) map[string]*kafkamsgbatch.KafkaMsg {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	index := r.Intn(3)
	m := make(map[string]*kafkamsgbatch.KafkaMsg, len)
	switch index {
	case 0:
		for k, v := range a.BuildByte(len, repeat) {
			m[k] = &kafkamsgbatch.KafkaMsg{Msg: v, Key: k, Topic: k}
		}
		return m
	case 1:
		for k, v := range a.BuildString(len, repeat) {
			m[k] = &kafkamsgbatch.KafkaMsg{Msg: []byte(v), Key: k, Topic: k}
		}
		return m
	}
	realLen := len - repeat
	for i := 1; i <= realLen; i++ {
		if repeat > 0 {
			k := fmt.Sprintf("a%d", repeat)
			v, _ := json.Marshal(map[string]interface{}{"id": k, "name": fmt.Sprintf("a%db%d", repeat, repeat)})
			m[k] = &kafkamsgbatch.KafkaMsg{Msg: v, Key: k, Topic: k}
			repeat--
		}
		k := fmt.Sprintf("a%d", i)
		v, _ := json.Marshal(map[string]interface{}{"id": fmt.Sprintf("a%d", i), "name": fmt.Sprintf("a%db%d", i, i)})
		m[k] = &kafkamsgbatch.KafkaMsg{Msg: v, Key: k, Topic: k}
	}
	return m
}

func (a *Build) BuildInterface(len, repeat int) map[string]interface{} {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	index := r.Intn(3)
	m := make(map[string]interface{}, len)
	switch index {
	case 0:
		for k, v := range a.BuildByte(len, repeat) {
			m[k] = v
		}
		return m
	case 1:
		for k, v := range a.BuildString(len, repeat) {
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

func (a *Build) BuildSaramaBatch(len, repeat int) map[string]sarama.ProducerMessage {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	m := make(map[string]sarama.ProducerMessage, len)
	for s, msg := range a.BuildKafkaMsgBatch(len, repeat) {
		m[s] = sarama.ProducerMessage{
			Topic:     "test_topic",
			Key:       sarama.StringEncoder(msg.Key),
			Value:     sarama.ByteEncoder(msg.Msg),
			Partition: int32(r.Intn(3)),
		}
	}
	return m
}
