package main

import (
	"fmt"
	"math"
	"runtime"
	"strings"
	"time"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}

	return lim
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func mySqrt(x float64) float64 {
	z := 1.0

	for i := 0; i < 10; i++ {
		a := z*z - x
		b := 2 * z
		c := a / b

		tempResult := z - c

		if tempResult == z {
			break
		}

		z = tempResult
	}

	return z
}

func checkRuntime() {
	fmt.Println(runtime.GOOS)
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("MAC")

	case "linux":
		fmt.Println("LINUX")
	default:
		fmt.Printf("%s.\n", os)
	}

}

func getSaturday() {
	fmt.Println("Когда суббота?")
	today := time.Now().Weekday()
	fmt.Println(today)
	switch time.Saturday {
	case today + 0:
		fmt.Println("Седне")
	case today + 1:
		fmt.Println("Завтра")
	case today + 2:
		fmt.Println("Послезавтра")
	case today + 3:
		fmt.Println("Послепослезавтра")
	case today + 4:
		fmt.Println("Послепослепослезавтра")
	case today + 5:
		fmt.Println("Послепослепослепослезавтра")
	default:
		fmt.Println("что такое суббота")
	}
}

func getGreetings() {
	currHour := time.Now().Hour()

	switch {
	case currHour < 12:
		fmt.Println("доброе утро")
	case currHour < 17:
		fmt.Println("добрый день")
	default:
		fmt.Println("добрый вечер")
	}
}

func deferFirst() {
	defer fmt.Println("hey you")
	defer fmt.Println("world")
	defer fmt.Println("please")
	fmt.Println("Hello")
}

func pointers() {
	i, j := 42, 2701

	p := &i
	fmt.Println(*p)
	*p = 84
	fmt.Println(i)

	p = &j
	fmt.Println(*p)
	*p = *p / 37
	fmt.Println(j)
}

type Vertex struct {
	X int
	Y int
}

func getStructs() {
	a := Vertex{X: 1, Y: 2}
	a.X = 16
	fmt.Println(a)

	v := Vertex{3, 4}
	p := &v
	p.X = 1e9
	fmt.Println(v)

	v1 := Vertex{5, 6}
	v2 := Vertex{X: 5}
	v3 := Vertex{}
	pp := &Vertex{1, 2}
	ppp := &v1

	fmt.Println(v1, v2, v3, *pp, *ppp)
}

func getArrays() {
	var arr [3]string
	arr[0] = "Hello"
	arr[1] = "World"

	fmt.Println(arr)

	primes := [5]int{2, 3, 5, 7, 11}

	fmt.Println(primes)
}

func getSlices() {
	arr := [4]int{1, 2, 3, 4}

	var slice []int = arr[0:2]
	fmt.Println(slice)

	names := [5]string{"foo", "bar", "baz", "koz", "boz"}
	fmt.Println(names)

	var slice1 []string = names[0:2]
	var slice2 []string = names[1:3]
	fmt.Println(slice1, slice2)

	slice1[1] = "hui"

	fmt.Println("after slice change")
	fmt.Println(names)
	fmt.Println(slice1, slice2)
}

func slicesLikeArrays() {
	q := []int{1, 2, 3, 4, 5}
	fmt.Println(q)

	str := []struct {
		x int
		y bool
	}{
		{2, false},
		{1, true},
	}
	fmt.Println(str)
}

func defaultSlices() {
	arr := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(arr[:])
	fmt.Println(arr[:10])
	fmt.Println(arr[0:])
	fmt.Println(arr[0:10])
}

func lenAndCapSlices() {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	printSliceInfo(arr)

	arr = arr[:4]
	printSliceInfo(arr)

	arr = arr[:0]
	printSliceInfo(arr)

	arr = arr[:5]
	printSliceInfo(arr)

	arr = arr[1:]
	printSliceInfo(arr)

	arr = arr[:9]
	printSliceInfo(arr)
}

func printSliceInfo(givenSlice []int) {
	fmt.Printf("len=%d, cap=%d %v \n", len(givenSlice), cap(givenSlice), givenSlice)
}

