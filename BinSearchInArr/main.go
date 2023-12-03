package main

import (
	"fmt"
	"math"
)

// Sorted Insert Position
// You are given a sorted array A of size N and a target value B.
// Your task is to find the index (0-based indexing) of the target value in the array.
// If the target value is present, return its index.
// If the target value is not found, return the index of least element greater than equal to B.
// Your solution should have a time complexity of O(log(N)).
func SortedInsertPosition(A []int, B int) int {
	n := len(A)
	s := 0
	e := n - 1
	res := -1
	nextRes := -1
	for s <= e {
		mid := (s + e) / 2
		if A[mid] == B {
			res = mid
			e = mid - 1
		} else if A[mid] > B {
			nextRes = mid
			e = mid - 1
		} else {
			nextRes = mid + 1
			s = mid + 1
		}
	}
	//	fmt.Println("res : ", res, "nex :", nextRes)
	if res >= 0 {
		return res
	}
	return nextRes
}

// Search for a Range
func searchRange(A []int, B int) (C []int) {
	leftMostIdx := firstOrLastOccurranceInDuplicateArrOpt(A, B, false) // left most idx
	rightMostIdx := firstOrLastOccurranceInDuplicateArrOpt(A, B, true) // right most idx
	C = append(C, leftMostIdx, rightMostIdx)
	return
}
func firstOrLastOccurranceInDuplicateArrOpt(A []int, target int, islastOccur bool) (idx int) {
	n := len(A)
	s := 0
	e := n - 1
	res := -1
	for s <= e {
		mid := (s + e) / 2
		if A[mid] == target && islastOccur {
			res = mid
			s = mid + 1
		} else if A[mid] == target && !islastOccur {
			res = mid
			e = mid - 1
		} else if A[mid] > target {
			e = mid - 1
		} else {
			s = mid + 1
		}

	}
	return res
}

// Single Element in Sorted Array
func SingleElementSortedArray(A []int) int {
	n := len(A)
	s := 0
	e := n - 1
	if n == 1 {
		return A[0]
	}
	if A[s] != A[s+1] {
		return A[s]
	}
	if A[e] != A[e-1] {
		return A[e]
	}

	// Binary Search
	for s <= e {
		mid := (s + e) / 2
		// CASE 1
		if (A[mid] != A[mid-1]) && (A[mid] != A[mid+1]) {
			return A[mid]
		} else if (A[mid] == A[mid+1]) && (mid%2 == 0) || (A[mid] == A[mid-1]) && (mid%2 != 0) {
			// CASE 2 and CASE 3
			s = mid + 1
		} else {
			// CASE 4 and CASE 5
			e = mid - 1
		}
	}
	// If  element not found
	return -1
}

// Matrix Search
func searchMatrix(A [][]int, B int) int {
	nRow := len(A)
	nCol := len(A[0])
	if A == nil {
		return 0
	}
	row := 0
	col := nCol - 1
	res := 0
	for row < nRow && col >= 0 {
		if A[row][col] == B {
			res = 1
			col--
		} else if B < A[row][col] {
			col--
		} else {
			row++
		}
	}
	return res
}

// sort in O(n) , single loop
func sortArrInON(A []int) []int {
	for i := 0; i < len(A)-1; i++ {
		//swaping
		if A[i] > A[i+1] {
			A[i], A[i+1] = A[i+1], A[i]
			i = -1 // again it make 0 by i++ in loop
		}
	}
	return A
}

// product of 3 smallest in array
func minProdunctOf3Elem(A []int) int {
	sortArrInON(A)
	n := len(A)
	//all positive
	if A[0] >= 0 {
		return A[0] * A[1] * A[2]
	}
	// in negative
	min1 := A[0] * A[1] * A[2]
	min2 := A[0] * A[n-1] * A[n-2]
	ans := min1
	if min2 < min1 {
		ans = min2
	}
	return ans
}

// good pair (i <= j && A[i] >= A[j])
func countGoodPair(A, B string) int {

	i, j, res := 0, 0, 0
	m := len(A)
	n := len(B)
	for i < m && j < n {
		if A[i] >= B[j] {
			val := int(math.Abs(float64(j - i)))
			if val > res {
				res = val
			}
			j++
		} else {
			i++
		}
		if i > j {
			j = i
		}

	}
	return res
}
func main() {
	// a := []int{8, 45, 96, 1, 5}
	// res := sortArrInON(a)
	// fmt.Println("res :", res)

	// a := []int{-8, -6, -5, -4, -2, 1, 10, 3}
	// a := []int{-4, -3, -2, 1, 0, 0, 0}
	// a := []int{1, 2, -3, 4, 5, 0}
	// res := minProdunctOf3Elem(a)
	// fmt.Println("res :", res)

	res := countGoodPair("aabcccccd", "aabbbb")
	fmt.Println("res :", res)
}
