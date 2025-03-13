// Package normalbatch
//
// ----------------develop info----------------
//
//	@Author xunmuhuang@rastar.com
//	@DateTime 2025-1-7 17:19
//
// --------------------------------------------
package normalbatch

type SliceStringBatch struct {
	list            []string
	mapUuidWithList map[string]int
	mapWithList     map[int]struct{}
	len             int
}

// NewSliceStringBatch
//
//	@Description:获取一个实例
//	@param len int
//	@return *SliceStringBatch
//
// ----------------develop info----------------
//
//	@Author:		xunmuhuang@rastar.com
//	@DateTime:		2025-01-07 17:29:46
//
// --------------------------------------------
func NewSliceStringBatch(len int) *SliceStringBatch {
	return &SliceStringBatch{len: len, list: make([]string, 0, len), mapUuidWithList: make(map[string]int, len), mapWithList: make(map[int]struct{}, len)}
}

// Push
//
//	@Description: 插入list并且返回长度(如果有相同的uuid，则先到先得)
//	@receiver: a *SliceStringBatch
//	@receiver a
//	@param msg string
//	@param uuid string
//	@return int
//
// ----------------develop info----------------
//
//	@Author:		huang_calvin@163.com
//	@DateTime:		2024-09-07 11:31:41
//
// --------------------------------------------
func (a *SliceStringBatch) Push(msg string, uuid string) int {
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
//	@receiver: a *SliceStringBatch
//	@receiver a
//	@return []string
//
// ----------------develop info----------------
//
//	@Author:		huang_calvin@163.com
//	@DateTime:		2024-09-07 14:57:40
//
// --------------------------------------------
func (a *SliceStringBatch) GetClean() []string {
	defer func() {
		a.list = make([]string, 0, a.len)
		a.mapUuidWithList = make(map[string]int, a.len)
		a.mapWithList = make(map[int]struct{}, a.len)
	}()

	l := len(a.list)
	dataList := make([]string, 0, l)
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
//	@receiver: a *SliceStringBatch
//	@receiver a
//	@param uuid string
//	@return bool
//
// ----------------develop info----------------
//
//	@Author:		huang_calvin@163.com
//	@DateTime:		2024-09-07 14:58:35
//
// --------------------------------------------
func (a *SliceStringBatch) Filter(uuid string) bool {
	if _, ok := a.mapUuidWithList[uuid]; ok {
		delete(a.mapWithList, a.mapUuidWithList[uuid])
		return true
	}
	return false
}

// FilterMulti
//
//	@Description: 批量过滤
//	@receiver: a *SliceStringBatch
//	@receiver a
//	@param filter []string
//	@return []string
//
// ----------------develop info----------------
//
//	@Author:		huang_calvin@163.com
//	@DateTime:		2024-09-07 11:38:31
//
// --------------------------------------------
func (a *SliceStringBatch) FilterMulti(filter []string) []string {
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
//	@receiver: a *SliceStringBatch
//	@receiver a
//	@return []string
//
// ----------------develop info----------------
//
//	@Author:		huang_calvin@163.com
//	@DateTime:		2024-09-07 14:21:39
//
// --------------------------------------------
func (a *SliceStringBatch) GetUuidList() []string {
	list := make([]string, 0, len(a.list))
	for s := range a.mapUuidWithList {
		list = append(list, s)
	}
	return list
}
