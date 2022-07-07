package main

import (
	"fmt"
)

type Human struct {
	name string
	age  int
}

func (h Human) Say() {
	fmt.Printf("I can say, I am %s, %d years old\n", h.name, h.age)
}

func (h Human) Run() {
	fmt.Printf("%s is running\n", h.name)
}

type Singer struct {
	Human
	action string
}

func (s Singer) Sing() {
	fmt.Printf("%s can sing %s\n", s.name, s.action)
}

type Man interface {
	Say()
	Run()
}

func main() {
	tom := Singer{
		Human: Human{
			name: "tom",
			age:  12,
		},
		action: "sing",
	}

	divd := Human{
		name: "Divd",
		age:  12,
	}

	// tom.Say()
	// tom.Run()
	// tom.Sing()

	// divd.Say()
	// divd.Run()

	test := []Man{}
	test = append(test, tom, divd)

	for _, value := range test {
		value.Say()
		value.Run()
	}

	fmt.Println("-------------------")

	var foo Foo = 1000
	fmt.Printf("%d\n", foo)
	fmt.Println(foo)

	fmt.Println("-------------------")

	temp := []interface{}{}

	a := 1
	b := "foo"
	c := Human{
		name: "bar",
		age:  12,
	}
	temp = append(temp, a, b, c)

	for i, value := range temp {
		if v, ok := value.(int); ok {
			fmt.Printf("%d,value: %v\n", i, v)
		} else if v, ok := value.(string); ok {
			fmt.Printf("%d,value: %s\n", i, v)
		} else if v, ok := value.(Human); ok {
			fmt.Printf("%d,value: %s\n", i, v.name)
		}
	}

	for i, value := range temp {
		switch v := value.(type) {
		case int:
			fmt.Printf("%d,value: %v\n", i, v)
		case string:
			fmt.Printf("%d,value: %s\n", i, v)
		case Human:
			fmt.Printf("%d,value: %s\n", i, v.name)
		}
	}

}

type Foo int

func (f Foo) String() string {
	return "foo"
}
