/*
go setup

설치 및 확인
brew install go
go version

vscode 도구 업데이트
vscode extension 설치 후 Ctrl+Shift+P
Go: Install/Update tools

패키지 셋업 및 실행
go mod init tutorial
F5 -> 실행
*/
package main

import (
	"bufio"
	"fmt"
	"internal/greeting"
	"os"
)

func main() {
	fmt.Println("Hello Wolrd")

	// 동일 main package
	goodByeWorld()

	// 외부 패키지 함수 실행
	greeting.Hello("John")

	var student1 string = "John"

	// 아래 경우들은 compiler가 타입을 추론
	var student2 = "Jane"
	x := 2

	// 사용하지 않을 변수를 함수내에 선언하면 에러
	// var unusedVariable = ""

	fmt.Println(student1)
	fmt.Println(student2)
	fmt.Println(x)

	fmt.Println(student1, student2)
	fmt.Println(student1, "\n", student2)

	/*
		printf
		%v : value 출력
		%#v : go 포맷에 맞는 value 출력
		%T : 타입 출력
		%% : % 출력
	*/
	fmt.Printf("%v came to see %v\n", student1, student2)
	fmt.Printf("%v came to see %#v\n", student1, student2)
	fmt.Printf("%v came to see %T\n", student1, student2)

	// 초기화만 하는 것도 가능
	var a string
	var b int
	var c bool

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	a = "test"
	fmt.Println(a)

	fmt.Println(testVar())

	multipleVariableDeclaration()

	fmt.Println(PI)

	basicDataTypes()

	conditionsAndLoop()

	b, a = testFunction(1, "1")
	fmt.Println(b)
	fmt.Println(a)

	fmt.Println(factorial_recursion(4))

	testStruct()

	testMap()

	testPointer()

	testInterface()

	testEmptyInterface()

	testPanicRevocer()

	var weirdType test = "hello"

	test.testFunction(weirdType)

}

// var는 외부에서 선언 가능, := 의 경우 불가능
var a int

// b := 3

func testVar() int {
	a = 1
	return a
}

// 여러 변수 한줄에 선언 가능
func multipleVariableDeclaration() {
	var a, b, c, d int = 1, 3, 5, 7

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	var e, f = 6, "Hello"
	g, h := 7, "World!"

	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)

	var (
		i int    = 1
		j string = "hello"
	)

	fmt.Println(i)
	fmt.Println(j)
}

// typed constants
const A int = 1

// untyped constatns
const PI = 3.14

func basicDataTypes() {
	var a bool = true    // Boolean
	var b int = 5        // Integer
	var c float32 = 3.14 // Floating point number
	var d string = "Hi!" // String

	fmt.Printf("Boolean: %v, Type: %T\n", a, a)
	fmt.Printf("Integer: %v, Type: %T\n", b, b)
	fmt.Printf("Float:   %v, Type: %T\n", c, c)
	fmt.Printf("String:  %v, Type: %T\n", d, d)

	/*
		int
		-2147483648 to 2147483647 in 32 bit systems
		-9223372036854775808 to 9223372036854775807

		int8
		-128 to 127

		int16
		-32768 to 32767

		int32
		-2147483648 to 2147483647

		int64
		-9223372036854775808 to 9223372036854775807

	*/
	var x int = 500
	var y int16 = -4500
	fmt.Printf("Type: %T, value: %v\n", x, x)
	fmt.Printf("Type: %T, value: %v\n", y, y)

	/*
		float32
		-3.4e+38 to 3.4e+38

		float64
		-1.7e+308 to +1.7e+308
	*/
	var e float32 = 123.78
	var f float32 = 3.4e+38
	fmt.Printf("Type: %T, value: %v\n", e, e)
	fmt.Printf("Type: %T, value: %v\n", f, f)

	/*
		array
	*/
	// length의 경우 compiler가 추론
	var arr1 = [...]int{1, 2, 3}
	fmt.Println(arr1)

	// length 선언
	cars := [4]string{"Volvo", "BMW", "Ford", "Mazda"}
	fmt.Println(cars)
	fmt.Println(cars[0])

	/*
		slice
		array와 다르게 길이가 가변적
		직접 선언 혹은 make() 함수를 통한 선언
	*/
	arr2 := [4]string{"Volvo", "BMW", "Ford", "Mazda"}
	myslice := []string{"Go", "Slices", "Are", "Powerful"}
	fmt.Printf("myslice = %v\n", myslice)
	fmt.Printf("capacity = %d\n", cap(myslice))
	// 배열에서 자름
	myslice = arr2[2:4]

	fmt.Printf("myslice = %v\n", myslice)
	fmt.Printf("length = %d\n", len(myslice))
	// capacity가 변함
	fmt.Printf("capacity = %d\n", cap(myslice))

	// make 함수의 경우 type, length, capacity로 선언
	myslice1 := make([]int, 5, 10)
	slicePrintFormat("make myslice1", myslice1)

	// with omitted capacity
	myslice2 := make([]int, 5)
	slicePrintFormat("make myslice2", myslice2)

	myslice2 = append(myslice2, 3)
	slicePrintFormat("make myslice2", myslice2)

	myslice2 = append(myslice2, myslice1...)
	slicePrintFormat("make myslice2", myslice2)

	// copy : memory에 올라간 capacity를 줄이는 효과가 있음
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	// Original slice
	slicePrintFormat("original array", numbers)

	// Create copy with only needed numbers
	neededNumbers := numbers[:len(numbers)-10]
	slicePrintFormat("slice original array", neededNumbers)

	numbersCopy := make([]int, len(neededNumbers))
	copy(numbersCopy, neededNumbers)
	slicePrintFormat("copy! capacity가 변함", numbersCopy)
}

