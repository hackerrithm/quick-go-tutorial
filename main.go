/*

Hacker: Kemar Galloway
GitHub: www.github.com/hackerrithm
Note: This is my personal tutorial taken from "Learn X in Y minutes"
	  Where X=Go

Check out the tutorial at https://learnxinyminutes.com/docs/go/

*/

package main

import (
	"fmt"
	"os"
	m "math"
	"strconv"
)

func main() {
	fmt.Println("Great work is ... idk")
	beyondHello()
}

func beyondHello() {

	var x int = 6
	y := 4
	sum, prod := learnMultiple(x, y)
	fmt.Println("sum: ", sum, "product: ", prod)
	learnTypes()
}

func learnMultiple(x, y int) (sum, prod int) {
	return x + y, x * y
}

func learnTypes() {
	var arr4 [4]int
	arr3 := [...]int{112, 2334, 5435}

	slice1 := make([]int, 4)
	var slice2 [][]float64
	bs := []byte("a slice")

	s := []int{1, 2, 3}
	s = append(s, 4, 5, 6)
	fmt.Println(s)

	s = append(s, []int{7, 8, 9}...)
	fmt.Print(s, "second arg is a slice literal")

	m := map[string]int{"grade": 3, "followers": 123243}
	m["one"] = 1

	fmt.Println(m["one"])
	_, _, _, _, _ = arr4, arr3, slice1, slice2, bs

	file, _ := os.Create("kemarfile")
	fmt.Fprint(file, "You have done it")
	file.Close()


	learnFlowControl()
}

func expensiveComputation() float64 {
	return m.Exp(10)
}

func learnFlowControl() {
	if false {

	} else {

	}

	x := 42.0
	switch x {
	case 0:
	case 1:
	case 42:
	case 43:
	default:

	}

	for x := 0; x < 3; x++ {
		fmt.Println("iteration", x)
	}

	for key, value := range map[string]int{"one": 1, "two": 2, "three": 3} {
		fmt.Printf("key=%s, value=%d\n", key, value)
	}

	for _, name := range []string{"Kemar", "Ashley", "Emma"} {
		fmt.Printf("Hey %s\n", name)
	}

	if y := expensiveComputation(); y > x {
		x = y
	}

	xBig := func() bool {
		return x > 10000
	}

	x = 99999
	fmt.Println("xBig:", xBig())
	x = 1.3e3
	fmt.Println("xBig:", xBig())


	goto love

love:
	learnFunctionFactory()
	leranDefer()
	learnInterfaces()
}

func learnFunctionFactory() {
	fmt.Println(sentenceFactory("summer")("A beautiful", "day!!"))

	d := sentenceFactory("summer")
	fmt.Println(d("A beautiful", "day"))
	fmt.Println(d("A lazy", "afternoon!"))


}



func sentenceFactory(mystring string) func(before, after string) string {
	return func(before, after string) string {
		return fmt.Sprintf("%s %s %s", before, mystring, after)
	}
}

func leranDefer() (ok bool) {
	defer fmt.Println("LIFO")
	defer fmt.Println("Defer is ")

	return true
}

type Stringer interface {
	String() string
}

type pair struct {
	x, y int
}

func (p pair) String() string {
	return fmt.Sprintf("(%d %d)", p.x, p.y)
}

func learnInterfaces() {
	p := pair{3, 4}
	fmt.Println(p.String())
	var i Stringer
	i = p
	fmt.Println(i.String())

	fmt.Println(p)
	fmt.Println(i)

	learnVariadicParams("kemar", "jane", "kerry")

}

func learnVariadicParams(mystring ...interface{}) {

	for _, param := range mystring {
		fmt.Println("param:", param)
	}

	fmt.Println("params: ", fmt.Sprintln(mystring...))

	learnErrorHandling()
}

func learnErrorHandling() {
	m := map[int]string {3 : "three", 4: "four"}

	if x, ok := m[1]; !ok {
		fmt.Println("No one there")
	} else {
		fmt.Println(x)
	}

	if _, err := strconv.Atoi("non-int"); err != nil {
		fmt.Println(err)
	}
	fmt.Println("ended function----")
	learnConcurrency()
}


func inc(i int, c chan int) {
	c <- i + 1
}

func learnConcurrency() {
	c := make(chan int)

	go inc(0, c)
	go inc(10, c)
	go inc(-100, c)

	fmt.Println(<-c, <-c, <-c)

	cs := make(chan string)
	ccs := make(chan string)

	go func() { c <- 84}()
	go func() { ccs <- "kemoG"}()


	select {
		case i := <-c:
			fmt.Printf("it's a %T", i)
		case <- cs:
			fmt.Println("Its a string")
		case <- ccs:
			fmt.Println("This will not happen as its an empty channel. It is not ready for communication")
	}
}
