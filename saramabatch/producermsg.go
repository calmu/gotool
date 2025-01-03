// Package saramabatch
//
// ----------------develop info----------------
//
//	@Author huang_calvin@163.com
//	@DateTime 2025-1-3 11:00
//
// --------------------------------------------
package saramabatch

import "github.com/IBM/sarama"

// BatchMsgList
// @Description: 存储msg
type BatchMsgList struct {
	list            []*sarama.ProducerMessage
	mapUuidWithList map[string]int
	mapWithList     map[int]struct{}
	len             int
}

// NewBatchMsgList
//
//	@Description: 获取一个实例
//	@param len int
//	@return *BatchMsgList
//
// ----------------develop info----------------
//
//	@Author:		huang_calvin@163.com
//	@DateTime:		2024-09-03 11:12:08
//
// --------------------------------------------
func NewBatchMsgList(len int) *BatchMsgList {
	return &BatchMsgList{len: len, list: make([]*sarama.ProducerMessage, 0, len), mapUuidWithList: make(map[string]int, len), mapWithList: make(map[int]struct{}, len)}
}

// Push
//
//	@Description: 插入list并且返回长度(如果有相同的uuid，则先到先得)
//	@receiver: a *BatchMsgList
//	@receiver a
//	@param msg *sarama.ProducerMessage
//	@param uuid string
//	@return int
//
// ----------------develop info----------------
//
//	@Author:		huang_calvin@163.com
//	@DateTime:		2024-09-03 11:31:41
//
// --------------------------------------------
func (a *BatchMsgList) Push(msg *sarama.ProducerMessage, uuid string) int {
	if _, ok := a.mapUuidWithList[uuid]; ok {
		return len(a.list)
	}
	a.list = append(a.list, msg)
	l := len(a.list)
	a.mapUuidWithList[uuid] = l - 1
	a.mapWithList[l-1] = struct{}{}
	return l
}

// GetClean
//
//	@Description: 返回过滤后的列表并且清理
//	@receiver: a *BatchMsgList
//	@receiver a
//	@return []*sarama.ProducerMessage
//
// ----------------develop info----------------
//
//	@Author:		huang_calvin@163.com
//	@DateTime:		2024-09-03 14:57:40
//
// --------------------------------------------
func (a *BatchMsgList) GetClean() []*sarama.ProducerMessage {
	defer func() {
		a.list = a.list[:0]
		a.mapUuidWithList = make(map[string]int, a.len)
		a.mapWithList = make(map[int]struct{}, a.len)
	}()

	l := len(a.list)
	dataList := make([]*sarama.ProducerMessage, 0, l)
	for i := 0; i < l; i++ {
		if _, ok := a.mapWithList[i]; ok {
			dataList = append(dataList, a.list[i])
		}
	}
	return dataList
}

// Filter
//
//	@Description: 过滤单条
//	@receiver: a *BatchMsgList
//	@receiver a
//	@param uuid string
//	@return *BatchMsgList
//
// ----------------develop info----------------
//
//	@Author:		huang_calvin@163.com
//	@DateTime:		2024-09-03 14:58:35
//
// --------------------------------------------
func (a *BatchMsgList) Filter(uuid string) *BatchMsgList {
	delete(a.mapUuidWithList, uuid)
	return a
}

// FilterMulti
//
//	@Description: 批量过滤
//	@receiver: a *BatchMsgList
//	@receiver a
//	@param filter []string
//	@return BatchMsgList
//
// ----------------develop info----------------
//
//	@Author:		huang_calvin@163.com
//	@DateTime:		2024-09-03 11:38:31
//
// --------------------------------------------
func (a *BatchMsgList) FilterMulti(filter []string) *BatchMsgList {
	if len(filter) == 0 {
		return a
	}
	for _, uuid := range filter {
		delete(a.mapWithList, a.mapUuidWithList[uuid])
	}

	return a
}

// GetUuidList
//
//	@Description: 获得批次的uuid切片
//	@receiver: a *BatchMsgList
//	@receiver a
//	@return []string
//
// ----------------develop info----------------
//
//	@Author:		huang_calvin@163.com
//	@DateTime:		2024-09-03 14:21:39
//
// --------------------------------------------
func (a *BatchMsgList) GetUuidList() []string {
	list := make([]string, 0, len(a.list))
	for s := range a.mapUuidWithList {
		list = append(list, s)
	}
	return list
}
