package main

import (
	"math"
	"sort"
)

// Merge Two Sorted Arrays
func MergeTwoSortedArrays(A []int, B []int) (C []int) {
	nA := len(A)
	nB := len(B)
	//temp := make([]int, nA+nB)
	i, j, k := 0, 0, 0
	for i < nA && j < nB {
		if A[i] <= B[j] {
			// temp[k] = A[i]
			C = append(C, A[i])
			k++
			i++
		} else {
			//A[i] > B[j]
			//temp[k] = B[j]
			C = append(C, B[j])
			k++
			j++
		}
	}
	//looping for remains
	for i < nA {
		//temp[k] = A[i]
		C = append(C, A[i])
		k++
		i++
	}
	for j < nB {
		//temp[k] = A[j]
		C = append(C, B[j])
		k++
		j++
	}
	return
}

// Inversion count in an array
// Given an array of integers A. If i < j and A[i] > A[j],
//
//	then the pair (i, j) is called an inversion of A. Find the total
//	 number of inversions of A modulo (109 + 7).
func Inversioncountinarray(A []int) int {
	n := len(A)
	l := 0
	r := n - 1
	x := inverseCount(A, l, r)
	return x % 1000000007
}
func inverseCount(A []int, l, r int) int {
	if l == r {
		return 0
	}
	inv := 0

	mid := (l + r) / 2
	inv += inverseCount(A, l, mid) % 1000000007
	inv += inverseCount(A, mid+1, r) % 1000000007
	inv += countPairGivenIndices(A, l, mid, r) % 1000000007
	return inv % 1000000007
}
func countPairGivenIndices(A []int, l, y, r int) int {
	C := make([]int, r-l+1)
	i, j, k := l, y+1, 0
	count := 0
	for i <= y && j <= r {
		if A[i] <= A[j] {
			C[k] = A[i]
			k++
			i++
		} else {
			//A[i] > A[j]
			C[k] = A[j]
			count += y - i + 1
			k++
			j++
		}
	}
	//remain values bind
	for i <= y {
		C[k] = A[i]
		k++
		i++
	}
	for j <= r {
		C[k] = A[j]
		k++
		j++
	}
	//copy c array
	i = l
	k = 0
	for i <= r {
		A[i] = C[k]
		i++
		k++
	}
	return count
}

// Minimum Absolute Difference
// Given an array of integers A, find and return the minimum value
// of | A [ i ] - A [ j ] | where i != j and |x| denotes the absolute
// value of x.
func MinimumAbsoluteDifference(A []int) int {
	n := len(A)
	def := 0
	ans := math.MaxInt
	sort.IntSlice(A).Sort()
	for i := 1; i < n; i++ {
		def = A[i] - A[i-1]
		if def < ans {
			ans = def
		}
	}
	return ans
}

// Max Chunks To Make Sorted
func MaxChunksToMakeSorted(A []int) int {
	n := len(A)
	ans := 0
	currMax := A[0]
	for i := 0; i < n; i++ {
		if A[i] > currMax {
			currMax = A[i]
		}
		if currMax == i {
			ans++
		}
	}
	return ans
}

func main() {

}
