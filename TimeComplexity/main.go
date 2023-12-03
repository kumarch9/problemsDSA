package main

//# remains Assignment
/*
 //Find Time Complexity - 4
 int func(int n){
int s = 0;
for(int i = 1 ; i*i*i <= n ; i++){
s = s + i;
}
return s;
}
// ans :O(n^(1/3))
*/
/*
// Time Complexity-i
int ans = 0
for (i ==; i<n ;i++):
	ans+= i*i
return ans

//ans :O(n)
*/

/*
//Find Time Complexity - 9
for(int i = 1 ; i <= n ; i++){
for(int j = 1 ; j <= 3^i ; j++){
	print(i + j);
}
}
//asn : O(3^n)
*/
/*
//Time-Complexity-8
If an algorithm has a time complexity of O(1), then the complexity of it is ?
ans : constant

 //Time-Complexity-9
 If for an algorithm time complexity is given by O(log2n) then complexity will:
 ans : logarithmic

 //Time-Complexity-10
If an algorithm has a time complexity of O(n), then the complexity of it is ?
ans : linear

//Time-Complexity-12
If for an algorithm time complexity is given by O((3/2)^n) then complexity will:
ans :exponential

//NESTED_CMPL
//What is the time, space complexity of following code :
int a = 0, b = 0;
for (i = 0; i < N; i++) {
for (j=0; j < N; j++) {
	a = a + j;
}
}
for (k = 0; k < N; k++) { b = b+k;
}
//ans :O(N * N) time, O(1) space


// Time Complexity - M4
void solve(int N) {
for(int i = 0 i < pow(2, N) ;i++) {
int j = 1;
while (j > 0) \
j = 1; -=
}
}
//ans :O(4^N)

//Find Time Complexity
What will be the time complexity of the above function where n is a positive integer?
void function (int n) {
while (n > 0) {
n++;
n -= 2;
}
}
//ans :O(n)

//Time Complexity-iii
What is the time complexity of the following code snippet?
for(int i = 0; i < n; i++){
for(int j = i- 1; j >= 0; j++){
ans += i + j;
}
}
//ans : Code will run indefinitely

//Time Complexity - 8
What is the time complexity of the following code snippet?
sum = 0;
for(int i = 0; i < N; i++){
for(int j = i; j <= N && j > i; j++){
sum += 1;
}
}
//ans : O(N)

// Find Time Complexity - 1
What is the time complexity of the following code snippet
for(int i = 1 ; i <= n ; i+=2){
print(i);
}
ans : O(n)

//Find Time Complexity - 2
What is the time complexity of the following code snippet
void solve(int N, int M){
for(int i = 1 ; i <= N ; i++){
if(N % i == 0)
print(i);
}
for(int i = 1 ; i <= M ; i++){
if(M % i == 0)
print(i);
}
}
ans :O(N+M)

// Find Time Complexity - 3
What is the time complexity of the following code snippet
int func(int n){
int s = 0;
for(int i = 1 ; i <= 100 ; i++){
s = s + i;
}
return s;
}
ans : O(1)

// Find Time Complexity - 7
What is the time complexity of the following code snippet
for(int i = 0 ; i < n ; i++){
for(int j = 0 ; j <= i ; j++){
print(i+j);
}
}
ans :O(n^2)

//Find Time Complexity - 8
What is the time complexity of the following code snippet
for(int i = 1 ; i <= n ; i*=2){
for(int j = 1 ; j <= n ; j++){
print(i + j);
}
}
ans :O(nlogn)

//Time-Complexity-5
What is the time complexity of following code:
int a = 0, i = N;
while (i > 0) {
    a += i;
    i /= 2;
}
ans :O(log N)

//. Find Time Complexity - 6
What is the time complexity of the following code snippet
for(int i = 1 ; i <= 100 ; i*=2){
for(int j = 1 ; j <= n ; j++){
print(i + j);
}
}
ans : O(n)

// Find Time Complexity - 5
What is the time complexity of the following code snippet
int func(int n){
int s = 0;
for(int i = 0 ; i < n ; i = i * 2){
s = s + i;
}
return s;
}

ans : O(∞)

//Time Complexity Easy 01
int count =0; while(N) {
count++;
N/=3;
}
ans : O(log(N)) {Base 3}

//Time Complexity - 3B
void solve() {
int i = 1;
while (i < n) {
int x = 1;
while (x--) {
//0(1) operation
}
i++;
}
}
ans :O(n^2)

// Time Complexity Easy 02
for (i = 0; i < N; i++) {
for (j=i; j < N; j++) { break;
}
}
ans :O(N)

// NESTED_CMPL2
int a = 0
for ( i = 0; i < N ;i++)\
for (j = N; j> i; j--)
a = a + i + j;
}
}
ams :O(N*N)

*/
func main() {

}
