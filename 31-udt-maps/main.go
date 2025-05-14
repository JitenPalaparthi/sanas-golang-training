package main

import "fmt"

func main() {

	var m1 MyMap
	m1 = make(MyMap)

	m1["1"] = 10
	m1["2"] = "Hello World"
	m1["3"] = func() {
		println("Hello World")
	}

	keys := m1.GetKeys()
	fmt.Println(keys)
	keys, values := m1.GetKeysValues()
	fmt.Println(keys, values)

	var m2 map[string]any

	m2 = make(map[string]any)

	m2["1"] = 10
	m2["2"] = "Hello World"
	m2["3"] = func() {
		println("Hello World")
	}

	keys = (*MyMap)(&m2).GetKeys()
	fmt.Println(keys)
	keys, values = MyMap(m2).GetKeysValues()
	fmt.Println(keys, values)

}

type MyMap map[string]any

// type myslice []int

func (m *MyMap) GetKeys() []string {
	keys := make([]string, len(*m))
	i := 0
	for key := range *m {
		keys[i] = key
		i++
	}
	return keys
}

func (m MyMap) GetValues() []any {
	var values []any // nil
	for _, value := range m {
		values = append(values, value)
	}
	return values
}

func (m MyMap) GetKeysValues() ([]string, []any) {
	var values []any // nil
	var keys []string
	for key, value := range m {
		values = append(values, value)
		keys = append(keys, key)
	}
	return keys, values
}

func (m MyMap) Delete(key string) error {
	_, ok := m[key]
	if ok {
		delete(m, key)
		return nil
	}
	return fmt.Errorf("key:%s does not exist in the map", key)
}
