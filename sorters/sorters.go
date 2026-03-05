package sorters

// BubbleSort - сортировка пузырьком
//
//	Description:
//	Каждый элемент сравнивается со следующим и если он больше, то элементы меняются местами
func BubbleSort(a []int) []int {
	i := len(a) - 1

	for i > 0 {
		for j := 0; j < i; j++ {
			if a[j] > a[j+1] {
				swap(a, j, j+1)
			}
		}
		i -= 1
	}

	return a
}

// InsertionSort - сортировка вставкой
//
//	Description:
//	Массив делится на отсортированную (слева) и неотсортированную части.
//	На каждом шаге элемент из неотсортированной части "вставляется" в правильное место отсортированной, сдвигая большие элементы вправо.
func InsertionSort(a []int) []int {
	var n = len(a)

	for i := 1; i < n; i++ {
		num := a[i]
		j := i - 1
		for j >= 0 && num < a[j] {
			a[j+1] = a[j]
			j--
		}
		a[j+1] = num
	}

	return a
}

// SelectionSort - сортировка выбором
//
//	Description:
//	Ищет минимальный элемент в неотсортированной части и меняет его местами с первым элементом этой части.
func SelectionSort(a []int) []int {
	var n = len(a)

	for i := 0; i < n; i++ {
		// minIndex - индекс минимального элемента
		minIndex := i

		for j := i + 1; j < n; j++ {
			// сравниваем элементы, из них запоминаем индекс меньшего элемента
			if a[j] < a[minIndex] {
				minIndex = j
			}
		}

		if minIndex != i {
			swap(a, i, minIndex)
		}
	}

	return a
}

func MergeSort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	mid := len(a) / 2
	left := MergeSort(a[:mid])
	right := MergeSort(a[mid:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	res := make([]int, 0)
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			res = append(res, left[i])
			i++
		} else {
			res = append(res, right[j])
			j++
		}
	}

	res = append(res, left[i:]...)
	res = append(res, right[j:]...)
	return res
}

func QuickLomut(a []int, low, high int) []int {
	// проверяем что массив состоит не из одного элемента или не пуст
	if low < high {
		// разделение массива
		pivotIndex := partition(a, low, high)

		// сортировка обоих частей
		QuickLomut(a, low, pivotIndex-1)
		QuickLomut(a, pivotIndex+1, high)
	}

	return a
}

func swap(a []int, i, j int) {
	a[i], a[j] = a[j], a[i]
}

func partition(a []int, low, high int) int {
	// в качестве опорного элемента выбираем медианный
	mid := (low + high) / 2
	pivot := a[mid]

	// закинем опорный элемент в конец массива для стандартного алгоритма
	a[mid], a[high] = a[high], a[mid]

	// индекс для элементов меньше опорного
	i := low - 1
	for j := low; j < high; j++ {
		if a[j] <= pivot {
			i++
			swap(a, i, j)
		}
	}

	swap(a, i+1, high)
	return i + 1
}

func QuickHoar(a []int, low, high int) {
	if low < high {
		pivotIndex := partHoar(a, low, high)

		QuickHoar(a, low, pivotIndex)
		QuickHoar(a, pivotIndex+1, high)
	}
}

func partHoar(a []int, low, high int) int {
	pivot := a[low]
	i := low - 1
	j := high + 1

	for {
		i++
		for a[i] < pivot {
			i++
		}
		j--
		for a[j] > pivot {
			j--
		}

		if i >= j {
			return j
		}

		swap(a, i, j)
	}
}
