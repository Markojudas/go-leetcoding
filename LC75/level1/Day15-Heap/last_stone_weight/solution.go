package main

func lastStoneWeight(stones []int) int {

	for {
		if len(stones) == 0 {
			return 0
		}
		if len(stones) == 1 {
			return stones[0]
		}

		x, y, index := 0, 0, 0

		for idx, val := range stones {
			if val > y {
				y = val
				index = idx
			}
		}
		stones = append(stones[:index], stones[index+1:]...)

		for idx, val := range stones {
			if val > x {
				x = val
				index = idx
			}
		}
		stones = append(stones[:index], stones[index+1:]...)

		if y != x {
			stones = append(stones, y-x)
		}
	}
}
