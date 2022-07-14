package main

func twoSum(nums []int, target int) []int {
	//Create a map. Key = nums[i]; Val = i
	hmap := map[int]int{}

	//range over the slice
	for idx, val := range nums {

		//check if the key (compliment = target - value): value has been mapped
		if _, mapped := hmap[target-val]; mapped {
			//return the nums[idx] and the value of the map as int slice - go ok idiom
			return []int{idx, hmap[target-val]}
		}

		// if the key (target - value): value hasn't been mapped, then map it!
		hmap[val] = idx
	}

	//if not found return an empty int slice
	return []int{}
}
