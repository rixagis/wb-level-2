// Демонстрация паттерна "стратегия" на примере сортировщика, использующего разные алгоритмы сортировки
package pattern

import (
)

// SortingStrategy - интерфейс стратегии сортировки
type SortingStrategy interface {
	Sort([]int)
}

// QuickSortStrategy - стратегия быстрой сортировки
type QuickSortStrategy struct {

}

// Sort сортирует слайс быстрой сортировкой
func (qss QuickSortStrategy) Sort(arr []int) {
	qss.quicksortRecurse(arr, 0, len(arr) - 1)
}

// partition - разделение (часть алгоритма)
func (qss QuickSortStrategy) partition(arr []int, low, high int) int {
	var pivot = arr[(low + high)/2]

	var left = low - 1
	var right = high + 1

	for {
		left++
		for arr[left] < pivot {
			left++
		}

		right--
		for arr[right] > pivot {
			right--
		}

		if left >= right {
			return right
		}
		arr[left], arr[right] = arr[right], arr[left]
	}
}

// quicksortRecurse - рекурсивная функция сортировки (часть алгоритма)
func (qss QuickSortStrategy) quicksortRecurse(arr []int, low, high int) {
	if low < high {
		var pivot = qss.partition(arr, low, high)
		qss.quicksortRecurse(arr, low, pivot)
		qss.quicksortRecurse(arr, pivot + 1, high)
	}
}

// MergeSortStrategy - стратегия сортировки слиянием
type MergeSortStrategy struct {

}

// Sort сортирует слайс сортировкой слиянием
func (mss MergeSortStrategy) Sort(A []int) {
	var B = make([]int, len(A))
	copy(B, A)
	mss.splitMerge(B, 0, len(A), A)
}

// splitMerge рекурсивно разделяет и сливает слайсы (часть алгоритма)
func (mss MergeSortStrategy) splitMerge(B []int, begin, end int, A []int) {
	if end - 1 <= begin {
		return
	}

	var middle = (begin + end) / 2

	mss.splitMerge(A, begin, middle, B)
	mss.splitMerge(A, middle, end, B)
	mss.merge(B, begin, middle, end, A)
}

// merge сливает упорядоченно отсортированные слайсы (часть алгоритма)
func (mss MergeSortStrategy) merge(A []int, begin, middle, end int, B []int) {
	var i = begin
	var j = middle
	for k := begin; k < end; k++ {
		if (i < middle) && (j >= end || A[i] <= A[j]) {
			B[k] = A[i]
			i++
		} else {
			B[k] = A[j]
			j++
		}
	}
}


// HeapSortStrategy - стратегия сортировки кучей
type HeapSortStrategy struct {

}

// Sort сортирует слайс сортировкой кучей
func (hss HeapSortStrategy) Sort(arr []int) {
	var length = len(arr)
	for i := length / 2 - 1; i > -1; i-- {
		hss.heapify(arr, length, i)
	}

	for i := length - 1; i > 0; i-- {
		arr[i], arr[0] = arr[0], arr[i]
		hss.heapify(arr, i, 0)
	}
}

// heapify делает кучу из слайса с вершиной в root (часть алгоритма)
func (hss HeapSortStrategy) heapify(arr []int, heapSize, root int) {
	var (
		largest = root
		left = 2 * root + 1
		right = 2 * root + 2
	)

	if left < heapSize && arr[left] > arr[largest] {
		largest = left
	}
	if right < heapSize && arr[right] > arr[largest] {
		largest = right
	}

	if largest != root {
		arr[largest], arr[root] = arr[root], arr[largest]
		hss.heapify(arr, heapSize, largest)
	}
}


// Sorter - сортировщик, применяющий разные стратегии сортировки
type Sorter struct {
	strategy SortingStrategy
}

// NewSorter - конструктор Sorter
func NewSorter() *Sorter {
	return &Sorter{strategy: QuickSortStrategy{}}	// стратегия по умолчанию
}

// SetStrategy устанавливает стратегию сортировки
func (s *Sorter) SetStrategy(strategy SortingStrategy) {
	s.strategy = strategy
}

// Sort сортирует слайс с помощью установленной стратегии сортировки
func (s *Sorter) Sort(arr []int) {
	s.strategy.Sort(arr)
}

// Пример использования
/*func main() {
	var (
		sorter = Sorter{}
		quickStrategy = QuickSortStrategy{}
		mergeStrategy = MergeSortStrategy{}
		heapStrategy = HeapSortStrategy{}
		arr1 = []int{12, 4, 1234, 5, 53, 5, 74, 45}
		arr2 = []int{645, 3, 2345, 5, 62, 233, 6236}
		arr3 = []int{545, 243, 5345, 234523, 23, 66, 23, 67, 54}
	)

	fmt.Println("First array (unsorted):", arr1)
	sorter.SetStrategy(quickStrategy)
	sorter.Sort(arr1)
	fmt.Println("First array (sorted by quicksort):", arr1)

	fmt.Println("Second array (unsorted):", arr2)
	sorter.SetStrategy(mergeStrategy)
	sorter.Sort(arr2)
	fmt.Println("Second array (sorted by merge sort):", arr2)

	fmt.Println("Third array (unsorted):", arr3)
	sorter.SetStrategy(heapStrategy)
	sorter.Sort(arr3)
	fmt.Println("Third array (sorted by heap sort):", arr3)


}*/