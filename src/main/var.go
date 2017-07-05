package main


//import "fmt"
import (
	"os"
)

func fun1() {
	/*Note
		全局变量和局部变量的地址是不一样的
	*/
	x := "abc"
	println(&x, x)
}

func fun2()  {

	x := 100
	println(&x)

	/*Note := 会退化成赋值操作 一般情况下不会这样用*/
	x, y := 200, "abc"

	println(&x, y)
	println(y)

}

func fun3()  {

	x := 100
	println(&x)

	{
		/*Note 不同域中的变量地址也是不同的*/
		x, y := 200, 300
		println(&x, x, &y)
	}

}

func fun4()  {
	f, err := os.Open("./hello.go1")

	buf := make([]byte, 1024)
	n, err := f.Read(buf)

	println(n, err, buf[1])

}

func fun5() {

	const (
		x = iota
		y
		z
	)

	println(x, y, z)
}

func fun6()  {
	const (
		_ = iota
		KB = 1 << (10 * iota)
		MB
		GB
	)

	println(KB, MB, GB)

}

type color byte

func test(c color)  {
	println(c)
}

func fun7()  {
	const (
		balck color = iota
		red
		blue
	)

	test(red)
	test(100)

	/*Note 参数类型不一致*/
	//x := 2
	//test(x)
}

func fun8()  {

	x := 10
	println(&x, x)
	var p *int = &x

	*p += 20
	/*Note p是一个指针类型的变量，这一点和C语言很像*/
	println(&x, x)
	println(p, *p)

}

func fun9()  {

	m := map[string]int{"a":1}
	println(m["a"])
	//println(&m["a"])

}

func fun10()  {
	a := struct {
		x int
	}{}

	a.x = 100

	p := &a
	p.x += 100

	println(p.x, a.x)
}

func fun11()  {
	x := 3

	if x > 5 {
		println("a")
	} else if x < 5 && x > 0 {
		println("b")
	} else {
		println("c")
	}

}

func fun12()  {
	x := 10
	/*Note Go一种特别的语法，分支语句支持预处理*/
	if a, b := x+1, x+2; a > 10 {
		println("a")
	} else if a < 20 {
		println(b)
	} else {
		print(a, b)
	}

	//println(a, b)

}

func main() {
	//fun1()
	//fun2()
	//fun3()
	//fun4()
	//fun5()
	//fun6()
	//fun7()
	//fun8()
	//fun9()
	//fun10()
	//fun11()
	fun12()
}

