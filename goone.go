package main

import (
	"math/cmplx"
	"fmt"
	"math"
)

func add(x int, y int) int {
	return x + y
}

func multiplication(x, y int) int {
	return x * y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 21)
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {
	//fmt.Println("Hello, World!")
	//fmt.Println("The time is ", time.Now())
	//fmt.Printf("Now you have %g problems", math.Nextafter(2, 3))
	//fmt.Println(math.Pi)
	//fmt.Println(add(2, 2));
	//fmt.Println(multiplication(2, 5))
	//a, b := swap("A", "B")
	//fmt.Println(a, b)
	//fmt.Println(swap("Chris", "Tan"))
	//fmt.Println(split(45))
	//const f = "%T(%v)\n"
	//fmt.Println(f, z, z)
	//fmt.Println(z)
	//var i int
	//var f float64
	//var b bool
	//var s string
	//fmt.Printf("%v %v %v %q\n", i, f, b, s)

	//var x, y int = 3, 4
	//var f float64 = math.Sqrt(float64(x*x + y*y))
	//var z int = int(f)
	//fmt.Println(x, y, z)

	//var sum int = 0
	//sum := 0
	//for i := 0; i < 10; i++ {
	//	sum += 1
	//}

	//fmt.Println(sum)

	//fmt.Println(pow(3, 2, 10))
	//fmt.Println(pow(3, 3, 20))
	//fmt.Println(MySqrt(4.0))

	//defer fmt.Println("world")
	//fmt.Println("Hello")
	//
	//fmt.Println("counting")
	//
	//for i := 0; i < 10; i++ {
	//	defer fmt.Println(i)
	//}
	//
	//fmt.Println("Done")

	//i, j := 23, 232
	//	//p := &i
	//	//fmt.Println(*p)
	//	//fmt.Println(p)
	//	//*p = 21
	//	//fmt.Println(i)
	//	//p = &j
	//	//*p = int(*p / 37)
	//	//fmt.Println(j)
	//v := Vertex{1, 2}
	//p := &v
	//p.X = 1e9
	//fmt.Println(v)
	//sliceM()

	//m = make(map[string]Demo)
	//m["test"] = Demo{
	//	12.123213, 123123.123123,
	//}
	//fmt.Println(m["test"])
	//
	//tm := make(map[string]int)
	//
	//tm["test"] = 2
	//fmt.Println(tm["test"])
	//
	//tm["test"] = 90
	//fmt.Println(tm["test"])
	//
	//delete(tm, "test")
	//fmt.Println(tm["test"])
	//
	//v, ok := tm["test"]
	//
	//fmt.Println("The value : ", v, " Prsent?", ok)

	//pos, neg := adder(), adder()
	//
	//for i := 0; i < 10; i++ {
	//	fmt.Println(
	//		pos(i),
	//		neg(-2*i),
	//	)
	//}
	v := &Demo2{3, 4}
	fmt.Println(v.Abs())
	v.Scale(5)
	fmt.Println(v, v.Abs())
}

type Demo2 struct {
	X, Y float64
}

func (v *Demo2) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Demo2) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

type Demo struct {
	Lat, Long float64
}

var m map[string]Demo

func sliceM() {

	p := [] int{1, 2, 2, 2, 2, 2, 2, 2,}
	fmt.Println("p == ", len(p))
	p = append(p, 2312321)

	for i := 0; i < len(p); i++ {
		fmt.Println(i)
		fmt.Printf("p[%d] == %d\n", i, p[i])
	}

	for i, v := range p {
		fmt.Printf("p[%d] == %d\n", i, v)
	}

	fmt.Println("p[0:3] ===", p[:3])
}

type Vertex struct {
	X int
	Y int
}

func MySqrt(x float64) float64 {
	z := 1.0
	for {
		tmp := z - (z*z-x)/(2*z)
		fmt.Println(tmp)
		if tmp == z || math.Abs(tmp-z) < 0.000000000001 {
			break
		}
		z = tmp
	}
	return z
}
