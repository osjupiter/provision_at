package main

import(
        "fmt"
        "net/http"
        "io/ioutil"
        "golang.org/x/net/html"
        "io"
      )

func getHtml(_url string) io.Reader{
    var resp,_=http.Get(_url)
    defer resp.Body.Close()
    // body, _ := ioutil.ReadAll(resp.Body)
    return resp.Body 
}
func pauseContest(reader Reader){

   z := html.NewTokenizer(reader )
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
func main(){
    url:="http://atcoder.jp"
    pauseContest(getHtml(url))

    //fmt.Println(string(body));
}
