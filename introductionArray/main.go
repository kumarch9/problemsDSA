package main

// Good Pair
// Problem Description
// Given an array A and an integer B. A pair(i, j) in the array is a good
// pair if i != j and (A[i] + A[j] == B). Check if any good pair exist or not.
func GoodPair(A []int, B int) int {
	for i := 0; i < len(A); i++ {
		for j := i + 1; j < len(A); j++ {
			if A[i]+A[j] == B {
				return 1
			}
		}
	}
	return 0
}

//Reverse in a range
// Given an array A of N integers and also given two integers B and C.
//  Reverse the elements of the array A within the given inclusive range [B, C].
// Input 1:
// A = [1, 2, 3, 4]
// B = 2
// C = 3
// Input 2:

// A = [2, 5, 6]
// B = 0
// C = 2

// Output 1:[1, 2, 4, 3]
// Output 2: [6, 5, 2]
func ReverseInRange(A []int, B int, C int) []int {
	//B C are indices
	if B >= len(A) || B < 0 || C >= len(A) || C < 0 {
		return A
	}
	for i, j := B, C; i < j; i, j = i+1, j-1 {
		A[i], A[j] = A[j], A[i]
	}
	return A

}

// Array Rotation
// Problem Description:
// Given an integer array A of size N and an integer B, you have to
// return the same array after rotating it B times towards the right.
// Input 1:
// A = [1, 2, 3, 4]
// B = 2

// Input 2:
// A = [2, 5, 6]
// B = 1

// Output 1:[3, 4, 1, 2]
// Output 2:[6, 2, 5]
func ArrayRotation(A []int, B int) []int {
	clock := B % len(A)
	for i := 1; i <= clock; i++ {
		A = oneRotateLR(A)
	}
	return A
} //helper oneRotateLR()
func oneRotateLR(ar []int) []int {
	lastElem := len(ar) - 1
	temp := ar[lastElem]
	for i := len(ar) - 2; i >= 0; i-- {
		ar[i+1] = ar[i] //one forwarding
	}
	ar[0] = temp
	return ar
}

// Time Complexity - Arrays
// What is the time complexity for inserting/deleting at the beginning of the array?
// ans :O(N)

//Max Min of an Array
// Problem Description:
// Given an array A of size N. You need to find the sum of Maximum and Minimum
// element in the given array.
// Input 1:
// A = [-2, 1, -4, 5, 3]

// Input 2:
// A = [1, 3, 4, 1]

// Output 1: 1
// Output 2: 5
func MaxMinArray(A []int) int {
	max, min := A[0], A[0]
	for _, val := range A {
		if val > max {
			max = val
		}
		if val < min {
			min = val
		}
	}
	return (max + min)
}

// Linear Search - Multiple Occurences  // given frequency
// Problem Description:
// Given an array A and an integer B, find the number of occurrences of B in A.
// Input 1:
//  A = [1, 2, 2], B = 2

// Input 2:
//  A = [1, 2, 1], B = 3

//  Output 1:
//  2
// Output 2:
//  0
func LinearSearchMultipleOccurences(A []int, B int) int {
	count := 0
	for _, val := range A {
		if val == B {
			count++
		}
	}
	return count
}

//Second Largest
// Problem Description
// You are given an integer array A. You have to find the second largest
// element/value in the array or report that no such element exists.
// Input 1:

//  A = [2, 1, 2]
// Input 2:
//  A = [2]

//  Output 1: 1
// Output 2: -1
func SecondLargest(A []int) int {
	if len(A) <= 1 {
		return -1
	}
	max := A[0]
	lessMax := A[0]
	for _, val := range A {
		if val > max {
			max = val
		}
		if val < lessMax {
			lessMax = val
		}
	}
	for _, val := range A {
		if val > lessMax && val < max {
			lessMax = val
		}
	}
	if max == 0 || lessMax == 0 || max == lessMax {
		return -1
	}
	return lessMax
}

//Time to equality
// Problem Description
// Given an integer array A of size N. In one second, you can increase the value of one element by 1.
// Find the minimum time in seconds to make all elements of the array equal.
// Example Input :
// A = [2, 4, 1, 3, 2]

// Example Output : 8
func TimeToQuality(A []int) int {
	max := 0
	sumUnit := 0
	for _, val := range A {
		if val > max {
			max = val
		}
	}
	for _, val := range A {
		sumUnit += max - val
	}
	return sumUnit
}

//Problems on 1D Arrays-2 MCQ D
// What will be the output of the following code?
// class Main {
//    static void fun(int[]arr) {
//        arr[3] = 98;
//        return;
//    }
//    public static void main(String args[]) {
//        int[]arr = {10,20,30,40,50};
//        fun(arr);
//        System.out.println(arr[3]);
//    }
// }
// ans : 98

//Count of elements
// Problem Description
// Given an array A of N integers.
// Count the number of elements that have at least 1 elements greater than itself.
// Input 1:
// A = [3, 1, 2]
// Input 2:
// A = [5, 5, 3]

// Output 1:
// 2
// Output 2:
// 1
func CountOfElementsGtItself(A []int) int {
	max := A[0]
	count := 0
	for _, val := range A {
		if val > max {
			max = val
		}
	}
	for _, val := range A {
		if val == max {
			count++
		}
	}
	return (len(A) - count)
}

func main() {

}
