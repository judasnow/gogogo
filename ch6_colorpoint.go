package main

import "image/color"

type Point struct{ X, Y float64 }


// 大写字母开头的变量会被导出 小写字母开头的则 不会
// 这就是 go 提供的唯一一种封装的方法

// 类型的扩展
// 对于内嵌的类型 可以直接获取被扩展类型的内部成员
// 被扩展类型的方法也会被 "继承"
// method 的调用可以分为两步进行 第一步叫选择器 感觉是函数部分应用(或者是闭包)的一种
// 体现
type ColoredPoint struct {
    Point
    Color color.RGBA
}


func main() {

}
