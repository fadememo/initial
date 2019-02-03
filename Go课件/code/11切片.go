package main

import (
	"fmt"
)

func main1() {
	           //0，1，2，3，4
	data:= []int{1,2,3,4,5}
	fmt.Println(data)
	fmt.Println(data[0:3]) //切片，0-3
	fmt.Println(data[3:]) //索引是3的到最后一个
	fmt.Println(data[:2])//索引从第一个到小于索引2
	fmt.Println(data[:])//全部
	fmt.Println(len(data))//长度是5
	fmt.Println(cap(data))//最大的切片长度
}

func showslice(x [] int)  {
	fmt.Println(len(x),cap(x),x)
}

func main2() {
	//data:= []int{1,2,3,4,5}
	//fmt.Println(data)

	var  myslice [] int=make([]int,4,5) //make创造用于切片的数组
	fmt.Println(myslice,myslice[:]  )
	showslice(myslice)
}

func main3() {
	var  data []int
	if data==nil{
		fmt.Println("空切片")
	}

	fmt.Println(data)
}

func main4() {
	var  num  []int
	num=append(num,1) //追加
	fmt.Println(num)
	num=append(num,2)
	fmt.Println(num)
	num=append(num,3,4,5)
	fmt.Println(num)
	var  data  []int =make([]int ,10,10) //开辟空间
	fmt.Println(data)
	copy(data,num) //拷贝数据
	fmt.Println(data)
}
//锯齿数组，
func main() {
	data :=make([][]int ,3)
	fmt.Println(data)
	for i:=0;i<3;i++{
		length:=i+1
		data[i]  = make([]int ,length,length)
		for j:=0;j<length;j++{
			data[i][j]=i+j
		}
	}
	fmt.Println(data)

}