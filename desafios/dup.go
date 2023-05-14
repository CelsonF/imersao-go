package main

/*
Dado um array inteiro nums, retorne true se qualquer valor aparecer pelo menos duas vezes no array e retorne false se cada elemento for distinto.
*/
func contemDuplicado(nums []int) bool {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				return true
			}
		}
	}
	return false
}
