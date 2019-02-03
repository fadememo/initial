package main

import "fmt"

var  谭祎 string  //全局变量 ,函数外部，多个函数都可以调用

func nimei(){
	var 马龙飞="逗逼" //局部变量，仅仅作用于函数内部
	fmt.Println(马龙飞)
	谭祎="tw  go"
}


func main(){
	var x,y,z int
	x=10  //局部变量
	y=20
	z=30
	fmt.Println(x,y,z)
	//fmt.Println(马龙飞)
	nimei()
	fmt.Println(谭祎)
	var  谭祎 string="你妹的谭祎"
	fmt.Println(谭祎) //局部变量屏蔽同名的全局变量
	nimei()
}