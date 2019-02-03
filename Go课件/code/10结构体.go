package main

import "fmt"


//备胎的结构体
type beitai struct {
	name string
	age  int
	money  int
	tall   int
	handsome int
	score int
}

func showbeitai  ( llbeitai beitai)  {
	var zhanghongjie= llbeitai
	fmt.Println(zhanghongjie)
	fmt.Println(zhanghongjie.name)
	fmt.Println(zhanghongjie.age)
	fmt.Println(zhanghongjie.score)
}

func main1() {
	var  zhanghongjie beitai
	zhanghongjie.money=1000
	zhanghongjie.name="张宏杰"
	zhanghongjie.handsome=73
	zhanghongjie.tall=154
	zhanghongjie.age=32
	zhanghongjie.score=77
	showbeitai  ( zhanghongjie)
	var  xiaozhang *beitai=&zhanghongjie //指针指向结构体
	fmt.Println("xiaozhang",xiaozhang.name)


}

func main() {
	//var data[10] int
	var  xiaozhang[3] beitai  //备胎数组
	xiaozhang[0].name="zhanghongjie1"
	xiaozhang[1].name="zhanghongjie2"
	xiaozhang[2].name="zhanghongjie3"
	fmt.Println(xiaozhang)
}


