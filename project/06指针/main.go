package main

import "fmt"

////=====================================================================================================

// 指针 1.题目 ：编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值。 考察点 ：指针的使用、值传递与引用传递的区别。
func setAndGet(p *int) int {
	*p += 10

	return *p
}

// 指针 2.题目 ：实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。考察点 ：指针运算、切片操作。
func setAndGetPtr(p []*int) []*int {
	for k, v := range p {
		fmt.Println("第", k, "个元素", *v)
		*v *= 2
		fmt.Println("第", k, "个元素", *v)
	}

	return p
}

func main() {

	// 指针 1.题目 ：编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值。 考察点 ：指针的使用、值传递与引用传递的区别。
	g := 15
	fmt.Println(setAndGet(&g))
	fmt.Println(g)

	// 指针 2.题目 ：实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。考察点 ：指针运算、切片操作。
	var a = 1
	var b = 2
	var c = 3
	var d = 4
	var e = 5
	var f = []*int{&a, &b, &c, &d, &e}
	fmt.Println(setAndGetPtr(f))
	fmt.Println(f)
	for k, v := range f {
		fmt.Println("第", k, "个元素", *v)
	}

}