func slicePrintFormat(title string, slice []int) {
	fmt.Printf("TITLE : %v\n", title)
	fmt.Printf("slice = %v\n", slice)
	fmt.Printf("length = %d\n", len(slice))
	fmt.Printf("capacity = %d\n", cap(slice))
}

func conditionsAndLoop() {
	// if
	num := 20
	if num >= 10 {
		fmt.Println("Num is more than 10.")
		if num > 15 {
			fmt.Println("Num is also more than 15.")
		}
	} else {
		fmt.Println("Num is less than 10.")
	}

	// switch
	day := 5

	switch day {
	case 1, 3, 5:
		fmt.Println("Odd weekday")
	case 2, 4:
		fmt.Println("Even weekday")
	case 6, 7:
		fmt.Println("Weekend")
	default:
		fmt.Println("Invalid day of day number")
	}

	// for
	adj := [2]string{"big", "tasty"}
	fruits := [3]string{"apple", "orange", "banana"}
	for i := 0; i < len(adj); i++ {
		for j := 0; j < len(fruits); j++ {
			fmt.Println(adj[i], fruits[j])
		}
	}

	// range
	// index, value return
	for idx, val := range fruits {
		fmt.Printf("%v\t%v\n", idx, val)
	}
}

// function
// return할 변수를 미리 정하고 return 명령어만으로 return 가능하며, 여러 parameter를 return 가능
// defer : 함수 호출이 끝나고 실행
func testFunction(x int, y string) (result int, txt1 string) {
	defer fmt.Println("함수 종료")
	fmt.Println("함수 시작")
	result = x + x
	txt1 = y + " World!"
	fmt.Println(result, txt1)
	return
}

// recursion
func factorial_recursion(x int) (y int) {
	if x > 0 {
		y = x * factorial_recursion(x-1)
	} else {
		y = 1
	}
	return
}

// 구조체 Struct
type Person struct {
	name   string
	age    int
	job    string
	salary int
}

func testStruct() {
	var pers1 Person
	var pers2 Person

	// Pers1 specification
	pers1.name = "Hege"
	pers1.age = 45
	pers1.job = "Teacher"
	pers1.salary = 6000

	// Pers2 specification
	pers2.name = "Cecilie"
	pers2.age = 24
	pers2.job = "Marketing"
	pers2.salary = 4500

	// Access and print Pers1 info
	fmt.Println("Name: ", pers1.name)
	fmt.Println("Age: ", pers1.age)
	fmt.Println("Job: ", pers1.job)
	fmt.Println("Salary: ", pers1.salary)

	// Access and print Pers2 info
	fmt.Println("Name: ", pers2.name)
	fmt.Println("Age: ", pers2.age)
	fmt.Println("Job: ", pers2.job)
	fmt.Println("Salary: ", pers2.salary)
}

