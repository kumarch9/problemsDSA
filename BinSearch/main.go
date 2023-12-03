package main

import (
	"math"
	"sort"
)

// Painter's Partition Problem
func paint(A int, B int, C []int) int {
	size := len(C)
	//BS on time
	maxElem := C[0]
	sumArr := 0
	for i := 0; i < len(C); i++ {
		if C[i] > maxElem {
			maxElem = C[i]
		}
		sumArr += C[i]
	}
	if A > size {
		return (maxElem * B) % 10000003
	}
	s := maxElem
	e := sumArr
	ans := 0
	for s <= e {
		mid := s + (e-s)/2
		if checkIsAPaint(mid, C, size, A) {
			ans = mid % 10000003
			e = mid - 1 //left

		} else {
			s = mid + 1 //right
		}
	}
	return (ans * B) % 10000003
}
func checkIsAPaint(mid int, board []int, size int, k int) bool {
	cur := 1
	work := 0
	for i := 0; i < size; i++ {
		if (work + board[i]) <= mid {
			work = work + board[i]
		} else {
			cur++
			work = board[i]
		}
	}
	return cur <= k
}

// Aggressive cows
func Aggressivecows(A []int, B int) int {
	sort.SliceStable(A, func(i, j int) bool { return A[i] < A[j] })
	N := len(A)
	minDist := math.MaxInt
	for i := 1; i < N; i++ {
		min := A[i] - A[i-1]
		if min < minDist {
			minDist = min
		}
	}
	ans := 0
	s := minDist
	e := A[N-1] - A[0]
	for s <= e {
		mid := s + (e-s)/2
		//can cow place  atleast mid distance apart
		if checkCowInMinDist(mid, A, N, B) {
			ans = mid
			s = mid + 1 // right
		} else {
			e = mid - 1 // left
		}
	}
	return ans
}
func checkCowInMinDist(mid int, A []int, N int, M int) bool {
	cur := 1
	position := A[0]
	for i := 0; i < N; i++ {
		if A[i]-position >= mid {
			cur++
			position = A[i]
			if cur == M {
				return true
			}
		}
	}
	return false
}

// Square Root of Integer
func sqrt(A int) int {
	if A == 0 || A == 1 {
		return A
	}
	s := 1
	e := A
	ans := 0
	for s <= e {
		mid := s + (e-s)/2
		if (mid * mid) == A {
			return mid
		} else if mid*mid < A {
			ans = mid
			s = mid + 1
		} else {
			//mid * mid > A
			e = mid - 1
		}
	}
	return ans
}

// Rotated Sorted Array Search
func RotatedSortedArraySearch(A []int, B int) int {
	n := len(A)
	s := 0
	e := n - 1
	for s <= e {
		mid := s + ((e - s) / 2)
		if A[mid] == B {
			return mid
		}
		if A[s] < A[mid] { //mean we assume  left side  is sort
			if B >= A[s] && B <= A[mid-1] {
				e = mid - 1 //backward left side
			} else {
				s = mid + 1 //backward left side
			}
		} else { // right side is sorted
			//A[s] >= A[mid]
			if B >= A[mid+1] && B <= A[e] {
				s = mid + 1 //forward right side
			} else {
				e = mid - 1 //backward left side
			}
		}
	}
	return -1
}

// Ath Magical Number
// You are given three positive integers, A, B, and C.
// Any positive integer is magical if divisible by either B or C.
// Return the Ath smallest magical number. Since the answer may be very large,
// return modulo 109 + 7.
// Note: Ensure to prevent integer overflow while calculating.
func helperGcd(a int64, b int64) int64 {
	if a == 0 {
		return b
	}
	val := helperGcd(b%a, a)
	return val
}
func helperLcm(a, b int64) int64 {
	gcdVal := helperGcd(a, b)
	val := (a * b) / gcdVal
	return val
}
func AthMagicalNumber(A int, B int, C int) int {
	mod := 1e9 + 7
	var s int64 = 1
	var e int64 = 1e17
	var lcmbc int64 = helperLcm(int64(B), int64(C))
	for s < e {
		mid := s + (e-s)/2
		target := (mid / int64(B)) + (mid / int64(C)) - mid/lcmbc
		if target < int64(A) {
			s = mid + 1
		} else {
			e = mid
		}
	}
	return int(e) % int(mod)
}

func main() {

}
