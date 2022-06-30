package main

import (
    "flag"
    "os"
    "fmt"
    "io"
    "net/http"
    "encoding/json"
)

func main() {
    flag.Parse()
    args := flag.Args()

    var num string
    if len(args) == 1 {
        num = args[0]
    }else{
        os.Exit(0)
        //num = "13"
    }
    url := "https://www.jma.go.jp/bosai/forecast/data/overview_forecast/" + num + "0000.json" 

    resp, err := http.Get( url )
    if err != nil {
        os.Exit(0)
    }

    defer resp.Body.Close()
    body, _ := io.ReadAll( resp.Body )

    // JSONを構造体にエンコード
    var X strc
    json.Unmarshal(body, &X)

    //fmt.Println( X.PublishingOffice)
    //fmt.Println( X.ReportDatetime)
    //fmt.Println( X.TargetArea)
    fmt.Println( X.Text)
}
type strc struct {
    // 気象台
    PublishingOffice    string `json:"publishingOffice"`
    // 20xx-xx-xxT16:32:00+09:00
    ReportDatetime      string `json:"reportDatetime"`
    // 県
    TargetArea          string `json:"targetArea"` 
    // ?
    HeadlineText        string `json:"headlineText"`
    // 本文
    Text                string `json:"text"`
}






