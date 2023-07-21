package main

import (
    "io"
    "net/http"
    "encoding/json"
    "time"
)

var NewsPaper   string = ""
var OldPaper    string = ""

type source struct{
    Id      string      `json:"id"`
    Name    string      `json:"name"` 
}


type article struct{
   Source       source  `json:"source"`
   Author       string  `json:"author"`
   Title        string  `json:"title"`
   Description  string  `json:"description"`
   Url          string  `json:"url"`
   UrlToImage   string  `json:"urlToImage"`
   PublishedAt  string  `json:"publishedAt"`
   Content      string  `json:"content"`
}


type x struct{
    Status          string  `json:"status"`
    TotalResults    int     `json:"totalResults"`
    //Articles        string  `-`
    Articles        []article `json:"articles"`
}

func GetNews() {
    apiKey  := Config.API 
    url     := "https://newsapi.org/v2/everything"

    req, _ := http.NewRequest("GET", url , nil)
    req.Header.Add("Content-Type", "application/json")
    req.Header.Add("X-Api-Key", apiKey )

    q := req.URL.Query()
    q.Add("sortBy"  , "publishedAt" )

    //ロイター
    //q.Add("sources"  , "reuters"         )
    q.Add("sources"  , "bbc-news"         )

    //検索ワード
    //q.Add("q"       , ""    )
    q.Add("q"       , Config.SearchWord    )

    nowUTC := time.Now().UTC().Format(time.RFC3339)
    q.Add("to"  , nowUTC )

    //一昨日の今の時間から昨日の今の時間までを検索対象に。
    hourUTC := time.Now().UTC().AddDate( 0, 0, -2).Format(time.RFC3339)
    q.Add("from"  , hourUTC )

    //q.Add("domains"  , "bbc.co.uk" )
    //q.Add("excludeDomains"  , "asahi.com,nhk.or.jp" )
    //q.Add("pageSize", "1"         )
    q.Add("pageSize", "100"         )

    req.URL.RawQuery = q.Encode()
    var client *http.Client = &http.Client{}
    resp, _ := client.Do( req )
    defer resp.Body.Close()

    var t x
    body, _ := io.ReadAll(resp.Body)
    //全文
    //fmt.Println( string( body ) )

    json.Unmarshal( body , &t)
    //fmt.Println( t.Status )
    //fmt.Println( len( t.Articles))
    _i := 0
    NewsPaper = ""
    for( _i < len( t.Articles ) ){
        NewsPaper = NewsPaper + "\\_a[OnNewsPaperOpenLink," + t.Articles[_i].Url + "]" + t.Articles[_i].Title + "\\_a\\n\\n"
        //fmt.Print( t.Articles[_i].PublishedAt )
        _i++
    }
    if NewsPaper != "" {
        NewsPaper = "\\0\\b[2]" + NewsPaper 
    }
}











