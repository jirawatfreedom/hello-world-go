package main

import (
	"fmt"

	"github.com/jirawatfreedom/hello-world-go"
)

type Person struct {
	firstName string
	lastName  string
	age       int
}

func (p *Person) getAge() int {
	return p.age
}
func (p *Person) getLastname() string {
	return p.lastName
}

func main() {
	a := 1
	b := 2
	if a == b {
		fmt.Println("a == b is true")
	} else {
		fmt.Println("a == b is false")
	}
	var array_a [2]string
	array_a[0] = "one"
	array_a[1] = "two"
	fmt.Println(array_a)

	var colors = make([]string, 0)
	colors = append(colors, "ONE")
	colors = append(colors, "TWO")
	colors = append(colors, "THREE")
	colors = append(colors, "FOUR")
	colors = append(colors, "FIVE")
	colors = append(colors, "SIX")
	printSlice("colors", colors)

	fmt.Println(sum(a, b))

	p1 := Person{
		firstName: "Jirawat",
		lastName:  "Muabkhuntod",
		age:       24,
	}
	fmt.Println(p1)
	fmt.Println(p1.firstName)
	fmt.Println(p1.lastName)
	fmt.Println(p1.age)
	fmt.Println(p1.getAge())

	animals := map[string]string{
		"D1": "Dog",
		"D2": "Cat",
		"D3": "Chopper",
	}
	fmt.Println(animals)
	fmt.Println(animals["D1"])

	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	x := 1
	switch x {
	case 1:
		fmt.Println("x = 1")
	case 2:
		fmt.Println("x = 2")
	default:
		fmt.Println("x = unknow")
	}

	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	fmt.Println(color.getRed())

}
func printSlice(s string, x []string) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
func sum(a int, b int) int {
	return a + b
}
