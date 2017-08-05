package helpers

// https://stackoverflow.com/a/15312097/2908123
func CmpIntSlices(a, b []int) bool {
	if a == nil && b == nil {

		return true
	}

	if a == nil || b == nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func CmpStringSlices(a, b []string) bool {
	if a == nil && b == nil {

		return true
	}

	if a == nil || b == nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
