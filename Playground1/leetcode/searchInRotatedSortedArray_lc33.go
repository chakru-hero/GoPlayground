package leetcode

func searchInRotatedSortedArray(arr []int, target int) int{
	left,right := 0, len(arr)-1;
	for left<=right {
		mid:= (left + right)/2;
		//check if mid is target
		if arr[mid] == target {
			return mid;
		}
		if arr[left]<=arr[mid]{	// get in the sorted half of the array
			if arr[left] <= target && target < arr[mid] {	// check if target is in left part of the array
				right = mid -1;		// if yes, then ignore the right half.
			} else {
				left = mid + 1; // else ignore the left half 
			}
		} else {  // get in the rotated half of the array 
			if arr[mid] < target && target <= arr[right] {
				left = mid + 1;
			}else {
				right = mid -1;
			}
		}
	}
	return -1;
}
