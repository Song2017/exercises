package advance

import (
	"fmt"
)

func TestFlowControl() {
	fmt.Println("\nTestFlowControl")

	for_range()
	switch_case()
	if_condition()
}

func for_range() {
	numbers := []int{1, 2, 11, 3, 4, 5, 6, 7}
	for i, v := range numbers {
		fmt.Print(i, ", ", v, "; ")
	}

	nums2 := [...]int{1, 2, 3, 4, 5}
	val, err := interface{}(nums2).([5]int)
	fmt.Println("[...]int{1,2,3,4,5}", val, err)
	maxIndex2 := len(nums2) - 1
	for i, v := range nums2 {
		if i == maxIndex2 {
			nums2[0] += v
		} else {
			nums2[i+1] += v
		}
	}
	fmt.Println(nums2)

	numbers3 := []int{1, 2, 3, 4, 5, 6, 7}
	maxIndex3 := len(numbers3) - 1
	for i, v := range numbers3 {
		fmt.Print(i, ",", numbers3, "; ")
		if i == maxIndex3 {
			numbers3[0] += v
		} else {
			numbers3[i+1] += v
		}
	}
	fmt.Println(numbers3)
}

func switch_case() {
	nums := [...]int8{0, 1, 2, 3, 4, 5, 6}

	switch nums[2] {
	case nums[0], nums[1]:
		fmt.Println("0 or 1")
	case 2, nums[3]:
		fmt.Println("2 or 3")
	case 4, 5, 6:
		fmt.Println("4 or 5 or 6")
	}

	switch nums[2] {
	case nums[0], nums[1], nums[2]:
		fmt.Println("0 or 1 or 2")
	case 2, nums[3]:
		fmt.Println("2 or 3")
	case 4, 5, 6:
		fmt.Println("4 or 5 or 6")
	default:
		fmt.Println("default")
	}
}

func if_condition() {
	if ii := 1; ii < 0 {
		fmt.Println("if ", ii)
	} else if ii == 1 {
		fmt.Println("if == 1", ii)
	} else {
		fmt.Println("if default", ii)
	}
}
