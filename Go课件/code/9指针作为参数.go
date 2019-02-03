package main

import "fmt"

//函数参数具备副本机制，
func mvnum(num int ){
	num=1000
	fmt.Println("mvnum",num)
}
//根据地址修改原来的数据
func mvpnum(pnum *int) {
	*pnum=10
}

func main() {
	var num=100
	mvpnum(&num)
	fmt.Println("main",num)
}