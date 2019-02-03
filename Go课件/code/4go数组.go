package main

import "fmt"

func main1() {
	var  ids= [5]int{1,2,3,4,5} //定义一个数组有5个int元素
	                    //  0      1          2        3           4
	var names=[5]string{"马龙飞 ","李朋 ","桂升斌 ","谭祎 ","关俊杰"}
	fmt.Println(ids)
	fmt.Println(names)
	fmt.Println(names[4]) //f访问呢数组元素
}

func main2()  {

	var lelebeitai=[5]string{"马龙飞 ","李朋 ","桂升斌 ","谭祎 ","关俊杰"}
	for  i:=0;i<5;i++{
		fmt.Println(i,lelebeitai[i]) //编号，名字
	}



}

func main3(){
	var data[10] int //定义数组有100个int 元素
	fmt.Println(data) //默认初始化0
	for  i:=0;i<10;i++{
		data[i]=i*100  //初始化每个数组元素
	}
	for  i:=0;i<10;i++{
		fmt.Println(data[i]) //显示
	}

}

func main4() {
	var matrix[3][4] int
	fmt.Println(matrix) //2维数组
	var matrixplus[3][4][5] int
	fmt.Println(matrixplus) //3维数组
}

func main5() {
	//matrix:=[3][2] int{{1,2},{3,4},{5,6}} 2维数组
	var  matrix=[3][2] int{{1,2},{3,4},{5,6}}
	for  i:=0 ;i<3;i++{
		for j:=0;j<2;j++{
			fmt.Print(matrix[i][j],"  ") //隔开
		}
		fmt.Println() //换行
	}
	fmt.Println(matrix)
}
func main()  {
	//3维数组
	var  matrix=[3][2][2] int{{{1,2},{3,4}},{{1,2},{3,4}},{{1,2},{3,4}}}
	fmt.Println(matrix)



}