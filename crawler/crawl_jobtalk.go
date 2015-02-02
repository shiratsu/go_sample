package main
 
import (
    "fmt"
    "net/http"
    "io/ioutil"
    "os"
    ""
    //"net/url"
    )

type Result struct {
     Url string
     star float64
}

func main() {
	crawl("http://jobtalk.jp/company/index_1.html")

}


/**
 *	実際にクロールする関数
 */
func crawl(url string){
	response,  err := http.Get(url)
    if err != nil {
        fmt.Printf("%s", err)
        os.Exit(1)
    } else {
        defer response.Body.Close()
        contents, err := ioutil.ReadAll(response.Body)
        if err != nil {
            fmt.Printf("%s", err)
            os.Exit(1)
        }
        fmt.Printf("%s\n", string(contents))
    }
}