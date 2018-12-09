package array

func Unique(a []interface{}) []interface{} {
	m := make(map[interface{}]struct{})
	u := []interface{}{}
	for _, e := range a {
		if _, in := m[e]; in {
			continue
		}
		m[e] = struct{}{}
		u = append(u, e)
	}
	return u
}
