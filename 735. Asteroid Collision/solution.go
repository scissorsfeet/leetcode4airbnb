package _35__Asteroid_Collision

func asteroidCollision(asteroids []int) []int {
	res := make([]int, len(asteroids))
	pos := 0

	for _, num := range asteroids {
		if pos == 0 {
			res[pos] = num
			pos++
		} else {
			if num > 0 {
				res[pos] = num
				pos++
			} else {
				if res[pos-1] < 0 {
					res[pos] = num
					pos++
				} else {
					pre := pos - 1
					outfromneg := false
					for pre >= 0 && res[pre] > 0 && num < 0 {
						if abs(num) > res[pre] {
							res[pre] = num
							pre--
							outfromneg = true
						} else {
							if abs(num) == res[pre] {
								pre--
							}
							outfromneg = false
							break
						}
					}
					if outfromneg {
						pos = pre + 2
					} else {
						pos = pre + 1
					}
				}
			}
		}
	}
	return res[0:pos]
}

func abs(x int) int {
	if x < 0 {
		return 0 - x
	}
	return x
}
