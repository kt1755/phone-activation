package tiki

func QuickSort(data []PhoneUsage) []PhoneUsage {
	if len(data) < 2 {
		return data
	}

	left, right := 0, len(data)-1

	var pivot int
	if (right+left)%2 == 0 {
		pivot = (right + left) / 2
	} else {
		pivot = (right + left - 1) / 2
	}

	data[pivot], data[right] = data[right], data[pivot]

	for i, _ := range data {
		if data[i].Duration.ActivationDate.Before(data[right].Duration.ActivationDate) {
			data[left], data[i] = data[i], data[left]
			left++
		}
	}

	data[left], data[right] = data[right], data[left]
	QuickSort(data[:left])
	QuickSort(data[left+1:])

	return data
}
