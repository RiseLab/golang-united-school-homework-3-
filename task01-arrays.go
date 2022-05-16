package homework

func average(input [15]float32) (result float32) {
	var sum, num float32

	for _, val := range input {
		if val == 0 {
			break
		}
		sum += val
		num++
	}

	if num != 0 {
		result = sum / num
	}

	return
}
