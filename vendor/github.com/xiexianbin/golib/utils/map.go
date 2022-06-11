// Copyright 2020 xiexianbin<me@xiexianbin.cn>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// https://github.com/xiexianbin/gsync/blob/7431fdacf3/utils/map_utils.go

package utils

// UnionMap return union map
func UnionMap(m1, m2 map[string]string) map[string]string {
	m := make(map[string]string)
	for k1, v1 := range m1 {
		for k2, v2 := range m2 {
			if k1 == k2 && v1 == v2 {
				m[k1] = v1
			}
		}
	}

	return m
}

// DiffMap two map diff
func DiffMap(m1, m2 map[string]interface{}) (map[string]interface{}, map[string]interface{}, map[string]interface{}) {
	justM1 := make(map[string]interface{})
	diffM1AndM2 := make(map[string]interface{})
	justM2 := make(map[string]interface{})
	for k1, v1 := range m1 {
		v2 := m2[k1]
		// just in m1
		if v2 == nil {
			justM1[k1] = v1
		}

		// key both in m1 and m2, but value is diff
		if v2 != nil && v1 != v2 {
			diffM1AndM2[k1] = v1
		}
	}
	for k2, v2 := range m2 {
		v1 := m1[k2]
		// just in m2
		if v1 == nil {
			justM2[k2] = v2
		}
	}

	return justM1, justM2, diffM1AndM2
}
