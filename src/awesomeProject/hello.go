package main

/*
#include <stdio.h>
#include <stdlib.h>
*/

import "fmt"
import "C"

func sum(values [] int, resultChan chan int){
	sum := 0
	for _, value := range values{
		sum += value
	}
	resultChan <- sum //将计算结果发送到channel中
}

func main(){
	fmt.Println("Hello World!")

	// 并行计算实例
	values := [] int {1,2,3,4,5,6,7,8,9,10}

	resultChan := make (chan int, 2)
	go sum(values[:len(values)/2], resultChan)
	go sum(values[len(values)/2:], resultChan)
	sum1, sum2 := <- resultChan, <- resultChan // 接收结果

	fmt.Println("result:", sum1, sum2, sum1 + sum2)
	// 实例到此结束

	// Cgo
	/*
	cstr := C.CString("HELLO WORLD C!")
	C.puts(cstr)
	C.free(unsafe.Pointer(cstr)
	 */




}
