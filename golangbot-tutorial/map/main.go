package main

import (
	"fmt"
)

func main() {
	personSalary := make(map[string]int)
	personSalary["steve"] = 12000
	personSalary["jamie"] = 15000
	personSalary["mike"] = 9000
	fmt.Println("personSalary map contents:", personSalary)

	personSalary2 := map[string]int{
		"steve": 12000,
		"jamie": 15000,
	}
	personSalary2["mike"] = 9000
	fmt.Println("personSalary2 map contents:", personSalary2)

	employee := "jamie"
	fmt.Println("salary of", employee, "is", personSalary[employee])

	newEmployee := "joe"
	fmt.Println("salary of", newEmployee, "is", personSalary[newEmployee])
	value, ok := personSalary[newEmployee]
	if ok {
		fmt.Println("salary of", newEmployee, "is", value)
	} else {
		fmt.Println(newEmployee, "not found")
	}

	fmt.Println("All items of a map")
	for key, value := range personSalary {
		fmt.Printf("personSalary[%s] = %d\n", key, value)
	}

	delete(personSalary, "steve")
	fmt.Println("map after deletion", personSalary)
	fmt.Println("length is", len(personSalary))

	mapReference := personSalary
	mapReference["mike"] = 18000
	fmt.Println("person salary changes", personSalary)

	map1 := map[string]string{
		"algeria":  "algiers",
		"botswana": "gaborone",
		"chad":     "n'djamena",
		"djibouti": "djibouti",
	}
	map2 := map[string]string{
		"armenia": "yerevan",
		"belgium": "brussels",
		"croatia": "zagreb",
	}
	map3 := copyStringMap(map1)
	fmt.Println("map1 == map2", areStringMapsEqual(map1, map2))
	fmt.Println("map1 == map3", areStringMapsEqual(map1, map3))
	fmt.Println("map2 == map3", areStringMapsEqual(map2, map3))
}

func areStringMapsEqual(first map[string]string, second map[string]string) bool {
	if first == nil || second == nil {
		return first == nil && second == nil
	}

	if len(first) != len(second) {
		return false
	}

	for key, value := range first {
		if value != second[key] {
			return false
		}
	}
	return true
}

func copyStringMap(source map[string]string) map[string]string {
	target := make(map[string]string)
	for key, value := range source {
		target[key] = value
	}
	return target
}
