package structure

import (
	"fmt"
	"math/rand"
	"reflect"
	"sync"
	"time"
)

var (
	_queue_lock = sync.Mutex{}
)

type queue struct {
	list []reflect.Value
	lock *sync.Mutex
}

func (q *queue) initIfNeeded() {
	_queue_lock.Lock()
	defer _queue_lock.Unlock()
	if q.lock == nil {
		q.lock = &sync.Mutex{}
	}
	if q.list == nil {
		q.list = []reflect.Value{}
	}
}

func (q *queue) Push(ele interface{}) {
	q.initIfNeeded()
	q.lock.Lock()
	defer q.lock.Unlock()

	q.list = append(q.list, reflect.ValueOf(ele))
}

func (q *queue) Pop() interface{} {
	q.initIfNeeded()

	q.lock.Lock()
	defer q.lock.Unlock()

	if len(q.list) == 0 {
		return nil
	}

	ele := q.list[0]
	q.list = q.list[1:]

	return ele.Interface()
}

func (q *queue) Desc() string {
	ret := []interface{}{}
	for _, v := range q.list {
		ret = append(ret, v.Interface())
	}
	return fmt.Sprintf("%v", ret)
}

func TestQueue() {
	q := queue{}
	for i := 0; i < 100; i++ {
		rand.Seed(time.Now().UnixNano())
		ele := rand.Intn(1000)
		if ele%2 == 0 {
			q.Push(ele)
			fmt.Printf("推入%v\n", ele)
		} else {
			e := q.Pop()
			fmt.Printf("弹出%v\n", e)
		}
		fmt.Printf("结果%v\n", q.Desc())
	}
}
