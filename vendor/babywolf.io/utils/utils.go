package utils

type Any interface{}

func GetValue(array [1][52]map[string]Any, i, j int, key string) (interface{}, bool) {
	if i < 0 || i >= 1 || j < 0 || j >= 52 {
		return nil, false
	}
	value, exists := array[i][j][key]
	return value, exists
}
