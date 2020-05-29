package sort

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func GenerateRanomArr(l int) []int {
	arr := []int{}
	for i := 0; i < l; i++ {
		rand.Seed(time.Now().UnixNano())
		arr = append(arr, rand.Intn(100))
	}
	return arr
}

func TestAll() {
	for k, v := range map[string]func(*[]int){
		// "冒泡": BubbleSort,
		// "选择": SelectSort,
		// "插入": InsertSort,
		// "中值": MidSort,
		"堆": HeapSort,
	} {
		v(&[]int{})
		for i := 0; i < 1000; i++ {
			rand.Seed(time.Now().UnixNano())
			n := rand.Intn(100)
			arr := GenerateRanomArr(n)
			fmt.Println("测试" + k + "排序, 长度" + strconv.Itoa(len(arr)) + " ...")
			v(&arr)
			unsortedIndex := CheckSort(arr)
			if unsortedIndex >= 0 {
				fmt.Println(arr)
				panic(errors.New(fmt.Sprintf("%v排序失败，位置%v", k, unsortedIndex)))
			}
		}
	}
}
