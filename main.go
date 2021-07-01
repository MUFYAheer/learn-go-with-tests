package main

import (
	"fmt"

	arraysandslices "github.com/MUFYAheer/learn-go-with-tests/arrays_and_slices"
	"github.com/MUFYAheer/learn-go-with-tests/integers"
	"github.com/MUFYAheer/learn-go-with-tests/iteration"
	structsmethodsandinterfaces "github.com/MUFYAheer/learn-go-with-tests/structs_methods_and_interfaces"
)

const (
	spanish = "Spanish"
	french  = "French"
)

const (
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
)

// Hello takes `name` and `language` and returns Greeting to `name` in `language`
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("World", ""))
	fmt.Println("2 + 3 =", integers.Add(2, 3))
	fmt.Println("Repeat `b` 5 times:", iteration.Repeat("b", 5))

	numbers := []int{1, 2, 3, 4}
	fmt.Printf("Sum of %v is %d\n", numbers, arraysandslices.Sum(numbers))
	fmt.Println(arraysandslices.SumAll([]int{3, 4}, []int{1, 2}, []int{1, 2, 3, 4}))

	rectangle := structsmethodsandinterfaces.Rectangle{Width: 4, Height: 5}
	fmt.Printf("Area of rectangle %v is %.2f\n", rectangle, rectangle.Area())
	circle := structsmethodsandinterfaces.Circle{Radius: 10}
	fmt.Printf("Area of circle %v is %.2f\n", circle, circle.Area())
	triangle := structsmethodsandinterfaces.Triangle{Base: 4, Height: 5}
	fmt.Printf("Area of triangle %v is %.2f\n", triangle, triangle.Area())
}
