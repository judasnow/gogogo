package main

import (
    "errors"
    "fmt"
    "net/http"
    "net/rpc"
)

// 为啥 rpc server 中必须是一个方法 ???
// 函数

type Args struct {
    A int
    B int
}

type Quotient struct {
    Quo int
    Rem int
}

type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
    *reply = args.A * args.B
    return nil
}

func (t *Arith) Divide(args *Args, quo *Quotient) error {
    if args.B == 0 {
        return errors.New("divide by zero")
    } else {
        quo.Quo = args.A / args.B
        quo.Rem = args.A / args.B
        return nil
    }
}

func main() {
    arith := new(Arith)
    rpc.Register(arith)
    rpc.HandleHTTP()

    err := http.ListenAndServe(":1234", nil)
    if err != nil {
        fmt.Println(err.Error())
    }
}
