package proAt
//テストできるものを羅列する

import(
    "testing"
    "strings"
    "io/ioutil"
     "fmt"
    _ "os"
)

//指定URLのhtmlを取得する
func TestGetHtml(t *testing.T){
    url:="https://www.example.com"
    html:=GetHtml(url)
    if(strings.Index(html,"<h1>Example Domain</h1>")==-1){
     t.Error("error")
    }
}


//topのhtmlを受け取ってurlリストを取得する
func TestParseAtcoder(t *testing.T){
    dat, _ := ioutil.ReadFile("./testfiles/index.html")
    lists:=ParseContest(string(dat))
    if(len(lists)!=320){
        t.Error("error")
    }
}
//urlリストからコンテストだけを抜き出す
func TestFilterContest(t *testing.T){
    list:=[]string{
        "http://abc038.contest.atcoder.jp/data/abc/038/editorial.pdf",
        "http://jag2013summer-day3.contest.atcoder.jp/",
        "http://www.teikoku-vol.com",
    }
    contests:=FilterContest(list)
    if(len(contests)!=1){
        t.Error("cant filter only contests")
    }
}
//コンテスト一覧を取得する
func TestGetContests(t *testing.T){
    dat, _ := ioutil.ReadFile("./testfiles/index.html")
    contests:=GetContests(string(dat))
    //fmt.Println(contests)
    if len(contests)==0{
        t.Error("something")
    }
}
//コンテストのうち最新5件を表示する
func TestNamedContests(t *testing.T){
    dat, _ := ioutil.ReadFile("./testfiles/index.html")
    contests:=GetNamedContests(string(dat),"abc",0)
    fmt.Println(contests)
    if len(contests)!=5{
        t.Error("cant filter only 5")
    }
}
