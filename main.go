package main

import (
    "fmt"
    "io"
    "net/http"
    "encoding/json"
)

//input     :都道府県コード Int 2桁。
//output    :テキスト。表示方法を考える。配列として渡してパースしてもよい。




func main() {
    url := "https://www.jma.go.jp/bosai/forecast/data/overview_forecast/280000.json" 
    //url := "https://www.jma.go.jp/bosai/forecast/data/overview_forecast/130000.json" 

    resp, err := http.Get( url )
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()
    body, _ := io.ReadAll( resp.Body )

    // JSONを構造体にエンコード
    var X strc
    json.Unmarshal(body, &X)

    fmt.Println( X.PublishingOffice)
    fmt.Println( X.ReportDatetime)
    fmt.Println( X.TargetArea)
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
