package main

import (
	"encoding/hex"
	"errors"
	"fmt"
	"unsafe"
)

// xu ly phan thap phan ntn
// float64

// convert between type

type StructA struct {
	name string
}

func sayHelloValue(name StructA) {
	fmt.Printf("Hello %s from value\n", name.name)
	name.name = "Dave"
}

func sayHelloReference(name *StructA) {
	fmt.Println(name)
	fmt.Printf("Hello %s from reference\n", name.name)
	name.name = "Dave"
}

func floatingPointer() {
	a := 2.1
	b := 4.2

	c := a + b
	d := 2.1 + 4.2
	e := float64(2.1) + float64(4.2)

	fmt.Println(c == d)
	fmt.Println(c == e)
}

// interface
type Eyes struct {
	Color string
	Shape string
}

type Human struct {
	Eyes
}

type Dog struct {
	Eyes Eyes
}

type Printer struct{}

type Walker interface {
	Walk()
}

type OrderStatus string

const (
	OrderCreated OrderStatus = "created" 
)
func checkStatus(order OrderStatus) {

}

// receiver + function
// receiver
func (d *Dog) Walk() {
	checkStatus(OrderCreated)
	d.Eyes.Color = "Black"
}

func (d Dog) walk1() {
	fmt.Println(d.Eyes.Color)
}

func structComposition() {
	var human Human
	var dog Dog

	// promoted field access
	human.Color = "Blue"
	human.Shape = "Round"

	// long form field access
	human.Eyes.Color = "Brown"
	fmt.Println(human, dog)
}

func helloRune() {
	greet1 := "hello"
	greet2 := "hello🚀"

	fmt.Println(len(greet1))
	fmt.Println(len(greet2))

	myString := "🚀"
	fmt.Println("Unicode codepoint represented by rune:", []rune(myString))
	fmt.Println("UTF-8 code represented by up to 4 bytes:", []byte(myString))
	fmt.Println("UTF-8 code represented as Hexadecimal:", hex.EncodeToString([]byte(myString)))
	fmt.Println("length", len(myString))
}

type Suboptimal struct {
	bool1 bool
	int1  int
	bool2 bool
	int2  int
	bool3 bool
	int3  int
	bool4 bool
	int4  int
}

type Optimal struct {
	bool1 bool
	bool2 bool
	bool3 bool
	bool4 bool
	int1  int
	int2  int
	int3  int
	int4  int
}

func padding() {
	subOptimal := [1000]Suboptimal{}
	sizeSuboptimal := float64(unsafe.Sizeof(subOptimal))

	optimal := [1000]Optimal{}
	sizeOptimal := float64(unsafe.Sizeof(optimal))

	fmt.Printf("Size of suboptimal array = %v bytes\n", sizeSuboptimal)
	fmt.Printf("Size of optimal array = %v bytes\n", sizeOptimal)

	diff := (sizeSuboptimal - sizeOptimal) / sizeSuboptimal * 100
	fmt.Printf("Padding means we're using %d%% more memory...", int(diff))
}

func mapFn() {
	// declaration followed my initialisation with make
	var myMap map[string]string

	myMap = make(map[string]string)

	// initialisation with make
	myMap2 := make(map[string]string)

	// initialisation as empty map literal
	myMap3 := map[string]string{}

	// initialisation as map literal with key/values
	myMap4 := map[string]string{
		"key": "value",
	}

	fmt.Println(myMap, myMap2, myMap3, myMap4)
}

func mapAccess() {
	var myMap = make(map[string]string)

	// create three keys in the map
	myMap["key1"] = "value1"
	myMap["key2"] = "value2"
	myMap["key3"] = "value3"
	myMap["key4"] = "value4"
	myMap["key5"] = "value5"

	// obtain the number of key/values
	fmt.Println("Map keys:", len(myMap))

	// access the value stored for a key
	fmt.Println(myMap["key1"])

	// iterate over the map with range and print each key/value
	for key, value := range myMap {
		fmt.Println(key, value)
	}
	// psedu random -> random access

	// delete key3
	delete(myMap, "key3")

	fmt.Println(myMap)
}

// mutex
func safetyAccessMap() {
	activeUser := map[string]bool{
		"Joe Blogs":  true,
		"Dave Blogs": false,
	}

	// keys exists no issue
	fmt.Println("Joe is active:", activeUser["Joe Blogs"])
	fmt.Println("Dave is active:", activeUser["Dave Blogs"])
	// key does not exist
	fmt.Println("Trevor is active:", activeUser["Trevor Blogs"])

	// check key active
	if active, ok := activeUser["Trevor Blogs"]; !ok {
		fmt.Println("User does not exist")
	} else {
		fmt.Println("Trevor is active:", active)
	}
}

