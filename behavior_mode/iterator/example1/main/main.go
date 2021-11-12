package main

import iterator2 "go-patterns/behavior_mode/iterator"

func main() {
	teacher := new(iterator2.Teacher)
	analysis := new(iterator2.Analysis)
	//迭代器
	iterator := iterator2.NewIterator()
	iterator.Add(teacher)
	iterator.Add(analysis)

	for iterator.HasNext() {
		iterator.Next().Visit()
	}
}
