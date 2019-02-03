package main

import "fmt"

func main1() {
	myarr:= []int{1,2,3}
	for  i:=0;i<3;i++{
		fmt.Println(i,myarr[i],&myarr[i]) //数组顺序存放
	}
}

func main() {
	var myarr= []int{10,2,3}
	var myparr[3]* int
	for i:=0;i<3;i++{
		myparr[i]=&myarr[i]
	}
	for i:=0;i<3;i++{
		fmt.Println(*myparr[i],myparr[i])
	}


}