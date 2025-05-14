package main

import (
	"fmt"
)

func main() {
	var map1 map[string]any // declared but not instantiated
	// can map be nil -->
	if map1 == nil {
		println("nil map")
	}

	map1 = make(map[string]any) // how many buckets are created, most probably 2

	map1["560086"] = "Bangalore-Yashvantpur"
	map1["560096"] = "Bangalore-Rajajinagar"
	map1["560034"] = "Bangalore-Bellandur"
	map1["522001"] = "Guntur-1"
	map1["690051"] = "Trivandrum - Medical College"

	fmt.Println(map1)

	city := map1["690051"] // retrive a value from a map
	println(city)

	map2 := make(map[string]bool)
	map2["Jiten"] = true
	map2["Priya"] = true
	map2["Raju"] = false
	map2["john"] = true

	v := map2["Priya"]
	println(v)

	// v = map2["ANil"]

	v, ok := map2["Rahim"]
	if ok {
		println(v)
	} else {
		println("No value for the key")
	}

	for key, value := range map1 {
		println("Key:", key, "Value:", value)
	}

	fmt.Println(GetKeys(map1))
	fmt.Println(GetValues(map1))

	fmt.Println(GetKeysValues(map1))

	//clear(map1)

	fmt.Println(map1)

	delete(map1, "560086")
	fmt.Println(map1)
	if err := Delete(map1, "560086"); err != nil {
		println(err.Error())
	}

}

// collisions
// overflow
// load factor
// bucket
// hashfunc

func GetKeys(m map[string]any) []string {
	keys := make([]string, len(m))
	i := 0
	for key := range m {
		keys[i] = key
		i++
	}
	return keys
}

func GetValues(m map[string]any) []any {
	var values []any // nil
	for _, value := range m {
		values = append(values, value)
	}
	return values
}

func GetKeysValues(m map[string]any) ([]string, []any) {
	var values []any // nil
	var keys []string
	for key, value := range m {
		values = append(values, value)
		keys = append(keys, key)
	}
	return keys, values
}

func Delete(m map[string]any, key string) error {
	_, ok := m[key]
	if ok {
		delete(m, key)
		return nil
	}
	return fmt.Errorf("key:%s does not exist in the map", key)
}
