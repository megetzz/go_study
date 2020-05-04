//开发hello.go, 输出helloworld
//go文件的后缀是 .go
// package main  表示该 hello.go 文件所在的包是main
// 在go中,每个文件都必须归属一个包
//3. import "fmt"
 // 表示 引入一个包 报名位fmt,引入改包后,就可以使用fmt包的函数 eg:fmt.Println
// 4. func main(){}
//  func是一个关键字,表示一个函数  
//  main 是函数名,是一个主函数,即我们程序的入口
//5. fmt.Println("hello")
// 表示调用 fmt包的函数 Println输出helloworld

//dos下
// 编译 go build xxx.go  对go文件进行编译,生成 xxx.exe 文件 
// 运行 xxx.exe 

//注意:go run 命令可以直接执行xxx.go 程序,(类似执行一个脚本文件)
package main
import "fmt"

func main() {
	fmt.Println("hello,world!!!")
}

