package iterator

import "fmt"

// 迭代器接口
type Iterator interface {
	HasNext() bool
	Next() interface{}
}

// 集合接口
type Container interface {
	GetIterator() Iterator
}

// 数组迭代器
type ArrayIterator struct {
	currentIndex int
	ac           *ArrayContainer
}

func (ai *ArrayIterator) HasNext() bool {
	if ai.ac.arrayData != nil && ai.currentIndex < len(ai.ac.arrayData) {
		return true
	}
	return false
}

func (ai *ArrayIterator) Next() interface{} {
	if ai.HasNext() {
		defer func() { ai.currentIndex++ }()
		return ai.ac.arrayData[ai.currentIndex]

	}
	return nil
}

// 数组集合
type ArrayContainer struct {
	arrayData []interface{}
}

func (ac *ArrayContainer) GetIterator() Iterator {
	return &ArrayIterator{currentIndex: 0, ac: ac}
}

func Run() {
	arr := []interface{}{"a", "b", "c", "d"}
	arrayContainer := &ArrayContainer{arrayData: arr}
	iterator := arrayContainer.GetIterator()
	for iterator.HasNext() {
		fmt.Println(iterator.Next().(string))
	}
}
