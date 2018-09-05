package main

import (
	"fmt"
	"strconv"
	"math"
	"github.com/pkg/errors"
	"os"
	"io"
)

func sum(values [] int, resultChan chan int) {
	sum := 0
	for _, value := range values {
		sum += value
	}
	resultChan <- sum
}

func paracalc() {
	values := [] int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	resultChan := make(chan int, 2)
	go sum(values[:len(values)/2], resultChan)
	go sum(values[len(values)/2:], resultChan)
	sum1, sum2 := <-resultChan, <-resultChan

	fmt.Println("Result:", sum1, sum2, sum1+sum2)
}

func run(arg string) {
	fmt.Println(arg)
}

func getName() (firstName, middleName, lastName, nickName string) {
	return "tan", "zhong", "tao", "chris";
}

type Bird struct {
	Name     string
	Category string
	Weight   int
}

type IFly interface {
	Fly()
}

func (b *Bird) Fly() {
	println("bird fly! name:" + b.Name + " weight:" + strconv.Itoa(b.Weight))
}

func IsEqual(f1, f2, p float64) bool {
	return math.Pow(f1, f2) < p
}

const (
	c0 = iota
	c1 = iota
	c2 = iota
)

func SliceMethon() bool {
	var myArray [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var mySlice []int = myArray[:5]
	fmt.Println("Elements of myArray : ")
	for _, v := range myArray {
		fmt.Print(v, " ")
	}
	fmt.Println("\nElements of mySlice: ")
	for _, v := range mySlice {
		fmt.Print(v, " ")
	}
	return true
}

func ForSlice(arr [] int) bool {
	for _, v := range arr {
		print(v, " ")
	}
	fmt.Println()
	return true
}

type PersonInfo struct {
	ID      string
	Name    string
	Address string
}

func myFunc() {
	i := 0
HERE:
	fmt.Println(i)
	i++
	if i < 10 {
		goto HERE
	}
}

func Add(a int, b int) (ret int, err error) {
	if a < 10 || b < 0 {
		err = errors.New("this is test error")
		return
	}
	return a + b, nil
}

func MyFunc2(args ...int) {
	for _, arg := range args {
		fmt.Println(arg)
	}
}

type PathError struct {
	Op   string
	Path string
	Err  error
}

//接口继承
func (e *PathError) Error() string {
	return e.Op + " " + e.Path + " " + e.Err.Error()
}

func CopyFile(dstName, src string) (w int64, err error) {
	srcFile, err := os.Open(src)
	if err != nil {
		return
	}
	defer func() {
		srcFile.Close()
	}()

	dstFile, err := os.Create(dstName)
	if err != nil {
		return
	}
	defer func() {
		dstFile.Close()
	}()

	return io.Copy(dstFile, srcFile)
}

type Integer int

func (a Integer) Less(b Integer) bool {
	return a < b
}

type Rect struct {
	x, y          float64
	width, height float64
}

func (r *Rect) Area() float64 {
	return r.width * r.height
}

func NewRect(x, y, width, height float64) *Rect {
	return &Rect{x, y, width, height}
}

type Base struct {
	Name string
}

func (base *Base) Foo() {
	return
}

func (base *Base) Bar() {
	return
}

type Logger struct {
	Level int
}

type Y struct {
	*Logger
	Name string
}

//实现一个File接口
type File struct{}

func (f *File) Read(buf []byte) (n int, err error) {
	return 0, nil
}

func (f *File) Write(buf []byte) (n int, err error) {
	return 0, nil
}

func (f *File) Seek(off int64, whence int) (pos int64, err error) {
	return 0, nil
}

func (f *File) Colse() (err error) {
	return nil
}



func main() {

	var v1 interface{} = 1
	var v2 interface{} = "abc"

	fmt.Println(v1, v2)
	//rect1 := new(Rect)
	//rect2 := &Rect{}
	//var a Integer = 1
	//if a.Less(2) {
	//	fmt.Println(a, "Less 2")
	//}

	//rst := func(a, b int, z float64) bool {
	//	return a*b < int(z)
	//}(1,2,3.0)
	//fmt.Println(rst)
	//MyFunc2(1,2,2,2,2,2,2)

	//value, err := Add(100, 20)
	//if err == nil {
	//	fmt.Println(value)
	//}
	//myFunc()
	//iSwitchNum := 1
	//switch iSwitchNum {
	//case 0:
	//	fmt.Println("0")
	//case 1:
	//	fmt.Println("1")
	//case 2:
	//	fallthrough
	//case 3:
	//	fmt.Println("3")
	//}

	//sum := 0
	//for {
	//	sum++
	//	if sum > 100 {
	//		break
	//	}
	//}
	//fmt.Println(sum)

	//var personDB map[string] PersonInfo
	//personDB := make(map[string]PersonInfo)
	//personDB["2"] = PersonInfo{"123", "Tom", "Room 203"}
	//personDB["3"] = PersonInfo{"1234", "Terry1", "Room 204"}
	//personDB["4"] = PersonInfo{"12345", "Terry2", "Room 205"}
	//delete(personDB, "2")
	//
	//person, ok := personDB["2"]
	//if ok {
	//	fmt.Println(person.ID)
	//} else {
	//	fmt.Println("Did not find person with ID 2")
	//}
	//mySlice1 := make([] int, 5)
	//mySlice2 := make([] int, 5, 10)
	//mySlice3 := []int{1, 2, 3, 4, 5}
	//mySlice2 = append(mySlice2, 1,2,3,4,5,6,7)

	//fmt.Println("len (mySlice1) : " , len(mySlice2))
	//fmt.Println("cap (mySlice1) : " , cap(mySlice2))
	//
	//println(mySlice1)
	//ForSlice(mySlice1)
	//println(mySlice2)
	//println(mySlice3)

	//SliceMethon();
	//fn, _, _, _ := getName()
	//println(fn)
	//
	//f := func(x, y int) int {
	//	return x + y
	//}
	//println(f(1, 2))
	//
	//var b *Bird = new(Bird)
	//b.Category = "category"
	//b.Weight = 100
	//b.Name = "terry"
	//b.Fly()
	//
	//go run("chris1")
	//go run("chris2")
	//go run("chris3")
	//fmt.Print("This is test!")
	//time.Sleep(time.Duration(2) * time.Second)
	//
	//print(os.Getenv("Home"))
	//
	//println(c0)
	//println(c1)
	//
	//time.Sleep(time.Duration(2) * time.Second)
	//
	//var v1 bool
	//v1 = true
	//v2 := (1 == 2)
	//fmt.Println(v1)
	//fmt.Println(v2)
	//
	//var value2 int32
	//value1 := 64
	//value2 = int32(value1)
	//fmt.Println(value1)
	//fmt.Println(value2)
	//
	//i, j := 1, 2
	//if i == j {
	//	fmt.Println("i and j are equal.")
	//} else {
	//	fmt.Println("i and j is not equal.")
	//}
	//
	//var va1 complex64
	//va1 = 3.2 + 12i
	//va2 := 23.2 + 12i
	//va3 := complex(3.2 , 12)
	//fmt.Println(va1)
	//fmt.Println(va2)
	//fmt.Println(va3)
	//fmt.Println(real(va3))
	//
	//var str string
	//str = "Hello World"
	//ch := str[0]
	//fmt.Printf( "%c" ,ch)
	//
	//str1 := "Hello,世界"
	//n :=len(str1)
	//for i :=0; i<n ; i++  {
	//	ch := str1[i]
	//	fmt.Println(i , ch)
	//}
	//
	//for i,  ch := range str1 {
	//	fmt.Println(i , ch)
	//}

}
