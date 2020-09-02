package repository

type stringDiff struct {
	Inc []string
	Dec []string
}

func stringArrayDiff(before []string, after []string) stringDiff {
	return stringDiff{
		Inc: stringArraySub(after, before),
		Dec: stringArraySub(before, after),
	}
}

func stringArraySub(a []string, b []string) []string {
	r := []string{}
	m := make(map[string]bool)
	for _, v := range b {
		m[v] = true
	}
	for _, v := range a {
		if _, ok := m[v]; !ok {
			r = append(r, v)
		}
	}
	return r
}
