/*
golang  执行流程分析
如果对源码编译后,再执行, Go的执行流程:
	go文件 --go build 编译--> 可执行文件(.exe或可执行文件) --运行-->结果
	源文件					可执行文件
	
go run 执行
	
	go文件  ----go run 编译运行一步---->  结果
	
	*/


package main
import "fmt"

func main() {
	fmt.Println("hello,world!")
}

