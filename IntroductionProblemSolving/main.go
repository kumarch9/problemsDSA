package main

func FindPerfectNumbers(A int) int {
	sum := 0
	for i := 1; i*i < A; i++ {
		if A%i == 0 {
			sum += i
			sum += A / i
		}
	}
	if sum/2 == A {
		return 1
	}
	return 0
}

func Countprime(A int) int {
	countNum := 0
	for i := 1; i <= A; i++ {
		if ok := checkPrime(i); ok == 1 {
			countNum++
		}
	}
	return countNum
}

func checkPrime(num int) int {
	if num <= 1 {
		return 0
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return 0
		}
	}
	return 1
}

//How many elements are there between [135, 246] (inclusive of them) ? // 112

/*
//Find Number of Iterations - 2
void solve(int N, int M)

	{
	    int count = 0;
	    for (int i = 1; i <= N; i++)
	    {

	        if (N % i == 0)
	            count++;
	    }

	    for (int i = 1; i <= M; i++)
	    {

	        if (M % i == 0)
	            count++;
	    }
	    print count;

ans :  N+M
}
*/

/*
//Find Number of Iterations - 9
for (int i = 1; i <= n; i++)

	{
	    for (int j = 1; j <= 3 ^ i; j++)
	    {
	        print(i + j);
	    }
	}

ans :3 + 9 + 27 + .... + 3^n times.
*/

// Count Factors - 2
func CountFactors(A int) int {
	count := 0
	for i := 1; i*i <= A; i++ {
		if A%i == 0 {
			if i == A/i {
				count++
			} else {
				count += 2
			}
		}
	}
	return count
}

// IsPrime
func IsPrime(A int) int {
	if A == 1 {
		return 0
	}
	for i := 2; i < A; i++ {
		if A%i == 0 {
			return 0
		}
	}
	return 1
}

/*
//Find Iterations
for i -> 1 to N:
if(i * i == N):
return i
//ans :sqrt(N)
*/

//Sum of N natural numbers
//ans : (n*(n+1))/2

/*
//Find Number of Iterations - 3
int func(int n){
    int s = 0;
    for(int i = 1 ; i <= 100 ; i++){
        s = s + i;
    }
    return s;
}

ans  : 100
*/

/*
//Find Number of Iterations - i
int ans = 0;
for (int i = 0; i < n; i++)

	{
	    ans += i * i;
	}

//ans : n
*/

/*
//Find Number of Iterations - 7

	for (int i = 0; i < n; i++) {
	    for (int j = 0; j <= i; j++) {
	        print(i + j);
	    }
	}

//ans : n*(n+1) / 2
*/

func main() {

}
