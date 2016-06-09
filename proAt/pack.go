package proAt

import(
        //"fmt"
        "net/http"
        "io/ioutil"
        "golang.org/x/net/html"
        "strings"
      )

func GetHtml(_url string) string{
    var resp,_=http.Get(_url)
    defer resp.Body.Close()
    body, _ := ioutil.ReadAll(resp.Body)
    //fmt.Println(string(body))
    return string(body)
}
func ParseContest(str string) []string{
    res:=make([]string,0)
    r := strings.NewReader(str)
    z := html.NewTokenizer(r)
    for {
    tt := z.Next()
        switch {
            case tt == html.ErrorToken:
                // End of the document, we're done
                return res
            case tt == html.StartTagToken:
            t := z.Token()
            isAnchor := t.Data == "a"
            if isAnchor {
                res=append(res, t.Attr[0].Val)
            }
        }
    }
}
func FilterContest(list []string)[]string {
    list=FilterByString(list,".contest.atcoder.jp")
    list=Filter(list,func(url string)bool{
        return strings.HasSuffix(url,".contest.atcoder.jp/")
    })
    return list 
}
func FilterByString(list []string, str string)[]string{
    return Filter(list,func(url string)bool{
        return strings.Index(url,str)!=-1
    })
}
func Filter(list []string, f func(string)bool)[]string{
    res:=make([]string,0)
    for _,url :=range(list){
        if f(url){
            res=append(res,url)
        } 
    }
    return res

}
func GetContests(html string)[]string{
    urls:=ParseContest(html)
    contests:=FilterContest(urls)
    return contests;
}
func GetNamedContests(html string,str string,page int)[]string{
    res:=make([]string,0) 
    contests:=GetContests(html)
    contests=FilterByString(contests,str)
    for i:=page*5;i<(page+1)*5&&i<len(contests);i++{
        res=append(res,contests[i])
    }
    return res
}
