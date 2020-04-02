package go808

import (
	"container/heap"
	"sync"
	"time"
)

type deadline struct {
	ID     uint64
	Index  int
	Expire int64
}

// 到期队列
type minHeap []*deadline

func (h minHeap) Len() int {
	return len(h)
}

func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].Index, h[j].Index = i, j
}

func (h minHeap) Less(i, j int) bool {
	return h[i].Expire < h[j].Expire
}

func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	item.Index = -1
	*h = old[0 : n-1]
	return item
}

func (h *minHeap) Push(x interface{}) {
	n := len(*h)
	item := x.(*deadline)
	item.Index = n
	*h = append(*h, item)
}

func (h *minHeap) Top() *deadline {
	if h.Len() == 0 {
		return nil
	}
	return (*h)[0]
}

func (h *minHeap) Update(item *deadline, expire int64) {
	item.Expire = expire
	heap.Fix(h, item.Index)
}

func (h *minHeap) Remove(item *deadline) {
	heap.Remove(h, item.Index)
}

// 保活计时器
type keepaliveTimer struct {
	keepalive int64
	h         *minHeap
	mutex     sync.Mutex
	table     map[uint64]*deadline
	callback  func(uint64)
}

// 创建计时器
func newKeepaliveTimer(keepalive int64, callback func(uint64)) *keepaliveTimer {
	h := make(minHeap, 0)
	timer := keepaliveTimer{
		h:         &h,
		callback:  callback,
		keepalive: keepalive,
		table:     make(map[uint64]*deadline),
	}
	go timer.start()
	return &timer
}

// 开始作业
func (timer *keepaliveTimer) start() {
	t := time.NewTicker(time.Second * 1)
	for {
		select {
		case <-t.C:
			ids := make([]uint64, 0)
			now := time.Now().Unix()

			timer.mutex.Lock()
			for {
				if timer.h.Len() == 0 {
					break
				}

				top := timer.h.Top()
				if top == nil || top.Expire > now {
					break
				}
				timer.h.Remove(top)
				delete(timer.table, top.ID)
				ids = append(ids, top.ID)
			}
			timer.mutex.Unlock()

			if timer.callback != nil {
				for _, id := range ids {
					timer.callback(id)
				}
			}
		}
	}
}

// 更新时间
func (timer *keepaliveTimer) update(id uint64) {
	timer.remove(id)
	d := deadline{
		ID:     id,
		Expire: time.Now().Unix() + timer.keepalive,
	}
	timer.mutex.Lock()
	defer timer.mutex.Unlock()
	heap.Push(timer.h, &d)
	timer.table[id] = &d
}

// 删除计时
func (timer *keepaliveTimer) remove(id uint64) {
	timer.mutex.Lock()
	defer timer.mutex.Unlock()
	deadline, ok := timer.table[id]
	if !ok {
		return
	}
	delete(timer.table, id)
	timer.h.Remove(deadline)
}
