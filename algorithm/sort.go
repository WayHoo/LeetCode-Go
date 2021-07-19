package algorithm

/*
 Reference: https://github.com/CyC2018/CS-Notes/blob/master/notes/%E7%AE%97%E6%B3%95%20-%20%E6%8E%92%E5%BA%8F.md
*/

import (
	"errors"
	"reflect"
)

// Bubble 冒泡排序
func Bubble(T []int) {
	isSorted := false
	for i := len(T) - 1; i > 0 && !isSorted; i-- {
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
	for i := 1; i < len(T); i++ {
		for j := i; j > 0 && T[j] < T[j-1]; j-- {
			T[j-1], T[j] = T[j], T[j-1]
		}
	}
}

// Shell 希尔排序
// 希尔排序增量序列简介：https://blog.csdn.net/Foliciatarier/article/details/53891144
func Shell(T []int) {
	N, h := len(T), 1
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
	var subSort func(l, r int)
	subSort = func(l, r int) {
		if l >= r {
			return
		}
		m := (l + r) / 2
		subSort(l, m)
		subSort(m+1, r)
		merge(T, l, m, r)
	}
	subSort(0, len(T)-1)
}

// Down2UpMergeSort 自底向上归并排序
func Down2UpMergeSort(T []int) {
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

// QuickSort 快速排序
func QuickSort(T []int) {
	// 选主元
	medium3 := func(l, r int) int {
		m := l + (r-l)/2
		if T[l] > T[m] {
			T[l], T[m] = T[m], T[l]
		}
		if T[l] > T[r] {
			T[l], T[r] = T[r], T[l]
		}
		if T[m] > T[r] {
			T[m], T[r] = T[r], T[m]
		}
        // 主元藏到 r-1 的位置
		T[m], T[r-1] = T[r-1], T[m]
		return T[r-1]
	}
	var quickSort func(l, r int)
	quickSort = func(l, r int) {
		if l >= r {
			return
		}
		pivot := medium3(l, r)
		i, j := l, r-1
		for i < j {
			for i++; T[i] < pivot; i++ {
			}
			for j--; T[j] > pivot; j-- {
			}
			if i < j {
				T[i], T[j] = T[j], T[i]
			}
		}
		T[i], T[r-1] = T[r-1], T[i]
		quickSort(l, i-1)
		quickSort(i+1, r)
	}
	quickSort(0, len(T)-1)
}

// HeapSort 堆排序
func HeapSort(T []int) {
    L := make([]int, 1)
    L = append(L, T...)
    sink := func(k, N int) {
        for 2*k <= N {
            j := 2*k
            if j < N && L[j] < L[j+1] {
                j++
            }
            if L[k] >= L[j] {
                break
            }
            L[k], L[j] = L[j], L[k]
            k = j
        }
    }
    N := len(T)
    for k := N/2; k >= 1; k-- {
        sink(k, N)
    }
    for N > 1 {
        L[1], L[N] = L[N], L[1]
        N--
        sink(1, N)
    }
    for i := 0; i < len(T); i++ {
        T[i] = L[i+1]
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
