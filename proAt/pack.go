package main

import(
        "fmt"
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
func PauseContest(str string){

    r := strings.NewReader(str)
   z := html.NewTokenizer(r)
    for {
    tt := z.Next()
        switch {
            case tt == html.ErrorToken:
                // End of the document, we're done
                return
            case tt == html.StartTagToken:
            t := z.Token()
            isAnchor := t.Data == "a"
            if isAnchor {
                fmt.Println("We found a link!")
            }
        }
    }
}
