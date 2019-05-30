package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

const (
	num      = 10000  // 测试数组长度
	rangeNum = 100000 // 数组元素大小范围
)

func main() {
	//arr := []int{6, 5, 7, 8, 4, 0, 1, 2, 9, 3}
	//org_arr := make([]int, 10)
	arr := GenerateRand()
	org_arr := make([]int, num)
	copy(org_arr, arr) //内存中属于不同的两块数组
	//冒泡排序
	//Bubble(arr)
	// 选择排序
	//SelectSort(arr)
	// 插入排序
	//InsertSort(arr)
	//快速排序
	//QuickSort(arr, 0, len(arr)-1)
	// 归并排序
	//MergeSort(arr, 0, len(arr)-1)
	// 堆排序
	HeapSort(arr)
	sort.Ints(org_arr) //sort.Ints(arr []int) 接受slice
	//fmt.Println(arr, org_arr, IsSame(arr, org_arr))
	//打印前15个数,并对比排序是否正确
	fmt.Println(arr[:15], org_arr[:15], IsSame(arr, org_arr))
}

func GenerateRand() []int {
	randSeed := rand.New(rand.NewSource(time.Now().Unix() + time.Now().UnixNano()))
	arr := make([]int, num)
	for i := 0; i < num; i++ {
		arr[i] = randSeed.Intn(rangeNum)
	}
	return arr
}

func IsSame(slice1, slice2 []int) bool {
	if len(slice1) != len(slice2) {
		return false
	}

	if (slice1 == nil) != (slice2 == nil) {
		return false
	}

	for i, num := range slice1 {
		if num != slice2[i] {
			return false
		}
	}
	return true
}

func Bubble(arr []int) {
	size := len(arr)
	var swapped bool
	for i := size - 1; i > 0; i-- {
		swapped = false
		for j := 0; j < i; j++ {
			if arr[j+1] < arr[j] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}
		if swapped != true {
			break
		}
	}

}

func SelectSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j <= len(arr)-1; j++ {
			if arr[j] < arr[i] {
				arr[j], arr[i] = arr[i], arr[j]
			}
		}
	}
}

func InsertSort(arr []int) {
	//不得不说 传统for循环更好用一些
	for i := 1; i <= len(arr)-1; i++ {
		for j := i; j > 0; j-- {
			if arr[j-1] > arr[j] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
			}
		}
	}
}

func QuickSort(arr []int, l, r int) {
	if l < r {
		pivot := arr[r]
		i := l - 1
		for j := l; j < r; j++ {
			if arr[j] <= pivot {
				i++
				arr[j], arr[i] = arr[i], arr[j]
			}
		}
		i++
		arr[r], arr[i] = arr[i], arr[r]
		QuickSort(arr, l, i-1)
		QuickSort(arr, i+1, r)
	}
}

func Merge(arr []int, l, mid, r int) {
	n1, n2 := mid-l+1, r-mid
	left, right := make([]int, n1), make([]int, n2)
	copy(left, arr[l:mid+1])
	copy(right, arr[mid+1:r+1])
	//fmt.Println("left:", left) // [6][5]-->[5][5]
	//fmt.Println("right:", right)
	i, j := 0, 0
	k := l
	for ; i < n1 && j < n2; k++ {
		if left[i] <= right[j] {
			arr[k] = left[i]
			//fmt.Println("left[i]:", left[i])
			i++
		} else {
			arr[k] = right[j]
			//fmt.Println("right[j]:", right[j])
			j++
		}
	}
	//fmt.Println("left[i]:", left)
	for ; i < n1; i++ {
		arr[k] = left[i]
		//fmt.Println("left[i]:", left[i])
		k++
	}
	for ; j < n2; j++ {
		arr[k] = right[j]
		//fmt.Println("right[j]:", right[j])
		k++
	}
}

func MergeSort(arr []int, l, r int) {
	if l < r {
		mid := (l + r - 1) / 2
		MergeSort(arr, l, mid)
		MergeSort(arr, mid+1, r)
		Merge(arr, l, mid, r)
	}
}

func adjust_heap(arr []int, i, size int) {
	if i <= (size-2)/2 {
		l, r := 2*i+1, 2*i+2
		m := i
		if l < size && arr[l] > arr[m] {
			m = l
		}
		if r < size && arr[r] > arr[m] {
			m = r
		}
		if m != i {
			arr[m], arr[i] = arr[i], arr[m]
			adjust_heap(arr, m, size)
		}
	}
}

func build_heap(arr []int) {
	size := len(arr)
	for i := (size - 2) / 2; i >= 0; i-- {
		adjust_heap(arr, i, size)
	}
}

func HeapSort(arr []int) {
	size := len(arr)
	build_heap(arr)
	for i := size - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		adjust_heap(arr, 0, i)
	}
}
