package main

import (
	"errors"
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

	fmt.Println("------------------")

	if _, err := checkUserNameExist("foo"); err != nil {
		if isErrUserNameExist(err) {
			fmt.Println(err)
		}
	}

	if _, err := checkUserNameExist("bar"); err != nil {
		if isErrUserNameExist(err) {
			fmt.Println(err)
		} else {
			fmt.Println("IsErrUserNameExist is false")
		}
	}

	fmt.Println("------------------")

	fooSlice := []string{"A", "B"}
	fmt.Println("before fooSlice", fooSlice)
	modify(fooSlice)
	fmt.Println("after fooSlice", fooSlice)

	fmt.Println("------------------")

	fooSlice = []string{"A", "B"}
	fmt.Println("before fooSlice", fooSlice)
	addValue(fooSlice)
	fmt.Println("after fooSlice", fooSlice)

	fmt.Println("------------------")

	fooSlice = []string{"A", "B"}
	bar1 := fooSlice[:1]
	fmt.Println("bar1:", bar1)
	s1 := append(bar1, "C")
	fmt.Println("fooSlice:", fooSlice)
	fmt.Println("s1:", s1)
	s2 := append(bar1, "D")
	fmt.Println("fooSlice:", fooSlice)
	fmt.Println("s2:", s2)
	s3 := append(bar1, "E", "F")
	fmt.Println("fooSlice:", fooSlice)
	fmt.Println("s3:", s3)
}

type TerminateRequestParam struct {
	Svcid string `json:"SVC_ID"`
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

func checkUserNameExist(username string) (bool, error) {
	if username == "foo" {
		return true, errUserNameExist{UserName: username}
	}

	if username == "bar" {
		return true, errors.New("username bar already exist")
	}
	return false, nil
}

type errUserNameExist struct {
	UserName string
}

func isErrUserNameExist(err error) bool {
	_, ok := err.(errUserNameExist)
	return ok
}

func (e errUserNameExist) Error() string {
	return fmt.Sprintf("username %s already exist", e.UserName)
}

func modify(fooSlice []string) {
	fooSlice[1] = "c"
	fmt.Println("modify fooSlice", fooSlice)
}

func addValue(fooSlice []string) {
	fooSlice = append(fooSlice, "C")
	fmt.Println("modify fooSlice", fooSlice)
}
