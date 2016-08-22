package main

import "time"
import (
    "encoding/json"
    "fmt"
    "net/http"
    "net/url"
    "strings"
)

const IssuesURL = "https://api.github.com/search/issues"

type IssuesSearchResult struct {
    TotalCount int `json:"total_count"`
    Items []*Issue
}

type Issue struct {
    Number int
    HTMLURL string `json:html_url`
    Title string
    State string
    User *User
    CreateAt time.Time
    Body string
}

type User struct {
    Login string
    HTMLURL string `json:"html_url"`
}

func main() {

}
