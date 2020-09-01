package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	// 声明变量
	var s string
	s = "10"
	fmt.Println(s)

	// 简短声明变量
	x := "20"
	x = "30"
	fmt.Println(x)

	// 简短声明一组变量
	y, z := 15, 20
	fmt.Println(y + z)

	// 用两个变量接收函数返回值
	f, err := os.Open("test")
	fmt.Println(f)
	fmt.Println(err)
	point()
	niltest()

	var server httpServer
	http.Handle("/", server)
}

func (server httpServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(r.URL.Path))
}

type httpServer struct {
}

func point() {
	// 简短声明了变量x
	x := 1
	// 获取了x变量的指针p
	p := &x
	// p指针保存了x变量的内存地址
	fmt.Println(p)
	// 通过p获取到了x的值
	fmt.Println(*p)
	// 通过p可以给x赋值
	*p = 2
	fmt.Println(x)
}

func incr(p *int) int {
	*p++ // 非常重要：只是增加p指向的变量的值，并不改变p指针！！！
	return *p
}

func niltest() {
	var x, y int
	// 两个未声明变量的指针不相同,且未声明指针和nil的相等测试为false
	fmt.Println(&x == &x, &x == &y, &x == nil) // "true false false"
}
