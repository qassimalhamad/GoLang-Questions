package piscine

func CanJump(nums []uint) bool {
	if len(nums) == 0 {
		return false
	}
	if len(nums) == 1 {
		return true
	}

	pos := 0

	for {
		// Get the exact step to move from current position
		steps := int(nums[pos])
		pos += steps

		if pos == len(nums)-1 {
			return true // reached the last index exactly
		}
		if pos >= len(nums) || steps == 0 {
			return false // out of bounds or stuck
		}
	}
}
