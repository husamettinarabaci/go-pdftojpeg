package tslice

import (
	"math"
	"sort"
)

type TSlice struct {
}

func (t TSlice) Abs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}

func (t TSlice) FindClosest(x int, data []int) int {
	closestIndex := 0
	closestValue := t.Abs(data[0] - x)
	for index, val := range data {
		currentValue := t.Abs(val - x)
		if currentValue < closestValue {
			closestValue = currentValue
			closestIndex = index
		}
	}
	return data[closestIndex]
}

func (t TSlice) Contains(data []int, value int) bool {
	for _, datum := range data {
		if datum == value {
			return true
		}
	}
	return false
}

func (t TSlice) Sum(data []int) int {
	result := 0
	for _, v := range data {
		result += v
	}
	return result
}

func (t TSlice) FindEdgest(arr []int, value int) (int, int) {
	sort.Ints(arr)
	var low = 0
	var high = len(arr) - 1
	var mid int
	for high-low > 1 {
		mid = int(math.Floor(float64((low + high) / 2)))
		if arr[mid] < value {
			low = mid
		} else {
			high = mid
		}
	}

	return arr[low], arr[high]
}
