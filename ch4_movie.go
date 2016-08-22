package main

import "fmt"
import "encoding/json"
import "log"


type Movie struct {
	Title string `json:"title"`
    // 以下是结构体的成员 tag
    // 是一种元数据在编译的时候绑定到相应的结构上
    // json 开头的指定了转换之后的名字
	Year int `json:"released"`
	Color bool `json:"color,omitempty"`
	Actors []string
}

func main() {
    movies := []Movie{
        {Title: "Casablanca", Year: 1942, Color: false,
            Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
        {Title: "Cool Hand Luke", Year: 1967, Color: true,
            Actors: []string{"Paul Newman"}},
        {Title: "Bullitt", Year: 1968, Color: true,
            Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
    }

    data, err := json.MarshalIndent(movies, "", " ")
    if err != nil {
        log.Fatalf("JSON marshaling failed: %s", err)
    }

    fmt.Printf("%s\n", data)

    // 一次性的结构体类型 不
    var titles []struct{ Title string }

    // 对于结构体的解构 只提取指定的键下面的值
    if err := json.Unmarshal(data, &titles); err != nil {
        log.Fatalf("JSON unmarshaling failed: %s", err)
    }
    fmt.Println(titles)
}
