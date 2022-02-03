package longestincreasingsubsequencegreedy

// LongestIncreasingSubsequenceGreedy is a function to find the longest increasing
// subsequence in a given array using a greedy approach.
// The dynamic programming approach is implemented in the dynamic programming directory.
// Worst Case Time Complexity: O(nlogn)
// Auxiliary Space: O(n), where n is the length of the array(slice).
// references: https://www.geeksforgeeks.org/construction-of-longest-monotonically-increasing-subsequence-n-log-n/
func LongestIncreasingSubsequenceGreedy(nums []int) int {
	longestIncreasingSubsequnce := make([]int, 0)

	for _, num := range nums {
		// find the leftmost index in longestIncreasingSubsequnce with value >= num
		leftmostIndex := lowerBound(longestIncreasingSubsequnce, num)

		if leftmostIndex == len(longestIncreasingSubsequnce) {
			longestIncreasingSubsequnce = append(longestIncreasingSubsequnce, num)
		} else {
			longestIncreasingSubsequnce[leftmostIndex] = num
		}
	}

	return len(longestIncreasingSubsequnce)
}

// Function to find the leftmost index in arr with value >= val, mimicking the inbuild lower_bound function in C++
// Time Complexity: O(logn)
// Auxiliary Space: O(1)
func lowerBound(arr []int, val int) int {
	searchWindowLeft, searchWindowRight := 0, len(arr)-1

	for searchWindowLeft <= searchWindowRight {
		middle := (searchWindowLeft + searchWindowRight) / 2

		if arr[middle] < val {
			searchWindowLeft = middle + 1
		} else {
			searchWindowRight = middle - 1
		}
	}

	return searchWindowRight + 1
}
