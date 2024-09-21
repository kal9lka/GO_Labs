package main

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

func main() {
	dataTime := time.Now()
	fmt.Println("Current data is " + dataTime.Format("02-01-2006 15:04:05"))
	fmt.Println("")

	age := 21
	name := "Daniil"
	var PI float64 = 3.1415926
	isVar := true
	fmt.Print("My name is: " + name + ". I am ")
	fmt.Println(age)
	fmt.Print("PI is ")
	fmt.Println(PI)
	fmt.Print("isVar is ")
	fmt.Println(isVar)
	fmt.Println("")

	var a int
	var b int
	fmt.Println("a: ")
	fmt.Scan(&a)
	fmt.Println("b: ")
	fmt.Scan(&b)
	fmt.Print("Результат cуммы: ")
	fmt.Println(a + b)
	fmt.Println("")

	var input string

	fmt.Print("aFloat: ")
	fmt.Scan(&input)
	aFloat, err := strconv.ParseFloat(input, 64)
	if err != nil {
		log.Fatalf("число указано не корректно: %v\n", err)
	}

	fmt.Print("bFloat: ")
	fmt.Scan(&input)
	bFloat, err := strconv.ParseFloat(input, 64)
	if err != nil {
		log.Fatalf("число указано не корректно: %v\n", err)
	}

	fmt.Println(aFloat)
	fmt.Printf("Результат cуммы: %.2f\n", sumFloat(aFloat, bFloat))
	fmt.Printf("Результат разности: %.2f\n", subFloat(aFloat, bFloat))
	fmt.Println("")

	var c int
	fmt.Println("a: ")
	fmt.Scan(&a)
	fmt.Println("b: ")
	fmt.Scan(&b)
	fmt.Println("c: ")
	fmt.Scan(&c)
	fmt.Print("Среднее a, b, c: ")
	fmt.Println(avg(a, b, c))
}

func sumFloat(a float64, b float64) float64 {
	return a + b
}
func subFloat(a float64, b float64) float64 {
	return a - b
}

func avg(a int, b int, c int) float64 {
	result := float64((a + b + c) / 3)
	return result
}
