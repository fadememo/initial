package main

import "fmt"

func main1() {
	var num=10
	fmt.Println(num,&num) //数据，数据的地址
}

func main2() {
	var  num=10
	var pnum *int =&num  //pnum是指针，存储地址
	fmt.Println(num,&num,pnum)
	*pnum=1//取出地址对应内容赋值，间接修改
	fmt.Println(num,*pnum)
	num=3 //直接修改
	fmt.Println(num,*pnum)
}

func main() {
	//var num=10
	//var pnum *int =&num//空指针
	var pnum *int
	fmt.Println(pnum)
	if pnum==nil{
		fmt.Println("这个指针妹子没有对象")
	}else {
		fmt.Println("这个指针妹子有对象")
	}

}