package structure

import (
	"fmt"
	"math/rand"
	"reflect"
	"sync"
	"time"
)

var (
	_stack_lock = sync.Mutex{}
)

type stack struct {
	list []reflect.Value
	lock *sync.Mutex
}

func (q *stack) initIfNeeded() {
	_stack_lock.Lock()
	defer _stack_lock.Unlock()

	if q.lock == nil {
		q.lock = &sync.Mutex{}
	}
	if q.list == nil {
		q.list = []reflect.Value{}
	}
}

func (s *stack) Pop() interface{} {
	s.initIfNeeded()
	if len(s.list) == 0 {
		return nil
	}

	s.lock.Lock()
	defer s.lock.Unlock()

	ele := s.list[len(s.list)-1]
	s.list = s.list[0 : len(s.list)-1]
	return ele.Interface()
}

func (s *stack) Push(ele interface{}) error {
	s.initIfNeeded()
	if ele == nil {
		return nil
	}

	s.lock.Lock()
	defer s.lock.Unlock()

	s.list = append(s.list, reflect.ValueOf(ele))

	return nil
}

func (s *stack) Desc() string {
	ret := []interface{}{}
	for _, v := range s.list {
		ret = append(ret, v.Interface())
	}
	return fmt.Sprintf("%v", ret)
}

func TestStack() {
	s := stack{}
	for i := 0; i < 100; i++ {
		rand.Seed(time.Now().UnixNano())
		ele := rand.Intn(1000)
		if ele%2 == 0 {
			s.Push(ele)
			fmt.Printf("加入元素%v\n", ele)
		} else {
			e := s.Pop()
			if nil != e {
				fmt.Printf("弹出元素%v\n", e)
			}
		}
		fmt.Printf("结果%v\n", s.Desc())
	}
}
