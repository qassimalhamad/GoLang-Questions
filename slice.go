package piscine

func Slice(a []string, nbrs ...int) []string {
	if len(nbrs) == 0 {
		return nil
	}

	start := nbrs[0]
	end := len(a)

	if len(nbrs) == 2 {
		end = nbrs[1]
	}

	if start < 0 {
		start = start + len(a)
	}

	if end < 0 {
		end = end + len(a)
	}

	if start < 0 || end < 0 || start > end || start > len(a) || end > len(a) {
		return nil
	}

	return a[start:end]
}