func creatingSlice() {
	var sl1 []int
	sl2 := []int{}
	sl3 := []int{1, 2, 3, 4, 5}
	sl4 := make([]int, 5)
	sl5 := make([]int, 1, 4)

	// check length
	fmt.Println(len(sl1), len(sl2), len(sl3), len(sl4), len(sl5))

	// check capacity
	fmt.Println(cap(sl1), cap(sl2), cap(sl3), cap(sl4), cap(sl5))

	// nilslice

	if sl1 == nil {
		fmt.Println("sl1 is nil!")
	}

	if sl2 != nil && sl3 != nil && sl4 != nil && sl5 != nil {
		fmt.Println("sl2, sl3, sl4 and sl5 are not nil!")
	}
}

func resizingSlice() {
	// https://github.com/golang/go/blob/master/src/runtime/slice.go
	// amend this start capacity to see the impact
	cp := 100
	sl := make([]int, 0, cp)
	els := 1000
	addr := fmt.Sprintf("%p", sl)

	for i := 0; i < els; i++ {
		sl = append(sl, i)
		if fmt.Sprintf("%p", sl) != addr {

			// when we detect address change inspect
			extra := (float64(cap(sl)) - float64(cp)) / float64(cp) * 100
			fmt.Printf("length: %d, capacity: %d, address: %p. Additional capacity: %0.2f%%\n", len(sl), cap(sl), sl, extra)
			cp = cap(sl)
			addr = fmt.Sprintf("%p", sl)
		}
	}

	fmt.Println()

	overCap := (float64(cap(sl)) - float64(els)) / float64(els) * 100
	fmt.Printf("Backing array over capacity: %0.2f%%\n", overCap)
}

func isEvenLessThan(n int, c int) (bool, bool) {
	var even, less bool

	if n%2 == 0 {
		even = true
	}
	if n < c {
		less = true
	}

	return even, less
}

func shortFormIf() {
	if even, _ := isEvenLessThan(2, 10); even {
		fmt.Println("Result was even")
	}
}

func switchCase() {
	myName := "Joe Blogs"

	switch myName {
	case "Joe Blogs":
		fmt.Println("Hi Joe!")
	case "Dave Blogs":
		fmt.Println("Hi Dave!")
	default:
		fmt.Println("Hi there!")
	}
}

func shortFormSwitchCase() {
	switch letter := "z"; letter {
	case "a", "e", "i", "o", "u":
		fmt.Println("letter was a vowel")
	default:
		fmt.Println("letter was not a vowel")
	}
}

func iterationLogic() {
	// this will timeout on the playground
	for {
		fmt.Println("infinite loop")
	}

	// three component
	for i := 0; i < 5; i++ {
		fmt.Println("iteration", i+1)
	}

	// while equivalent
	n := 1
	for n <= 5 {
		fmt.Println("Iteration", n)
		n++
	}

	// do while equivalent
	for ok := true; ok; ok = n != 5 {
		fmt.Println("Iteration", n)
		n++
	}

	// foreach
	sl := []int{1, 2}
	mp := map[string]string{"key1": "value1", "key2": "value2"}

	// slice/array
	for index, value := range sl {
		fmt.Printf("index: %d, value: %d\n", index, value)
	}

	// map
	for key, value := range mp {
		fmt.Printf("key: %s, value: %s\n", key, value)
	}

	_, er := dosomething()
	if er != nil {
		
	} 
}

func dummyFunctionError() {
	err := errors.New("dummy unhandlable error")
	if err != nil {
		panic(err)
	}
}

func panicRecover() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recovered from panic:", err)
		}
	}()

	dummyFunctionError()
}

type A struct {
	name string
}

type Test struct {
	A A       `json:a,omitempty`
	B string  `json:b,omitempty`
	C *string `json:c,omitempty`
}

// model, DTO

func main() {
	// sayHelloReference(&name)
	// fmt.Println(name)

	// sayHelloReference(&name)
	// fmt.Println(name)

	// floatingPointer()

	// structComposition()

	// padding()

	// mapFn()

	mapAccess()

	// creatingSlice()

	// resizingSlice()
}
