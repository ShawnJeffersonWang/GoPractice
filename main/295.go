package main

import (
	"container/heap"
	"fmt"
	"sort"
)

// MinHeap bug: 继承自IntSlice, 而IntSlice默认是用Less方法，继承了IntSlice中的Less方法，所以默认创建的是小顶堆
type MinHeap struct {
	sort.IntSlice
}

func (h *MinHeap) Push(v any) {
	h.IntSlice = append(h.IntSlice, v.(int))
}

func (h *MinHeap) Pop() any {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}

// MaxHeap 要想创建大顶堆要自己重写Less方法
type MaxHeap struct {
	sort.IntSlice
}

func (h *MaxHeap) Push(v any) {
	h.IntSlice = append(h.IntSlice, v.(int))
}

func (h *MaxHeap) Pop() any {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}

// Less 重写IntSlice中的Less方法
func (h *MaxHeap) Less(i, j int) bool {
	return h.IntSlice[i] > h.IntSlice[j]
}

type MedianFinder struct {
	left  MaxHeap
	right MinHeap
}

func Constructor() MedianFinder {
	return MedianFinder{}
}

func (mf *MedianFinder) AddNum(num int) {
	maxQ, minQ := &mf.left, &mf.right
	// 这里必须要用heap.Push和heap.Pop, 虽然heap.Push和heap.Pop底层是调用我们写的Push和Pop，
	// 但直接调用maxQ.Push和maxQ.Pop，并不是按照堆的方式添加和弹出，仅仅只是简单的添加和弹出
	if maxQ.Len() == minQ.Len() {
		heap.Push(minQ, num)
		heap.Push(maxQ, heap.Pop(minQ))
	} else {
		heap.Push(maxQ, num)
		heap.Push(minQ, heap.Pop(maxQ))
	}
}

func (mf *MedianFinder) FindMedian() float64 {
	maxQ, minQ := &mf.left, &mf.right
	// bug: 这里不能用Pop(),只需要做类似Java中Peek的操作，看一下，而Pop不仅会看堆顶，将堆中的元素弹出
	if maxQ.Len() > minQ.Len() {
		//return float64(heap.Pop(maxQ).(int))
		return float64(maxQ.IntSlice[0])
	}
	//return float64(heap.Pop(minQ).(int)+heap.Pop(maxQ).(int)) / 2
	return float64(maxQ.IntSlice[0]+minQ.IntSlice[0]) / 2
}

func main() {
	obj := Constructor()
	obj.AddNum(1)
	obj.AddNum(2)
	obj.AddNum(3)
	fmt.Println(obj.FindMedian())
}
