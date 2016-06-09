package main

import(
    "./proAt"
    "fmt"
)
func main(){
    contest:=proAt.GetNamedContests("abc",0)
    fmt.Println(contest)
    //fmt.Println(string(body));
}
