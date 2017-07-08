// 优先队列
package main

import (
	"container/heap"
	"fmt"
)

type Item struct {
	value    string
	priority int
	index    int
}

// 一个实现了 heap.interface 接口的优先队列
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	// Pop 返回最大值
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

// Push 复杂度 O(log(n))
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

// Pop 复杂度 O(log(n))
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // 为了安全考虑
	*pq = old[0 : n-1]

	return item
}

func (pq *PriorityQueue) update(item *Item, value string, priority int) {
	item.value = value
	item.priority = priority
	// O(log(n))
	heap.Fix(pq, item.index)
}

func main() {
	// 元素及优先级
	items := map[string]int{
		"banana": 3, "apple": 2, "pear": 4,
	}

	// 创建一个优先队列，并将上述元素放入队列中；然后初始化队列以满足堆的不变性、幂等性
	pq := make(PriorityQueue, len(items))
	i := 0
	for value, priority := range items {
		pq[i] = &Item{
			value:    value,
			priority: priority,
			index:    i,
		}
		i++
	}
	// O(n)
	heap.Init(&pq)

	// 队列插入一个元素，然后修改它的优先级
	item := &Item{
		value:    "orange",
		priority: 1,
	}
	heap.Push(&pq, item)
	pq.update(item, item.value, 5)

	// 降序打印队列中的元素
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		fmt.Printf("%.2d:%s ", item.priority, item.value)
	}
}
