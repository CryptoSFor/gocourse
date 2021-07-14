package arrays

func AverageValue(arr [5]float64) (sum float64) {
	for _, v := range arr {
		sum += v
	}
	length := float64(len(arr))
	sum = sum / length
	return
}
