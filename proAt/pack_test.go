package main
//テストできるものを羅列する

import(
    "testing"
    "strings"
)

//指定URLのhtmlを取得する
func TestGetHtml(t *testing.T){
    url:="https://www.example.com"
    html:=GetHtml(url)
    
    if(strings.Index(html,"<h1>Example Domain</h1>")==-1){
     t.Error("error")
    }
}


//topのhtmlを受け取ってコンテスト一覧を取得する
//コンテスト一覧を受け取ってコンテスト一覧を表示する

