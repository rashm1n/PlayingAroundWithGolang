package main

import (
	"fmt"
	"net/http"
)

// variable declaration
var i int16

// declaring a struct (type)
type Student struct {
	name string
	age int
	address string
}

// main function
func main() {
	fmt.Printf("Hello World\n")
	myFunc()

	// --------------------------------------------------------------------------------

	// using declared variable
	i += 1
	fmt.Printf("number %d\n",i)

	// declaring and assigning at the same time
	j := "rashmin"
	fmt.Printf("name %s\n",j)

	// --------------------------------------------------------------------------------

	// declaring a Student type
	var student Student

	student = Student{
		name:    "rashmin",
		age:     25,
		address: "Earth",
	}

	student2 := Student{
		name:    "rash2",
		age:     24,
		address: "Mars",
	}

	fmt.Printf("My name is %s\n", student.name)
	fmt.Printf("My name is %s\n", student2.name)

	// --------------------------------------------------------------------------------

	// anonymous struct
	teacher1 := struct {
		name    string
		age     int
		subject string
	}{name: "t1", subject: "maths", age: 29}

	fmt.Printf("Techer %s teachs %s\n\n", teacher1.name, teacher1.subject)

	// --------------------------------------------------------------------------------
	// pointers

	myNum := 3
	pointertest(2,&myNum)

	myNum2 := 5
	incrementNormal(myNum2)
	fmt.Printf("%d\n",myNum2)

	incrementReference(&myNum2)
	fmt.Printf("%d\n",myNum2)

	// --------------------------------------------------------------------------------
	// arrays

	arraytest()
	sliceTest()
	mapTest()

	// --------------------------------------------------------------------------------
	// methods
	application1 := Application{
		appId: 0,
		name:  "app1",
		owner: "rashmin",
	}

	fmt.Println(application1.getId())
}

// go function
func myFunc() {
	fmt.Printf("inside the function\n")
}

func pointertest(a int, b *int) {
	fmt.Printf("number %d\n",a)
	fmt.Printf("number %d\n",&a)
	fmt.Printf("number %d\n",*&a)
	fmt.Printf("number %d\n",b)
	fmt.Printf("number %d\n",*b)
}

// just increments the value
func incrementNormal(a int) {
	a++
}

// gets a pointer to the passed value , dereference it and increments the value in that memory address
func incrementReference(a *int) {
	*a++
}

func arraytest() {
	var myAr1 [10]int

	myAr1[0] = 10
	myAr1[1] = 12
	fmt.Println(myAr1)

	myAr2 := [5] int {1,2,3,4,5}
	fmt.Println(myAr2)
}

func sliceTest() {
	var mySli1 []int
	mySli1 = append(mySli1, 1)
	fmt.Println(mySli1)

	dynamicSlice := make([]int, 0)
	fmt.Println(dynamicSlice)
	dynamicSlice = append(dynamicSlice, 20,20,30)
	fmt.Println(dynamicSlice)
}

func mapTest() {
	var m map[string]string
	m = make(map[string]string)
	m["r0"] = "rash1"
	m["r1"] = "rash2"
	m["r2"] = "rash3"
	fmt.Println(m)

	delete(m,"r1")
	fmt.Println(m)
}

type Application struct {
	appId int
	name string
	owner string
}

func (a Application) getName() string {
	return a.name
}

func (a Application) getId() int {
	return a.appId
}

// dont have to be a struct to declare methods

type blahblah string

func (b blahblah) Handler(r http.ResponseWriter, req *http.Request) {
	fmt.Println("testing 1 2 3 .... \n")
}

// can use pointer recievers for setting values inside the intended type

func (a *Application) setName(n string)  {
	a.name = n
}