type test string

func (t test) testFunction() {
	fmt.Println(t)
}

func testMap() {
	// map[key 타입]value 타입 {}
	var a = map[string]string{
		"brand": "Ford",
		"model": "Mustang",
		"year":  "1964",
	}
	fmt.Printf("a\t%v\n", a)

	// make
	a = make(map[string]string)
	fmt.Printf("a\t%v\n", a)

	a["brand"] = "Ford"
	a["model"] = "Mustang"
	a["year"] = "1964"
	fmt.Printf("a\t%v\n", a)

	// update and add
	a["year"] = "1970"
	a["color"] = ""
	fmt.Printf("a\t%v\n", a)

	// delete
	delete(a, "year")
	fmt.Printf("a\t%v\n", a)

	// check
	val1, ok1 := a["brand"]
	val2, ok2 := a["test"]
	val3, ok3 := a["color"]
	_, ok4 := a["model"]

	fmt.Println(val1, ok1)
	fmt.Println(val2, ok2)
	fmt.Println(val3, ok3)
	fmt.Println(ok4)

	// reference
	b := a
	fmt.Println(a)
	fmt.Println(b)

	b["brand"] = "facebook"
	fmt.Println("After change to b:")

	fmt.Println(a)
	fmt.Println(b)

	// iterating
	for k, v := range a {
		fmt.Printf("%v : %v, ", k, v)
	}
	fmt.Println(b)
}

// 포인터
type Data struct {
	value int
	data  [200]int
}

func changeData(arg Data) {
	arg.value = 100
	arg.data[100] = 999
}

func changePData(arg *Data) {
	arg.value = 100
	arg.data[100] = 999
}

func testPointer() {
	var a int = 10
	var p *int

	fmt.Println(a)
	fmt.Println(p)

	p = &a
	fmt.Printf("%v\n", &a)
	fmt.Printf("%v\n", p)

	*p = 20
	fmt.Println(a)
	fmt.Println(*p)

	var data Data
	changeData(data)
	fmt.Printf("\nvalue = %d\n", data.value)
	fmt.Printf("data[100] = %d\n", data.data[100])

	changePData(&data)
	fmt.Printf("value = %d\n", data.value)
	fmt.Printf("data[100] = %d\n", data.data[100])
}

type Human interface {
	Walk() string
}

type Student struct {
	Name string
	Age  int
}

func (s Student) Walk() string {
	return fmt.Sprintf("%s can walk", s.Name)
}

func (s Student) GetAge() int {
	return s.Age
}

func testInterface() {
	s := Student{Name: "John", Age: 20}
	var h Human = s

	fmt.Println(h.Walk())
	// fmt.Println(h.GetAge()) // ERROR
}

func Print(v interface{}) {
	switch v := v.(type) {
	case int:
		fmt.Println("v is int", v)
	case float64:
		fmt.Println("v is float64", v)
	case string:
		fmt.Println("v is string", v)
	default:
		fmt.Printf("Not supported %T:%v\n", v, v)
	}
}

func testEmptyInterface() {
	Print(10)
	Print(3.14)
	Print("Hello word")
	Print(Student{Age: 10})
}

// 에러 핸들링
func ReadFile(filename string) (string, error) {
	file, err := os.Open(filename)

	if err != nil {
		return "", err
	}

	defer file.Close()

	rd := bufio.NewReader(file)
	line, _ := rd.ReadString('\n')

	return line, nil
}

// panic, recover
func firstCall() {
	fmt.Println("(2) firstCall function start")
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in firstCall", r)
		}
	}()

	group()
	fmt.Println("(2) firstCall function end")
}

func group() {
	fmt.Println("(3) group function start")
	fmt.Printf("4/2 = %d\n", divide(4, 2))
	fmt.Printf("4/0 = %d\n", divide(4, 0))
	fmt.Println("(3) group function end")
}

func divide(a, b int) int {
	if b == 0 {
		panic("divide by zero")
	}
	return a / b
}

func testPanicRevocer() {
	fmt.Println("(1) main function start")
	firstCall()
	fmt.Println("(1) main function end")
}
