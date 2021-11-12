package example1

import "fmt"

/*
	设计思想：
		1. Iterator结构体
			实现Next()  HasNext()方法
		2. Container容器
			容器实现添加 移除Visitor 和
*/
//创建迭代器
type Iterator struct {
	index int
	Container
}

func (i *Iterator) Next() Visitor {
	fmt.Println(i.index)
	visitor := i.List[i.index]
	i.index += 1
	return visitor
}

func (i *Iterator) HasNext() bool {
	if i.index >= len(i.List) {
		return false
	}
	return true
}

//创建容器
type Container struct {
	List []Visitor
}

func (c *Container) Add(visitor Visitor) {
	c.List = append(c.List, visitor)
}

func (c *Container) Remove(index int) {
	if index < 0 || index > len(c.List) {
		return
	}
	c.List = append(c.List[:index], c.List[index+1:]...)
}

//创建Visitor接口
type Visitor interface {
	Visit()
}

//创建具体的visitor对象
type Teacher struct{}

type Analysis struct{}

func (t *Teacher) Visit() {
	fmt.Println("this is teacher visitor")
}

func (a *Analysis) Visit() {
	fmt.Println("this is analysis visitor")
}

//工厂方法创建迭代器
func NewIterator() *Iterator {
	return &Iterator{
		index:     0,
		Container: Container{},
	}
}
