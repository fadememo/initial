package main

import "fmt"
//参数有一个副本机制，形式参数a,,b,c当作局部变量
func add(a int ,b int ,c int) {
	a=30
	b=40
	c=50
	fmt.Println(a,b,c)
	fmt.Println("go  add",a+b+c)
}
func main() {
	var x,y,z int  //x,y,z
	x=10
	y=20
	z=30
	add(x,y,z)
	fmt.Println(x,y,z)



}
