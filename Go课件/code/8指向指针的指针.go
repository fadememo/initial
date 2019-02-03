package main

import "fmt"

func main() {
	var x int =100
	var pint *int =&x
	var ppint **int=&pint
	fmt.Println(*ppint,pint,&x) //等价
	fmt.Println(**ppint,*pint,x)//等价
	var y=200
	*ppint=&y
	fmt.Println(*ppint,pint,&y) //等价
	fmt.Println(**ppint,*pint,y)//等价



}
