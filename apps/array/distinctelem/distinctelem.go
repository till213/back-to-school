package distinctelem

func atMostK(arr []int, k int) int {
	var count int
	var left, right int

	dc := make(map[int] int)
	for right < len(arr) {

		dc[arr[right]]++
		for len(dc) > k {
			dc[arr[left]]--
			if dc[arr[left]] == 0 {
				delete(dc, arr[left])
			}
			left++
		}

		count += right - left + 1
        right++
	}
	return count
}

// DistinctElements returns the number of subarrays in arr
// with *exactly* k distinct elements
// https://www.geeksforgeeks.org/count-of-subarrays-having-exactly-k-distinct-elements/
func DistinctElements(arr []int, k int) int {
	// Count of subarrays with exactly k distinct 
    // elements is equal to the difference of the 
    // count of subarrays with at most K distinct 
    // elements and the count of subararys with 
    // at most (K - 1) distinct elements 
	return atMostK(arr, k) - atMostK(arr, k-1)
}
