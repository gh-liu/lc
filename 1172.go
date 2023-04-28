package main

import "fmt"

// TODO: timeout
func main() {
	obj := Constructor(2)
	obj.Push(1)
	obj.Push(2)
	obj.Push(3)
	obj.Push(4)
	obj.Push(5)
	obj.Push(6)
	obj.Push(7)
	obj.Push(8)
	obj.Push(9)
	obj.Push(10)
	obj.Push(11)
	obj.Push(12)
	obj.Push(13)

	res := obj.Pop()
	fmt.Println(res == 13)
	res = obj.Pop()
	fmt.Println(res == 12)
	res = obj.Pop()
	fmt.Println(res == 11)
	res = obj.Pop()
	fmt.Println(res == 10)

	obj.Push(14)
	obj.Push(15)
	obj.Push(16)
	obj.Push(17)

	fmt.Println(obj.Elements)
	fmt.Println(obj.Indexs)

	res2 := obj.PopAtStack(0)
	fmt.Println(res2 == 2)
	res = obj.PopAtStack(0)
	fmt.Println(res == 1)
}

type DinnerPlates struct {
	Elements []int // index * capacity + offset(max capacity) -1
	Indexs   []int // val <= capacity
	capacity int
}

func Constructor(capacity int) DinnerPlates {
	return DinnerPlates{
		Elements: []int{},
		Indexs:   []int{},
		capacity: capacity,
	}
}

func (d *DinnerPlates) Push(val int) {
	// 1. zero stack
	// 2. find a slot in a stack
	// 3. full: new a stack
	idx, offset := -1, 1
	grow := true

	for i, v := range d.Indexs {
		if v < d.capacity {
			idx = i
			offset = v + 1
			grow = false
			break
		} else {
			idx++
		}
	}

	if grow {
		idx++
		d.Indexs = append(d.Indexs, 0)
		d.Elements = append(d.Elements, make([]int, d.capacity)...)
	}

	d.Indexs[idx] = offset

	d.Elements[d.capacity*idx+offset-1] = val
}

func (d *DinnerPlates) Pop() int {
	for i := len(d.Indexs) - 1; i >= 0; i-- {
		if d.Indexs[i] > 0 {
			d.Indexs[i]--
			return d.Elements[d.capacity*i+d.Indexs[i]]
		}
	}
	return -1
}

func (d *DinnerPlates) PopAtStack(index int) int {
	if len(d.Indexs) < index {
		return -1
	}
	idx := d.Indexs[index]
	if idx > 0 {
		d.Indexs[index] = idx - 1
		return d.Elements[d.capacity*index+idx-1]
	}
	return -1
}
