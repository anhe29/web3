package main

import "fmt"

// func main() {
// }

//指针
//题目 ：编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值。
//考察点 ：指针的使用、值传递与引用传递的区别。/
func task1(a *int) {
	*a += 10
	fmt.Println(*a)
}

func task2(sli []*int) {
	for _, Value := range sli {
		*Value *= 2
	}
}
