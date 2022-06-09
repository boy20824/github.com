package main

import (
	"fmt"
	"strings"

	"github.com/go-training/test/helloworld"
)

const (
	Monday = iota + 1
	Tuesday
	Wednesday
	Thusdaty
)

func main() {

	//變數宣告
	a := "123"
	var b string
	b = "456"
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println("------------------")
	fmt.Println(helloworld.HelloWorld())
	fmt.Println("------------------")
	fmt.Println(Monday)
	fmt.Println(Tuesday)
	fmt.Println(Wednesday)
	fmt.Println(Thusdaty)
	fmt.Println("@@@@@@@@@@@@@")
	//func 使用
	fmt.Println(add(1, 2))

	i := 1
	j := 2
	i, j = swap(i, j)
	fmt.Println("------------------")
	fmt.Println(i)
	fmt.Println(j)

	fmt.Println("------------------")
	i, j = j, i //這樣就可以交換
	fmt.Println(i)
	fmt.Println(j)

	fmt.Println("@@@@@@@@@@@@@")

	bar := foo()
	fmt.Printf("%T\n", bar)
	fmt.Println(bar())

	fmt.Println("------------------")

	bar2 := func(i, j float32) float32 {
		return i + j
	}
	fmt.Printf("%T\n", bar2)
	fmt.Println(bar2(1, 2))

	fmt.Println("------------------")

	foo := func() string {
		return "Hello World"
	}
	fmt.Println(foo())

	foo2 := func() {
		fmt.Println("Hello World 2")
	}

	foo2()

	func() {
		fmt.Println("Hello World 3")
	}()

	//goroutine
	go func(i, j int) {
		fmt.Println(i + j)
	}(1, 2)

	fmt.Println("------------------")

	fmt.Println(getUserListOptsSQL(searchOpts{
		username: "applyboy",
	}))
	fmt.Println(getUserListOptsSQL(searchOpts{
		username: "applyboy",
		email:    "test@gmail.com",
	}))

}

func add(i, j int) int {
	return i + j
}

func swap(i, j int) (int, int) {
	return j, i
}

func foo() func() int {
	return func() int {
		return 100
	}
}

type searchOpts struct {
	username string
	email    string
}

func getUserListOptsSQL(opts searchOpts) string {
	sql := "select * from user"
	where := []string{}

	if opts.username != "" {
		where = append(where, fmt.Sprintf("username ='%s'", opts.username))
	}

	if opts.email != "" {
		where = append(where, fmt.Sprintf("email ='%s'", opts.email))
	}

	return sql + "where" + strings.Join(where, " or ")
}
