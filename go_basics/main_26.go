package main

// func main() {

// 	var num = removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4})
// 	fmt.Print(num)
// }

func removeDuplicates(nums []int) int {

	slow := 1
	for fast := 1; fast < len(nums); fast++ {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}
