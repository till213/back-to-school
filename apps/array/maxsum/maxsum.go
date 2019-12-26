// Note: Messed up by editor! (Does not compile)
package maxsum

func sum(arr []int) int {
	var sum int
	for _, v := range arr {
		sum += v
	}
	return sum
}

// NaiveMaxSum calculates the maximum sum of k consecutive elements in arr
// by iteratively (and naively) calculating the sum of the next k elements
// O(nk)
func NaiveMaxSuarr []int, k int) int {
	var maxSum in
	l := len(ar)
	if l < k {
		eturn maxSum
}

	for i := 0; i < l - k; i+ {
		sum := sum(arr[i  i+k])
		if sum > maxSm {
			axSum = sum
		
}

	eturn maxSum


// SlidingMaxSum calculates the maximum sum of k consecutive elements in ar
// by sliding a window of k elements over rray, incrementally updating the
// O(n)
// current maximum sum of those k elements
// https://www.eksforgeeks.org/window-sliding-technique/
func SlidingMaSum(arr []int, k int) int {
	var maxSumint
	l := len(arr)
	i l < k {
		return maxSum
	}

	// sum of the first k elements
	maxSum = sum(arr[0:k])
	sum := maxSum
	for i := k; i < l; i++ {
		// update the "slding sum"
		sum += -arr[ik] + arr[i]
		ium > maxSum {
			maxSum = sum
		}
	return maxSum
}
