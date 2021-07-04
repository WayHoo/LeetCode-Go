package sort

/*
 Reference: https://github.com/CyC2018/CS-Notes/blob/master/notes/%E7%AE%97%E6%B3%95%20-%20%E6%8E%92%E5%BA%8F.md
*/

import (
	"errors"
	"reflect"
)

// Bubble 冒泡排序
func Bubble(T []int) {
	if len(T) <= 1 {
		return
	}
	N := len(T)
	isSorted := false
	for i := N - 1; i > 0 && !isSorted; i-- {
		isSorted = true
		for j := 0; j < i; j++ {
			if T[j] > T[j+1] {
				isSorted = false
				T[j], T[j+1] = T[j+1], T[j]
			}
		}
	}
}

// Insertion 插入排序
func Insertion(T []int) {
	if len(T) <= 1 {
		return
	}
	N := len(T)
	for i := 1; i < N; i++ {
		for j := i; j > 0 && T[j] < T[j-1]; j-- {
			T[j-1], T[j] = T[j], T[j-1]
		}
	}
}

// Shell 希尔排序
// 希尔排序增量序列简介：https://blog.csdn.net/Foliciatarier/article/details/53891144
func Shell(T []int) {
	if len(T) <= 1 {
		return
	}
	N := len(T)
	h := 1
	// Knuth 增量序列
	for h < N/3 {
		h = 3*h + 1
	}
	for h > 0 {
		for i := h; i < N; i++ {
			for j := i; j >= h && T[j] < T[j-h]; j -= h {
				T[j-h], T[j] = T[j], T[j-h]
			}
		}
		h /= 3
	}
}

func merge(T []int, l, m, r int) {
	if !(l >= 0 && l <= m && m <= r && r < len(T)) {
		return
	}
	var (
		tmp     []int
		i, j, k = 0, m + 1 - l, l
	)
	tmp = append(tmp, T[l:r+1]...)
	for ; i <= m-l && j <= r-l; k++ {
		if tmp[i] <= tmp[j] {
			T[k] = tmp[i]
			i++
		} else {
			T[k] = tmp[j]
			j++
		}
	}
	for ; i <= m-l; i, k = i+1, k+1 {
		T[k] = tmp[i]
	}
	for ; j <= r-l; j, k = j+1, k+1 {
		T[k] = tmp[j]
	}
}

// Up2DownMergeSort 自顶向下归并排序
func Up2DownMergeSort(T []int) {
	if len(T) <= 1 {
		return
	}
	up2DownSubSort(T, 0, len(T)-1)
}

func up2DownSubSort(T []int, l, r int) {
	if l == r {
		return
	}
	m := (l + r) / 2
	up2DownSubSort(T, l, m)
	up2DownSubSort(T, m+1, r)
	merge(T, l, m, r)
}

// Down2UpMergeSort 自底向上归并排序
func Down2UpMergeSort(T []int) {
	if len(T) <= 1 {
		return
	}
	N := len(T)
	for sz := 1; sz < N; sz <<= 1 {
		for i := 0; i+sz < N; i += sz + sz {
			l := i + sz + sz - 1
			if l >= N {
				l = N - 1
			}
			merge(T, i, i+sz-1, l)
		}
	}
}

// less judge if a < b
func less(a, b interface{}) (bool, error) {
	if reflect.TypeOf(a).Kind() != reflect.TypeOf(b).Kind() {
		return false, errors.New("two params with different types")
	}
	if !reflect.TypeOf(a).Comparable() {
		return false, errors.New("not comparable")
	}
	switch a.(type) {
	case int8, int16, int32, int64, int, uint8, uint16, uint32, uint64, uint:
		return reflect.ValueOf(a).Int() < reflect.ValueOf(b).Int(), nil
	case float32, float64:
		return reflect.ValueOf(a).Float() < reflect.ValueOf(b).Float(), nil
	case string:
		return reflect.ValueOf(a).String() < reflect.ValueOf(b).String(), nil
	default:
		return false, errors.New("type unsupported")
	}
}
