package main

import(
    "./proAt"
)
func main(){
    url:="http://atcoder.jp"
    str:=proAt.GetHtml(url)
    proAt.PauseContest(str)

    //fmt.Println(string(body));
}
