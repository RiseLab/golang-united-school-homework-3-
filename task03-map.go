package homework

import "sort"

func sortMapValues(input map[int]string) (result []string) {
	result = make([]string, len(input))

	keys := make([]int, len(input))
	i := 0

	for k := range input {
		keys[i] = k
		i++
	}

	sort.Ints(keys)

	for k, v := range keys {
		result[k] = input[v]
	}

	return
}
