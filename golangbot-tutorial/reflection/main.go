package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {
	i := 10
	fmt.Printf("%d %T\n", i, i)

	queryReflection()

	extractValueViaReflection()

	o := order{
		ordID:      456,
		customerID: 56,
	}
	createQuery(o)

	e := employee{
		name:    "Naveen",
		id:      565,
		address: "Coimbatore",
		salary:  90000,
		country: "India",
	}
	createQuery(e)

	v := 90
	createQuery(v)
}

type order struct {
	ordID      int
	customerID int
}

type employee struct {
	name    string
	id      int
	address string
	salary  int
	country string
}

func printReflectionData(q interface{}) {
	t := reflect.TypeOf(q)
	k := t.Kind()
	v := reflect.ValueOf(q)
	fmt.Println("Type ", t)
	fmt.Println("Kind ", k)
	fmt.Println("Value ", v)
}

func queryReflection() {
	o := order{
		ordID:      456,
		customerID: 56,
	}
	printReflectionData(o)
	listStructFieldTypes(o)
}

func createQuery(q interface{}) {
	if reflect.ValueOf(q).Kind() != reflect.Struct {
		fmt.Println("unsupported type")
		return
	}

	t := reflect.TypeOf(q)
	v := reflect.ValueOf(q)

	var names []string
	var values []string
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		names = append(names, t.Field(i).Name)
		switch v.Field(i).Kind() {
		case reflect.Int:
			values = append(values, fmt.Sprintf("%d", field.Int()))
		case reflect.String:
			values = append(values, fmt.Sprintf("\"%s\"", field.String()))
		default:
			fmt.Println("Unsupported type")
			return
		}
	}
	query := fmt.Sprintf("insert into %s(%s) values(%s)", t.Name(), strings.Join(names, ", "), strings.Join(values, ", "))
	query = fmt.Sprintf("%s)", query)
	fmt.Println(query)
}

func listStructFieldTypes(q interface{}) {
	if reflect.ValueOf(q).Kind() == reflect.Struct {
		v := reflect.ValueOf(q)
		fmt.Println("Number of fields", v.NumField())
		for i := 0; i < v.NumField(); i++ {
			fmt.Printf("Field:%d type:%T value:%v\n", i, v.Field(i), v.Field(i))
		}
	}
}

func extractValueViaReflection() {
	a := 56
	x := reflect.ValueOf(a).Int()
	fmt.Printf("type:%T value:%v\n", x, x)
	b := "Naveen"
	y := reflect.ValueOf(b).String()
	fmt.Printf("type:%T value:%v\n", y, y)
}
