// Package kafkamsgbatch
//
// ----------------develop info----------------
//
//	@Author xunmuhuang@rastar.com
//	@DateTime 2025-1-8 17:19
//
// --------------------------------------------
package kafkamsgbatch

type KafkaMsg struct {
	Topic string
	Key   string
	Msg   []byte
}

type SliceMsgBatch struct {
	list            []*KafkaMsg
	mapUuidWithList map[string]int
	mapWithList     map[int]struct{}
	len             int
}

// NewSliceMsgBatch
//
//	@Description:获取一个实例
//	@param len int
//	@return *SliceMsgBatch
//
// ----------------develop info----------------
//
//	@Author:		xunmuhuang@rastar.com
//	@DateTime:		2025-01-08 17:29:46
//
// --------------------------------------------
func NewSliceMsgBatch(len int) *SliceMsgBatch {
	return &SliceMsgBatch{len: len, list: make([]*KafkaMsg, 0, len), mapUuidWithList: make(map[string]int, len), mapWithList: make(map[int]struct{}, len)}
}

// Push
//
//	@Description: 插入list并且返回长度(如果有相同的uuid，则先到先得)
//	@receiver: a *SliceMsgBatch
//	@receiver a
//	@param msg *KafkaMsg
//	@param uuid string
//	@return int
//
// ----------------develop info----------------
//
//	@Author:		huang_calvin@163.com
//	@DateTime:		2024-09-08 11:31:41
//
// --------------------------------------------
func (a *SliceMsgBatch) Push(msg *KafkaMsg, uuid string) int {
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
//	@receiver: a *SliceMsgBatch
//	@receiver a
//	@return []*KafkaMsg
//
// ----------------develop info----------------
//
//	@Author:		huang_calvin@163.com
//	@DateTime:		2024-09-08 14:57:40
//
// --------------------------------------------
func (a *SliceMsgBatch) GetClean() []*KafkaMsg {
	defer func() {
		a.list = make([]*KafkaMsg, 0, a.len)
		a.mapUuidWithList = make(map[string]int, a.len)
		a.mapWithList = make(map[int]struct{}, a.len)
	}()

	l := len(a.list)
	dataList := make([]*KafkaMsg, 0, l)
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
//	@receiver: a *SliceMsgBatch
//	@receiver a
//	@param uuid string
//	@return bool
//
// ----------------develop info----------------
//
//	@Author:		huang_calvin@163.com
//	@DateTime:		2024-09-08 14:58:35
//
// --------------------------------------------
func (a *SliceMsgBatch) Filter(uuid string) bool {
	if _, ok := a.mapUuidWithList[uuid]; ok {
		delete(a.mapWithList, a.mapUuidWithList[uuid])
		return true
	}
	return false
}

// FilterMulti
//
//	@Description: 批量过滤
//	@receiver: a *SliceMsgBatch
//	@receiver a
//	@param filter []string
//	@return []string
//
// ----------------develop info----------------
//
//	@Author:		huang_calvin@163.com
//	@DateTime:		2024-09-08 11:38:31
//
// --------------------------------------------
func (a *SliceMsgBatch) FilterMulti(filter []string) []string {
	if len(filter) == 0 {
		return nil
	}
	res := make([]string, 0, len(filter))
	for _, uuid := range filter {
		if _, ok := a.mapUuidWithList[uuid]; ok {
			delete(a.mapWithList, a.mapUuidWithList[uuid])
			res = append(res, uuid)
		}
	}

	return res
}

// GetUuidList
//
//	@Description: 获得批次的uuid切片
//	@receiver: a *SliceMsgBatch
//	@receiver a
//	@return []string
//
// ----------------develop info----------------
//
//	@Author:		huang_calvin@163.com
//	@DateTime:		2024-09-08 14:21:39
//
// --------------------------------------------
func (a *SliceMsgBatch) GetUuidList() []string {
	list := make([]string, 0, len(a.list))
	for s := range a.mapUuidWithList {
		list = append(list, s)
	}
	return list
}
