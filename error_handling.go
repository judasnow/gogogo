package main

import (
    "os"
    "io"
    "fmt"
    "errors"
    "bytes"
)

// 用 error 接口类型标识错误 定义如下:
// type error interface {
//     Error() string
// }

// 使用 errors 的 New 方法来创建一个新的 errorString 值

type dbError struct {
    s string
}
func (e *dbError) Error() string {
    return e.s
}

func Sqrt(f float64) (float64, error) {
    if f < 0 {
        // 格式化信息 并直接生成一个 error 类型
        fmt.Errorf("math: square root of negative number %g", f)
        return 0, errors.New("math: square root of negative number")
    } else {
        return 1024, nil
    }
}

// 定义特殊的 Error 类型 以充分利用类型系统
type NegativeSqrtError float64
func (f NegativeSqrtError) Error() string {
    return fmt.Sprintf("math: square root of negative number %g", float64(f))
}

func main() {
    // nil 是有效的 error 类型值
    // var err = argError{}
    // err = nil
    err := errors.New("1024")
    fmt.Printf("%T\n", err)

    f, err := Sqrt(-1)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(f)
    }

    // 类型断言 使用在 接口值上
    // 接口不是一个具体的类型 接口的值应该有两个分量
    // 其一是具体类型 其二是此具体类型的值
    // 以下是类型的三个值
    var w io.Writer

    // 接口类型将被初始化为 nil
    // %T 不能检测出非具体类型
    fmt.Printf("%T\n", w)

    w = os.Stdout
    fmt.Printf("%T\n", w)

    w = new(bytes.Buffer)
    w = nil

    // var w io.Writer
    // w = os.Stdout
    // q := w.(*os.File)
    // c := w.(*bytes.Buffer)

    // 类型断言的格式是 x.(T)
    // x 是一个接口类型( expression x of interface type ) 而 T 是一个类型
    // 这里就分为两种不同的情况
    // 1 T 是具体类型
    // 2 T 是接口类型

	/*
	asserts that x is not nil and that the value stored in x is of type T.
	The notation x.(T) is called a type assertion.

 	More precisely, if T is not an interface type, x.(T) asserts that the
 		dynamic type of x is identical to the type T. In this case, T must implement the (interface) type of x;
 		otherwise the type assertion is invalid since it is not possible for x to store a value of type T.
 		 If T is an interface type, x.(T) asserts that the dynamic type of x implements the interface T.

	If the type assertion holds, the value of the expression is the
		 value stored in x and its type is T.
		 If the type assertion is false, a run-time panic occurs.
		 In other words, even though the dynamic type of x is known only at run time,
		 the type of x.(T) is known to be T in a correct program.
	*/

    // fmt.Print(q, c)
}
