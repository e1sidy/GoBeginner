// Checking Existence of keys and Deleting key value pairs

package main

func main() {

	map1 := make(map[string]int)
	map1["Delhi"] = 31
	map1["Lko"] = 24
	map1["Mbai"] = 29

	// Checking existence of a key and extracting a value
	val, isPresent := map1["Delhi"]
	println(val, isPresent)
	// if a value does not exist then the returns false, default value
	println(map1["Bejing"])
	// Deleting a key value pair
	delete(map1, "Delhi")
	println(map1["Delhi"])

}
