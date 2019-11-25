package main

import (
	"log"
	"os"

	_ "github.com/goinaction/code/chapter2/sample/matchers"
	"github.com/goinaction/code/chapter2/sample/search"
)

// init在main之前调用
func init() {
	// 将日志输出到标准输出
	log.SetOutput(os.Stdout)
}

// main是整个程序的入口
func main() {
	// 使用特定的项做搜索
	search.Run("president")
}

/**
	学习笔记：
	可以发现main函数保存在main的包里面。如果main函数不在main包里面，构建工具就不会生成可执行文件
	Go语言的每个代码文件都属于一个包，main.go也不理例外。
	包这个特性对于go语言来说很重要

	03--09行，这里申明了所有导入的项，import关键字的作用就是导入一段代码
	其中第四行和第五行为导入标准库，只需要写标准库的名称
	第七行第八行为导入第三方库，其中第七行前面加了_下划线，那是由于
	Go语言的严谨性，不允许导入在页面中没有使用的代码，但是我们可能需要用到那个库里面的init初始化部分，所有在引入的前面加入了下划线
 */
