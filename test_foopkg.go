package main

import (
    // 包的导入是到目录一级 而不是到文件一级
    "gogogo/foopkg"
)

// 这里不能直接给包外的类型添加新的方法
// func (p Pointfoo) foobar() error {
//    return nil
// }

func main() {
    var p = new(foopkg.Pointfoo)
    // 大写的结构体成员就可以导出
    p.Name = "1024"
    // 小写的结构体成员不能导出
    // p.x = 1024
    // 小写的方法名字也不会被导出
    // p.distance()
}
