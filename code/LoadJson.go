package main

import (
    "fmt"
    "os"
    "encoding/json"
)

type NewsAPIConfig struct {
    API             string
    SearchWord      string 
    Count           int
}
var Config NewsAPIConfig

func LoadJson(){
	JsonNewsAPI, err := os.Open( Directory + "/Config.json")
	if err != nil {
        fmt.Println( err )
	}
	defer JsonNewsAPI.Close()
    decoder := json.NewDecoder( JsonNewsAPI )
    err     = decoder.Decode( &Config )
	if err != nil {
        fmt.Println(  err  )
    }
}






