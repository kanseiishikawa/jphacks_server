package main

import (
	"os"
    "io/ioutil"
	"encoding/json"
	//"bytes"
	"fmt"
)

type Form_Day struct {
	Day string `json:"day"`
	Check int `json:check`
}

type Form struct {
	Area string `json:"area"`
	Genre string `json:"genre"`
	Free string `json:"free"`
	Day []Form_Day `json:"select_day"`
}

func main() {
	var form_storage []Form
	file_name := "form_read.json" 

	file, err := os.Open( file_name )

	if err != nil {
		fmt.Println( err )
		return
	}

	defer file.Close()

	// 一気に全部読み取り
    b, err := ioutil.ReadAll( file )

	err = json.Unmarshal( b, &form_storage )

	if err != nil {
		fmt.Println( err )
		return
	}

	fmt.Println( "succes" )
}

