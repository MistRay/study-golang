package main

import "fmt"

const boilingF = 212.0

//Go语言主要有四种类型的声明语句：var、const、type和func，分别对应变量、常量、类型和函数实体对象的声明。
func main() {
	var f = boilingF
	var c = computeC(f)
	fmt.Printf("boiling point = %g °F or %g°C \n", f, c)

}

// func 函数名 (入参名称 入参类型) 出参类型
func computeC(f float64) float64 {
	return (f - 32) * 5 / 9
}
