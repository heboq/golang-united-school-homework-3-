package homework

func average(input [15]float32) (result float32) {
	var sum, n float32
	for _, v := range input {
		sum += v
		if v != 0 {
			n++
		}
	}
	return sum / n
}
