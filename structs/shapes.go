package structs

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

// 返回矩形面积
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// 返回矩形周常
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

type Circle struct {
	Radius float64
}

// 返回圆的面积
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// 三角形
type Triangle struct {
	Base   float64
	Height float64
}

// 返回三角形的面积
func (c Triangle) Area() float64 {
	return (c.Base * c.Height) * 0.5
}
