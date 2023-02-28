package main

import "log"

type Set map[any]*bool

func main () {
	arr1A, arr1B := []interface{}{4,2,1,6}, []interface{}{3,6,9,2,10}
	intersection1 := intersection(arr1A, arr1B)
	log.Printf("The intersection of %v and %v is %v", arr1A, arr1B, intersection1)

	arr2A, arr2B := []interface{}{2,4,6}, []interface{}{4,2}
	intersection2 := intersection(arr2A, arr2B)
	log.Printf("The intersection of %v and %v is %v", arr2A, arr2B, intersection2)

	arr3A, arr3B := []interface{}{0,1,2}, []interface{}{10,11}
	intersection3 := intersection(arr3A, arr3B)
	log.Printf("The intersection of %v and %v is %v", arr3A, arr3B, intersection3)

	arr4A, arr4B := []interface{}{4,2,1}, []interface{}{1,2,4,6}
	intersection4 := intersection(arr4A, arr4B)
	log.Printf("The intersection of %v and %v is %v", arr2A, arr2B, intersection4)
}

// Write a function, intersection, that takes in two lists, a,b, as arguments. The function should return a new list containing elements that are in both of the two lists.
// By using a set-like structure, we reduce the runtime to O(n+m)
// intersection({4,2,1,6}, {3,6,9,2,10}) # -> {2,6}
// intersection([4,2,1], [1,2,4,6]) # -> [1,2,4]
// intersection([0,1,2], [10,11]) # -> []

func intersection(a []interface{}, b []interface{}) []interface{} {
	res := []interface{}{}
	setA := Set{}
	for _, el := range a {
		isIncluded := true
		setA[el] = &isIncluded
	}
	for _, el := range b {
		if setA[el] != nil && *setA[el] == true {
			res = append(res, el)
		}
	}
	return res
}