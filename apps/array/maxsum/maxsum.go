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
func NaiveMaxSum(arr []int, k int) int {
	var maxSum int
	l := len(arr)
	if l < k {
		return maxSum
	}

	for i := 0; i < l - k; i++ {
		sum := sum(arr[i:i+k])
		if sum > maxSum {
			maxSum = sum
		}
	}
	return maxSum
}

// SlidingMaxSum calculates the maximum sum of k consecutive elements in arr
// by sliding a window of k elements over array, incrementally updating the
// maximum sum
// O(n)
// current maximum sum of those k elements
// https://www.geeksforgeeks.org/window-sliding-technique/
func SlidingMaxSum(arr []int, k int) int {
	var maxSum int
	l := len(arr)
	if l < k {
		return maxSum
	}

	// sum of the first k elements
	maxSum = sum(arr[0:k])
	sum := maxSum
	for i := k; i < l; i++ {
		// update the "sliding sum"
		sum += -arr[i-k] + arr[i]
		if sum > maxSum {
			maxSum = sum
		}
	}
	return maxSum
}
