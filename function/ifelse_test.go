package function

func ifelse(start, end int) int {
	if distance := end - start; distance > 100 {
		return distance
	} else if distance > 60 {
		return distance
	} else {
		return start
	}
}
