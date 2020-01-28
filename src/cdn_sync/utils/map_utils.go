package utils

func UnionMap(m1, m2 map[string]string) (map[string]string, error) {
	m := make(map[string]string)
	for k1, v1 := range m1 {
		for k2, v2 := range m2 {
			if k1 == k2 && v1 == v2 {
				m[k1] = v1
			}
		}
	}

	return m, nil
}

// two map diff
func DiffMap(m1, m2 map[string]interface{}) (map[string]interface{}, map[string]interface{}, map[string]interface{}, error) {
	justM1 := make(map[string]interface{})
	diffM1AndM2 := make(map[string]interface{})
	justM2 := make(map[string]interface{})
	for k1, v1 := range m1 {
		v2 := m2[k1]
		// just in m1
		if v2 == "" {
			justM1[k1] = v1
		}

		// key both in m1 and m2, but value is diff
		if v2 != "" && v1 != v2 {
			diffM1AndM2[k1] = v1
		}
	}
	for k2, v2 := range m2 {
		v1 := m1[k2]
		// just in m2
		if v1 == "" {
			justM2[k2] = v2
		}
	}

	return justM1, justM2, diffM1AndM2, nil
}
