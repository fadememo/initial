package main

import "fmt"

//数组也有副本机制，
func showarr(para [5] string){
	para[0]="王保春"
	fmt.Println("showarr",para)
	for  i:=0;i<5;i++{
		fmt.Println(para[i],i)
	}
}

func main() {
	var names=[5]string{"马龙飞 ","李朋 ","桂升斌 ","谭祎 ","关俊杰"}
	showarr(names)
	fmt.Println("main",names)
}
