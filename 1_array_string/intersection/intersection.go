package main

type Set map[any]*bool

func main () {}

// Using Map - O(n+m)) Runtime Where n and m are the length of each slice
func intersection(a []interface{}, b []interface{}) []interface{} {
	res := []interface{}{}
	setA := Set{}
	for _, el := range a {
		isIncluded := true
		setA[el] = &isIncluded
	}
	for _, el := range b {
		if setA[el] != nil && *setA[el] {
			res = append(res, el)
		}
	}
	return res
}