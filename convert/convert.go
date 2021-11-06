package convert

import "strings"

func ToCamelCase(m map[string]string) map[string]string {
	if m == nil {
		return nil
	}
	p := make(map[string]string)
	for key, element := range m {
		n := ToCamel(key)
		p[n] = element
	}
	return p
}
func ToCamel(s string) string {
	s2 := strings.ToUpper(s)
	s1 := string(s[0])
	for i := 1; i < len(s); i++ {
		if string(s[i-1]) == "_" {
			s1 = s1[:len(s1)-1]
			s1 += string(s2[i])
		} else {
			s1 += string(s[i])
		}
	}
	return s1
}
