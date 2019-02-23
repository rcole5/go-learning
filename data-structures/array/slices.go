package slices

func Copy(source *[]string, target *[]string) {
	for i, v := range *source {
		(*target)[i] = v
	}
}

func Append(target *[]string, value string) {
	// Create new slice with of size len + 1
	t := make([]string, len(*target) + 1)

	// Copy values from old slice into new slice
	Copy(target, &t)

	// Add the new value
	t[len(t) - 1] = value

	// Set the target to the new slice
	*target = t
}

func Delete(target *[]string, index int) {
	// Create new slice with of size len - 1
	t := make([]string, len(*target) - 1)

	// Copy all values except the one we want to remove
	for i, v := range *target {
		if i == index {
			// Skip the value we want to remove
			continue
		} else if i < index {
			// If index is less than the removed index the index will be the same
			t[i] = v
		} else {
			// If index is greater than the removed index it will be offset by 1
			t[i - 1] = v
		}
	}

	// Set the target to the new slice
	*target = t
}
