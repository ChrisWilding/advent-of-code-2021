package sonarsweep

func SlidingWindow(depths []int) []int {
	var newDepths []int
	for i := 0; i < len(depths)-2; i++ {
		sum := depths[i] + depths[i+1] + depths[i+2]
		newDepths = append(newDepths, sum)
	}
	return newDepths
}

func SonarSweep(depths []int) int {
	var increases int
	previous := depths[0]
	for i := 1; i < len(depths); i++ {
		next := depths[i]
		if next > previous {
			increases++
		}
		previous = next
	}
	return increases
}
