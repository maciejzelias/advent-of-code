package main

func main() {
	time := []int{49, 97, 94, 94}
	distance := []int{263, 1532, 1378, 1851}
	multiply_sum := 1

	// part one
	for i := 0; i < len(time); i++ {
		var start int
		var end int
		for j := 0; j <= time[i]; j++ {
			left := time[i] - j

			if left*j > distance[i] {
				start = j
				break
			}
		}

		for j := time[i]; j > 0; j-- {
			left := time[i] - j

			if left*j > distance[i] {
				end = j
				break
			}
		}

		diff := end - start + 1
		multiply_sum *= diff
	}
	println(multiply_sum)

	// part 2

	combinedTime := 49979494
	combinedDistance := 263153213781851

	var start int
	var end int

	for j := 0; j <= combinedTime; j++ {
		left := combinedTime - j

		if left*j > combinedDistance {
			start = j
			break
		}
	}

	for j := combinedTime; j > 0; j-- {
		left := combinedTime - j

		if left*j > combinedDistance {
			end = j
			break
		}
	}

	diff := end - start + 1
	println(diff)
}
