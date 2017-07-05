package main

var x = 0x100
const y = 0x200

func main() {
	println(&x, x)

	/*Note 常量不同于变量在运行期分配存储状态，常量通常在编译器预处理阶段就直接展开，用作指令数据使用*/
	//println(&y, y)
}