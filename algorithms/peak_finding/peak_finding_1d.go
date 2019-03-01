// Algorithms MIT 6.006
// Peak finding algorithm

package algorithms

import (
	"fmt"
)

func peakFindAux(array []int, low int, high int, n int) int {
	var mid = low + (high-low)/2
	if (mid == 0 || array[mid-1] <= array[mid]) && (mid == n-1 || array[mid+1] <= array[mid]) {
		return mid
	} else if mid > 0 && array[mid-1] > array[mid] {
		return peakFindAux(array, low, mid-1, n)
	} else {
		return peakFindAux(array, mid+1, high, n)
	}
}

func peakFind(array []int, length int) int {
	return peakFindAux(array, 0, length-1, length)
}

func main() {
	arr := []int{1, 2, 3, 4, 3, 5}
	fmt.Println("Arr:", arr, "result", arr[peakFind(arr, len(arr))])

	arr = []int{110, 110}
	fmt.Println("Arr:", arr, "result", arr[peakFind(arr, len(arr))])
}