func printSliceInfoWithParam(s string, slice []int) {
	fmt.Printf("%s : len=%d, cap=%d %v \n", s, len(slice), cap(slice), slice)
}

func nilSlices() {
	var ns []int

	printSliceInfo(ns)
	fmt.Println(ns == nil)
}

func makeSlices() {
	a := make([]int, 5)
	printSliceInfoWithParam("a", a)

	b := make([]int, 0, 5)
	printSliceInfoWithParam("b", b)

	c := b[:2]
	printSliceInfoWithParam("c", c)

	d := c[2:5]
	printSliceInfoWithParam("d", d)

}

func Pic(dx, dy int) [][]uint8 {
	slice1 := make([][]uint8, dy)

	for y := range slice1 {
		slice1[y] = make([]uint8, dx)

		for x := range slice1[y] {
			val := (x + y) / 2
			slice1[y][x] = uint8(val)
		}
	}

	return slice1
}

type Maptex struct {
	Lat, Long float64
}

func firstMaps() {
	var m map[string]Maptex

	m = make(map[string]Maptex)

	m["Bell Labs"] = Maptex{
		40.68433, -74.39967,
	}

	fmt.Println(m)
	fmt.Println(m["Bell Labs"])

	mm := map[string]Maptex{
		"Bell Labs": Maptex{
			40.68433, -74.39967,
		},
		"Google": {
			37.42202, -122.08408,
		},
	}

	fmt.Println(mm)
}

func printMapValue(m map[string]int, val int) {
	fmt.Printf("Map is: %v, value: %d \n", m, val)
}

func mapMutations() {
	m := make(map[string]int)
	m["Answer"] = 42
	printMapValue(m, m["Answer"])

	m["New answer"] = 54
	printMapValue(m, m["New answer"])

	delete(m, "Answer")

	value, ok := m["Answer"]
	newValue, newOk := m["New answer"]
	fmt.Println("The value of answer:", value, "Present?", ok)
	fmt.Println("The value of new answer:", newValue, "Present?", newOk)
}

func WordsCount(s string) map[string]int {
	result := make(map[string]int)

	var words []string = strings.Fields(s)

	fmt.Println(result)

	for _, word := range words {
		_, ok := result[word]

		if ok == false {
			result[word] = 1
		} else {
			result[word] += 1
		}
	}

	return result
}

func adder() func(int) int {
	sum := 0

	return func(x int) int {
		sum += x
		return sum
	}
}

func closers() {
	pos, neg := adder(), adder()

	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i), neg(2*-i))
	}
}

func Fibbonaci() func() int {
	prev := 0
	cur := 1
	return func() int {

		result := prev

		prev, cur = cur, prev+cur

		return result
	}
}

func closers2() {
	f := Fibbonaci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

type Floatex struct {
	X, Y float64
}

func (v Floatex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Methods1() {
	v := Floatex{3, 4}
	vSquared := v.Abs()
	fmt.Println(vSquared)
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func Methods2() {
	myF := MyFloat(-math.Sqrt2)
	fmt.Println(myF.Abs())
}

///////////////////////

type Methodex struct {
	X, Y float64
}

func (v Methodex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (m *Methodex) Scale(f float64) {
	m.X = m.X * f
	m.Y = m.Y * f
}

func Methods3() {
	m := Methodex{3, 4}
	fmt.Println(m)
	m.Scale(10)
	fmt.Println(m)
	fmt.Println(m.Abs())
	fmt.Println(m)
}

// errors

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf(
		"at %v, %s",
		e.When, e.What,
	)
}

func run() error {
	return &MyError{
		time.Now(),
		"shit happened",
	}
}

// Sqrt with error

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func mySqrtWithError(x float64) (float64, error) {
	if x < 0 {
		return x, ErrNegativeSqrt(x)
	}

	z := 1.0

	for i := 0; i < 10; i++ {
		a := z*z - x
		b := 2 * z
		c := a / b

		tempResult := z - c

		if tempResult == z {
			break
		}

		z = tempResult
	}

	return z, nil
}

func main() {
	fmt.Println(mySqrtWithError(2))
	fmt.Println(mySqrtWithError(-2))

	fmt.Println(ErrNegativeSqrt(-2).Error())
}
